package keeper

import (
    "context"
	"fmt"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "micin/x/dex/types"
)

type msgServer struct {
    Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
    return &msgServer{Keeper: keeper}
}

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    pool, found := k.GetPool(ctx, msg.PoolId)
    if !found {
        return nil, fmt.Errorf("pool with id %d not found", msg.PoolId)
    }

    pool.ReserveA += msg.AmountA
    pool.ReserveB += msg.AmountB

    k.SetPool(ctx, pool)

    return &types.MsgAddLiquidityResponse{}, nil
}


// Pastikan msgServer mengimplementasikan interface types.MsgServer
var _ types.MsgServer = msgServer{}
