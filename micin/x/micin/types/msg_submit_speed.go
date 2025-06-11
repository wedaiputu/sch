package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSubmitSpeedTest defines a message for submitting internet speed

// NewMsgSubmitSpeedTest creates a new MsgSubmitSpeedTest instance
func NewMsgSubmitSpeedTest(creator string, speedMbps, timestamp int64) *MsgSubmitSpeedTest {
	return &MsgSubmitSpeedTest{
		Creator:   creator,
		SpeedMbps: speedMbps,
		Timestamp: timestamp,
	}
}

// Route returns the message route for routing
func (msg *MsgSubmitSpeedTest) Route() string {
	return RouterKey
}

// Type returns the message type
func (msg *MsgSubmitSpeedTest) Type() string {
	return "SubmitSpeedTest"
}

// GetSigners returns the addresses of signers that must sign
func (msg *MsgSubmitSpeedTest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the bytes for the message signer to sign on
func (msg *MsgSubmitSpeedTest) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(fmt.Errorf("failed to marshal msg: %w", err))
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does basic validation of fields
func (msg *MsgSubmitSpeedTest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}
	if msg.SpeedMbps <= 0 {
		return fmt.Errorf("invalid speed: must be > 0")
	}
	if msg.Timestamp <= 0 {
		return fmt.Errorf("invalid timestamp: must be > 0")
	}
	return nil
}
