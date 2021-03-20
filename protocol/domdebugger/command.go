// Code generated by cdpgen. DO NOT EDIT.

package domdebugger

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/runtime"
)

// GetEventListenersArgs represents the arguments for GetEventListeners in the DOMDebugger domain.
type GetEventListenersArgs struct {
	ObjectID runtime.RemoteObjectID `json:"objectId"`         // Identifier of the object to return listeners for.
	Depth    *int                   `json:"depth,omitempty"`  // The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Pierce   *bool                  `json:"pierce,omitempty"` // Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
}

// NewGetEventListenersArgs initializes GetEventListenersArgs with the required arguments.
func NewGetEventListenersArgs(objectID runtime.RemoteObjectID) *GetEventListenersArgs {
	args := new(GetEventListenersArgs)
	args.ObjectID = objectID
	return args
}

// SetDepth sets the Depth optional argument. The maximum depth at
// which Node children should be retrieved, defaults to 1. Use -1 for
// the entire subtree or provide an integer larger than 0.
func (a *GetEventListenersArgs) SetDepth(depth int) *GetEventListenersArgs {
	a.Depth = &depth
	return a
}

// SetPierce sets the Pierce optional argument. Whether or not iframes
// and shadow roots should be traversed when returning the subtree
// (default is false). Reports listeners for all contexts if pierce is
// enabled.
func (a *GetEventListenersArgs) SetPierce(pierce bool) *GetEventListenersArgs {
	a.Pierce = &pierce
	return a
}

// GetEventListenersReply represents the return values for GetEventListeners in the DOMDebugger domain.
type GetEventListenersReply struct {
	Listeners []EventListener `json:"listeners"` // Array of relevant listeners.
}

// RemoveDOMBreakpointArgs represents the arguments for RemoveDOMBreakpoint in the DOMDebugger domain.
type RemoveDOMBreakpointArgs struct {
	NodeID dom.NodeID        `json:"nodeId"` // Identifier of the node to remove breakpoint from.
	Type   DOMBreakpointType `json:"type"`   // Type of the breakpoint to remove.
}

// NewRemoveDOMBreakpointArgs initializes RemoveDOMBreakpointArgs with the required arguments.
func NewRemoveDOMBreakpointArgs(nodeID dom.NodeID, typ DOMBreakpointType) *RemoveDOMBreakpointArgs {
	args := new(RemoveDOMBreakpointArgs)
	args.NodeID = nodeID
	args.Type = typ
	return args
}

// RemoveEventListenerBreakpointArgs represents the arguments for RemoveEventListenerBreakpoint in the DOMDebugger domain.
type RemoveEventListenerBreakpointArgs struct {
	EventName string `json:"eventName"` // Event name.
	// TargetName EventTarget interface name.
	//
	// Note: This property is experimental.
	TargetName *string `json:"targetName,omitempty"`
}

// NewRemoveEventListenerBreakpointArgs initializes RemoveEventListenerBreakpointArgs with the required arguments.
func NewRemoveEventListenerBreakpointArgs(eventName string) *RemoveEventListenerBreakpointArgs {
	args := new(RemoveEventListenerBreakpointArgs)
	args.EventName = eventName
	return args
}

// SetTargetName sets the TargetName optional argument. EventTarget
// interface name.
//
// Note: This property is experimental.
func (a *RemoveEventListenerBreakpointArgs) SetTargetName(targetName string) *RemoveEventListenerBreakpointArgs {
	a.TargetName = &targetName
	return a
}

// RemoveInstrumentationBreakpointArgs represents the arguments for RemoveInstrumentationBreakpoint in the DOMDebugger domain.
type RemoveInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// NewRemoveInstrumentationBreakpointArgs initializes RemoveInstrumentationBreakpointArgs with the required arguments.
func NewRemoveInstrumentationBreakpointArgs(eventName string) *RemoveInstrumentationBreakpointArgs {
	args := new(RemoveInstrumentationBreakpointArgs)
	args.EventName = eventName
	return args
}

// RemoveXHRBreakpointArgs represents the arguments for RemoveXHRBreakpoint in the DOMDebugger domain.
type RemoveXHRBreakpointArgs struct {
	URL string `json:"url"` // Resource URL substring.
}

// NewRemoveXHRBreakpointArgs initializes RemoveXHRBreakpointArgs with the required arguments.
func NewRemoveXHRBreakpointArgs(url string) *RemoveXHRBreakpointArgs {
	args := new(RemoveXHRBreakpointArgs)
	args.URL = url
	return args
}

// SetBreakOnCSPViolationArgs represents the arguments for SetBreakOnCSPViolation in the DOMDebugger domain.
type SetBreakOnCSPViolationArgs struct {
	ViolationTypes []CSPViolationType `json:"violationTypes"` // CSP Violations to stop upon.
}

// NewSetBreakOnCSPViolationArgs initializes SetBreakOnCSPViolationArgs with the required arguments.
func NewSetBreakOnCSPViolationArgs(violationTypes []CSPViolationType) *SetBreakOnCSPViolationArgs {
	args := new(SetBreakOnCSPViolationArgs)
	args.ViolationTypes = violationTypes
	return args
}

// SetDOMBreakpointArgs represents the arguments for SetDOMBreakpoint in the DOMDebugger domain.
type SetDOMBreakpointArgs struct {
	NodeID dom.NodeID        `json:"nodeId"` // Identifier of the node to set breakpoint on.
	Type   DOMBreakpointType `json:"type"`   // Type of the operation to stop upon.
}

// NewSetDOMBreakpointArgs initializes SetDOMBreakpointArgs with the required arguments.
func NewSetDOMBreakpointArgs(nodeID dom.NodeID, typ DOMBreakpointType) *SetDOMBreakpointArgs {
	args := new(SetDOMBreakpointArgs)
	args.NodeID = nodeID
	args.Type = typ
	return args
}

// SetEventListenerBreakpointArgs represents the arguments for SetEventListenerBreakpoint in the DOMDebugger domain.
type SetEventListenerBreakpointArgs struct {
	EventName string `json:"eventName"` // DOM Event name to stop on (any DOM event will do).
	// TargetName EventTarget interface name to stop on. If equal to `"*"`
	// or not provided, will stop on any EventTarget.
	//
	// Note: This property is experimental.
	TargetName *string `json:"targetName,omitempty"`
}

// NewSetEventListenerBreakpointArgs initializes SetEventListenerBreakpointArgs with the required arguments.
func NewSetEventListenerBreakpointArgs(eventName string) *SetEventListenerBreakpointArgs {
	args := new(SetEventListenerBreakpointArgs)
	args.EventName = eventName
	return args
}

// SetTargetName sets the TargetName optional argument. EventTarget
// interface name to stop on. If equal to `"*"` or not provided, will
// stop on any EventTarget.
//
// Note: This property is experimental.
func (a *SetEventListenerBreakpointArgs) SetTargetName(targetName string) *SetEventListenerBreakpointArgs {
	a.TargetName = &targetName
	return a
}

// SetInstrumentationBreakpointArgs represents the arguments for SetInstrumentationBreakpoint in the DOMDebugger domain.
type SetInstrumentationBreakpointArgs struct {
	EventName string `json:"eventName"` // Instrumentation name to stop on.
}

// NewSetInstrumentationBreakpointArgs initializes SetInstrumentationBreakpointArgs with the required arguments.
func NewSetInstrumentationBreakpointArgs(eventName string) *SetInstrumentationBreakpointArgs {
	args := new(SetInstrumentationBreakpointArgs)
	args.EventName = eventName
	return args
}

// SetXHRBreakpointArgs represents the arguments for SetXHRBreakpoint in the DOMDebugger domain.
type SetXHRBreakpointArgs struct {
	URL string `json:"url"` // Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
}

// NewSetXHRBreakpointArgs initializes SetXHRBreakpointArgs with the required arguments.
func NewSetXHRBreakpointArgs(url string) *SetXHRBreakpointArgs {
	args := new(SetXHRBreakpointArgs)
	args.URL = url
	return args
}
