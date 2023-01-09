package peyote

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/warmage-sports/peyote/x/peyote/internal/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Initialise peyote
	for _, b := range data.Bonds {
		keeper.SetBond(ctx, b.Token, b)
	}

	// Initialise batches
	for _, b := range data.Batches {
		keeper.SetBatch(ctx, b.Token, b)
	}

	// Initialise params
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export peyote and batches
	var peyote []types.Bond
	var batches []types.Batch
	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bond := k.MustGetBondByKey(ctx, iterator.Key())
		batch := k.MustGetBatch(ctx, bond.Token)
		peyote = append(peyote, bond)
		batches = append(batches, batch)
	}

	// Export params
	params := k.GetParams(ctx)

	return GenesisState{
		Bonds:   peyote,
		Batches: batches,
		Params:  params,
	}
}
