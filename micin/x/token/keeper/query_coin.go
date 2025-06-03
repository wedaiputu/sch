package keeper

import (
	"context"

	"micin/x/token/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CoinAll(ctx context.Context, req *types.QueryAllCoinRequest) (*types.QueryAllCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var coins []types.Coin

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	coinStore := prefix.NewStore(store, types.KeyPrefix(types.CoinKey))

	pageRes, err := query.Paginate(coinStore, req.Pagination, func(key []byte, value []byte) error {
		var coin types.Coin
		if err := k.cdc.Unmarshal(value, &coin); err != nil {
			return err
		}

		coins = append(coins, coin)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoinResponse{Coin: coins, Pagination: pageRes}, nil
}

func (k Keeper) Coin(ctx context.Context, req *types.QueryGetCoinRequest) (*types.QueryGetCoinResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coin, found := k.GetCoin(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCoinResponse{Coin: coin}, nil
}
