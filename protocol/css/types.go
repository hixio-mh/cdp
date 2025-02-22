// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
)

// StyleSheetID
type StyleSheetID string

// StyleSheetOrigin Stylesheet type: "injected" for stylesheets injected via
// extension, "user-agent" for user-agent stylesheets, "inspector" for
// stylesheets created by the inspector (i.e. those holding the "via inspector"
// rules), "regular" for regular stylesheets.
type StyleSheetOrigin string

// StyleSheetOrigin as enums.
const (
	StyleSheetOriginNotSet    StyleSheetOrigin = ""
	StyleSheetOriginInjected  StyleSheetOrigin = "injected"
	StyleSheetOriginUserAgent StyleSheetOrigin = "user-agent"
	StyleSheetOriginInspector StyleSheetOrigin = "inspector"
	StyleSheetOriginRegular   StyleSheetOrigin = "regular"
)

func (e StyleSheetOrigin) Valid() bool {
	switch e {
	case "injected", "user-agent", "inspector", "regular":
		return true
	default:
		return false
	}
}

func (e StyleSheetOrigin) String() string {
	return string(e)
}

// PseudoElementMatches CSS rule collection for a single pseudo style.
type PseudoElementMatches struct {
	PseudoType dom.PseudoType `json:"pseudoType"` // Pseudo element type.
	Matches    []RuleMatch    `json:"matches"`    // Matches of CSS rules applicable to the pseudo style.
}

// InheritedStyleEntry Inherited CSS rule collection from ancestor node.
type InheritedStyleEntry struct {
	InlineStyle     *Style      `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []RuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// RuleMatch Match data for a CSS rule.
type RuleMatch struct {
	Rule              Rule  `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Value Data for a simple selector (these are delimited by commas in a
// selector list).
type Value struct {
	Text  string       `json:"text"`            // Value text.
	Range *SourceRange `json:"range,omitempty"` // Value range in the underlying resource (if available).
}

// SelectorList Selector list data.
type SelectorList struct {
	Selectors []Value `json:"selectors"` // Selectors in the list.
	Text      string  `json:"text"`      // Rule selector text.
}

// StyleSheetHeader CSS stylesheet metainformation.
type StyleSheetHeader struct {
	StyleSheetID  StyleSheetID       `json:"styleSheetId"`           // The stylesheet identifier.
	FrameID       page.FrameID       `json:"frameId"`                // Owner frame identifier.
	SourceURL     string             `json:"sourceURL"`              // Stylesheet resource URL. Empty if this is a constructed stylesheet created using new CSSStyleSheet() (but non-empty if this is a constructed sylesheet imported as a CSS module script).
	SourceMapURL  *string            `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin        StyleSheetOrigin   `json:"origin"`                 // Stylesheet origin.
	Title         string             `json:"title"`                  // Stylesheet title.
	OwnerNode     *dom.BackendNodeID `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled      bool               `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL  *bool              `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline      bool               `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	IsMutable     bool               `json:"isMutable"`              // Whether this stylesheet is mutable. Inline stylesheets become mutable after they have been modified via CSSOM API. <link> element's stylesheets become mutable only if DevTools modifies them. Constructed stylesheets (new CSSStyleSheet()) are mutable immediately after creation.
	IsConstructed bool               `json:"isConstructed"`          // True if this stylesheet is created through new CSSStyleSheet() or imported as a CSS module script.
	StartLine     float64            `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn   float64            `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length        float64            `json:"length"`                 // Size of the content (in characters).
	EndLine       float64            `json:"endLine"`                // Line offset of the end of the stylesheet within the resource (zero based).
	EndColumn     float64            `json:"endColumn"`              // Column offset of the end of the stylesheet within the resource (zero based).
}

// Rule CSS rule representation.
type Rule struct {
	StyleSheetID *StyleSheetID    `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList SelectorList     `json:"selectorList"`           // Rule selector data.
	Origin       StyleSheetOrigin `json:"origin"`                 // Parent stylesheet's origin.
	Style        Style            `json:"style"`                  // Associated style declaration.
	Media        []Media          `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
	// ContainerQueries Container query list array (for rules involving
	// container queries). The array enumerates container queries starting
	// with the innermost one, going outwards.
	//
	// Note: This property is experimental.
	ContainerQueries []ContainerQuery `json:"containerQueries,omitempty"`
}

// RuleUsage CSS coverage information.
type RuleUsage struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset  float64      `json:"startOffset"`  // Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset    float64      `json:"endOffset"`    // Offset of the end of the rule body from the beginning of the stylesheet.
	Used         bool         `json:"used"`         // Indicates whether the rule was actually used by some element in the page.
}

// SourceRange Text range within a resource. All numbers are zero-based.
type SourceRange struct {
	StartLine   int `json:"startLine"`   // Start line of range.
	StartColumn int `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int `json:"endLine"`     // End line of range
	EndColumn   int `json:"endColumn"`   // End column of range (exclusive).
}

// ShorthandEntry
type ShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important *bool  `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
}

// ComputedStyleProperty
type ComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// Style CSS style representation.
type Style struct {
	StyleSheetID     *StyleSheetID    `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CSSProperties    []Property       `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []ShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CSSText          *string          `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *SourceRange     `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// Property CSS property declaration data.
type Property struct {
	Name      string       `json:"name"`                // The property name.
	Value     string       `json:"value"`               // The property value.
	Important *bool        `json:"important,omitempty"` // Whether the property has "!important" annotation (implies `false` if absent).
	Implicit  *bool        `json:"implicit,omitempty"`  // Whether the property is implicit (implies `false` if absent).
	Text      *string      `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  *bool        `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies `true` if absent).
	Disabled  *bool        `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *SourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// Media CSS media rule descriptor.
type Media struct {
	Text string `json:"text"` // Media query text.
	// Source Source of the media query: "mediaRule" if specified by a
	// @media rule, "importRule" if specified by an @import rule,
	// "linkedSheet" if specified by a "media" attribute in a linked
	// stylesheet's LINK tag, "inlineSheet" if specified by a "media"
	// attribute in an inline stylesheet's STYLE tag.
	//
	// Values: "mediaRule", "importRule", "linkedSheet", "inlineSheet".
	Source       string        `json:"source"`
	SourceURL    *string       `json:"sourceURL,omitempty"`    // URL of the document containing the media query description.
	Range        *SourceRange  `json:"range,omitempty"`        // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetID *StyleSheetID `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList    []MediaQuery  `json:"mediaList,omitempty"`    // Array of media queries.
}

// MediaQuery Media query descriptor.
type MediaQuery struct {
	Expressions []MediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                   `json:"active"`      // Whether the media query condition is satisfied.
}

// MediaQueryExpression Media query expression descriptor.
type MediaQueryExpression struct {
	Value          float64      `json:"value"`                    // Media query expression value.
	Unit           string       `json:"unit"`                     // Media query expression units.
	Feature        string       `json:"feature"`                  // Media query expression feature.
	ValueRange     *SourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength *float64     `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// ContainerQuery CSS container query rule descriptor.
//
// Note: This type is experimental.
type ContainerQuery struct {
	Text         string        `json:"text"`                   // Container query text.
	Range        *SourceRange  `json:"range,omitempty"`        // The associated rule header range in the enclosing stylesheet (if available).
	StyleSheetID *StyleSheetID `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	Name         *string       `json:"name,omitempty"`         // Optional name for the container.
}

// PlatformFontUsage Information about amount of glyphs that were rendered
// with given font.
type PlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`   // Font's family name reported by platform.
	IsCustomFont bool    `json:"isCustomFont"` // Indicates if the font was downloaded or resolved locally.
	GlyphCount   float64 `json:"glyphCount"`   // Amount of glyphs that were rendered with this font.
}

// FontVariationAxis Information about font variation axes for variable fonts
type FontVariationAxis struct {
	Tag          string  `json:"tag"`          // The font-variation-setting tag (a.k.a. "axis tag").
	Name         string  `json:"name"`         // Human-readable variation name in the default language (normally, "en").
	MinValue     float64 `json:"minValue"`     // The minimum value (inclusive) the font supports for this tag.
	MaxValue     float64 `json:"maxValue"`     // The maximum value (inclusive) the font supports for this tag.
	DefaultValue float64 `json:"defaultValue"` // The default value.
}

// FontFace Properties of a web font:
// https://www.w3.org/TR/2008/REC-CSS2-20080411/fonts.html#font-descriptions
// and additional information such as platformFontFamily and fontVariationAxes.
type FontFace struct {
	FontFamily         string              `json:"fontFamily"`                  // The font-family.
	FontStyle          string              `json:"fontStyle"`                   // The font-style.
	FontVariant        string              `json:"fontVariant"`                 // The font-variant.
	FontWeight         string              `json:"fontWeight"`                  // The font-weight.
	FontStretch        string              `json:"fontStretch"`                 // The font-stretch.
	UnicodeRange       string              `json:"unicodeRange"`                // The unicode-range.
	Src                string              `json:"src"`                         // The src.
	PlatformFontFamily string              `json:"platformFontFamily"`          // The resolved platform font family
	FontVariationAxes  []FontVariationAxis `json:"fontVariationAxes,omitempty"` // Available variation settings (a.k.a. "axes").
}

// KeyframesRule CSS keyframes rule representation.
type KeyframesRule struct {
	AnimationName Value          `json:"animationName"` // Animation name.
	Keyframes     []KeyframeRule `json:"keyframes"`     // List of keyframes.
}

// KeyframeRule CSS keyframe rule representation.
type KeyframeRule struct {
	StyleSheetID *StyleSheetID    `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin       StyleSheetOrigin `json:"origin"`                 // Parent stylesheet's origin.
	KeyText      Value            `json:"keyText"`                // Associated key text.
	Style        Style            `json:"style"`                  // Associated style declaration.
}

// StyleDeclarationEdit A descriptor of operation to mutate style declaration
// text.
type StyleDeclarationEdit struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier.
	Range        SourceRange  `json:"range"`        // The range of the style text in the enclosing stylesheet.
	Text         string       `json:"text"`         // New style text.
}
