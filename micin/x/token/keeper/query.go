package keeper

import (
	"micin/x/token/types"
)

var _ types.QueryServer = Keeper{}
