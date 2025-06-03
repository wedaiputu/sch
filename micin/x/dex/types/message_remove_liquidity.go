package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveLiquidity{}

func NewMsgRemoveLiquidity(creator string, sender string, tokenA string, tokenB string, share string) *MsgRemoveLiquidity {
	return &MsgRemoveLiquidity{
		Creator: creator,
		Sender:  sender,
		TokenA:  tokenA,
		TokenB:  tokenB,
		Share:   share,
	}
}

func (msg *MsgRemoveLiquidity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
