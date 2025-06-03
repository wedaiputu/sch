package token_test

import (
	"testing"

	keepertest "micin/testutil/keeper"
	"micin/testutil/nullify"
	token "micin/x/token/module"
	"micin/x/token/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CoinList: []types.Coin{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CoinCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenKeeper(t)
	token.InitGenesis(ctx, k, genesisState)
	got := token.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CoinList, got.CoinList)
	require.Equal(t, genesisState.CoinCount, got.CoinCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
