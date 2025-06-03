package keeper

import (
	"context"
	"fmt"

	"micin/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreatePool membuat pool baru, meng-assign ID baru secara otomatis.
func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Buat objek Pool dengan data dari msg
	pool := types.Pool{
		Creator:  msg.Creator,
		TokenA:   msg.TokenA,
		TokenB:   msg.TokenB,
		ReserveA: msg.ReserveA,
		ReserveB: msg.ReserveB,
	}

	// Tambahkan pool ke store, dapatkan ID baru
	id := k.AppendPool(ctx, pool)

	return &types.MsgCreatePoolResponse{
		Id: id,
	}, nil
}

// UpdatePool update data pool, hanya bisa dilakukan oleh creator pool
func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Cek pool dengan ID yang dimaksud ada atau tidak
	existingPool, found := k.GetPool(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pool id %d doesn't exist", msg.Id))
	}

	// Validasi creator pesan sama dengan pemilik pool
	if msg.Creator != existingPool.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Update pool dengan data baru dari msg
	updatedPool := types.Pool{
		Id:       msg.Id,
		Creator:  msg.Creator,
		TokenA:   msg.TokenA,
		TokenB:   msg.TokenB,
		ReserveA: msg.ReserveA,
		ReserveB: msg.ReserveB,
	}

	k.SetPool(ctx, updatedPool)

	return &types.MsgUpdatePoolResponse{}, nil
}

// DeletePool menghapus pool, hanya bisa dilakukan oleh creator pool
func (k msgServer) DeletePool(goCtx context.Context, msg *types.MsgDeletePool) (*types.MsgDeletePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Cek pool ada atau tidak
	existingPool, found := k.GetPool(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("pool id %d doesn't exist", msg.Id))
	}

	// Validasi creator pesan sama dengan pemilik pool
	if msg.Creator != existingPool.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Hapus pool dari store
	k.RemovePool(ctx, msg.Id)

	return &types.MsgDeletePoolResponse{}, nil
}
