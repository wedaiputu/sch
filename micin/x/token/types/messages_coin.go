package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCoin{}

func NewMsgCreateCoin(creator string, name string, supply string, owner string) *MsgCreateCoin {
	return &MsgCreateCoin{
		Creator: creator,
		Name:    name,
		Supply:  supply,
		Owner:   owner,
	}
}

func (msg *MsgCreateCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCoin{}

func NewMsgUpdateCoin(creator string, id uint64, name string, supply string, owner string) *MsgUpdateCoin {
	return &MsgUpdateCoin{
		Id:      id,
		Creator: creator,
		Name:    name,
		Supply:  supply,
		Owner:   owner,
	}
}

func (msg *MsgUpdateCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCoin{}

func NewMsgDeleteCoin(creator string, id uint64) *MsgDeleteCoin {
	return &MsgDeleteCoin{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteCoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
