package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, tokenA string, tokenB string, reserveA uint64, reserveB uint64) *MsgCreatePool {
	return &MsgCreatePool{
		Creator:  creator,
		TokenA:   tokenA,
		TokenB:   tokenB,
		ReserveA: reserveA,
		ReserveB: reserveB,
	}
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(creator string, id uint64, tokenA string, tokenB string, reserveA uint64, reserveB uint64) *MsgUpdatePool {
	return &MsgUpdatePool{
		Id:       id,
		Creator:  creator,
		TokenA:   tokenA,
		TokenB:   tokenB,
		ReserveA: reserveA,
		ReserveB: reserveB,
	}
}

func (msg *MsgUpdatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePool{}

func NewMsgDeletePool(creator string, id uint64) *MsgDeletePool {
	return &MsgDeletePool{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeletePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
