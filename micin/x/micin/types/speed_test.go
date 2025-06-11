package types

import (
	"fmt"
)

type SpeedTest struct {
	Creator   string `json:"creator" yaml:"creator"`
	SpeedMbps int64  `json:"speed_mbps" yaml:"speed_mbps"`
	Timestamp int64  `json:"timestamp" yaml:"timestamp"`
}

func (s SpeedTest) Validate() error {
	if s.SpeedMbps <= 0 {
		return fmt.Errorf("speed must be > 0")
	}
	if s.Timestamp <= 0 {
		return fmt.Errorf("timestamp must be > 0")
	}
	return nil
}
