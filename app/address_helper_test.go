package app

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	nexusconfig "github.com/FDSLabs/Nexus/cmd/config"
)

func init() {
	config := sdk.GetConfig()
	nexusconfig.SetBech32Prefixes(config)
	nexusconfig.SetBip44CoinType(config)
	config.Seal()
}

// TestGenerateAddresses is a helper test that generates addresses with the correct prefix
// Run this test with: go test -v ./app -run TestGenerateAddresses
func TestGenerateAddresses(t *testing.T) {
	// Verify config is set correctly
	config := sdk.GetConfig()
	require.Equal(t, nexusconfig.Bech32PrefixAccAddr, config.GetBech32AccountAddrPrefix())
	require.Equal(t, nexusconfig.Bech32PrefixAccPub, config.GetBech32AccountPubPrefix())
	require.Equal(t, nexusconfig.Bech32PrefixValAddr, config.GetBech32ValidatorAddrPrefix())
	require.Equal(t, nexusconfig.Bech32PrefixValPub, config.GetBech32ValidatorPubPrefix())
	require.Equal(t, nexusconfig.Bech32PrefixConsAddr, config.GetBech32ConsensusAddrPrefix())
	require.Equal(t, nexusconfig.Bech32PrefixConsPub, config.GetBech32ConsensusPubPrefix())

	// 1. Generate regular bech32 addresses
	for i := 0; i < 5; i++ {
		privKey := ed25519.GenPrivKey()
		address := sdk.AccAddress(privKey.PubKey().Address())
		fmt.Printf("New bech32 address %d: %s\n", i+1, address.String())
		require.True(t, address.String()[:5] == "nexus")
	}

	// 2. Convert existing EVM addresses to bech32
	hexAddresses := []string{
		"0x925a4b1175Bdf6EBE517d68B8D64D392D238903D",
		"0x9dDc0FF8D6D4a81f73B8c7067E218f3A7e8d3675",
	}

	fmt.Printf("\nEVM address conversions:\n")
	for _, hexAddr := range hexAddresses {
		hexAddr = strings.TrimPrefix(hexAddr, "0x")
		addrBytes, err := hex.DecodeString(hexAddr)
		require.NoError(t, err)
		bech32Addr := sdk.AccAddress(addrBytes)
		fmt.Printf("EVM address: 0x%s\nBech32 address: %s\n\n", hexAddr, bech32Addr.String())
		require.True(t, bech32Addr.String()[:5] == "nexus")
	}

	// 3. Generate new EVM addresses and their bech32 equivalents
	fmt.Printf("\nNew EVM addresses and their bech32 equivalents:\n")
	for i := 0; i < 3; i++ {
		privKey := ed25519.GenPrivKey()
		evmAddress := common.BytesToAddress(privKey.PubKey().Address().Bytes())
		bech32Addr := sdk.AccAddress(evmAddress.Bytes())
		fmt.Printf("Pair %d:\n", i+1)
		fmt.Printf("  EVM: 0x%s\n", evmAddress.Hex()[2:])
		fmt.Printf("  Bech32: %s\n", bech32Addr.String())
		require.True(t, bech32Addr.String()[:5] == "nexus")
	}

	// 4. Generate validator addresses
	fmt.Printf("\nValidator addresses:\n")
	for i := 0; i < 2; i++ {
		privKey := ed25519.GenPrivKey()
		valAddr := sdk.ValAddress(privKey.PubKey().Address())
		fmt.Printf("Validator %d: %s\n", i+1, valAddr.String())
		require.Equal(t, nexusconfig.Bech32PrefixValAddr, valAddr.String()[:len(nexusconfig.Bech32PrefixValAddr)])
	}
}
