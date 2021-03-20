// Code generated by cdpgen. DO NOT EDIT.

package dom

import (
	"encoding/json"

	"github.com/mafredri/cdp/protocol/internal"
)

// NodeID Unique DOM node identifier.
type NodeID int

// BackendNodeID Unique DOM node identifier used to reference a node that may
// not have been pushed to the front-end.
type BackendNodeID int

// BackendNode Backend node with a friendly name.
type BackendNode struct {
	NodeType      int           `json:"nodeType"`      // `Node`'s nodeType.
	NodeName      string        `json:"nodeName"`      // `Node`'s nodeName.
	BackendNodeID BackendNodeID `json:"backendNodeId"` // No description.
}

// PseudoType Pseudo element type.
type PseudoType string

// PseudoType as enums.
const (
	PseudoTypeNotSet              PseudoType = ""
	PseudoTypeFirstLine           PseudoType = "first-line"
	PseudoTypeFirstLetter         PseudoType = "first-letter"
	PseudoTypeBefore              PseudoType = "before"
	PseudoTypeAfter               PseudoType = "after"
	PseudoTypeMarker              PseudoType = "marker"
	PseudoTypeBackdrop            PseudoType = "backdrop"
	PseudoTypeSelection           PseudoType = "selection"
	PseudoTypeTargetText          PseudoType = "target-text"
	PseudoTypeSpellingError       PseudoType = "spelling-error"
	PseudoTypeGrammarError        PseudoType = "grammar-error"
	PseudoTypeFirstLineInherited  PseudoType = "first-line-inherited"
	PseudoTypeScrollbar           PseudoType = "scrollbar"
	PseudoTypeScrollbarThumb      PseudoType = "scrollbar-thumb"
	PseudoTypeScrollbarButton     PseudoType = "scrollbar-button"
	PseudoTypeScrollbarTrack      PseudoType = "scrollbar-track"
	PseudoTypeScrollbarTrackPiece PseudoType = "scrollbar-track-piece"
	PseudoTypeScrollbarCorner     PseudoType = "scrollbar-corner"
	PseudoTypeResizer             PseudoType = "resizer"
	PseudoTypeInputListButton     PseudoType = "input-list-button"
)

func (e PseudoType) Valid() bool {
	switch e {
	case "first-line", "first-letter", "before", "after", "marker", "backdrop", "selection", "target-text", "spelling-error", "grammar-error", "first-line-inherited", "scrollbar", "scrollbar-thumb", "scrollbar-button", "scrollbar-track", "scrollbar-track-piece", "scrollbar-corner", "resizer", "input-list-button":
		return true
	default:
		return false
	}
}

func (e PseudoType) String() string {
	return string(e)
}

// ShadowRootType Shadow root type.
type ShadowRootType string

// ShadowRootType as enums.
const (
	ShadowRootTypeNotSet    ShadowRootType = ""
	ShadowRootTypeUserAgent ShadowRootType = "user-agent"
	ShadowRootTypeOpen      ShadowRootType = "open"
	ShadowRootTypeClosed    ShadowRootType = "closed"
)

func (e ShadowRootType) Valid() bool {
	switch e {
	case "user-agent", "open", "closed":
		return true
	default:
		return false
	}
}

func (e ShadowRootType) String() string {
	return string(e)
}

// Node DOM interaction is implemented in terms of mirror objects that
// represent the actual DOM nodes. DOMNode is a base node mirror type.
type Node struct {
	NodeID           NodeID                `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend will only push node with given `id` once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentID         *NodeID               `json:"parentId,omitempty"`         // The id of the parent node if any.
	BackendNodeID    BackendNodeID         `json:"backendNodeId"`              // The BackendNodeId for this node.
	NodeType         int                   `json:"nodeType"`                   // `Node`'s nodeType.
	NodeName         string                `json:"nodeName"`                   // `Node`'s nodeName.
	LocalName        string                `json:"localName"`                  // `Node`'s localName.
	NodeValue        string                `json:"nodeValue"`                  // `Node`'s nodeValue.
	ChildNodeCount   *int                  `json:"childNodeCount,omitempty"`   // Child count for `Container` nodes.
	Children         []Node                `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string              `json:"attributes,omitempty"`       // Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	DocumentURL      *string               `json:"documentURL,omitempty"`      // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL          *string               `json:"baseURL,omitempty"`          // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	PublicID         *string               `json:"publicId,omitempty"`         // `DocumentType`'s publicId.
	SystemID         *string               `json:"systemId,omitempty"`         // `DocumentType`'s systemId.
	InternalSubset   *string               `json:"internalSubset,omitempty"`   // `DocumentType`'s internalSubset.
	XMLVersion       *string               `json:"xmlVersion,omitempty"`       // `Document`'s XML version in case of XML documents.
	Name             *string               `json:"name,omitempty"`             // `Attr`'s name.
	Value            *string               `json:"value,omitempty"`            // `Attr`'s value.
	PseudoType       PseudoType            `json:"pseudoType,omitempty"`       // Pseudo element type for this node.
	ShadowRootType   ShadowRootType        `json:"shadowRootType,omitempty"`   // Shadow root type.
	FrameID          *internal.PageFrameID `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *Node                 `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []Node                `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *Node                 `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []Node                `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *Node                 `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []BackendNode         `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
	IsSVG            *bool                 `json:"isSVG,omitempty"`            // Whether the node is SVG.
}

// RGBA A structure holding an RGBA color.
type RGBA struct {
	R int      `json:"r"`           // The red component, in the [0-255] range.
	G int      `json:"g"`           // The green component, in the [0-255] range.
	B int      `json:"b"`           // The blue component, in the [0-255] range.
	A *float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// Quad An array of quad vertices, x immediately followed by y for each point,
// points clock-wise.
type Quad []float64

// BoxModel Box model.
type BoxModel struct {
	Content      Quad              `json:"content"`                // Content box
	Padding      Quad              `json:"padding"`                // Padding box
	Border       Quad              `json:"border"`                 // Border box
	Margin       Quad              `json:"margin"`                 // Margin box
	Width        int               `json:"width"`                  // Node width
	Height       int               `json:"height"`                 // Node height
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// ShapeOutsideInfo CSS Shape Outside details.
type ShapeOutsideInfo struct {
	Bounds      Quad              `json:"bounds"`      // Shape bounds
	Shape       []json.RawMessage `json:"shape"`       // Shape coordinate details
	MarginShape []json.RawMessage `json:"marginShape"` // Margin shape bounds
}

// Rect Rectangle.
type Rect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// CSSComputedStyleProperty
type CSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}
