// Code generated by cdpgen. DO NOT EDIT.

package emulation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ScreenOrientation Screen orientation.
type ScreenOrientation struct {
	// Type Orientation type.
	//
	// Values: "portraitPrimary", "portraitSecondary", "landscapePrimary", "landscapeSecondary".
	Type  string `json:"type"`
	Angle int    `json:"angle"` // Orientation angle.
}

// VirtualTimePolicy advance: If the scheduler runs out of immediate work, the virtual time base may fast forward to allow the next delayed task (if any) to run; pause: The virtual time base may not advance; pauseIfNetworkFetchesPending: The virtual time base may not advance if there are any pending resource fetches.
//
// Note: This type is experimental.
type VirtualTimePolicy int

// VirtualTimePolicy as enums.
const (
	VirtualTimePolicyNotSet VirtualTimePolicy = iota
	VirtualTimePolicyAdvance
	VirtualTimePolicyPause
	VirtualTimePolicyPauseIfNetworkFetchesPending
)

// Valid returns true if enum is set.
func (e VirtualTimePolicy) Valid() bool {
	return e >= 1 && e <= 3
}

func (e VirtualTimePolicy) String() string {
	switch e {
	case 0:
		return "VirtualTimePolicyNotSet"
	case 1:
		return "advance"
	case 2:
		return "pause"
	case 3:
		return "pauseIfNetworkFetchesPending"
	}
	return fmt.Sprintf("VirtualTimePolicy(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e VirtualTimePolicy) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("emulation.VirtualTimePolicy: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *VirtualTimePolicy) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"advance\"":
		*e = 1
	case "\"pause\"":
		*e = 2
	case "\"pauseIfNetworkFetchesPending\"":
		*e = 3
	default:
		return fmt.Errorf("emulation.VirtualTimePolicy: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}
