package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&Bond{}, "peyote/Bond", nil)
	cdc.RegisterConcrete(&FunctionParam{}, "peyote/FunctionParam", nil)
	cdc.RegisterConcrete(&Batch{}, "peyote/Batch", nil)
	cdc.RegisterConcrete(&BaseOrder{}, "peyote/BaseOrder", nil)
	cdc.RegisterConcrete(&BuyOrder{}, "peyote/BuyOrder", nil)
	cdc.RegisterConcrete(&SellOrder{}, "peyote/SellOrder", nil)
	cdc.RegisterConcrete(&SwapOrder{}, "peyote/SwapOrder", nil)
	cdc.RegisterConcrete(MsgCreateBond{}, "peyote/MsgCreateBond", nil)
	cdc.RegisterConcrete(MsgEditBond{}, "peyote/MsgEditBond", nil)
	cdc.RegisterConcrete(MsgBuy{}, "peyote/MsgBuy", nil)
	cdc.RegisterConcrete(MsgSell{}, "peyote/MsgSell", nil)
	cdc.RegisterConcrete(MsgSwap{}, "peyote/MsgSwap", nil)
	cdc.RegisterConcrete(MsgMakeOutcomePayment{}, "peyote/MsgMakeOutcomePayment", nil)
	cdc.RegisterConcrete(MsgWithdrawShare{}, "peyote/MsgWithdrawShare", nil)
}
