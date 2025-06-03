package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"micin/x/micin/codec" 
)

// MsgSubmitSpeedTest defines a message for submitting internet speed
type MsgSubmitSpeedTest struct {
	Creator    string
	SpeedMbps  int64
	Timestamp  int64
}

func NewMsgSubmitSpeedTest(creator string, speedMbps, timestamp int64) *MsgSubmitSpeedTest {
	return &MsgSubmitSpeedTest{
		Creator:   creator,
		SpeedMbps: speedMbps,
		Timestamp: timestamp,
	}
}

// Implement sdk.Msg interface
func (msg *MsgSubmitSpeedTest) Route() string { return RouterKey }
func (msg *MsgSubmitSpeedTest) Type() string  { return "SubmitSpeedTest" }

func (msg *MsgSubmitSpeedTest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSpeedTest) GetSignBytes() []byte {
	return sdk.MustSortJSON(codec.ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSubmitSpeedTest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return err
	}
	if msg.SpeedMbps <= 0 {
		return fmt.Errorf("invalid speed: must be > 0")
	}
	return nil
}
