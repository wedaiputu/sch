package keeper

import (
	"context"
	"encoding/binary"

	"micin/x/token/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetCoinCount get the total number of coin
func (k Keeper) GetCoinCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CoinCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCoinCount set the total number of coin
func (k Keeper) SetCoinCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CoinCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCoin appends a coin in the store with a new id and update the count
func (k Keeper) AppendCoin(
	ctx context.Context,
	coin types.Coin,
) uint64 {
	// Create the coin
	count := k.GetCoinCount(ctx)

	// Set the ID of the appended value
	coin.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CoinKey))
	appendedValue := k.cdc.MustMarshal(&coin)
	store.Set(GetCoinIDBytes(coin.Id), appendedValue)

	// Update coin count
	k.SetCoinCount(ctx, count+1)

	return count
}

// SetCoin set a specific coin in the store
func (k Keeper) SetCoin(ctx context.Context, coin types.Coin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CoinKey))
	b := k.cdc.MustMarshal(&coin)
	store.Set(GetCoinIDBytes(coin.Id), b)
}

// GetCoin returns a coin from its id
func (k Keeper) GetCoin(ctx context.Context, id uint64) (val types.Coin, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CoinKey))
	b := store.Get(GetCoinIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCoin removes a coin from the store
func (k Keeper) RemoveCoin(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CoinKey))
	store.Delete(GetCoinIDBytes(id))
}

// GetAllCoin returns all coin
func (k Keeper) GetAllCoin(ctx context.Context) (list []types.Coin) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CoinKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Coin
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCoinIDBytes returns the byte representation of the ID
func GetCoinIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.CoinKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
