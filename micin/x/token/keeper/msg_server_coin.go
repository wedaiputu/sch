package keeper

import (
	"context"
	"fmt"

	"micin/x/token/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCoin(goCtx context.Context, msg *types.MsgCreateCoin) (*types.MsgCreateCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var coin = types.Coin{
		Creator: msg.Creator,
		Name:    msg.Name,
		Supply:  msg.Supply,
		Owner:   msg.Owner,
	}

	id := k.AppendCoin(
		ctx,
		coin,
	)

	return &types.MsgCreateCoinResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCoin(goCtx context.Context, msg *types.MsgUpdateCoin) (*types.MsgUpdateCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var coin = types.Coin{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Supply:  msg.Supply,
		Owner:   msg.Owner,
	}

	// Checks that the element exists
	val, found := k.GetCoin(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCoin(ctx, coin)

	return &types.MsgUpdateCoinResponse{}, nil
}

func (k msgServer) DeleteCoin(goCtx context.Context, msg *types.MsgDeleteCoin) (*types.MsgDeleteCoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetCoin(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCoin(ctx, msg.Id)

	return &types.MsgDeleteCoinResponse{}, nil
}
