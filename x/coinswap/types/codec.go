package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec registers concrete types on the codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSwapOrder{}, "nexus/MsgSwapOrder", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "nexus/MsgAddLiquidity", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidity{}, "nexus/MsgRemoveLiquidity", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "nexus/x/coinswap/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&Params{}, "nexus/x/coinswap/Params", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwapOrder{},
		&MsgAddLiquidity{},
		&MsgRemoveLiquidity{},
		&MsgUpdateParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
