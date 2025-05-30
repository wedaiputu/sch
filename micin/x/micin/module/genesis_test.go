package micin_test

import (
	"testing"

	keepertest "micin/testutil/keeper"
	"micin/testutil/nullify"
	micin "micin/x/micin/module"
	"micin/x/micin/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MicinKeeper(t)
	micin.InitGenesis(ctx, k, genesisState)
	got := micin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
