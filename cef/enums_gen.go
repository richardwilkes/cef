// Code created from "enum.go.tmpl" - don't edit by hand

package cef

// AlphaType (cef_alpha_type_t from include/internal/cef_types.h)
// Describes how to interpret the alpha component of a pixel.
type AlphaType int

// Possible values for AlphaType
const (
	// Describes how to interpret the alpha component of a pixel.
	AlphaTypeOpaque AlphaType = 0 // CEF_ALPHA_TYPE_OPAQUE
	// Describes how to interpret the alpha component of a pixel.
	AlphaTypePremultiplied AlphaType = 1 // CEF_ALPHA_TYPE_PREMULTIPLIED
	// Describes how to interpret the alpha component of a pixel.
	AlphaTypePostmultiplied AlphaType = 2 // CEF_ALPHA_TYPE_POSTMULTIPLIED
)

// ButtonState (cef_button_state_t from include/internal/cef_types.h)
// Specifies the button display state.
type ButtonState int

// Possible values for ButtonState
const (
	// Specifies the button display state.
	ButtonStateNormal ButtonState = 0 // CEF_BUTTON_STATE_NORMAL
	// Specifies the button display state.
	ButtonStateHovered ButtonState = 1 // CEF_BUTTON_STATE_HOVERED
	// Specifies the button display state.
	ButtonStatePressed ButtonState = 2 // CEF_BUTTON_STATE_PRESSED
	// Specifies the button display state.
	ButtonStateDisabled ButtonState = 3 // CEF_BUTTON_STATE_DISABLED
)

// CdmRegistrationError (cef_cdm_registration_error_t from include/internal/cef_types.h)
// Error codes for CDM registration. See cef_web_plugin.h for details.
type CdmRegistrationError int

// Possible values for CdmRegistrationError
const (
	// Error codes for CDM registration. See cef_web_plugin.h for details.
	CdmRegistrationErrorNone CdmRegistrationError = 0 // CEF_CDM_REGISTRATION_ERROR_NONE
	// Error codes for CDM registration. See cef_web_plugin.h for details.
	CdmRegistrationErrorIncorrectContents CdmRegistrationError = 1 // CEF_CDM_REGISTRATION_ERROR_INCORRECT_CONTENTS
	// Error codes for CDM registration. See cef_web_plugin.h for details.
	CdmRegistrationErrorIncompatible CdmRegistrationError = 2 // CEF_CDM_REGISTRATION_ERROR_INCOMPATIBLE
	// Error codes for CDM registration. See cef_web_plugin.h for details.
	CdmRegistrationErrorNotSupported CdmRegistrationError = 3 // CEF_CDM_REGISTRATION_ERROR_NOT_SUPPORTED
)

// CertStatus (cef_cert_status_t from include/internal/cef_types.h)
// Supported certificate status code values. See net\cert\cert_status_flags.h
// for more information. CERT_STATUS_NONE is new in CEF because we use an
// enum while cert_status_flags.h uses a typedef and static const variables.
type CertStatus int

// Possible values for CertStatus
const (
	CertStatusNone              CertStatus = 0      // CERT_STATUS_NONE
	CertStatusCommonNameInvalid CertStatus = 1 << 0 // CERT_STATUS_COMMON_NAME_INVALID
	CertStatusDateInvalid       CertStatus = 1 << 1 // CERT_STATUS_DATE_INVALID
	CertStatusAuthorityInvalid  CertStatus = 1 << 2 // CERT_STATUS_AUTHORITY_INVALID
	// 1 << 3 is reserved for ERR_CERT_CONTAINS_ERRORS (not useful with WinHTTP).
	CertStatusNoRevocationMechanism   CertStatus = 1 << 4 // CERT_STATUS_NO_REVOCATION_MECHANISM
	CertStatusUnableToCheckRevocation CertStatus = 1 << 5 // CERT_STATUS_UNABLE_TO_CHECK_REVOCATION
	CertStatusRevoked                 CertStatus = 1 << 6 // CERT_STATUS_REVOKED
	CertStatusInvalid                 CertStatus = 1 << 7 // CERT_STATUS_INVALID
	CertStatusWeakSignatureAlgorithm  CertStatus = 1 << 8 // CERT_STATUS_WEAK_SIGNATURE_ALGORITHM
	// 1 << 9 was used for CERT_STATUS_NOT_IN_DNS
	CertStatusNonUniqueName CertStatus = 1 << 10 // CERT_STATUS_NON_UNIQUE_NAME
	CertStatusWeakKey       CertStatus = 1 << 11 // CERT_STATUS_WEAK_KEY
	// 1 << 12 was used for CERT_STATUS_WEAK_DH_KEY
	CertStatusPinnedKeyMissing        CertStatus = 1 << 13 // CERT_STATUS_PINNED_KEY_MISSING
	CertStatusNameConstraintViolation CertStatus = 1 << 14 // CERT_STATUS_NAME_CONSTRAINT_VIOLATION
	CertStatusValidityTooLong         CertStatus = 1 << 15 // CERT_STATUS_VALIDITY_TOO_LONG
	// Bits 16 to 31 are for non-error statuses.
	CertStatusIsEv               CertStatus = 1 << 16 // CERT_STATUS_IS_EV
	CertStatusRevCheckingEnabled CertStatus = 1 << 17 // CERT_STATUS_REV_CHECKING_ENABLED
	// Bit 18 was CERT_STATUS_IS_DNSSEC
	CertStatusSha1SignaturePresent CertStatus = 1 << 19 // CERT_STATUS_SHA1_SIGNATURE_PRESENT
	CertStatusCTComplianceFailed   CertStatus = 1 << 20 // CERT_STATUS_CT_COMPLIANCE_FAILED
)

// ChannelLayout (cef_channel_layout_t from include/internal/cef_types.h)
// Enumerates the various representations of the ordering of audio channels.
// Logged to UMA, so never reuse a value, always add new/greater ones!
// See media\base\channel_layout.h
type ChannelLayout int

// Possible values for ChannelLayout
const (
	ChannelLayoutNone        ChannelLayout = 0 // CEF_CHANNEL_LAYOUT_NONE
	ChannelLayoutUnsupported ChannelLayout = 1 // CEF_CHANNEL_LAYOUT_UNSUPPORTED
	// Front C
	ChannelLayoutMono ChannelLayout = 2 // CEF_CHANNEL_LAYOUT_MONO
	// Front L, Front R
	ChannelLayoutStereo ChannelLayout = 3 // CEF_CHANNEL_LAYOUT_STEREO
	// Front L, Front R, Back C
	ChannelLayout21 ChannelLayout = 4 // CEF_CHANNEL_LAYOUT_2_1
	// Front L, Front R, Front C
	ChannelLayoutSurround ChannelLayout = 5 // CEF_CHANNEL_LAYOUT_SURROUND
	// Front L, Front R, Front C, Back C
	ChannelLayout40 ChannelLayout = 6 // CEF_CHANNEL_LAYOUT_4_0
	// Front L, Front R, Side L, Side R
	ChannelLayout22 ChannelLayout = 7 // CEF_CHANNEL_LAYOUT_2_2
	// Front L, Front R, Back L, Back R
	ChannelLayoutQuad ChannelLayout = 8 // CEF_CHANNEL_LAYOUT_QUAD
	// Front L, Front R, Front C, Side L, Side R
	ChannelLayout50 ChannelLayout = 9 // CEF_CHANNEL_LAYOUT_5_0
	// Front L, Front R, Front C, LFE, Side L, Side R
	ChannelLayout51 ChannelLayout = 10 // CEF_CHANNEL_LAYOUT_5_1
	// Front L, Front R, Front C, Back L, Back R
	ChannelLayout50Back ChannelLayout = 11 // CEF_CHANNEL_LAYOUT_5_0_BACK
	// Front L, Front R, Front C, LFE, Back L, Back R
	ChannelLayout51Back ChannelLayout = 12 // CEF_CHANNEL_LAYOUT_5_1_BACK
	// Front L, Front R, Front C, Side L, Side R, Back L, Back R
	ChannelLayout70 ChannelLayout = 13 // CEF_CHANNEL_LAYOUT_7_0
	// Front L, Front R, Front C, LFE, Side L, Side R, Back L, Back R
	ChannelLayout71 ChannelLayout = 14 // CEF_CHANNEL_LAYOUT_7_1
	// Front L, Front R, Front C, LFE, Side L, Side R, Front LofC, Front RofC
	ChannelLayout71Wide ChannelLayout = 15 // CEF_CHANNEL_LAYOUT_7_1_WIDE
	// Stereo L, Stereo R
	ChannelLayoutStereoDownmix ChannelLayout = 16 // CEF_CHANNEL_LAYOUT_STEREO_DOWNMIX
	// Stereo L, Stereo R, LFE
	ChannelLayout2point1 ChannelLayout = 17 // CEF_CHANNEL_LAYOUT_2POINT1
	// Stereo L, Stereo R, Front C, LFE
	ChannelLayout31 ChannelLayout = 18 // CEF_CHANNEL_LAYOUT_3_1
	// Stereo L, Stereo R, Front C, Rear C, LFE
	ChannelLayout41 ChannelLayout = 19 // CEF_CHANNEL_LAYOUT_4_1
	// Stereo L, Stereo R, Front C, Side L, Side R, Back C
	ChannelLayout60 ChannelLayout = 20 // CEF_CHANNEL_LAYOUT_6_0
	// Stereo L, Stereo R, Side L, Side R, Front LofC, Front RofC
	ChannelLayout60Front ChannelLayout = 21 // CEF_CHANNEL_LAYOUT_6_0_FRONT
	// Stereo L, Stereo R, Front C, Rear L, Rear R, Rear C
	ChannelLayoutHexagonal ChannelLayout = 22 // CEF_CHANNEL_LAYOUT_HEXAGONAL
	// Stereo L, Stereo R, Front C, LFE, Side L, Side R, Rear Center
	ChannelLayout61 ChannelLayout = 23 // CEF_CHANNEL_LAYOUT_6_1
	// Stereo L, Stereo R, Front C, LFE, Back L, Back R, Rear Center
	ChannelLayout61Back ChannelLayout = 24 // CEF_CHANNEL_LAYOUT_6_1_BACK
	// Stereo L, Stereo R, Side L, Side R, Front LofC, Front RofC, LFE
	ChannelLayout61Front ChannelLayout = 25 // CEF_CHANNEL_LAYOUT_6_1_FRONT
	// Front L, Front R, Front C, Side L, Side R, Front LofC, Front RofC
	ChannelLayout70Front ChannelLayout = 26 // CEF_CHANNEL_LAYOUT_7_0_FRONT
	// Front L, Front R, Front C, LFE, Back L, Back R, Front LofC, Front RofC
	ChannelLayout71WideBack ChannelLayout = 27 // CEF_CHANNEL_LAYOUT_7_1_WIDE_BACK
	// Front L, Front R, Front C, Side L, Side R, Rear L, Back R, Back C.
	ChannelLayoutOctagonal ChannelLayout = 28 // CEF_CHANNEL_LAYOUT_OCTAGONAL
	// Channels are not explicitly mapped to speakers.
	ChannelLayoutDiscrete ChannelLayout = 29 // CEF_CHANNEL_LAYOUT_DISCRETE
	// Front L, Front R, Front C. Front C contains the keyboard mic audio. This
	// layout is only intended for input for WebRTC. The Front C channel
	// is stripped away in the WebRTC audio input pipeline and never seen outside
	// of that.
	ChannelLayoutStereoAndKeyboardMic ChannelLayout = 30 // CEF_CHANNEL_LAYOUT_STEREO_AND_KEYBOARD_MIC
	// Front L, Front R, Side L, Side R, LFE
	ChannelLayout41QuadSide ChannelLayout = 31 // CEF_CHANNEL_LAYOUT_4_1_QUAD_SIDE
	// Actual channel layout is specified in the bitstream and the actual channel
	// count is unknown at Chromium media pipeline level (useful for audio
	// pass-through mode).
	ChannelLayoutBitstream ChannelLayout = 32 // CEF_CHANNEL_LAYOUT_BITSTREAM
	// Max value, must always equal the largest entry ever logged.
	ChannelLayoutMax ChannelLayout = ChannelLayoutBitstream // CEF_CHANNEL_LAYOUT_MAX
)

// ColorModel (cef_color_model_t from include/internal/cef_types.h)
// Print job color mode values.
type ColorModel int

// Possible values for ColorModel
const (
	// Print job color mode values.
	ColorModelUnknown ColorModel = 0 // COLOR_MODEL_UNKNOWN
	// Print job color mode values.
	ColorModelGray ColorModel = 1 // COLOR_MODEL_GRAY
	// Print job color mode values.
	ColorModelColor ColorModel = 2 // COLOR_MODEL_COLOR
	// Print job color mode values.
	ColorModelCmyk ColorModel = 3 // COLOR_MODEL_CMYK
	// Print job color mode values.
	ColorModelCmy ColorModel = 4 // COLOR_MODEL_CMY
	// Print job color mode values.
	ColorModelKcmy ColorModel = 5 // COLOR_MODEL_KCMY
	// Print job color mode values.
	ColorModelCmyK ColorModel = 6 // COLOR_MODEL_CMY_K
	// Print job color mode values.
	ColorModelBlack ColorModel = 7 // COLOR_MODEL_BLACK
	// Print job color mode values.
	ColorModelGrayscale ColorModel = 8 // COLOR_MODEL_GRAYSCALE
	// Print job color mode values.
	ColorModelRgb ColorModel = 9 // COLOR_MODEL_RGB
	// Print job color mode values.
	ColorModelRgb16 ColorModel = 10 // COLOR_MODEL_RGB16
	// Print job color mode values.
	ColorModelRgba ColorModel = 11 // COLOR_MODEL_RGBA
	// Print job color mode values.
	ColorModelColormodeColor ColorModel = 12 // COLOR_MODEL_COLORMODE_COLOR
	// Print job color mode values.
	ColorModelColormodeMonochrome ColorModel = 13 // COLOR_MODEL_COLORMODE_MONOCHROME
	// Print job color mode values.
	ColorModelHpColorColor ColorModel = 14 // COLOR_MODEL_HP_COLOR_COLOR
	// Print job color mode values.
	ColorModelHpColorBlack ColorModel = 15 // COLOR_MODEL_HP_COLOR_BLACK
	// Print job color mode values.
	ColorModelPrintoutmodeNormal ColorModel = 16 // COLOR_MODEL_PRINTOUTMODE_NORMAL
	// Print job color mode values.
	ColorModelPrintoutmodeNormalGray ColorModel = 17 // COLOR_MODEL_PRINTOUTMODE_NORMAL_GRAY
	// Print job color mode values.
	ColorModelProcesscolormodelCmyk ColorModel = 18 // COLOR_MODEL_PROCESSCOLORMODEL_CMYK
	// Print job color mode values.
	ColorModelProcesscolormodelGreyscale ColorModel = 19 // COLOR_MODEL_PROCESSCOLORMODEL_GREYSCALE
	// Print job color mode values.
	ColorModelProcesscolormodelRgb ColorModel = 20 // COLOR_MODEL_PROCESSCOLORMODEL_RGB
)

// ColorType (cef_color_type_t from include/internal/cef_types.h)
// Describes how to interpret the components of a pixel.
type ColorType int

// Possible values for ColorType
const (
	// Describes how to interpret the components of a pixel.
	ColorTypeRgba8888 ColorType = 0 // CEF_COLOR_TYPE_RGBA_8888
	// Describes how to interpret the components of a pixel.
	ColorTypeBgra8888 ColorType = 1 // CEF_COLOR_TYPE_BGRA_8888
)

// COMInitMode (cef_com_init_mode_t from include/internal/cef_types.h)
// Windows COM initialization mode. Specifies how COM will be initialized for a
// new thread.
type COMInitMode int

// Possible values for COMInitMode
const (
	// Windows COM initialization mode. Specifies how COM will be initialized for a
	// new thread.
	COMInitModeNone COMInitMode = 0 // COM_INIT_MODE_NONE
	// Windows COM initialization mode. Specifies how COM will be initialized for a
	// new thread.
	COMInitModeSta COMInitMode = 1 // COM_INIT_MODE_STA
	// Windows COM initialization mode. Specifies how COM will be initialized for a
	// new thread.
	COMInitModeMta COMInitMode = 2 // COM_INIT_MODE_MTA
)

// ContextMenuEditStateFlags (cef_context_menu_edit_state_flags_t from include/internal/cef_types.h)
// Supported context menu edit state bit flags.
type ContextMenuEditStateFlags int

// Possible values for ContextMenuEditStateFlags
const (
	CmEditflagNone         ContextMenuEditStateFlags = 0      // CM_EDITFLAG_NONE
	CmEditflagCanUndo      ContextMenuEditStateFlags = 1 << 0 // CM_EDITFLAG_CAN_UNDO
	CmEditflagCanRedo      ContextMenuEditStateFlags = 1 << 1 // CM_EDITFLAG_CAN_REDO
	CmEditflagCanCut       ContextMenuEditStateFlags = 1 << 2 // CM_EDITFLAG_CAN_CUT
	CmEditflagCanCopy      ContextMenuEditStateFlags = 1 << 3 // CM_EDITFLAG_CAN_COPY
	CmEditflagCanPaste     ContextMenuEditStateFlags = 1 << 4 // CM_EDITFLAG_CAN_PASTE
	CmEditflagCanDelete    ContextMenuEditStateFlags = 1 << 5 // CM_EDITFLAG_CAN_DELETE
	CmEditflagCanSelectAll ContextMenuEditStateFlags = 1 << 6 // CM_EDITFLAG_CAN_SELECT_ALL
	CmEditflagCanTranslate ContextMenuEditStateFlags = 1 << 7 // CM_EDITFLAG_CAN_TRANSLATE
)

// ContextMenuMediaStateFlags (cef_context_menu_media_state_flags_t from include/internal/cef_types.h)
// Supported context menu media state bit flags.
type ContextMenuMediaStateFlags int

// Possible values for ContextMenuMediaStateFlags
const (
	CmMediaflagNone               ContextMenuMediaStateFlags = 0      // CM_MEDIAFLAG_NONE
	CmMediaflagError              ContextMenuMediaStateFlags = 1 << 0 // CM_MEDIAFLAG_ERROR
	CmMediaflagPaused             ContextMenuMediaStateFlags = 1 << 1 // CM_MEDIAFLAG_PAUSED
	CmMediaflagMuted              ContextMenuMediaStateFlags = 1 << 2 // CM_MEDIAFLAG_MUTED
	CmMediaflagLoop               ContextMenuMediaStateFlags = 1 << 3 // CM_MEDIAFLAG_LOOP
	CmMediaflagCanSave            ContextMenuMediaStateFlags = 1 << 4 // CM_MEDIAFLAG_CAN_SAVE
	CmMediaflagHasAudio           ContextMenuMediaStateFlags = 1 << 5 // CM_MEDIAFLAG_HAS_AUDIO
	CmMediaflagHasVideo           ContextMenuMediaStateFlags = 1 << 6 // CM_MEDIAFLAG_HAS_VIDEO
	CmMediaflagControlRootElement ContextMenuMediaStateFlags = 1 << 7 // CM_MEDIAFLAG_CONTROL_ROOT_ELEMENT
	CmMediaflagCanPrint           ContextMenuMediaStateFlags = 1 << 8 // CM_MEDIAFLAG_CAN_PRINT
	CmMediaflagCanRotate          ContextMenuMediaStateFlags = 1 << 9 // CM_MEDIAFLAG_CAN_ROTATE
)

// ContextMenuMediaType (cef_context_menu_media_type_t from include/internal/cef_types.h)
// Supported context menu media types.
type ContextMenuMediaType int

// Possible values for ContextMenuMediaType
const (
	// Supported context menu media types.
	CmMediatypeNone ContextMenuMediaType = 0 // CM_MEDIATYPE_NONE
	// Supported context menu media types.
	CmMediatypeImage ContextMenuMediaType = 1 // CM_MEDIATYPE_IMAGE
	// Supported context menu media types.
	CmMediatypeVideo ContextMenuMediaType = 2 // CM_MEDIATYPE_VIDEO
	// Supported context menu media types.
	CmMediatypeAudio ContextMenuMediaType = 3 // CM_MEDIATYPE_AUDIO
	// Supported context menu media types.
	CmMediatypeFile ContextMenuMediaType = 4 // CM_MEDIATYPE_FILE
	// Supported context menu media types.
	CmMediatypePlugin ContextMenuMediaType = 5 // CM_MEDIATYPE_PLUGIN
)

// ContextMenuTypeFlags (cef_context_menu_type_flags_t from include/internal/cef_types.h)
// Supported context menu type flags.
type ContextMenuTypeFlags int

// Possible values for ContextMenuTypeFlags
const (
	// No node is selected.
	CmTypeflagNone ContextMenuTypeFlags = 0 // CM_TYPEFLAG_NONE
	// The top page is selected.
	CmTypeflagPage ContextMenuTypeFlags = 1 << 0 // CM_TYPEFLAG_PAGE
	// A subframe page is selected.
	CmTypeflagFrame ContextMenuTypeFlags = 1 << 1 // CM_TYPEFLAG_FRAME
	// A link is selected.
	CmTypeflagLink ContextMenuTypeFlags = 1 << 2 // CM_TYPEFLAG_LINK
	// A media node is selected.
	CmTypeflagMedia ContextMenuTypeFlags = 1 << 3 // CM_TYPEFLAG_MEDIA
	// There is a textual or mixed selection that is selected.
	CmTypeflagSelection ContextMenuTypeFlags = 1 << 4 // CM_TYPEFLAG_SELECTION
	// An editable element is selected.
	CmTypeflagEditable ContextMenuTypeFlags = 1 << 5 // CM_TYPEFLAG_EDITABLE
)

// CrossAxisAlignment (cef_cross_axis_alignment_t from include/internal/cef_types.h)
// Specifies where along the cross axis the CefBoxLayout child views should be
// laid out.
type CrossAxisAlignment int

// Possible values for CrossAxisAlignment
const (
	// Specifies where along the cross axis the CefBoxLayout child views should be
	// laid out.
	CrossAxisAlignmentStretch CrossAxisAlignment = 0 // CEF_CROSS_AXIS_ALIGNMENT_STRETCH
	// Specifies where along the cross axis the CefBoxLayout child views should be
	// laid out.
	CrossAxisAlignmentStart CrossAxisAlignment = 1 // CEF_CROSS_AXIS_ALIGNMENT_START
	// Specifies where along the cross axis the CefBoxLayout child views should be
	// laid out.
	CrossAxisAlignmentCenter CrossAxisAlignment = 2 // CEF_CROSS_AXIS_ALIGNMENT_CENTER
	// Specifies where along the cross axis the CefBoxLayout child views should be
	// laid out.
	CrossAxisAlignmentEnd CrossAxisAlignment = 3 // CEF_CROSS_AXIS_ALIGNMENT_END
)

// CursorType (cef_cursor_type_t from include/internal/cef_types.h)
// Cursor type values.
type CursorType int

// Possible values for CursorType
const (
	CTPointer                  CursorType = 0  // CT_POINTER
	CTCross                    CursorType = 1  // CT_CROSS
	CTHand                     CursorType = 2  // CT_HAND
	CTIbeam                    CursorType = 3  // CT_IBEAM
	CTWait                     CursorType = 4  // CT_WAIT
	CTHelp                     CursorType = 5  // CT_HELP
	CTEastresize               CursorType = 6  // CT_EASTRESIZE
	CTNorthresize              CursorType = 7  // CT_NORTHRESIZE
	CTNortheastresize          CursorType = 8  // CT_NORTHEASTRESIZE
	CTNorthwestresize          CursorType = 9  // CT_NORTHWESTRESIZE
	CTSouthresize              CursorType = 10 // CT_SOUTHRESIZE
	CTSoutheastresize          CursorType = 11 // CT_SOUTHEASTRESIZE
	CTSouthwestresize          CursorType = 12 // CT_SOUTHWESTRESIZE
	CTWestresize               CursorType = 13 // CT_WESTRESIZE
	CTNorthsouthresize         CursorType = 14 // CT_NORTHSOUTHRESIZE
	CTEastwestresize           CursorType = 15 // CT_EASTWESTRESIZE
	CTNortheastsouthwestresize CursorType = 16 // CT_NORTHEASTSOUTHWESTRESIZE
	CTNorthwestsoutheastresize CursorType = 17 // CT_NORTHWESTSOUTHEASTRESIZE
	CTColumnresize             CursorType = 18 // CT_COLUMNRESIZE
	CTRowresize                CursorType = 19 // CT_ROWRESIZE
	CTMiddlepanning            CursorType = 20 // CT_MIDDLEPANNING
	CTEastpanning              CursorType = 21 // CT_EASTPANNING
	CTNorthpanning             CursorType = 22 // CT_NORTHPANNING
	CTNortheastpanning         CursorType = 23 // CT_NORTHEASTPANNING
	CTNorthwestpanning         CursorType = 24 // CT_NORTHWESTPANNING
	CTSouthpanning             CursorType = 25 // CT_SOUTHPANNING
	CTSoutheastpanning         CursorType = 26 // CT_SOUTHEASTPANNING
	CTSouthwestpanning         CursorType = 27 // CT_SOUTHWESTPANNING
	CTWestpanning              CursorType = 28 // CT_WESTPANNING
	CTMove                     CursorType = 29 // CT_MOVE
	CTVerticaltext             CursorType = 30 // CT_VERTICALTEXT
	CTCell                     CursorType = 31 // CT_CELL
	CTContextmenu              CursorType = 32 // CT_CONTEXTMENU
	CTAlias                    CursorType = 33 // CT_ALIAS
	CTProgress                 CursorType = 34 // CT_PROGRESS
	CTNodrop                   CursorType = 35 // CT_NODROP
	CTCopy                     CursorType = 36 // CT_COPY
	CTNone                     CursorType = 37 // CT_NONE
	CTNotallowed               CursorType = 38 // CT_NOTALLOWED
	CTZoomin                   CursorType = 39 // CT_ZOOMIN
	CTZoomout                  CursorType = 40 // CT_ZOOMOUT
	CTGrab                     CursorType = 41 // CT_GRAB
	CTGrabbing                 CursorType = 42 // CT_GRABBING
	CTCustom                   CursorType = 43 // CT_CUSTOM
)

// DOMDocumentType (cef_dom_document_type_t from include/internal/cef_types.h)
// DOM document types.
type DOMDocumentType int

// Possible values for DOMDocumentType
const (
	DOMDocumentTypeUnknown DOMDocumentType = 0 // DOM_DOCUMENT_TYPE_UNKNOWN
	DOMDocumentTypeHTML    DOMDocumentType = 1 // DOM_DOCUMENT_TYPE_HTML
	DOMDocumentTypeXhtml   DOMDocumentType = 2 // DOM_DOCUMENT_TYPE_XHTML
	DOMDocumentTypePlugin  DOMDocumentType = 3 // DOM_DOCUMENT_TYPE_PLUGIN
)

// DOMEventCategory (cef_dom_event_category_t from include/internal/cef_types.h)
// DOM event category flags.
type DOMEventCategory int

// Possible values for DOMEventCategory
const (
	DOMEventCategoryUnknown                DOMEventCategory = 0x0    // DOM_EVENT_CATEGORY_UNKNOWN
	DOMEventCategoryUI                     DOMEventCategory = 0x1    // DOM_EVENT_CATEGORY_UI
	DOMEventCategoryMouse                  DOMEventCategory = 0x2    // DOM_EVENT_CATEGORY_MOUSE
	DOMEventCategoryMutation               DOMEventCategory = 0x4    // DOM_EVENT_CATEGORY_MUTATION
	DOMEventCategoryKeyboard               DOMEventCategory = 0x8    // DOM_EVENT_CATEGORY_KEYBOARD
	DOMEventCategoryText                   DOMEventCategory = 0x10   // DOM_EVENT_CATEGORY_TEXT
	DOMEventCategoryComposition            DOMEventCategory = 0x20   // DOM_EVENT_CATEGORY_COMPOSITION
	DOMEventCategoryDrag                   DOMEventCategory = 0x40   // DOM_EVENT_CATEGORY_DRAG
	DOMEventCategoryClipboard              DOMEventCategory = 0x80   // DOM_EVENT_CATEGORY_CLIPBOARD
	DOMEventCategoryMessage                DOMEventCategory = 0x100  // DOM_EVENT_CATEGORY_MESSAGE
	DOMEventCategoryWheel                  DOMEventCategory = 0x200  // DOM_EVENT_CATEGORY_WHEEL
	DOMEventCategoryBeforeTextInserted     DOMEventCategory = 0x400  // DOM_EVENT_CATEGORY_BEFORE_TEXT_INSERTED
	DOMEventCategoryOverflow               DOMEventCategory = 0x800  // DOM_EVENT_CATEGORY_OVERFLOW
	DOMEventCategoryPageTransition         DOMEventCategory = 0x1000 // DOM_EVENT_CATEGORY_PAGE_TRANSITION
	DOMEventCategoryPopstate               DOMEventCategory = 0x2000 // DOM_EVENT_CATEGORY_POPSTATE
	DOMEventCategoryProgress               DOMEventCategory = 0x4000 // DOM_EVENT_CATEGORY_PROGRESS
	DOMEventCategoryXmlhttprequestProgress DOMEventCategory = 0x8000 // DOM_EVENT_CATEGORY_XMLHTTPREQUEST_PROGRESS
)

// DOMEventPhase (cef_dom_event_phase_t from include/internal/cef_types.h)
// DOM event processing phases.
type DOMEventPhase int

// Possible values for DOMEventPhase
const (
	DOMEventPhaseUnknown   DOMEventPhase = 0 // DOM_EVENT_PHASE_UNKNOWN
	DOMEventPhaseCapturing DOMEventPhase = 1 // DOM_EVENT_PHASE_CAPTURING
	DOMEventPhaseAtTarget  DOMEventPhase = 2 // DOM_EVENT_PHASE_AT_TARGET
	DOMEventPhaseBubbling  DOMEventPhase = 3 // DOM_EVENT_PHASE_BUBBLING
)

// DOMNodeType (cef_dom_node_type_t from include/internal/cef_types.h)
// DOM node types.
type DOMNodeType int

// Possible values for DOMNodeType
const (
	DOMNodeTypeUnsupported            DOMNodeType = 0 // DOM_NODE_TYPE_UNSUPPORTED
	DOMNodeTypeElement                DOMNodeType = 1 // DOM_NODE_TYPE_ELEMENT
	DOMNodeTypeAttribute              DOMNodeType = 2 // DOM_NODE_TYPE_ATTRIBUTE
	DOMNodeTypeText                   DOMNodeType = 3 // DOM_NODE_TYPE_TEXT
	DOMNodeTypeCdataSection           DOMNodeType = 4 // DOM_NODE_TYPE_CDATA_SECTION
	DOMNodeTypeProcessingInstructions DOMNodeType = 5 // DOM_NODE_TYPE_PROCESSING_INSTRUCTIONS
	DOMNodeTypeComment                DOMNodeType = 6 // DOM_NODE_TYPE_COMMENT
	DOMNodeTypeDocument               DOMNodeType = 7 // DOM_NODE_TYPE_DOCUMENT
	DOMNodeTypeDocumentType           DOMNodeType = 8 // DOM_NODE_TYPE_DOCUMENT_TYPE
	DOMNodeTypeDocumentFragment       DOMNodeType = 9 // DOM_NODE_TYPE_DOCUMENT_FRAGMENT
)

// DragOperationsMask (cef_drag_operations_mask_t from include/internal/cef_types.h)
// "Verb" of a drag-and-drop operation as negotiated between the source and
// destination. These constants match their equivalents in WebCore's
// DragActions.h and should not be renumbered.
type DragOperationsMask uint

// Possible values for DragOperationsMask
const (
	DragOperationNone    DragOperationsMask = 0  // DRAG_OPERATION_NONE
	DragOperationCopy    DragOperationsMask = 1  // DRAG_OPERATION_COPY
	DragOperationLink    DragOperationsMask = 2  // DRAG_OPERATION_LINK
	DragOperationGeneric DragOperationsMask = 4  // DRAG_OPERATION_GENERIC
	DragOperationPrivate DragOperationsMask = 8  // DRAG_OPERATION_PRIVATE
	DragOperationMove    DragOperationsMask = 16 // DRAG_OPERATION_MOVE
	DragOperationDelete  DragOperationsMask = 32 // DRAG_OPERATION_DELETE
	DragOperationEvery   DragOperationsMask = 33 // DRAG_OPERATION_EVERY
)

// DuplexMode (cef_duplex_mode_t from include/internal/cef_types.h)
// Print job duplex mode values.
type DuplexMode int

// Possible values for DuplexMode
const (
	DuplexModeUnknown   DuplexMode = -1 // DUPLEX_MODE_UNKNOWN
	DuplexModeSimplex   DuplexMode = 0  // DUPLEX_MODE_SIMPLEX
	DuplexModeLongEdge  DuplexMode = 1  // DUPLEX_MODE_LONG_EDGE
	DuplexModeShortEdge DuplexMode = 2  // DUPLEX_MODE_SHORT_EDGE
)

// Errorcode (cef_errorcode_t from include/internal/cef_types.h)
// Supported error code values. See net\base\net_error_list.h for complete
// descriptions of the error codes.
type Errorcode int

// Possible values for Errorcode
const (
	ErrNone                        Errorcode = 0                        // ERR_NONE
	ErrFailed                      Errorcode = -2                       // ERR_FAILED
	ErrAborted                     Errorcode = -3                       // ERR_ABORTED
	ErrInvalidArgument             Errorcode = -4                       // ERR_INVALID_ARGUMENT
	ErrInvalidHandle               Errorcode = -5                       // ERR_INVALID_HANDLE
	ErrFileNotFound                Errorcode = -6                       // ERR_FILE_NOT_FOUND
	ErrTimedOut                    Errorcode = -7                       // ERR_TIMED_OUT
	ErrFileTooBig                  Errorcode = -8                       // ERR_FILE_TOO_BIG
	ErrUnexpected                  Errorcode = -9                       // ERR_UNEXPECTED
	ErrAccessDenied                Errorcode = -10                      // ERR_ACCESS_DENIED
	ErrNotImplemented              Errorcode = -11                      // ERR_NOT_IMPLEMENTED
	ErrConnectionClosed            Errorcode = -100                     // ERR_CONNECTION_CLOSED
	ErrConnectionReset             Errorcode = -101                     // ERR_CONNECTION_RESET
	ErrConnectionRefused           Errorcode = -102                     // ERR_CONNECTION_REFUSED
	ErrConnectionAborted           Errorcode = -103                     // ERR_CONNECTION_ABORTED
	ErrConnectionFailed            Errorcode = -104                     // ERR_CONNECTION_FAILED
	ErrNameNotResolved             Errorcode = -105                     // ERR_NAME_NOT_RESOLVED
	ErrInternetDisconnected        Errorcode = -106                     // ERR_INTERNET_DISCONNECTED
	ErrSSLProtocolError            Errorcode = -107                     // ERR_SSL_PROTOCOL_ERROR
	ErrAddressInvalid              Errorcode = -108                     // ERR_ADDRESS_INVALID
	ErrAddressUnreachable          Errorcode = -109                     // ERR_ADDRESS_UNREACHABLE
	ErrSSLClientAuthCertNeeded     Errorcode = -110                     // ERR_SSL_CLIENT_AUTH_CERT_NEEDED
	ErrTunnelConnectionFailed      Errorcode = -111                     // ERR_TUNNEL_CONNECTION_FAILED
	ErrNoSSLVersionsEnabled        Errorcode = -112                     // ERR_NO_SSL_VERSIONS_ENABLED
	ErrSSLVersionOrCipherMismatch  Errorcode = -113                     // ERR_SSL_VERSION_OR_CIPHER_MISMATCH
	ErrSSLRenegotiationRequested   Errorcode = -114                     // ERR_SSL_RENEGOTIATION_REQUESTED
	ErrCertCommonNameInvalid       Errorcode = -200                     // ERR_CERT_COMMON_NAME_INVALID
	ErrCertBegin                   Errorcode = ErrCertCommonNameInvalid // ERR_CERT_BEGIN
	ErrCertDateInvalid             Errorcode = -201                     // ERR_CERT_DATE_INVALID
	ErrCertAuthorityInvalid        Errorcode = -202                     // ERR_CERT_AUTHORITY_INVALID
	ErrCertContainsErrors          Errorcode = -203                     // ERR_CERT_CONTAINS_ERRORS
	ErrCertNoRevocationMechanism   Errorcode = -204                     // ERR_CERT_NO_REVOCATION_MECHANISM
	ErrCertUnableToCheckRevocation Errorcode = -205                     // ERR_CERT_UNABLE_TO_CHECK_REVOCATION
	ErrCertRevoked                 Errorcode = -206                     // ERR_CERT_REVOKED
	ErrCertInvalid                 Errorcode = -207                     // ERR_CERT_INVALID
	ErrCertWeakSignatureAlgorithm  Errorcode = -208                     // ERR_CERT_WEAK_SIGNATURE_ALGORITHM
	// -209 is available: was ERR_CERT_NOT_IN_DNS.
	ErrCertNonUniqueName           Errorcode = -210                   // ERR_CERT_NON_UNIQUE_NAME
	ErrCertWeakKey                 Errorcode = -211                   // ERR_CERT_WEAK_KEY
	ErrCertNameConstraintViolation Errorcode = -212                   // ERR_CERT_NAME_CONSTRAINT_VIOLATION
	ErrCertValidityTooLong         Errorcode = -213                   // ERR_CERT_VALIDITY_TOO_LONG
	ErrCertEnd                     Errorcode = ErrCertValidityTooLong // ERR_CERT_END
	ErrInvalidURL                  Errorcode = -300                   // ERR_INVALID_URL
	ErrDisallowedURLScheme         Errorcode = -301                   // ERR_DISALLOWED_URL_SCHEME
	ErrUnknownURLScheme            Errorcode = -302                   // ERR_UNKNOWN_URL_SCHEME
	ErrTooManyRedirects            Errorcode = -310                   // ERR_TOO_MANY_REDIRECTS
	ErrUnsafeRedirect              Errorcode = -311                   // ERR_UNSAFE_REDIRECT
	ErrUnsafePort                  Errorcode = -312                   // ERR_UNSAFE_PORT
	ErrInvalidResponse             Errorcode = -320                   // ERR_INVALID_RESPONSE
	ErrInvalidChunkedEncoding      Errorcode = -321                   // ERR_INVALID_CHUNKED_ENCODING
	ErrMethodNotSupported          Errorcode = -322                   // ERR_METHOD_NOT_SUPPORTED
	ErrUnexpectedProxyAuth         Errorcode = -323                   // ERR_UNEXPECTED_PROXY_AUTH
	ErrEmptyResponse               Errorcode = -324                   // ERR_EMPTY_RESPONSE
	ErrResponseHeadersTooBig       Errorcode = -325                   // ERR_RESPONSE_HEADERS_TOO_BIG
	ErrCacheMiss                   Errorcode = -400                   // ERR_CACHE_MISS
	ErrInsecureResponse            Errorcode = -501                   // ERR_INSECURE_RESPONSE
)

// EventFlags (cef_event_flags_t from include/internal/cef_types.h)
// Supported event bit flags.
type EventFlags int

// Possible values for EventFlags
const (
	EventflagNone              EventFlags = 0      // EVENTFLAG_NONE
	EventflagCapsLockOn        EventFlags = 1 << 0 // EVENTFLAG_CAPS_LOCK_ON
	EventflagShiftDown         EventFlags = 1 << 1 // EVENTFLAG_SHIFT_DOWN
	EventflagControlDown       EventFlags = 1 << 2 // EVENTFLAG_CONTROL_DOWN
	EventflagAltDown           EventFlags = 1 << 3 // EVENTFLAG_ALT_DOWN
	EventflagLeftMouseButton   EventFlags = 1 << 4 // EVENTFLAG_LEFT_MOUSE_BUTTON
	EventflagMiddleMouseButton EventFlags = 1 << 5 // EVENTFLAG_MIDDLE_MOUSE_BUTTON
	EventflagRightMouseButton  EventFlags = 1 << 6 // EVENTFLAG_RIGHT_MOUSE_BUTTON
	// Mac OS-X command key.
	EventflagCommandDown EventFlags = 1 << 7  // EVENTFLAG_COMMAND_DOWN
	EventflagNumLockOn   EventFlags = 1 << 8  // EVENTFLAG_NUM_LOCK_ON
	EventflagIsKeyPad    EventFlags = 1 << 9  // EVENTFLAG_IS_KEY_PAD
	EventflagIsLeft      EventFlags = 1 << 10 // EVENTFLAG_IS_LEFT
	EventflagIsRight     EventFlags = 1 << 11 // EVENTFLAG_IS_RIGHT
)

// FileDialogMode (cef_file_dialog_mode_t from include/internal/cef_types.h)
// Supported file dialog modes.
type FileDialogMode int

// Possible values for FileDialogMode
const (
	// Requires that the file exists before allowing the user to pick it.
	FileDialogOpen FileDialogMode = 0 // FILE_DIALOG_OPEN
	// Requires that the file exists before allowing the user to pick it.
	FileDialogOpenMultiple FileDialogMode = 1 // FILE_DIALOG_OPEN_MULTIPLE
	// Requires that the file exists before allowing the user to pick it.
	FileDialogOpenFolder FileDialogMode = 2 // FILE_DIALOG_OPEN_FOLDER
	// Requires that the file exists before allowing the user to pick it.
	FileDialogSave FileDialogMode = 3 // FILE_DIALOG_SAVE
	// General mask defining the bits used for the type values.
	FileDialogTypeMask FileDialogMode = 0xFF // FILE_DIALOG_TYPE_MASK
	// Prompt to overwrite if the user selects an existing file with the Save
	// dialog.
	FileDialogOverwritepromptFlag FileDialogMode = 0x01000000 // FILE_DIALOG_OVERWRITEPROMPT_FLAG
	// Do not display read-only files.
	FileDialogHidereadonlyFlag FileDialogMode = 0x02000000 // FILE_DIALOG_HIDEREADONLY_FLAG
)

// FocusSource (cef_focus_source_t from include/internal/cef_types.h)
// Focus sources.
type FocusSource int

// Possible values for FocusSource
const (
	// The source is explicit navigation via the API (LoadURL(), etc).
	FocusSourceNavigation FocusSource = 0 // FOCUS_SOURCE_NAVIGATION
	// The source is explicit navigation via the API (LoadURL(), etc).
	FocusSourceSystem FocusSource = 1 // FOCUS_SOURCE_SYSTEM
)

// HorizontalAlignment (cef_horizontal_alignment_t from include/internal/cef_types.h)
// Specifies the horizontal text alignment mode.
type HorizontalAlignment int

// Possible values for HorizontalAlignment
const (
	// Specifies the horizontal text alignment mode.
	HorizontalAlignmentLeft HorizontalAlignment = 0 // CEF_HORIZONTAL_ALIGNMENT_LEFT
	// Specifies the horizontal text alignment mode.
	HorizontalAlignmentCenter HorizontalAlignment = 1 // CEF_HORIZONTAL_ALIGNMENT_CENTER
	// Specifies the horizontal text alignment mode.
	HorizontalAlignmentRight HorizontalAlignment = 2 // CEF_HORIZONTAL_ALIGNMENT_RIGHT
)

// JsdialogType (cef_jsdialog_type_t from include/internal/cef_types.h)
// Supported JavaScript dialog types.
type JsdialogType int

// Possible values for JsdialogType
const (
	JsdialogtypeAlert   JsdialogType = 0 // JSDIALOGTYPE_ALERT
	JsdialogtypeConfirm JsdialogType = 1 // JSDIALOGTYPE_CONFIRM
	JsdialogtypePrompt  JsdialogType = 2 // JSDIALOGTYPE_PROMPT
)

// JSONParserError (cef_json_parser_error_t from include/internal/cef_types.h)
// Error codes that can be returned from CefParseJSONAndReturnError.
type JSONParserError int

// Possible values for JSONParserError
const (
	JSONNoError                 JSONParserError = 0 // JSON_NO_ERROR
	JSONInvalidEscape           JSONParserError = 1 // JSON_INVALID_ESCAPE
	JSONSyntaxError             JSONParserError = 2 // JSON_SYNTAX_ERROR
	JSONUnexpectedToken         JSONParserError = 3 // JSON_UNEXPECTED_TOKEN
	JSONTrailingComma           JSONParserError = 4 // JSON_TRAILING_COMMA
	JSONTooMuchNesting          JSONParserError = 5 // JSON_TOO_MUCH_NESTING
	JSONUnexpectedDataAfterRoot JSONParserError = 6 // JSON_UNEXPECTED_DATA_AFTER_ROOT
	JSONUnsupportedEncoding     JSONParserError = 7 // JSON_UNSUPPORTED_ENCODING
	JSONUnquotedDictionaryKey   JSONParserError = 8 // JSON_UNQUOTED_DICTIONARY_KEY
	JSONParseErrorCount         JSONParserError = 9 // JSON_PARSE_ERROR_COUNT
)

// JSONParserOptions (cef_json_parser_options_t from include/internal/cef_types.h)
// Options that can be passed to CefParseJSON.
type JSONParserOptions int

// Possible values for JSONParserOptions
const (
	// Parses the input strictly according to RFC 4627. See comments in Chromium's
	// base/json/json_reader.h file for known limitations/deviations from the RFC.
	JSONParserRfc JSONParserOptions = 0 // JSON_PARSER_RFC
	// Allows commas to exist after the last element in structures.
	JSONParserAllowTrailingCommas JSONParserOptions = 1 << 0 // JSON_PARSER_ALLOW_TRAILING_COMMAS
)

// JSONWriterOptions (cef_json_writer_options_t from include/internal/cef_types.h)
// Options that can be passed to CefWriteJSON.
type JSONWriterOptions int

// Possible values for JSONWriterOptions
const (
	// Default behavior.
	JSONWriterDefault JSONWriterOptions = 0 // JSON_WRITER_DEFAULT
	// This option instructs the writer that if a Binary value is encountered,
	// the value (and key if within a dictionary) will be omitted from the
	// output, and success will be returned. Otherwise, if a binary value is
	// encountered, failure will be returned.
	JSONWriterOmitBinaryValues JSONWriterOptions = 1 << 0 // JSON_WRITER_OMIT_BINARY_VALUES
	// This option instructs the writer to write doubles that have no fractional
	// part as a normal integer (i.e., without using exponential notation
	// or appending a '.0') as long as the value is within the range of a
	// 64-bit int.
	JSONWriterOmitDoubleTypePreservation JSONWriterOptions = 1 << 1 // JSON_WRITER_OMIT_DOUBLE_TYPE_PRESERVATION
	// Return a slightly nicer formatted json string (pads with whitespace to
	// help with readability).
	JSONWriterPrettyPrint JSONWriterOptions = 1 << 2 // JSON_WRITER_PRETTY_PRINT
)

// KeyEventType (cef_key_event_type_t from include/internal/cef_types.h)
// Key event types.
type KeyEventType int

// Possible values for KeyEventType
const (
	// Notification that a key transitioned from "up" to "down".
	KeyeventRawkeydown KeyEventType = 0 // KEYEVENT_RAWKEYDOWN
	// Notification that a key transitioned from "up" to "down".
	KeyeventKeydown KeyEventType = 1 // KEYEVENT_KEYDOWN
	// Notification that a key transitioned from "up" to "down".
	KeyeventKeyup KeyEventType = 2 // KEYEVENT_KEYUP
	// Notification that a key transitioned from "up" to "down".
	KeyeventChar KeyEventType = 3 // KEYEVENT_CHAR
)

// LogSeverity (cef_log_severity_t from include/internal/cef_types.h)
// Log severity levels.
type LogSeverity int

// Possible values for LogSeverity
const (
	// Log severity levels.
	LogseverityDefault LogSeverity = 0 // LOGSEVERITY_DEFAULT
	// Log severity levels.
	LogseverityVerbose LogSeverity = 1 // LOGSEVERITY_VERBOSE
	// DEBUG logging.
	LogseverityDebug LogSeverity = LogseverityVerbose // LOGSEVERITY_DEBUG
	// DEBUG logging.
	LogseverityInfo LogSeverity = (LogseverityVerbose) + 1 // LOGSEVERITY_INFO
	// DEBUG logging.
	LogseverityWarning LogSeverity = ((LogseverityVerbose) + 1) + 1 // LOGSEVERITY_WARNING
	// DEBUG logging.
	LogseverityError LogSeverity = (((LogseverityVerbose) + 1) + 1) + 1 // LOGSEVERITY_ERROR
	// DEBUG logging.
	LogseverityFatal LogSeverity = ((((LogseverityVerbose) + 1) + 1) + 1) + 1 // LOGSEVERITY_FATAL
	// Disable logging to file for all messages, and to stderr for messages with
	// severity less than FATAL.
	LogseverityDisable LogSeverity = 99 // LOGSEVERITY_DISABLE
)

// MainAxisAlignment (cef_main_axis_alignment_t from include/internal/cef_types.h)
// Specifies where along the main axis the CefBoxLayout child views should be
// laid out.
type MainAxisAlignment int

// Possible values for MainAxisAlignment
const (
	// Specifies where along the main axis the CefBoxLayout child views should be
	// laid out.
	MainAxisAlignmentStart MainAxisAlignment = 0 // CEF_MAIN_AXIS_ALIGNMENT_START
	// Specifies where along the main axis the CefBoxLayout child views should be
	// laid out.
	MainAxisAlignmentCenter MainAxisAlignment = 1 // CEF_MAIN_AXIS_ALIGNMENT_CENTER
	// Specifies where along the main axis the CefBoxLayout child views should be
	// laid out.
	MainAxisAlignmentEnd MainAxisAlignment = 2 // CEF_MAIN_AXIS_ALIGNMENT_END
)

// MenuAnchorPosition (cef_menu_anchor_position_t from include/internal/cef_types.h)
// Specifies how a menu will be anchored for non-RTL languages. The opposite
// position will be used for RTL languages.
type MenuAnchorPosition int

// Possible values for MenuAnchorPosition
const (
	// Specifies how a menu will be anchored for non-RTL languages. The opposite
	// position will be used for RTL languages.
	MenuAnchorTopleft MenuAnchorPosition = 0 // CEF_MENU_ANCHOR_TOPLEFT
	// Specifies how a menu will be anchored for non-RTL languages. The opposite
	// position will be used for RTL languages.
	MenuAnchorTopright MenuAnchorPosition = 1 // CEF_MENU_ANCHOR_TOPRIGHT
	// Specifies how a menu will be anchored for non-RTL languages. The opposite
	// position will be used for RTL languages.
	MenuAnchorBottomcenter MenuAnchorPosition = 2 // CEF_MENU_ANCHOR_BOTTOMCENTER
)

// MenuColorType (cef_menu_color_type_t from include/internal/cef_types.h)
// Supported color types for menu items.
type MenuColorType int

// Possible values for MenuColorType
const (
	// Supported color types for menu items.
	MenuColorText MenuColorType = 0 // CEF_MENU_COLOR_TEXT
	// Supported color types for menu items.
	MenuColorTextHovered MenuColorType = 1 // CEF_MENU_COLOR_TEXT_HOVERED
	// Supported color types for menu items.
	MenuColorTextAccelerator MenuColorType = 2 // CEF_MENU_COLOR_TEXT_ACCELERATOR
	// Supported color types for menu items.
	MenuColorTextAcceleratorHovered MenuColorType = 3 // CEF_MENU_COLOR_TEXT_ACCELERATOR_HOVERED
	// Supported color types for menu items.
	MenuColorBackground MenuColorType = 4 // CEF_MENU_COLOR_BACKGROUND
	// Supported color types for menu items.
	MenuColorBackgroundHovered MenuColorType = 5 // CEF_MENU_COLOR_BACKGROUND_HOVERED
	// Supported color types for menu items.
	MenuColorCount MenuColorType = 6 // CEF_MENU_COLOR_COUNT
)

// MenuID (cef_menu_id_t from include/internal/cef_types.h)
// Supported menu IDs. Non-English translations can be provided for the
// IDS_MENU_* strings in CefResourceBundleHandler::GetLocalizedString().
type MenuID int

// Possible values for MenuID
const (
	// Navigation.
	MenuIDBack          MenuID = 100 // MENU_ID_BACK
	MenuIDForward       MenuID = 101 // MENU_ID_FORWARD
	MenuIDReload        MenuID = 102 // MENU_ID_RELOAD
	MenuIDReloadNocache MenuID = 103 // MENU_ID_RELOAD_NOCACHE
	MenuIDStopload      MenuID = 104 // MENU_ID_STOPLOAD
	// Editing.
	MenuIDUndo      MenuID = 110 // MENU_ID_UNDO
	MenuIDRedo      MenuID = 111 // MENU_ID_REDO
	MenuIDCut       MenuID = 112 // MENU_ID_CUT
	MenuIDCopy      MenuID = 113 // MENU_ID_COPY
	MenuIDPaste     MenuID = 114 // MENU_ID_PASTE
	MenuIDDelete    MenuID = 115 // MENU_ID_DELETE
	MenuIDSelectAll MenuID = 116 // MENU_ID_SELECT_ALL
	// Miscellaneous.
	MenuIDFind       MenuID = 130 // MENU_ID_FIND
	MenuIDPrint      MenuID = 131 // MENU_ID_PRINT
	MenuIDViewSource MenuID = 132 // MENU_ID_VIEW_SOURCE
	// Spell checking word correction suggestions.
	MenuIDSpellcheckSuggestion0    MenuID = 200 // MENU_ID_SPELLCHECK_SUGGESTION_0
	MenuIDSpellcheckSuggestion1    MenuID = 201 // MENU_ID_SPELLCHECK_SUGGESTION_1
	MenuIDSpellcheckSuggestion2    MenuID = 202 // MENU_ID_SPELLCHECK_SUGGESTION_2
	MenuIDSpellcheckSuggestion3    MenuID = 203 // MENU_ID_SPELLCHECK_SUGGESTION_3
	MenuIDSpellcheckSuggestion4    MenuID = 204 // MENU_ID_SPELLCHECK_SUGGESTION_4
	MenuIDSpellcheckSuggestionLast MenuID = 204 // MENU_ID_SPELLCHECK_SUGGESTION_LAST
	MenuIDNoSpellingSuggestions    MenuID = 205 // MENU_ID_NO_SPELLING_SUGGESTIONS
	MenuIDAddToDictionary          MenuID = 206 // MENU_ID_ADD_TO_DICTIONARY
	// Custom menu items originating from the renderer process. For example,
	// plugin placeholder menu items or Flash menu items.
	MenuIDCustomFirst MenuID = 220 // MENU_ID_CUSTOM_FIRST
	MenuIDCustomLast  MenuID = 250 // MENU_ID_CUSTOM_LAST
	// All user-defined menu IDs should come between MENU_ID_USER_FIRST and
	// MENU_ID_USER_LAST to avoid overlapping the Chromium and CEF ID ranges
	// defined in the tools/gritsettings/resource_ids file.
	MenuIDUserFirst MenuID = 26500 // MENU_ID_USER_FIRST
	MenuIDUserLast  MenuID = 28500 // MENU_ID_USER_LAST
)

// MenuItemType (cef_menu_item_type_t from include/internal/cef_types.h)
// Supported menu item types.
type MenuItemType int

// Possible values for MenuItemType
const (
	// Supported menu item types.
	MenuitemtypeNone MenuItemType = 0 // MENUITEMTYPE_NONE
	// Supported menu item types.
	MenuitemtypeCommand MenuItemType = 1 // MENUITEMTYPE_COMMAND
	// Supported menu item types.
	MenuitemtypeCheck MenuItemType = 2 // MENUITEMTYPE_CHECK
	// Supported menu item types.
	MenuitemtypeRadio MenuItemType = 3 // MENUITEMTYPE_RADIO
	// Supported menu item types.
	MenuitemtypeSeparator MenuItemType = 4 // MENUITEMTYPE_SEPARATOR
	// Supported menu item types.
	MenuitemtypeSubmenu MenuItemType = 5 // MENUITEMTYPE_SUBMENU
)

// MessageLoopType (cef_message_loop_type_t from include/internal/cef_types.h)
// Message loop types. Indicates the set of asynchronous events that a message
// loop can process.
type MessageLoopType int

// Possible values for MessageLoopType
const (
	// Message loop types. Indicates the set of asynchronous events that a message
	// loop can process.
	MlTypeDefault MessageLoopType = 0 // ML_TYPE_DEFAULT
	// Message loop types. Indicates the set of asynchronous events that a message
	// loop can process.
	MlTypeUI MessageLoopType = 1 // ML_TYPE_UI
	// Message loop types. Indicates the set of asynchronous events that a message
	// loop can process.
	MlTypeIO MessageLoopType = 2 // ML_TYPE_IO
)

// MouseButtonType (cef_mouse_button_type_t from include/internal/cef_types.h)
// Mouse button types.
type MouseButtonType int

// Possible values for MouseButtonType
const (
	MbtLeft   MouseButtonType = 0 // MBT_LEFT
	MbtMiddle MouseButtonType = 1 // MBT_MIDDLE
	MbtRight  MouseButtonType = 2 // MBT_RIGHT
)

// NavigationType (cef_navigation_type_t from include/internal/cef_types.h)
// Navigation types.
type NavigationType int

// Possible values for NavigationType
const (
	NavigationLinkClicked     NavigationType = 0 // NAVIGATION_LINK_CLICKED
	NavigationFormSubmitted   NavigationType = 1 // NAVIGATION_FORM_SUBMITTED
	NavigationBackForward     NavigationType = 2 // NAVIGATION_BACK_FORWARD
	NavigationReload          NavigationType = 3 // NAVIGATION_RELOAD
	NavigationFormResubmitted NavigationType = 4 // NAVIGATION_FORM_RESUBMITTED
	NavigationOther           NavigationType = 5 // NAVIGATION_OTHER
)

// PaintElementType (cef_paint_element_type_t from include/internal/cef_types.h)
// Paint element types.
type PaintElementType int

// Possible values for PaintElementType
const (
	PetView  PaintElementType = 0 // PET_VIEW
	PetPopup PaintElementType = 1 // PET_POPUP
)

// PathKey (cef_path_key_t from include/internal/cef_types.h)
// Path key values.
type PathKey int

// Possible values for PathKey
const (
	// Path key values.
	PkDirCurrent PathKey = 0 // PK_DIR_CURRENT
	// Path key values.
	PkDirExe PathKey = 1 // PK_DIR_EXE
	// Path key values.
	PkDirModule PathKey = 2 // PK_DIR_MODULE
	// Path key values.
	PkDirTemp PathKey = 3 // PK_DIR_TEMP
	// Path key values.
	PkFileExe PathKey = 4 // PK_FILE_EXE
	// Path key values.
	PkFileModule PathKey = 5 // PK_FILE_MODULE
	// Path key values.
	PkLocalAppData PathKey = 6 // PK_LOCAL_APP_DATA
	// Path key values.
	PkUserData PathKey = 7 // PK_USER_DATA
	// Path key values.
	PkDirResources PathKey = 8 // PK_DIR_RESOURCES
)

// PDFPrintMarginType (cef_pdf_print_margin_type_t from include/internal/cef_types.h)
// Margin type for PDF printing.
type PDFPrintMarginType int

// Possible values for PDFPrintMarginType
const (
	// Margin type for PDF printing.
	PDFPrintMarginDefault PDFPrintMarginType = 0 // PDF_PRINT_MARGIN_DEFAULT
	// Margin type for PDF printing.
	PDFPrintMarginNone PDFPrintMarginType = 1 // PDF_PRINT_MARGIN_NONE
	// Margin type for PDF printing.
	PDFPrintMarginMinimum PDFPrintMarginType = 2 // PDF_PRINT_MARGIN_MINIMUM
	// Margin type for PDF printing.
	PDFPrintMarginCustom PDFPrintMarginType = 3 // PDF_PRINT_MARGIN_CUSTOM
)

// PluginPolicy (cef_plugin_policy_t from include/internal/cef_types.h)
// Plugin policies supported by CefRequestContextHandler::OnBeforePluginLoad.
type PluginPolicy int

// Possible values for PluginPolicy
const (
	// Plugin policies supported by CefRequestContextHandler::OnBeforePluginLoad.
	PluginPolicyAllow PluginPolicy = 0 // PLUGIN_POLICY_ALLOW
	// Plugin policies supported by CefRequestContextHandler::OnBeforePluginLoad.
	PluginPolicyDetectImportant PluginPolicy = 1 // PLUGIN_POLICY_DETECT_IMPORTANT
	// Plugin policies supported by CefRequestContextHandler::OnBeforePluginLoad.
	PluginPolicyBlock PluginPolicy = 2 // PLUGIN_POLICY_BLOCK
	// Plugin policies supported by CefRequestContextHandler::OnBeforePluginLoad.
	PluginPolicyDisable PluginPolicy = 3 // PLUGIN_POLICY_DISABLE
)

// PointerType (cef_pointer_type_t from include/internal/cef_types.h)
// The device type that caused the event.
type PointerType int

// Possible values for PointerType
const (
	PointerTypeTouch   PointerType = 0 // CEF_POINTER_TYPE_TOUCH
	PointerTypeMouse   PointerType = 1 // CEF_POINTER_TYPE_MOUSE
	PointerTypePen     PointerType = 2 // CEF_POINTER_TYPE_PEN
	PointerTypeEraser  PointerType = 3 // CEF_POINTER_TYPE_ERASER
	PointerTypeUnknown PointerType = 4 // CEF_POINTER_TYPE_UNKNOWN
)

// PostdataelementType (cef_postdataelement_type_t from include/internal/cef_types.h)
// Post data elements may represent either bytes or files.
type PostdataelementType int

// Possible values for PostdataelementType
const (
	PdeTypeEmpty PostdataelementType = 0 // PDE_TYPE_EMPTY
	PdeTypeBytes PostdataelementType = 1 // PDE_TYPE_BYTES
	PdeTypeFile  PostdataelementType = 2 // PDE_TYPE_FILE
)

// ProcessID (cef_process_id_t from include/internal/cef_types.h)
// Existing process IDs.
type ProcessID int

// Possible values for ProcessID
const (
	// Existing process IDs.
	PidBrowser ProcessID = 0 // PID_BROWSER
	// Existing process IDs.
	PidRenderer ProcessID = 1 // PID_RENDERER
)

// ReferrerPolicy (cef_referrer_policy_t from include/internal/cef_types.h)
// Policy for how the Referrer HTTP header value will be sent during navigation.
// If the `--no-referrers` command-line flag is specified then the policy value
// will be ignored and the Referrer value will never be sent.
// Must be kept synchronized with net::URLRequest::ReferrerPolicy from Chromium.
type ReferrerPolicy int

// Possible values for ReferrerPolicy
const (
	// Policy for how the Referrer HTTP header value will be sent during navigation.
	// If the `--no-referrers` command-line flag is specified then the policy value
	// will be ignored and the Referrer value will never be sent.
	// Must be kept synchronized with net::URLRequest::ReferrerPolicy from Chromium.
	ReferrerPolicyClearReferrerOnTransitionFromSecureToInsecure    ReferrerPolicy = 0 // REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	ReferrerPolicyDefault                                          ReferrerPolicy = 1 // REFERRER_POLICY_DEFAULT
	ReferrerPolicyReduceReferrerGranularityOnTransitionCrossOrigin ReferrerPolicy = 2 // REFERRER_POLICY_REDUCE_REFERRER_GRANULARITY_ON_TRANSITION_CROSS_ORIGIN
	ReferrerPolicyOriginOnlyOnTransitionCrossOrigin                ReferrerPolicy = 3 // REFERRER_POLICY_ORIGIN_ONLY_ON_TRANSITION_CROSS_ORIGIN
	ReferrerPolicyNeverClearReferrer                               ReferrerPolicy = 4 // REFERRER_POLICY_NEVER_CLEAR_REFERRER
	ReferrerPolicyOrigin                                           ReferrerPolicy = 5 // REFERRER_POLICY_ORIGIN
	ReferrerPolicyClearReferrerOnTransitionCrossOrigin             ReferrerPolicy = 6 // REFERRER_POLICY_CLEAR_REFERRER_ON_TRANSITION_CROSS_ORIGIN
	ReferrerPolicyOriginClearOnTransitionFromSecureToInsecure      ReferrerPolicy = 7 // REFERRER_POLICY_ORIGIN_CLEAR_ON_TRANSITION_FROM_SECURE_TO_INSECURE
	ReferrerPolicyNoReferrer                                       ReferrerPolicy = 8 // REFERRER_POLICY_NO_REFERRER
	// Always the last value in this enumeration.
	ReferrerPolicyLastValue ReferrerPolicy = ReferrerPolicyNoReferrer // REFERRER_POLICY_LAST_VALUE
)

// ResourceType (cef_resource_type_t from include/internal/cef_types.h)
// Resource type for a request.
type ResourceType int

// Possible values for ResourceType
const (
	// Top level page.
	RtMainFrame ResourceType = 0 // RT_MAIN_FRAME
	// Top level page.
	RtSubFrame ResourceType = 1 // RT_SUB_FRAME
	// Top level page.
	RtStylesheet ResourceType = 2 // RT_STYLESHEET
	// Top level page.
	RtScript ResourceType = 3 // RT_SCRIPT
	// Top level page.
	RtImage ResourceType = 4 // RT_IMAGE
	// Top level page.
	RtFontResource ResourceType = 5 // RT_FONT_RESOURCE
	// Top level page.
	RtSubResource ResourceType = 6 // RT_SUB_RESOURCE
	// Top level page.
	RtObject ResourceType = 7 // RT_OBJECT
	// Top level page.
	RtMedia ResourceType = 8 // RT_MEDIA
	// Top level page.
	RtWorker ResourceType = 9 // RT_WORKER
	// Top level page.
	RtSharedWorker ResourceType = 10 // RT_SHARED_WORKER
	// Top level page.
	RtPrefetch ResourceType = 11 // RT_PREFETCH
	// Top level page.
	RtFavicon ResourceType = 12 // RT_FAVICON
	// Top level page.
	RtXhr ResourceType = 13 // RT_XHR
	// Top level page.
	RtPing ResourceType = 14 // RT_PING
	// Top level page.
	RtServiceWorker ResourceType = 15 // RT_SERVICE_WORKER
	// Top level page.
	RtCspReport ResourceType = 16 // RT_CSP_REPORT
	// Top level page.
	RtPluginResource ResourceType = 17 // RT_PLUGIN_RESOURCE
)

// ResponseFilterStatus (cef_response_filter_status_t from include/internal/cef_types.h)
// Return values for CefResponseFilter::Filter().
type ResponseFilterStatus int

// Possible values for ResponseFilterStatus
const (
	// Return values for CefResponseFilter::Filter().
	ResponseFilterNeedMoreData ResponseFilterStatus = 0 // RESPONSE_FILTER_NEED_MORE_DATA
	// Return values for CefResponseFilter::Filter().
	ResponseFilterDone ResponseFilterStatus = 1 // RESPONSE_FILTER_DONE
	// Return values for CefResponseFilter::Filter().
	ResponseFilterError ResponseFilterStatus = 2 // RESPONSE_FILTER_ERROR
)

// ReturnValue (cef_return_value_t from include/internal/cef_types.h)
// Return value types.
type ReturnValue int

// Possible values for ReturnValue
const (
	// Cancel immediately.
	RvCancel ReturnValue = 0 // RV_CANCEL
	// Cancel immediately.
	RvContinue ReturnValue = 1 // RV_CONTINUE
	// Cancel immediately.
	RvContinueAsync ReturnValue = 2 // RV_CONTINUE_ASYNC
)

// ScaleFactor (cef_scale_factor_t from include/internal/cef_types.h)
// Supported UI scale factors for the platform. SCALE_FACTOR_NONE is used for
// density independent resources such as string, html/js files or an image that
// can be used for any scale factors (such as wallpapers).
type ScaleFactor int

// Possible values for ScaleFactor
const (
	ScaleFactorNone ScaleFactor = 0 // SCALE_FACTOR_NONE
	ScaleFactor100p ScaleFactor = 1 // SCALE_FACTOR_100P
	ScaleFactor125p ScaleFactor = 2 // SCALE_FACTOR_125P
	ScaleFactor133p ScaleFactor = 3 // SCALE_FACTOR_133P
	ScaleFactor140p ScaleFactor = 4 // SCALE_FACTOR_140P
	ScaleFactor150p ScaleFactor = 5 // SCALE_FACTOR_150P
	ScaleFactor180p ScaleFactor = 6 // SCALE_FACTOR_180P
	ScaleFactor200p ScaleFactor = 7 // SCALE_FACTOR_200P
	ScaleFactor250p ScaleFactor = 8 // SCALE_FACTOR_250P
	ScaleFactor300p ScaleFactor = 9 // SCALE_FACTOR_300P
)

// SchemeOptions (cef_scheme_options_t from include/internal/cef_types.h)
//
// Configuration options for registering a custom scheme.
// These values are used when calling AddCustomScheme.
//
type SchemeOptions int

// Possible values for SchemeOptions
const (
	SchemeOptionNone SchemeOptions = 0 // CEF_SCHEME_OPTION_NONE
	// If CEF_SCHEME_OPTION_STANDARD is set the scheme will be treated as a
	// standard scheme. Standard schemes are subject to URL canonicalization and
	// parsing rules as defined in the Common Internet Scheme Syntax RFC 1738
	// Section 3.1 available at http://www.ietf.org/rfc/rfc1738.txt
	//
	// In particular, the syntax for standard scheme URLs must be of the form:
	// <pre>
	//  [scheme]://[username]:[password]@[host]:[port]/[url-path]
	// </pre> Standard scheme URLs must have a host component that is a fully
	// qualified domain name as defined in Section 3.5 of RFC 1034 [13] and
	// Section 2.1 of RFC 1123. These URLs will be canonicalized to
	// "scheme://host/path" in the simplest case and
	// "scheme://username:password@host:port/path" in the most explicit case. For
	// example, "scheme:host/path" and "scheme:///host/path" will both be
	// canonicalized to "scheme://host/path". The origin of a standard scheme URL
	// is the combination of scheme, host and port (i.e., "scheme://host:port" in
	// the most explicit case).
	//
	// For non-standard scheme URLs only the "scheme:" component is parsed and
	// canonicalized. The remainder of the URL will be passed to the handler as-
	// is. For example, "scheme:///some%20text" will remain the same. Non-standard
	// scheme URLs cannot be used as a target for form submission.
	SchemeOptionStandard SchemeOptions = 1 << 0 // CEF_SCHEME_OPTION_STANDARD
	// If CEF_SCHEME_OPTION_LOCAL is set the scheme will be treated with the same
	// security rules as those applied to "file" URLs. Normal pages cannot link to
	// or access local URLs. Also, by default, local URLs can only perform
	// XMLHttpRequest calls to the same URL (origin + path) that originated the
	// request. To allow XMLHttpRequest calls from a local URL to other URLs with
	// the same origin set the CefSettings.file_access_from_file_urls_allowed
	// value to true (1). To allow XMLHttpRequest calls from a local URL to all
	// origins set the CefSettings.universal_access_from_file_urls_allowed value
	// to true (1).
	SchemeOptionLocal SchemeOptions = 1 << 1 // CEF_SCHEME_OPTION_LOCAL
	// If CEF_SCHEME_OPTION_DISPLAY_ISOLATED is set the scheme can only be
	// displayed from other content hosted with the same scheme. For example,
	// pages in other origins cannot create iframes or hyperlinks to URLs with the
	// scheme. For schemes that must be accessible from other schemes don't set
	// this, set CEF_SCHEME_OPTION_CORS_ENABLED, and use CORS
	// "Access-Control-Allow-Origin" headers to further restrict access.
	SchemeOptionDisplayIsolated SchemeOptions = 1 << 2 // CEF_SCHEME_OPTION_DISPLAY_ISOLATED
	// If CEF_SCHEME_OPTION_SECURE is set the scheme will be treated with the same
	// security rules as those applied to "https" URLs. For example, loading this
	// scheme from other secure schemes will not trigger mixed content warnings.
	SchemeOptionSecure SchemeOptions = 1 << 3 // CEF_SCHEME_OPTION_SECURE
	// If CEF_SCHEME_OPTION_CORS_ENABLED is set the scheme can be sent CORS
	// requests. This value should be set in most cases where
	// CEF_SCHEME_OPTION_STANDARD is set.
	SchemeOptionCorsEnabled SchemeOptions = 1 << 4 // CEF_SCHEME_OPTION_CORS_ENABLED
	// If CEF_SCHEME_OPTION_CSP_BYPASSING is set the scheme can bypass Content-
	// Security-Policy (CSP) checks. This value should not be set in most cases
	// where CEF_SCHEME_OPTION_STANDARD is set.
	SchemeOptionCspBypassing SchemeOptions = 1 << 5 // CEF_SCHEME_OPTION_CSP_BYPASSING
	// If CEF_SCHEME_OPTION_FETCH_ENABLED is set the scheme can perform Fetch API
	// requests.
	SchemeOptionFetchEnabled SchemeOptions = 1 << 6 // CEF_SCHEME_OPTION_FETCH_ENABLED
)

// SSLContentStatus (cef_ssl_content_status_t from include/internal/cef_types.h)
// Supported SSL content status flags. See content/public/common/ssl_status.h
// for more information.
type SSLContentStatus int

// Possible values for SSLContentStatus
const (
	SSLContentNormalContent            SSLContentStatus = 0      // SSL_CONTENT_NORMAL_CONTENT
	SSLContentDisplayedInsecureContent SSLContentStatus = 1 << 0 // SSL_CONTENT_DISPLAYED_INSECURE_CONTENT
	SSLContentRanInsecureContent       SSLContentStatus = 1 << 1 // SSL_CONTENT_RAN_INSECURE_CONTENT
)

// SSLVersion (cef_ssl_version_t from include/internal/cef_types.h)
// Supported SSL version values. See net/ssl/ssl_connection_status_flags.h
// for more information.
type SSLVersion int

// Possible values for SSLVersion
const (
	SSLConnectionVersionUnknown SSLVersion = 0 // SSL_CONNECTION_VERSION_UNKNOWN
	SSLConnectionVersionSSL2    SSLVersion = 1 // SSL_CONNECTION_VERSION_SSL2
	SSLConnectionVersionSSL3    SSLVersion = 2 // SSL_CONNECTION_VERSION_SSL3
	SSLConnectionVersionTLS1    SSLVersion = 3 // SSL_CONNECTION_VERSION_TLS1
	SSLConnectionVersionTLS11   SSLVersion = 4 // SSL_CONNECTION_VERSION_TLS1_1
	SSLConnectionVersionTLS12   SSLVersion = 5 // SSL_CONNECTION_VERSION_TLS1_2
	// Reserve 6 for TLS 1.3.
	SSLConnectionVersionQuic SSLVersion = 7 // SSL_CONNECTION_VERSION_QUIC
)

// State (cef_state_t from include/internal/cef_types.h)
// Represents the state of a setting.
type State int

// Possible values for State
const (
	// Use the default state for the setting.
	StateDefault State = 0 // STATE_DEFAULT
	// Use the default state for the setting.
	StateEnabled State = 1 // STATE_ENABLED
	// Use the default state for the setting.
	StateDisabled State = 2 // STATE_DISABLED
)

// StorageType (cef_storage_type_t from include/internal/cef_types.h)
// Storage types.
type StorageType int

// Possible values for StorageType
const (
	STLocalstorage   StorageType = 0 // ST_LOCALSTORAGE
	STSessionstorage StorageType = 1 // ST_SESSIONSTORAGE
)

// TerminationStatus (cef_termination_status_t from include/internal/cef_types.h)
// Process termination status values.
type TerminationStatus int

// Possible values for TerminationStatus
const (
	// Process termination status values.
	TSAbnormalTermination TerminationStatus = 0 // TS_ABNORMAL_TERMINATION
	// Process termination status values.
	TSProcessWasKilled TerminationStatus = 1 // TS_PROCESS_WAS_KILLED
	// Process termination status values.
	TSProcessCrashed TerminationStatus = 2 // TS_PROCESS_CRASHED
	// Process termination status values.
	TSProcessOom TerminationStatus = 3 // TS_PROCESS_OOM
)

// TextInputMode (cef_text_input_mode_t from include/internal/cef_types.h)
// Input mode of a virtual keyboard. These constants match their equivalents
// in Chromium's text_input_mode.h and should not be renumbered.
// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
type TextInputMode int

// Possible values for TextInputMode
const (
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeDefault TextInputMode = 0 // CEF_TEXT_INPUT_MODE_DEFAULT
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeNone TextInputMode = 1 // CEF_TEXT_INPUT_MODE_NONE
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeText TextInputMode = 2 // CEF_TEXT_INPUT_MODE_TEXT
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeTel TextInputMode = 3 // CEF_TEXT_INPUT_MODE_TEL
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeURL TextInputMode = 4 // CEF_TEXT_INPUT_MODE_URL
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeEmail TextInputMode = 5 // CEF_TEXT_INPUT_MODE_EMAIL
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeNumeric TextInputMode = 6 // CEF_TEXT_INPUT_MODE_NUMERIC
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeDecimal TextInputMode = 7 // CEF_TEXT_INPUT_MODE_DECIMAL
	// Input mode of a virtual keyboard. These constants match their equivalents
	// in Chromium's text_input_mode.h and should not be renumbered.
	// See https://html.spec.whatwg.org/#input-modalities:-the-inputmode-attribute
	TextInputModeSearch TextInputMode = 8                   // CEF_TEXT_INPUT_MODE_SEARCH
	TextInputModeMax    TextInputMode = TextInputModeSearch // CEF_TEXT_INPUT_MODE_MAX
)

// TextStyle (cef_text_style_t from include/internal/cef_types.h)
// Text style types. Should be kepy in sync with gfx::TextStyle.
type TextStyle int

// Possible values for TextStyle
const (
	// Text style types. Should be kepy in sync with gfx::TextStyle.
	TextStyleBold TextStyle = 0 // CEF_TEXT_STYLE_BOLD
	// Text style types. Should be kepy in sync with gfx::TextStyle.
	TextStyleItalic TextStyle = 1 // CEF_TEXT_STYLE_ITALIC
	// Text style types. Should be kepy in sync with gfx::TextStyle.
	TextStyleStrike TextStyle = 2 // CEF_TEXT_STYLE_STRIKE
	// Text style types. Should be kepy in sync with gfx::TextStyle.
	TextStyleDiagonalStrike TextStyle = 3 // CEF_TEXT_STYLE_DIAGONAL_STRIKE
	// Text style types. Should be kepy in sync with gfx::TextStyle.
	TextStyleUnderline TextStyle = 4 // CEF_TEXT_STYLE_UNDERLINE
)

// ThreadID (cef_thread_id_t from include/internal/cef_types.h)
// Existing thread IDs.
type ThreadID int

// Possible values for ThreadID
const (
	// Existing thread IDs.
	TIDUI ThreadID = 0 // TID_UI
	// Existing thread IDs.
	TIDFileBackground   ThreadID = 1                                               // TID_FILE_BACKGROUND
	TIDFile             ThreadID = TIDFileBackground                               // TID_FILE
	TIDFileUserVisible  ThreadID = (TIDFileBackground) + 1                         // TID_FILE_USER_VISIBLE
	TIDFileUserBlocking ThreadID = ((TIDFileBackground) + 1) + 1                   // TID_FILE_USER_BLOCKING
	TIDProcessLauncher  ThreadID = (((TIDFileBackground) + 1) + 1) + 1             // TID_PROCESS_LAUNCHER
	TIDIO               ThreadID = ((((TIDFileBackground) + 1) + 1) + 1) + 1       // TID_IO
	TIDRenderer         ThreadID = (((((TIDFileBackground) + 1) + 1) + 1) + 1) + 1 // TID_RENDERER
)

// ThreadPriority (cef_thread_priority_t from include/internal/cef_types.h)
// Thread priority values listed in increasing order of importance.
type ThreadPriority int

// Possible values for ThreadPriority
const (
	// Thread priority values listed in increasing order of importance.
	TPBackground ThreadPriority = 0 // TP_BACKGROUND
	// Thread priority values listed in increasing order of importance.
	TPNormal ThreadPriority = 1 // TP_NORMAL
	// Thread priority values listed in increasing order of importance.
	TPDisplay ThreadPriority = 2 // TP_DISPLAY
	// Thread priority values listed in increasing order of importance.
	TPRealtimeAudio ThreadPriority = 3 // TP_REALTIME_AUDIO
)

// TouchEventType (cef_touch_event_type_t from include/internal/cef_types.h)
// Touch points states types.
type TouchEventType int

// Possible values for TouchEventType
const (
	TetReleased  TouchEventType = 0 // CEF_TET_RELEASED
	TetPressed   TouchEventType = 1 // CEF_TET_PRESSED
	TetMoved     TouchEventType = 2 // CEF_TET_MOVED
	TetCancelled TouchEventType = 3 // CEF_TET_CANCELLED
)

// TransitionType (cef_transition_type_t from include/internal/cef_types.h)
// Transition type for a request. Made up of one source value and 0 or more
// qualifiers.
type TransitionType uint

// Possible values for TransitionType
const (
	// Source is a link click or the JavaScript window.open function. This is
	// also the default value for requests like sub-resource loads that are not
	// navigations.
	TTLink TransitionType = 0 // TT_LINK
	// Source is some other "explicit" navigation action such as creating a new
	// browser or using the LoadURL function. This is also the default value
	// for navigations where the actual type is unknown.
	TTExplicit TransitionType = 1 // TT_EXPLICIT
	// Source is a subframe navigation. This is any content that is automatically
	// loaded in a non-toplevel frame. For example, if a page consists of several
	// frames containing ads, those ad URLs will have this transition type.
	// The user may not even realize the content in these pages is a separate
	// frame, so may not care about the URL.
	TTAutoSubframe TransitionType = 3 // TT_AUTO_SUBFRAME
	// Source is a subframe navigation explicitly requested by the user that will
	// generate new navigation entries in the back/forward list. These are
	// probably more important than frames that were automatically loaded in
	// the background because the user probably cares about the fact that this
	// link was loaded.
	TTManualSubframe TransitionType = 4 // TT_MANUAL_SUBFRAME
	// Source is a form submission by the user. NOTE: In some situations
	// submitting a form does not result in this transition type. This can happen
	// if the form uses a script to submit the contents.
	TTFormSubmit TransitionType = 7 // TT_FORM_SUBMIT
	// Source is a "reload" of the page via the Reload function or by re-visiting
	// the same URL. NOTE: This is distinct from the concept of whether a
	// particular load uses "reload semantics" (i.e. bypasses cached data).
	TTReload TransitionType = 8 // TT_RELOAD
	// General mask defining the bits used for the source values.
	TTSourceMask TransitionType = 0xFF // TT_SOURCE_MASK
	// Attempted to visit a URL but was blocked.
	TTBlockedFlag TransitionType = 0x00800000 // TT_BLOCKED_FLAG
	// Used the Forward or Back function to navigate among browsing history.
	TTForwardBackFlag TransitionType = 0x01000000 // TT_FORWARD_BACK_FLAG
	// The beginning of a navigation chain.
	TTChainStartFlag TransitionType = 0x10000000 // TT_CHAIN_START_FLAG
	// The last transition in a redirect chain.
	TTChainEndFlag TransitionType = 0x20000000 // TT_CHAIN_END_FLAG
	// Redirects caused by JavaScript or a meta refresh tag on the page.
	TTClientRedirectFlag TransitionType = 0x40000000 // TT_CLIENT_REDIRECT_FLAG
	// Redirects sent from the server by HTTP headers.
	TTServerRedirectFlag TransitionType = 0x80000000 // TT_SERVER_REDIRECT_FLAG
	// Used to test whether a transition involves a redirect.
	TTIsRedirectMask TransitionType = 0xC0000000 // TT_IS_REDIRECT_MASK
	// General mask defining the bits used for the qualifiers.
	TTQualifierMask TransitionType = 0xFFFFFF00 // TT_QUALIFIER_MASK
)

// URIUnescapeRule (cef_uri_unescape_rule_t from include/internal/cef_types.h)
// URI unescape rules passed to CefURIDecode().
type URIUnescapeRule int

// Possible values for URIUnescapeRule
const (
	// Don't unescape anything at all.
	UuNone URIUnescapeRule = 0 // UU_NONE
	// Don't unescape anything special, but all normal unescaping will happen.
	// This is a placeholder and can't be combined with other flags (since it's
	// just the absence of them). All other unescape rules imply "normal" in
	// addition to their special meaning. Things like escaped letters, digits,
	// and most symbols will get unescaped with this mode.
	UuNormal URIUnescapeRule = 1 << 0 // UU_NORMAL
	// Convert %20 to spaces. In some places where we're showing URLs, we may
	// want this. In places where the URL may be copied and pasted out, then
	// you wouldn't want this since it might not be interpreted in one piece
	// by other applications.
	UuSpaces URIUnescapeRule = 1 << 1 // UU_SPACES
	// Unescapes '/' and '\\'. If these characters were unescaped, the resulting
	// URL won't be the same as the source one. Moreover, they are dangerous to
	// unescape in strings that will be used as file paths or names. This value
	// should only be used when slashes don't have special meaning, like data
	// URLs.
	UuPathSeparators URIUnescapeRule = 1 << 2 // UU_PATH_SEPARATORS
	// Unescapes various characters that will change the meaning of URLs,
	// including '%', '+', '&', '#'. Does not unescape path separators.
	// If these characters were unescaped, the resulting URL won't be the same
	// as the source one. This flag is used when generating final output like
	// filenames for URLs where we won't be interpreting as a URL and want to do
	// as much unescaping as possible.
	UuURLSpecialCharsExceptPathSeparators URIUnescapeRule = 1 << 3 // UU_URL_SPECIAL_CHARS_EXCEPT_PATH_SEPARATORS
	// Unescapes characters that can be used in spoofing attempts (such as LOCK)
	// and control characters (such as BiDi control characters and %01).  This
	// INCLUDES NULLs.  This is used for rare cases such as data: URL decoding
	// where the result is binary data.
	//
	// DO NOT use UU_SPOOFING_AND_CONTROL_CHARS if the URL is going to be
	// displayed in the UI for security reasons.
	UuSpoofingAndControlChars URIUnescapeRule = 1 << 4 // UU_SPOOFING_AND_CONTROL_CHARS
	// URL queries use "+" for space. This flag controls that replacement.
	UuReplacePlusWithSpace URIUnescapeRule = 1 << 5 // UU_REPLACE_PLUS_WITH_SPACE
)

// UrlrequestFlags (cef_urlrequest_flags_t from include/internal/cef_types.h)
// Flags used to customize the behavior of CefURLRequest.
type UrlrequestFlags int

// Possible values for UrlrequestFlags
const (
	// Default behavior.
	UrFlagNone UrlrequestFlags = 0 // UR_FLAG_NONE
	// If set the cache will be skipped when handling the request. Setting this
	// value is equivalent to specifying the "Cache-Control: no-cache" request
	// header. Setting this value in combination with UR_FLAG_ONLY_FROM_CACHE will
	// cause the request to fail.
	UrFlagSkipCache UrlrequestFlags = 1 << 0 // UR_FLAG_SKIP_CACHE
	// If set the request will fail if it cannot be served from the cache (or some
	// equivalent local store). Setting this value is equivalent to specifying the
	// "Cache-Control: only-if-cached" request header. Setting this value in
	// combination with UR_FLAG_SKIP_CACHE or UR_FLAG_DISABLE_CACHE will cause the
	// request to fail.
	UrFlagOnlyFromCache UrlrequestFlags = 1 << 1 // UR_FLAG_ONLY_FROM_CACHE
	// If set the cache will not be used at all. Setting this value is equivalent
	// to specifying the "Cache-Control: no-store" request header. Setting this
	// value in combination with UR_FLAG_ONLY_FROM_CACHE will cause the request to
	// fail.
	UrFlagDisableCache UrlrequestFlags = 1 << 2 // UR_FLAG_DISABLE_CACHE
	// If set user name, password, and cookies may be sent with the request, and
	// cookies may be saved from the response.
	UrFlagAllowStoredCredentials UrlrequestFlags = 1 << 3 // UR_FLAG_ALLOW_STORED_CREDENTIALS
	// If set upload progress events will be generated when a request has a body.
	UrFlagReportUploadProgress UrlrequestFlags = 1 << 4 // UR_FLAG_REPORT_UPLOAD_PROGRESS
	// If set the CefURLRequestClient::OnDownloadData method will not be called.
	UrFlagNoDownloadData UrlrequestFlags = 1 << 5 // UR_FLAG_NO_DOWNLOAD_DATA
	// If set 5XX redirect errors will be propagated to the observer instead of
	// automatically re-tried. This currently only applies for requests
	// originated in the browser process.
	UrFlagNoRetryOn5xx UrlrequestFlags = 1 << 6 // UR_FLAG_NO_RETRY_ON_5XX
	// If set 3XX responses will cause the fetch to halt immediately rather than
	// continue through the redirect.
	UrFlagStopOnRedirect UrlrequestFlags = 1 << 7 // UR_FLAG_STOP_ON_REDIRECT
)

// UrlrequestStatus (cef_urlrequest_status_t from include/internal/cef_types.h)
// Flags that represent CefURLRequest status.
type UrlrequestStatus int

// Possible values for UrlrequestStatus
const (
	// Unknown status.
	UrUnknown UrlrequestStatus = 0 // UR_UNKNOWN
	// Unknown status.
	UrSuccess UrlrequestStatus = 1 // UR_SUCCESS
	// Unknown status.
	UrIOPending UrlrequestStatus = 2 // UR_IO_PENDING
	// Unknown status.
	UrCanceled UrlrequestStatus = 3 // UR_CANCELED
	// Unknown status.
	UrFailed UrlrequestStatus = 4 // UR_FAILED
)

// V8Accesscontrol (cef_v8_accesscontrol_t from include/internal/cef_types.h)
// V8 access control values.
type V8Accesscontrol int

// Possible values for V8Accesscontrol
const (
	V8AccessControlDefault              V8Accesscontrol = 0      // V8_ACCESS_CONTROL_DEFAULT
	V8AccessControlAllCanRead           V8Accesscontrol = 1      // V8_ACCESS_CONTROL_ALL_CAN_READ
	V8AccessControlAllCanWrite          V8Accesscontrol = 1 << 1 // V8_ACCESS_CONTROL_ALL_CAN_WRITE
	V8AccessControlProhibitsOverwriting V8Accesscontrol = 1 << 2 // V8_ACCESS_CONTROL_PROHIBITS_OVERWRITING
)

// V8Propertyattribute (cef_v8_propertyattribute_t from include/internal/cef_types.h)
// V8 property attribute values.
type V8Propertyattribute int

// Possible values for V8Propertyattribute
const (
	V8PropertyAttributeNone V8Propertyattribute = 0 // V8_PROPERTY_ATTRIBUTE_NONE
	//   Configurable
	V8PropertyAttributeReadonly   V8Propertyattribute = 1 << 0 // V8_PROPERTY_ATTRIBUTE_READONLY
	V8PropertyAttributeDontenum   V8Propertyattribute = 1 << 1 // V8_PROPERTY_ATTRIBUTE_DONTENUM
	V8PropertyAttributeDontdelete V8Propertyattribute = 1 << 2 // V8_PROPERTY_ATTRIBUTE_DONTDELETE
)

// ValueType (cef_value_type_t from include/internal/cef_types.h)
// Supported value types.
type ValueType int

// Possible values for ValueType
const (
	VtypeInvalid    ValueType = 0 // VTYPE_INVALID
	VtypeNull       ValueType = 1 // VTYPE_NULL
	VtypeBool       ValueType = 2 // VTYPE_BOOL
	VtypeInt        ValueType = 3 // VTYPE_INT
	VtypeDouble     ValueType = 4 // VTYPE_DOUBLE
	VtypeString     ValueType = 5 // VTYPE_STRING
	VtypeBinary     ValueType = 6 // VTYPE_BINARY
	VtypeDictionary ValueType = 7 // VTYPE_DICTIONARY
	VtypeList       ValueType = 8 // VTYPE_LIST
)

// WindowOpenDisposition (cef_window_open_disposition_t from include/internal/cef_types.h)
// The manner in which a link click should be opened. These constants match
// their equivalents in Chromium's window_open_disposition.h and should not be
// renumbered.
type WindowOpenDisposition int

// Possible values for WindowOpenDisposition
const (
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodUnknown WindowOpenDisposition = 0 // WOD_UNKNOWN
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodCurrentTab WindowOpenDisposition = 1 // WOD_CURRENT_TAB
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodSingletonTab WindowOpenDisposition = 2 // WOD_SINGLETON_TAB
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodNewForegroundTab WindowOpenDisposition = 3 // WOD_NEW_FOREGROUND_TAB
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodNewBackgroundTab WindowOpenDisposition = 4 // WOD_NEW_BACKGROUND_TAB
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodNewPopup WindowOpenDisposition = 5 // WOD_NEW_POPUP
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodNewWindow WindowOpenDisposition = 6 // WOD_NEW_WINDOW
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodSaveToDisk WindowOpenDisposition = 7 // WOD_SAVE_TO_DISK
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodOffTheRecord WindowOpenDisposition = 8 // WOD_OFF_THE_RECORD
	// The manner in which a link click should be opened. These constants match
	// their equivalents in Chromium's window_open_disposition.h and should not be
	// renumbered.
	WodIgnoreAction WindowOpenDisposition = 9 // WOD_IGNORE_ACTION
)

// XMLEncodingType (cef_xml_encoding_type_t from include/internal/cef_types.h)
// Supported XML encoding types. The parser supports ASCII, ISO-8859-1, and
// UTF16 (LE and BE) by default. All other types must be translated to UTF8
// before being passed to the parser. If a BOM is detected and the correct
// decoder is available then that decoder will be used automatically.
type XMLEncodingType int

// Possible values for XMLEncodingType
const (
	XMLEncodingNone    XMLEncodingType = 0 // XML_ENCODING_NONE
	XMLEncodingUtf8    XMLEncodingType = 1 // XML_ENCODING_UTF8
	XMLEncodingUtf16le XMLEncodingType = 2 // XML_ENCODING_UTF16LE
	XMLEncodingUtf16be XMLEncodingType = 3 // XML_ENCODING_UTF16BE
	XMLEncodingASCII   XMLEncodingType = 4 // XML_ENCODING_ASCII
)

// XMLNodeType (cef_xml_node_type_t from include/internal/cef_types.h)
// XML node types.
type XMLNodeType int

// Possible values for XMLNodeType
const (
	XMLNodeUnsupported           XMLNodeType = 0  // XML_NODE_UNSUPPORTED
	XMLNodeProcessingInstruction XMLNodeType = 1  // XML_NODE_PROCESSING_INSTRUCTION
	XMLNodeDocumentType          XMLNodeType = 2  // XML_NODE_DOCUMENT_TYPE
	XMLNodeElementStart          XMLNodeType = 3  // XML_NODE_ELEMENT_START
	XMLNodeElementEnd            XMLNodeType = 4  // XML_NODE_ELEMENT_END
	XMLNodeAttribute             XMLNodeType = 5  // XML_NODE_ATTRIBUTE
	XMLNodeText                  XMLNodeType = 6  // XML_NODE_TEXT
	XMLNodeCdata                 XMLNodeType = 7  // XML_NODE_CDATA
	XMLNodeEntityReference       XMLNodeType = 8  // XML_NODE_ENTITY_REFERENCE
	XMLNodeWhitespace            XMLNodeType = 9  // XML_NODE_WHITESPACE
	XMLNodeComment               XMLNodeType = 10 // XML_NODE_COMMENT
)
