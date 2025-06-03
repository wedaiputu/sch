package keeper

import (
	"context"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
	"micin/x/micin/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority  string
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,
		bankKeeper:   bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}
func (k Keeper) MintTokenBySpeed(ctx sdk.Context, creator string, speedMbps int64) error {
	if speedMbps < 0 {
		return fmt.Errorf("speedMbps must be positive")
	}

	amount := math.NewInt(speedMbps * 1000) // ✅ gunakan math.NewInt
	coin := sdk.NewCoin("micin", amount)

	addr, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin)); err != nil {
		return err
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(coin)); err != nil {
		return err
	}

	return nil
}



// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k msgServer) SubmitSpeedTest(goCtx context.Context, msg *types.MsgSubmitSpeedTest) (*types.MsgSubmitSpeedTestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Mint token berdasarkan kecepatan Mbps
	err := k.Keeper.MintTokenBySpeed(ctx, msg.Creator, msg.SpeedMbps)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitSpeedTestResponse{}, nil
}
