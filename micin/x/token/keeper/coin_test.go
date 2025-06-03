package keeper_test

import (
	"context"
	"testing"

	keepertest "micin/testutil/keeper"
	"micin/testutil/nullify"
	"micin/x/token/keeper"
	"micin/x/token/types"

	"github.com/stretchr/testify/require"
)

func createNCoin(keeper keeper.Keeper, ctx context.Context, n int) []types.Coin {
	items := make([]types.Coin, n)
	for i := range items {
		items[i].Id = keeper.AppendCoin(ctx, items[i])
	}
	return items
}

func TestCoinGet(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNCoin(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCoin(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCoinRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNCoin(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCoin(ctx, item.Id)
		_, found := keeper.GetCoin(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCoinGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNCoin(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCoin(ctx)),
	)
}

func TestCoinCount(t *testing.T) {
	keeper, ctx := keepertest.TokenKeeper(t)
	items := createNCoin(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCoinCount(ctx))
}
