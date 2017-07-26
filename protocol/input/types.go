// Code generated by cdpgen. DO NOT EDIT.

package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TouchPoint
//
// Note: This type is experimental.
type TouchPoint struct {
	// State State of the touch point.
	//
	// Values: "touchPressed", "touchReleased", "touchMoved", "touchStationary", "touchCancelled".
	State         string   `json:"state"`
	X             int      `json:"x"`                       // X coordinate of the event relative to the main frame's viewport.
	Y             int      `json:"y"`                       // Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX       *int     `json:"radiusX,omitempty"`       // X radius of the touch area (default: 1).
	RadiusY       *int     `json:"radiusY,omitempty"`       // Y radius of the touch area (default: 1).
	RotationAngle *float64 `json:"rotationAngle,omitempty"` // Rotation angle (default: 0.0).
	Force         *float64 `json:"force,omitempty"`         // Force (default: 1.0).
	ID            *float64 `json:"id,omitempty"`            // Identifier used to track touch sources between events, must be unique within an event.
}

// GestureSourceType
//
// Note: This type is experimental.
type GestureSourceType int

// GestureSourceType as enums.
const (
	GestureSourceTypeNotSet GestureSourceType = iota
	GestureSourceTypeDefault
	GestureSourceTypeTouch
	GestureSourceTypeMouse
)

// Valid returns true if enum is set.
func (e GestureSourceType) Valid() bool {
	return e >= 1 && e <= 3
}

func (e GestureSourceType) String() string {
	switch e {
	case 0:
		return "GestureSourceTypeNotSet"
	case 1:
		return "default"
	case 2:
		return "touch"
	case 3:
		return "mouse"
	}
	return fmt.Sprintf("GestureSourceType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e GestureSourceType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("input.GestureSourceType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *GestureSourceType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"default\"":
		*e = 1
	case "\"touch\"":
		*e = 2
	case "\"mouse\"":
		*e = 3
	default:
		return fmt.Errorf("input.GestureSourceType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64

// String calls (time.Time).String().
func (t TimeSinceEpoch) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t TimeSinceEpoch) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("input.TimeSinceEpoch: " + err.Error())
	}
	*t = TimeSinceEpoch(f)
	return nil
}

var _ json.Marshaler = (*TimeSinceEpoch)(nil)
var _ json.Unmarshaler = (*TimeSinceEpoch)(nil)
