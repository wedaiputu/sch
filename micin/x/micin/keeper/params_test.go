package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "micin/testutil/keeper"
	"micin/x/micin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MicinKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
