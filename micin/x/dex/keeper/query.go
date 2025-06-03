package keeper

import (
	"micin/x/dex/types"
)

var _ types.QueryServer = Keeper{}
