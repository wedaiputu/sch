package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"micin/x/token/keeper"
	"micin/x/token/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the coin
	for _, elem := range genState.CoinList {
		k.SetCoin(ctx, elem)
	}

	// Set coin count
	k.SetCoinCount(ctx, genState.CoinCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CoinList = k.GetAllCoin(ctx)
	genesis.CoinCount = k.GetCoinCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
