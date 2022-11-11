package accessibility

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"fmt"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// NodeID unique accessibility node identifier.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXNodeId
type NodeID string

// String returns the NodeID as string value.
func (t NodeID) String() string {
	return string(t)
}

// ValueType enum of possible property types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXValueType
type ValueType string

// String returns the ValueType as string value.
func (t ValueType) String() string {
	return string(t)
}

// ValueType values.
const (
	ValueTypeBoolean            ValueType = "boolean"
	ValueTypeTristate           ValueType = "tristate"
	ValueTypeBooleanOrUndefined ValueType = "booleanOrUndefined"
	ValueTypeIdref              ValueType = "idref"
	ValueTypeIdrefList          ValueType = "idrefList"
	ValueTypeInteger            ValueType = "integer"
	ValueTypeNode               ValueType = "node"
	ValueTypeNodeList           ValueType = "nodeList"
	ValueTypeNumber             ValueType = "number"
	ValueTypeString             ValueType = "string"
	ValueTypeComputedString     ValueType = "computedString"
	ValueTypeToken              ValueType = "token"
	ValueTypeTokenList          ValueType = "tokenList"
	ValueTypeDomRelation        ValueType = "domRelation"
	ValueTypeRole               ValueType = "role"
	ValueTypeInternalRole       ValueType = "internalRole"
	ValueTypeValueUndefined     ValueType = "valueUndefined"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ValueType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ValueType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ValueType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch ValueType(v) {
	case ValueTypeBoolean:
		*t = ValueTypeBoolean
	case ValueTypeTristate:
		*t = ValueTypeTristate
	case ValueTypeBooleanOrUndefined:
		*t = ValueTypeBooleanOrUndefined
	case ValueTypeIdref:
		*t = ValueTypeIdref
	case ValueTypeIdrefList:
		*t = ValueTypeIdrefList
	case ValueTypeInteger:
		*t = ValueTypeInteger
	case ValueTypeNode:
		*t = ValueTypeNode
	case ValueTypeNodeList:
		*t = ValueTypeNodeList
	case ValueTypeNumber:
		*t = ValueTypeNumber
	case ValueTypeString:
		*t = ValueTypeString
	case ValueTypeComputedString:
		*t = ValueTypeComputedString
	case ValueTypeToken:
		*t = ValueTypeToken
	case ValueTypeTokenList:
		*t = ValueTypeTokenList
	case ValueTypeDomRelation:
		*t = ValueTypeDomRelation
	case ValueTypeRole:
		*t = ValueTypeRole
	case ValueTypeInternalRole:
		*t = ValueTypeInternalRole
	case ValueTypeValueUndefined:
		*t = ValueTypeValueUndefined

	default:
		in.AddError(fmt.Errorf("unknown ValueType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ValueType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ValueSourceType enum of possible property sources.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXValueSourceType
type ValueSourceType string

// String returns the ValueSourceType as string value.
func (t ValueSourceType) String() string {
	return string(t)
}

// ValueSourceType values.
const (
	ValueSourceTypeAttribute      ValueSourceType = "attribute"
	ValueSourceTypeImplicit       ValueSourceType = "implicit"
	ValueSourceTypeStyle          ValueSourceType = "style"
	ValueSourceTypeContents       ValueSourceType = "contents"
	ValueSourceTypePlaceholder    ValueSourceType = "placeholder"
	ValueSourceTypeRelatedElement ValueSourceType = "relatedElement"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ValueSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ValueSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ValueSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch ValueSourceType(v) {
	case ValueSourceTypeAttribute:
		*t = ValueSourceTypeAttribute
	case ValueSourceTypeImplicit:
		*t = ValueSourceTypeImplicit
	case ValueSourceTypeStyle:
		*t = ValueSourceTypeStyle
	case ValueSourceTypeContents:
		*t = ValueSourceTypeContents
	case ValueSourceTypePlaceholder:
		*t = ValueSourceTypePlaceholder
	case ValueSourceTypeRelatedElement:
		*t = ValueSourceTypeRelatedElement

	default:
		in.AddError(fmt.Errorf("unknown ValueSourceType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ValueSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ValueNativeSourceType enum of possible native property sources (as a
// subtype of a particular AXValueSourceType).
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXValueNativeSourceType
type ValueNativeSourceType string

// String returns the ValueNativeSourceType as string value.
func (t ValueNativeSourceType) String() string {
	return string(t)
}

// ValueNativeSourceType values.
const (
	ValueNativeSourceTypeDescription    ValueNativeSourceType = "description"
	ValueNativeSourceTypeFigcaption     ValueNativeSourceType = "figcaption"
	ValueNativeSourceTypeLabel          ValueNativeSourceType = "label"
	ValueNativeSourceTypeLabelfor       ValueNativeSourceType = "labelfor"
	ValueNativeSourceTypeLabelwrapped   ValueNativeSourceType = "labelwrapped"
	ValueNativeSourceTypeLegend         ValueNativeSourceType = "legend"
	ValueNativeSourceTypeRubyannotation ValueNativeSourceType = "rubyannotation"
	ValueNativeSourceTypeTablecaption   ValueNativeSourceType = "tablecaption"
	ValueNativeSourceTypeTitle          ValueNativeSourceType = "title"
	ValueNativeSourceTypeOther          ValueNativeSourceType = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ValueNativeSourceType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ValueNativeSourceType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ValueNativeSourceType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch ValueNativeSourceType(v) {
	case ValueNativeSourceTypeDescription:
		*t = ValueNativeSourceTypeDescription
	case ValueNativeSourceTypeFigcaption:
		*t = ValueNativeSourceTypeFigcaption
	case ValueNativeSourceTypeLabel:
		*t = ValueNativeSourceTypeLabel
	case ValueNativeSourceTypeLabelfor:
		*t = ValueNativeSourceTypeLabelfor
	case ValueNativeSourceTypeLabelwrapped:
		*t = ValueNativeSourceTypeLabelwrapped
	case ValueNativeSourceTypeLegend:
		*t = ValueNativeSourceTypeLegend
	case ValueNativeSourceTypeRubyannotation:
		*t = ValueNativeSourceTypeRubyannotation
	case ValueNativeSourceTypeTablecaption:
		*t = ValueNativeSourceTypeTablecaption
	case ValueNativeSourceTypeTitle:
		*t = ValueNativeSourceTypeTitle
	case ValueNativeSourceTypeOther:
		*t = ValueNativeSourceTypeOther

	default:
		in.AddError(fmt.Errorf("unknown ValueNativeSourceType value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ValueNativeSourceType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ValueSource a single source for a computed AX property.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXValueSource
type ValueSource struct {
	Type              ValueSourceType       `json:"type"`                        // What type of source this is.
	Value             *Value                `json:"value,omitempty"`             // The value of this property source.
	Attribute         string                `json:"attribute,omitempty"`         // The name of the relevant attribute, if any.
	AttributeValue    *Value                `json:"attributeValue,omitempty"`    // The value of the relevant attribute, if any.
	Superseded        bool                  `json:"superseded,omitempty"`        // Whether this source is superseded by a higher priority source.
	NativeSource      ValueNativeSourceType `json:"nativeSource,omitempty"`      // The native markup source for this value, e.g. a <label> element.
	NativeSourceValue *Value                `json:"nativeSourceValue,omitempty"` // The value, such as a node or node list, of the native source.
	Invalid           bool                  `json:"invalid,omitempty"`           // Whether the value for this property is invalid.
	InvalidReason     string                `json:"invalidReason,omitempty"`     // Reason for the value being invalid, if it is.
}

// RelatedNode [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXRelatedNode
type RelatedNode struct {
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId"` // The BackendNodeId of the related DOM node.
	Idref            string            `json:"idref,omitempty"`  // The IDRef value provided, if any.
	Text             string            `json:"text,omitempty"`   // The text alternative of this node in the current context.
}

// Property [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXProperty
type Property struct {
	Name  PropertyName `json:"name"`  // The name of this property.
	Value *Value       `json:"value"` // The value of this property.
}

// Value a single computed AX property.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXValue
type Value struct {
	Type         ValueType           `json:"type"`                   // The type of this value.
	Value        easyjson.RawMessage `json:"value,omitempty"`        // The computed value of this property.
	RelatedNodes []*RelatedNode      `json:"relatedNodes,omitempty"` // One or more related nodes, if applicable.
	Sources      []*ValueSource      `json:"sources,omitempty"`      // The sources which contributed to the computation of this property.
}

// PropertyName values of AXProperty name: - from 'busy' to
// 'roledescription': states which apply to every AX node - from 'live' to
// 'root': attributes which apply to nodes in live regions - from 'autocomplete'
// to 'valuetext': attributes which apply to widgets - from 'checked' to
// 'selected': states which apply to widgets - from 'activedescendant' to 'owns'
// - relationships between elements other than parent/child/sibling.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXPropertyName
type PropertyName string

// String returns the PropertyName as string value.
func (t PropertyName) String() string {
	return string(t)
}

// PropertyName values.
const (
	PropertyNameBusy             PropertyName = "busy"
	PropertyNameDisabled         PropertyName = "disabled"
	PropertyNameEditable         PropertyName = "editable"
	PropertyNameFocusable        PropertyName = "focusable"
	PropertyNameFocused          PropertyName = "focused"
	PropertyNameHidden           PropertyName = "hidden"
	PropertyNameHiddenRoot       PropertyName = "hiddenRoot"
	PropertyNameInvalid          PropertyName = "invalid"
	PropertyNameKeyshortcuts     PropertyName = "keyshortcuts"
	PropertyNameSettable         PropertyName = "settable"
	PropertyNameRoledescription  PropertyName = "roledescription"
	PropertyNameLive             PropertyName = "live"
	PropertyNameAtomic           PropertyName = "atomic"
	PropertyNameRelevant         PropertyName = "relevant"
	PropertyNameRoot             PropertyName = "root"
	PropertyNameAutocomplete     PropertyName = "autocomplete"
	PropertyNameHasPopup         PropertyName = "hasPopup"
	PropertyNameLevel            PropertyName = "level"
	PropertyNameMultiselectable  PropertyName = "multiselectable"
	PropertyNameOrientation      PropertyName = "orientation"
	PropertyNameMultiline        PropertyName = "multiline"
	PropertyNameReadonly         PropertyName = "readonly"
	PropertyNameRequired         PropertyName = "required"
	PropertyNameValuemin         PropertyName = "valuemin"
	PropertyNameValuemax         PropertyName = "valuemax"
	PropertyNameValuetext        PropertyName = "valuetext"
	PropertyNameChecked          PropertyName = "checked"
	PropertyNameExpanded         PropertyName = "expanded"
	PropertyNameModal            PropertyName = "modal"
	PropertyNamePressed          PropertyName = "pressed"
	PropertyNameSelected         PropertyName = "selected"
	PropertyNameActivedescendant PropertyName = "activedescendant"
	PropertyNameControls         PropertyName = "controls"
	PropertyNameDescribedby      PropertyName = "describedby"
	PropertyNameDetails          PropertyName = "details"
	PropertyNameErrormessage     PropertyName = "errormessage"
	PropertyNameFlowto           PropertyName = "flowto"
	PropertyNameLabelledby       PropertyName = "labelledby"
	PropertyNameOwns             PropertyName = "owns"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PropertyName) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PropertyName) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PropertyName) UnmarshalEasyJSON(in *jlexer.Lexer) {
	v := in.String()
	switch PropertyName(v) {
	case PropertyNameBusy:
		*t = PropertyNameBusy
	case PropertyNameDisabled:
		*t = PropertyNameDisabled
	case PropertyNameEditable:
		*t = PropertyNameEditable
	case PropertyNameFocusable:
		*t = PropertyNameFocusable
	case PropertyNameFocused:
		*t = PropertyNameFocused
	case PropertyNameHidden:
		*t = PropertyNameHidden
	case PropertyNameHiddenRoot:
		*t = PropertyNameHiddenRoot
	case PropertyNameInvalid:
		*t = PropertyNameInvalid
	case PropertyNameKeyshortcuts:
		*t = PropertyNameKeyshortcuts
	case PropertyNameSettable:
		*t = PropertyNameSettable
	case PropertyNameRoledescription:
		*t = PropertyNameRoledescription
	case PropertyNameLive:
		*t = PropertyNameLive
	case PropertyNameAtomic:
		*t = PropertyNameAtomic
	case PropertyNameRelevant:
		*t = PropertyNameRelevant
	case PropertyNameRoot:
		*t = PropertyNameRoot
	case PropertyNameAutocomplete:
		*t = PropertyNameAutocomplete
	case PropertyNameHasPopup:
		*t = PropertyNameHasPopup
	case PropertyNameLevel:
		*t = PropertyNameLevel
	case PropertyNameMultiselectable:
		*t = PropertyNameMultiselectable
	case PropertyNameOrientation:
		*t = PropertyNameOrientation
	case PropertyNameMultiline:
		*t = PropertyNameMultiline
	case PropertyNameReadonly:
		*t = PropertyNameReadonly
	case PropertyNameRequired:
		*t = PropertyNameRequired
	case PropertyNameValuemin:
		*t = PropertyNameValuemin
	case PropertyNameValuemax:
		*t = PropertyNameValuemax
	case PropertyNameValuetext:
		*t = PropertyNameValuetext
	case PropertyNameChecked:
		*t = PropertyNameChecked
	case PropertyNameExpanded:
		*t = PropertyNameExpanded
	case PropertyNameModal:
		*t = PropertyNameModal
	case PropertyNamePressed:
		*t = PropertyNamePressed
	case PropertyNameSelected:
		*t = PropertyNameSelected
	case PropertyNameActivedescendant:
		*t = PropertyNameActivedescendant
	case PropertyNameControls:
		*t = PropertyNameControls
	case PropertyNameDescribedby:
		*t = PropertyNameDescribedby
	case PropertyNameDetails:
		*t = PropertyNameDetails
	case PropertyNameErrormessage:
		*t = PropertyNameErrormessage
	case PropertyNameFlowto:
		*t = PropertyNameFlowto
	case PropertyNameLabelledby:
		*t = PropertyNameLabelledby
	case PropertyNameOwns:
		*t = PropertyNameOwns

	default:
		in.AddError(fmt.Errorf("unknown PropertyName value: %v", v))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PropertyName) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Node a node in the accessibility tree.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Accessibility#type-AXNode
type Node struct {
	NodeID           NodeID            `json:"nodeId"`                     // Unique identifier for this node.
	Ignored          bool              `json:"ignored"`                    // Whether this node is ignored for accessibility
	IgnoredReasons   []*Property       `json:"ignoredReasons,omitempty"`   // Collection of reasons why this node is hidden.
	Role             *Value            `json:"role,omitempty"`             // This Node's role, whether explicit or implicit.
	ChromeRole       *Value            `json:"chromeRole,omitempty"`       // This Node's Chrome raw role.
	Name             *Value            `json:"name,omitempty"`             // The accessible name for this Node.
	Description      *Value            `json:"description,omitempty"`      // The accessible description for this Node.
	Value            *Value            `json:"value,omitempty"`            // The value for this Node.
	Properties       []*Property       `json:"properties,omitempty"`       // All other properties
	ParentID         NodeID            `json:"parentId,omitempty"`         // ID for this node's parent.
	ChildIDs         []NodeID          `json:"childIds,omitempty"`         // IDs for each of this node's child nodes.
	BackendDOMNodeID cdp.BackendNodeID `json:"backendDOMNodeId,omitempty"` // The backend ID for the associated DOM node, if any.
	FrameID          cdp.FrameID       `json:"frameId,omitempty"`          // The frame ID for the frame associated with this nodes document.
}
