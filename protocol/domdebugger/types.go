// Code generated by cdpgen. DO NOT EDIT.

package domdebugger

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/runtime"
)

// DOMBreakpointType DOM breakpoint type.
type DOMBreakpointType int

// DOMBreakpointType as enums.
const (
	DOMBreakpointTypeNotSet DOMBreakpointType = iota
	DOMBreakpointTypeSubtreeModified
	DOMBreakpointTypeAttributeModified
	DOMBreakpointTypeNodeRemoved
)

// Valid returns true if enum is set.
func (e DOMBreakpointType) Valid() bool {
	return e >= 1 && e <= 3
}

func (e DOMBreakpointType) String() string {
	switch e {
	case 0:
		return "DOMBreakpointTypeNotSet"
	case 1:
		return "subtree-modified"
	case 2:
		return "attribute-modified"
	case 3:
		return "node-removed"
	}
	return fmt.Sprintf("DOMBreakpointType(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e DOMBreakpointType) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("domdebugger.DOMBreakpointType: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *DOMBreakpointType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"subtree-modified\"":
		*e = 1
	case "\"attribute-modified\"":
		*e = 2
	case "\"node-removed\"":
		*e = 3
	default:
		return fmt.Errorf("domdebugger.DOMBreakpointType: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// EventListener Object event listener.
//
// Note: This type is experimental.
type EventListener struct {
	Type            string                `json:"type"`                      // EventListener's type.
	UseCapture      bool                  `json:"useCapture"`                // EventListener's useCapture.
	Passive         bool                  `json:"passive"`                   // EventListener's passive flag.
	Once            bool                  `json:"once"`                      // EventListener's once flag.
	ScriptID        runtime.ScriptID      `json:"scriptId"`                  // Script id of the handler code.
	LineNumber      int                   `json:"lineNumber"`                // Line number in the script (0-based).
	ColumnNumber    int                   `json:"columnNumber"`              // Column number in the script (0-based).
	Handler         *runtime.RemoteObject `json:"handler,omitempty"`         // Event handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"` // Event original handler function value.
	BackendNodeID   *dom.BackendNodeID    `json:"backendNodeId,omitempty"`   // Node the listener is added to (if any).
}
