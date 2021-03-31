package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgBuyName{}, "cbdc/BuyName", nil)
	cdc.RegisterConcrete(MsgSetName{}, "cbdc/SetName", nil)
	cdc.RegisterConcrete(MsgDeleteName{}, "cbdc/DeleteName", nil)
	cdc.RegisterConcrete(MsgPayment{}, "cbdc/payment", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
