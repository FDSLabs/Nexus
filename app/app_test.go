package app

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	dbm "github.com/cosmos/cosmos-db"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/state"
	cbfttypes "github.com/cometbft/cometbft/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store/iavl"
	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/migrations/v1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"

	"github.com/FDSLabs/Nexus/types"
	coinswaptypes "github.com/FDSLabs/Nexus/x/coinswap/types"
)

func TestNexusExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewNexus(
		log.NewLogger(os.Stdout),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		0,
		false,
		simtestutil.NewAppOptionsWithFlagHome(DefaultNodeHome),
		baseapp.SetChainID(types.TestnetChainID+"-1"),
	)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		&abci.RequestInitChain{
			ChainId:         types.TestnetChainID + "-1",
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)
	app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: app.LastBlockHeight() + 1,
	})
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewNexus(
		log.NewLogger(os.Stdout),
		db,
		nil,
		true,
		map[int64]bool{},
		DefaultNodeHome,
		0,
		false,
		simtestutil.NewAppOptionsWithFlagHome(DefaultNodeHome),
		baseapp.SetChainID(types.TestnetChainID+"-1"),
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestPrintModuleAddresses(t *testing.T) {
	PrintModuleAddresses()
}

// TestWorkingHash tests that the working hash of the IAVL store is calculated correctly during the initialization phase of the genesis, given the initial height specified in the GenesisDoc.
func TestWorkingHash(t *testing.T) {
	gdoc, err := state.MakeGenesisDocFromFile("height4-genesis.json")
	require.NoError(t, err)

	gs, err := state.MakeGenesisState(gdoc)
	require.NoError(t, err)

	fmt.Printf("Genesis Time: %v\n", gdoc.GenesisTime)
	fmt.Printf("Chain ID: %s\n", gdoc.ChainID)
	fmt.Printf("Initial Height: %d\n", gdoc.InitialHeight)

	tmpDir := "test-working-hash"
	db, err := dbm.NewGoLevelDB("test", tmpDir, nil)
	require.NoError(t, err)
	app := NewNexus(log.NewNopLogger(), db, nil, true, map[int64]bool{}, DefaultNodeHome, 0, false, simtestutil.NewAppOptionsWithFlagHome(DefaultNodeHome), baseapp.SetChainID(gdoc.ChainID))

	// delete tmpDir
	defer require.NoError(t, os.RemoveAll(tmpDir))

	pbparams := gdoc.ConsensusParams.ToProto()

	// Use genesis time for all time-related fields to ensure determinism
	blockTime := gdoc.GenesisTime
	fmt.Printf("Block Time: %v\n", blockTime)

	// Initialize the chain
	_, err = app.InitChain(&abci.RequestInitChain{
		Time:            blockTime,
		ChainId:         gdoc.ChainID,
		ConsensusParams: &pbparams,
		Validators:      cbfttypes.TM2PB.ValidatorUpdates(gs.Validators),
		AppStateBytes:   gdoc.AppState,
		InitialHeight:   gdoc.InitialHeight,
	})
	require.NoError(t, err)

	// Call FinalizeBlock to calculate each module's working hash.
	// Without calling this, all module's root node will have empty hash.
	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: gdoc.InitialHeight,
		Time:   blockTime,
	})
	require.NoError(t, err)

	storeKeys := app.GetStoreKeys()
	// Print all store keys first
	fmt.Println("\nAll Store Keys:")
	for _, key := range storeKeys {
		if key != nil {
			fmt.Printf("- %s\n", key.Name())
		}
	}

	fmt.Println("\nWorking Hashes:")
	// deterministicKeys are module keys which has always same working hash whenever run this test. (non deterministic module: staking, epoch, inflation)
	deterministicKeys := []string{
		authtypes.StoreKey, banktypes.StoreKey, capabilitytypes.StoreKey, coinswaptypes.StoreKey,
		consensustypes.StoreKey, crisistypes.StoreKey, distrtypes.StoreKey,
		evmtypes.StoreKey, feemarkettypes.StoreKey, govtypes.StoreKey, ibctransfertypes.StoreKey,
		paramstypes.StoreKey, slashingtypes.StoreKey, upgradetypes.StoreKey}

	// workingHashWithZeroInitialHeight is the working hash of the IAVL store with initial height 0 with given genesis.
	workingHashWithZeroInitialHeight := map[string]string{
		authtypes.StoreKey:        "d5a1c439fcdc7ead036e0d7e4bad83e03dc110de7ea458c1530f0ffafd028fd6",
		banktypes.StoreKey:        "13e1f59bca5c07be9ee4083b4948b19d755dfd21515e4f23235461a4aed1afc0",
		capabilitytypes.StoreKey:  "e9261548b1c687638721f75920e1c8e3f4f52cbd7ab3aeddc6c626cd8abc8718",
		coinswaptypes.StoreKey:    "364e8d434e86b3a798d9f47fd76c01cd0b36c2ce9a1f730ac5578237468cde29",
		consensustypes.StoreKey:   "2573460b2ef3a08c825bdfb485e51680038530f70c45f0b723167ae09599761c",
		crisistypes.StoreKey:      "5d2ad1048f5362960fd8687c7ca139232c7424f05089132246b51170d4b01e7b",
		distrtypes.StoreKey:       "8a0cd15b50e41f1e6e27e478cd572f6bc70a030cc9c1df19a7ba2c41406ddbd7",
		evmtypes.StoreKey:         "eb7310a9b1b0769be17af2b5c4da320c470e5838b69ee0055c35bc6f012132b3",
		feemarkettypes.StoreKey:   "b8a66cce8e7809f521db9fdd71bfeb980966f1b74ef252bda65804d8b89da7de",
		govtypes.StoreKey:         "18d5736ff5b173e9dfefcd2a8df73e068f1d405d055bd5876db70ba0c3450e13",
		ibctransfertypes.StoreKey: "3ffd548eb86288efc51964649e36dc710f591c3d60d6f9c1b42f2a4d17870904",
		paramstypes.StoreKey:      "75d8cb34ec47d90b5167052ca285df18789fbb7f08cc8fa55a370b5729491b89",
		slashingtypes.StoreKey:    "4144393b0227e2b46247ca33f38f4a5e9694973c4738cd6d6b5464bd949468fa",
		upgradetypes.StoreKey:     "9677219870ca98ba9868589ccdcd97411d9b82abc6a7aa1910016457b4730a93",
	}

	matchAny := func(key string) bool {
		for _, dk := range deterministicKeys {
			if dk == key {
				return true
			}
		}
		return false
	}

	for _, key := range storeKeys {
		if key != nil && matchAny(key.Name()) {
			kvstore := app.CommitMultiStore().GetCommitKVStore(key)
			require.Equal(t, storetypes.StoreTypeIAVL, kvstore.GetStoreType())
			iavlStore, ok := kvstore.(*iavl.Store)
			require.True(t, ok)
			workingHash := hex.EncodeToString(iavlStore.WorkingHash())
			fmt.Printf("Module: %s\n", key.Name())
			fmt.Printf("  Expected: %s\n", workingHashWithZeroInitialHeight[key.Name()])
			fmt.Printf("  Actual:   %s\n", workingHash)
			require.Equal(t, workingHashWithZeroInitialHeight[key.Name()], workingHash, key.Name())
		}
	}
}

func TestPrintEthAccountAddress(t *testing.T) {
	// The old address with invalid checksum
	oldAddress := "nexus1jfdykyt4hhmwhegh669c6exnjtfr3yparej8y8"

	// Extract the human readable part and data part
	parts := strings.Split(oldAddress, "1")
	require.Equal(t, 2, len(parts))

	// The data part without checksum
	data := parts[1][:len(parts[1])-6]
	fmt.Printf("Data without checksum: %s\n", data)

	// Construct the correct address
	correctAddress := fmt.Sprintf("nexus1%s3wne3u", data)
	fmt.Printf("Correct address: %s\n", correctAddress)

	// Verify the address is valid
	_, err := sdk.AccAddressFromBech32(correctAddress)
	require.NoError(t, err)
}

func TestPrintEVMAddresses(t *testing.T) {
	// Convert hex addresses to bech32
	hexAddresses := []string{
		"0x925a4b1175Bdf6EBE517d68B8D64D392D238903D",
		"0x9dDc0FF8D6D4a81f73B8c7067E218f3A7e8d3675",
	}

	for _, hexAddr := range hexAddresses {
		// Remove 0x prefix
		hexAddr = strings.TrimPrefix(hexAddr, "0x")
		// Convert to bytes
		addrBytes, err := hex.DecodeString(hexAddr)
		require.NoError(t, err)
		// Convert to bech32
		bech32Addr := sdk.AccAddress(addrBytes)
		fmt.Printf("Hex address: %s\nBech32 address: %s\n\n", hexAddr, bech32Addr.String())
	}
}

func TestPrintConsensusAddress(t *testing.T) {
	// Convert hex consensus address to bech32
	hexAddr := "CFF57A76F8200F0358989D9ABED90F9B8E2F5D80"
	addrBytes, err := hex.DecodeString(hexAddr)
	require.NoError(t, err)
	// Convert to bech32
	bech32Addr := sdk.ConsAddress(addrBytes)
	fmt.Printf("Hex address: %s\nBech32 address: %s\n\n", hexAddr, bech32Addr.String())
}
