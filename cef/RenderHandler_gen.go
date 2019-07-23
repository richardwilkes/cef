// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "RenderHandler_gen.h"
	"C"
)

// RenderHandlerProxy defines methods required for using RenderHandler.
type RenderHandlerProxy interface {
	GetAccessibilityHandler(self *RenderHandler) *AccessibilityHandler
	GetRootScreenRect(self *RenderHandler, browser *Browser, rect *Rect) int32
	GetViewRect(self *RenderHandler, browser *Browser, rect *Rect)
	GetScreenPoint(self *RenderHandler, browser *Browser, viewX, viewY int32, screenX, screenY *int32) int32
	GetScreenInfo(self *RenderHandler, browser *Browser, screen_info *ScreenInfo) int32
	OnPopupShow(self *RenderHandler, browser *Browser, show int32)
	OnPopupSize(self *RenderHandler, browser *Browser, rect *Rect)
	OnPaint(self *RenderHandler, browser *Browser, type_r PaintElementType, dirtyRectsCount uint64, dirtyRects *Rect, buffer unsafe.Pointer, width, height int32)
	OnAcceleratedPaint(self *RenderHandler, browser *Browser, type_r PaintElementType, dirtyRectsCount uint64, dirtyRects *Rect, shared_handle unsafe.Pointer)
	OnCursorChange(self *RenderHandler, browser *Browser, cursor unsafe.Pointer, type_r CursorType, custom_cursor_info *CursorInfo)
	StartDragging(self *RenderHandler, browser *Browser, drag_data *DragData, allowed_ops DragOperationsMask, x, y int32) int32
	UpdateDragCursor(self *RenderHandler, browser *Browser, operation DragOperationsMask)
	OnScrollOffsetChanged(self *RenderHandler, browser *Browser, x, y float64)
	OnImeCompositionRangeChanged(self *RenderHandler, browser *Browser, selected_range *Range, character_boundsCount uint64, character_bounds *Rect)
	OnTextSelectionChanged(self *RenderHandler, browser *Browser, selected_text string, selected_range *Range)
	OnVirtualKeyboardRequested(self *RenderHandler, browser *Browser, input_mode TextInputMode)
}

// RenderHandler (cef_render_handler_t from include/capi/cef_render_handler_capi.h)
// Implement this structure to handle events when window rendering is disabled.
// The functions of this structure will be called on the UI thread.
type RenderHandler C.cef_render_handler_t

// NewRenderHandler creates a new RenderHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRenderHandler(proxy RenderHandlerProxy) *RenderHandler {
	result := (*RenderHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_render_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_render_handler_proxy(result.toNative())
	}
	return result
}

func (d *RenderHandler) toNative() *C.cef_render_handler_t {
	return (*C.cef_render_handler_t)(d)
}

func lookupRenderHandlerProxy(obj *BaseRefCounted) RenderHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RenderHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RenderHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RenderHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetAccessibilityHandler (get_accessibility_handler)
// Return the handler for accessibility notifications. If no handler is
// provided the default implementation will be used.
func (d *RenderHandler) GetAccessibilityHandler() *AccessibilityHandler {
	return lookupRenderHandlerProxy(d.Base()).GetAccessibilityHandler(d)
}

//nolint:gocritic
//export gocef_render_handler_get_accessibility_handler
func gocef_render_handler_get_accessibility_handler(self *C.cef_render_handler_t) *C.cef_accessibility_handler_t {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	return (proxy__.GetAccessibilityHandler(me__)).toNative()
}

// GetRootScreenRect (get_root_screen_rect)
// Called to retrieve the root window rectangle in screen coordinates. Return
// true (1) if the rectangle was provided. If this function returns false (0)
// the rectangle from GetViewRect will be used.
func (d *RenderHandler) GetRootScreenRect(browser *Browser, rect *Rect) int32 {
	return lookupRenderHandlerProxy(d.Base()).GetRootScreenRect(d, browser, rect)
}

//nolint:gocritic
//export gocef_render_handler_get_root_screen_rect
func gocef_render_handler_get_root_screen_rect(self *C.cef_render_handler_t, browser *C.cef_browser_t, rect *C.cef_rect_t) C.int {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	rect_ := rect.toGo()
	return C.int(proxy__.GetRootScreenRect(me__, (*Browser)(browser), rect_))
}

// GetViewRect (get_view_rect)
// Called to retrieve the view rectangle which is relative to screen
// coordinates. This function must always provide a non-NULL rectangle.
func (d *RenderHandler) GetViewRect(browser *Browser, rect *Rect) {
	lookupRenderHandlerProxy(d.Base()).GetViewRect(d, browser, rect)
}

//nolint:gocritic
//export gocef_render_handler_get_view_rect
func gocef_render_handler_get_view_rect(self *C.cef_render_handler_t, browser *C.cef_browser_t, rect *C.cef_rect_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	rect_ := rect.toGo()
	proxy__.GetViewRect(me__, (*Browser)(browser), rect_)
}

// GetScreenPoint (get_screen_point)
// Called to retrieve the translation from view coordinates to actual screen
// coordinates. Return true (1) if the screen coordinates were provided.
func (d *RenderHandler) GetScreenPoint(browser *Browser, viewX, viewY int32, screenX, screenY *int32) int32 {
	return lookupRenderHandlerProxy(d.Base()).GetScreenPoint(d, browser, viewX, viewY, screenX, screenY)
}

//nolint:gocritic
//export gocef_render_handler_get_screen_point
func gocef_render_handler_get_screen_point(self *C.cef_render_handler_t, browser *C.cef_browser_t, viewX C.int, viewY C.int, screenX *C.int, screenY *C.int) C.int {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	return C.int(proxy__.GetScreenPoint(me__, (*Browser)(browser), int32(viewX), int32(viewY), (*int32)(screenX), (*int32)(screenY)))
}

// GetScreenInfo (get_screen_info)
// Called to allow the client to fill in the CefScreenInfo object with
// appropriate values. Return true (1) if the |screen_info| structure has been
// modified.
//
// If the screen info rectangle is left NULL the rectangle from GetViewRect
// will be used. If the rectangle is still NULL or invalid popups may not be
// drawn correctly.
func (d *RenderHandler) GetScreenInfo(browser *Browser, screen_info *ScreenInfo) int32 {
	return lookupRenderHandlerProxy(d.Base()).GetScreenInfo(d, browser, screen_info)
}

//nolint:gocritic
//export gocef_render_handler_get_screen_info
func gocef_render_handler_get_screen_info(self *C.cef_render_handler_t, browser *C.cef_browser_t, screen_info *C.cef_screen_info_t) C.int {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	screen_info_ := screen_info.toGo()
	return C.int(proxy__.GetScreenInfo(me__, (*Browser)(browser), screen_info_))
}

// OnPopupShow (on_popup_show)
// Called when the browser wants to show or hide the popup widget. The popup
// should be shown if |show| is true (1) and hidden if |show| is false (0).
func (d *RenderHandler) OnPopupShow(browser *Browser, show int32) {
	lookupRenderHandlerProxy(d.Base()).OnPopupShow(d, browser, show)
}

//nolint:gocritic
//export gocef_render_handler_on_popup_show
func gocef_render_handler_on_popup_show(self *C.cef_render_handler_t, browser *C.cef_browser_t, show C.int) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	proxy__.OnPopupShow(me__, (*Browser)(browser), int32(show))
}

// OnPopupSize (on_popup_size)
// Called when the browser wants to move or resize the popup widget. |rect|
// contains the new location and size in view coordinates.
func (d *RenderHandler) OnPopupSize(browser *Browser, rect *Rect) {
	lookupRenderHandlerProxy(d.Base()).OnPopupSize(d, browser, rect)
}

//nolint:gocritic
//export gocef_render_handler_on_popup_size
func gocef_render_handler_on_popup_size(self *C.cef_render_handler_t, browser *C.cef_browser_t, rect *C.cef_rect_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	rect_ := rect.toGo()
	proxy__.OnPopupSize(me__, (*Browser)(browser), rect_)
}

// OnPaint (on_paint)
// Called when an element should be painted. Pixel values passed to this
// function are scaled relative to view coordinates based on the value of
// CefScreenInfo.device_scale_factor returned from GetScreenInfo. |type|
// indicates whether the element is the view or the popup widget. |buffer|
// contains the pixel data for the whole image. |dirtyRects| contains the set
// of rectangles in pixel coordinates that need to be repainted. |buffer| will
// be |width|*|height|*4 bytes in size and represents a BGRA image with an
// upper-left origin. This function is only called when
// cef_window_tInfo::shared_texture_enabled is set to false (0).
func (d *RenderHandler) OnPaint(browser *Browser, type_r PaintElementType, dirtyRectsCount uint64, dirtyRects *Rect, buffer unsafe.Pointer, width, height int32) {
	lookupRenderHandlerProxy(d.Base()).OnPaint(d, browser, type_r, dirtyRectsCount, dirtyRects, buffer, width, height)
}

//nolint:gocritic
//export gocef_render_handler_on_paint
func gocef_render_handler_on_paint(self *C.cef_render_handler_t, browser *C.cef_browser_t, type_r C.cef_paint_element_type_t, dirtyRectsCount C.size_t, dirtyRects *C.cef_rect_t, buffer unsafe.Pointer, width C.int, height C.int) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	dirtyRects_ := dirtyRects.toGo()
	proxy__.OnPaint(me__, (*Browser)(browser), PaintElementType(type_r), uint64(dirtyRectsCount), dirtyRects_, buffer, int32(width), int32(height))
}

// OnAcceleratedPaint (on_accelerated_paint)
// Called when an element has been rendered to the shared texture handle.
// |type| indicates whether the element is the view or the popup widget.
// |dirtyRects| contains the set of rectangles in pixel coordinates that need
// to be repainted. |shared_handle| is the handle for a D3D11 Texture2D that
// can be accessed via ID3D11Device using the OpenSharedResource function.
// This function is only called when cef_window_tInfo::shared_texture_enabled
// is set to true (1), and is currently only supported on Windows.
func (d *RenderHandler) OnAcceleratedPaint(browser *Browser, type_r PaintElementType, dirtyRectsCount uint64, dirtyRects *Rect, shared_handle unsafe.Pointer) {
	lookupRenderHandlerProxy(d.Base()).OnAcceleratedPaint(d, browser, type_r, dirtyRectsCount, dirtyRects, shared_handle)
}

//nolint:gocritic
//export gocef_render_handler_on_accelerated_paint
func gocef_render_handler_on_accelerated_paint(self *C.cef_render_handler_t, browser *C.cef_browser_t, type_r C.cef_paint_element_type_t, dirtyRectsCount C.size_t, dirtyRects *C.cef_rect_t, shared_handle unsafe.Pointer) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	dirtyRects_ := dirtyRects.toGo()
	proxy__.OnAcceleratedPaint(me__, (*Browser)(browser), PaintElementType(type_r), uint64(dirtyRectsCount), dirtyRects_, shared_handle)
}

// OnCursorChange (on_cursor_change)
// Called when the browser's cursor has changed. If |type| is CT_CUSTOM then
// |custom_cursor_info| will be populated with the custom cursor information.
func (d *RenderHandler) OnCursorChange(browser *Browser, cursor unsafe.Pointer, type_r CursorType, custom_cursor_info *CursorInfo) {
	lookupRenderHandlerProxy(d.Base()).OnCursorChange(d, browser, cursor, type_r, custom_cursor_info)
}

//nolint:gocritic
//export gocef_render_handler_on_cursor_change
func gocef_render_handler_on_cursor_change(self *C.cef_render_handler_t, browser *C.cef_browser_t, cursor unsafe.Pointer, type_r C.cef_cursor_type_t, custom_cursor_info *C.cef_cursor_info_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	custom_cursor_info_ := custom_cursor_info.toGo()
	proxy__.OnCursorChange(me__, (*Browser)(browser), cursor, CursorType(type_r), custom_cursor_info_)
}

// StartDragging (start_dragging)
// Called when the user starts dragging content in the web view. Contextual
// information about the dragged content is supplied by |drag_data|. (|x|,
// |y|) is the drag start location in screen coordinates. OS APIs that run a
// system message loop may be used within the StartDragging call.
//
// Return false (0) to abort the drag operation. Don't call any of
// cef_browser_host_t::DragSource*Ended* functions after returning false (0).
//
// Return true (1) to handle the drag operation. Call
// cef_browser_host_t::DragSourceEndedAt and DragSourceSystemDragEnded either
// synchronously or asynchronously to inform the web view that the drag
// operation has ended.
func (d *RenderHandler) StartDragging(browser *Browser, drag_data *DragData, allowed_ops DragOperationsMask, x, y int32) int32 {
	return lookupRenderHandlerProxy(d.Base()).StartDragging(d, browser, drag_data, allowed_ops, x, y)
}

//nolint:gocritic
//export gocef_render_handler_start_dragging
func gocef_render_handler_start_dragging(self *C.cef_render_handler_t, browser *C.cef_browser_t, drag_data *C.cef_drag_data_t, allowed_ops C.cef_drag_operations_mask_t, x C.int, y C.int) C.int {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	return C.int(proxy__.StartDragging(me__, (*Browser)(browser), (*DragData)(drag_data), DragOperationsMask(allowed_ops), int32(x), int32(y)))
}

// UpdateDragCursor (update_drag_cursor)
// Called when the web view wants to update the mouse cursor during a drag &
// drop operation. |operation| describes the allowed operation (none, move,
// copy, link).
func (d *RenderHandler) UpdateDragCursor(browser *Browser, operation DragOperationsMask) {
	lookupRenderHandlerProxy(d.Base()).UpdateDragCursor(d, browser, operation)
}

//nolint:gocritic
//export gocef_render_handler_update_drag_cursor
func gocef_render_handler_update_drag_cursor(self *C.cef_render_handler_t, browser *C.cef_browser_t, operation C.cef_drag_operations_mask_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	proxy__.UpdateDragCursor(me__, (*Browser)(browser), DragOperationsMask(operation))
}

// OnScrollOffsetChanged (on_scroll_offset_changed)
// Called when the scroll offset has changed.
func (d *RenderHandler) OnScrollOffsetChanged(browser *Browser, x, y float64) {
	lookupRenderHandlerProxy(d.Base()).OnScrollOffsetChanged(d, browser, x, y)
}

//nolint:gocritic
//export gocef_render_handler_on_scroll_offset_changed
func gocef_render_handler_on_scroll_offset_changed(self *C.cef_render_handler_t, browser *C.cef_browser_t, x C.double, y C.double) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	proxy__.OnScrollOffsetChanged(me__, (*Browser)(browser), float64(x), float64(y))
}

// OnImeCompositionRangeChanged (on_ime_composition_range_changed)
// Called when the IME composition range has changed. |selected_range| is the
// range of characters that have been selected. |character_bounds| is the
// bounds of each character in view coordinates.
func (d *RenderHandler) OnImeCompositionRangeChanged(browser *Browser, selected_range *Range, character_boundsCount uint64, character_bounds *Rect) {
	lookupRenderHandlerProxy(d.Base()).OnImeCompositionRangeChanged(d, browser, selected_range, character_boundsCount, character_bounds)
}

//nolint:gocritic
//export gocef_render_handler_on_ime_composition_range_changed
func gocef_render_handler_on_ime_composition_range_changed(self *C.cef_render_handler_t, browser *C.cef_browser_t, selected_range *C.cef_range_t, character_boundsCount C.size_t, character_bounds *C.cef_rect_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	selected_range_ := selected_range.toGo()
	character_bounds_ := character_bounds.toGo()
	proxy__.OnImeCompositionRangeChanged(me__, (*Browser)(browser), selected_range_, uint64(character_boundsCount), character_bounds_)
}

// OnTextSelectionChanged (on_text_selection_changed)
// Called when text selection has changed for the specified |browser|.
// |selected_text| is the currently selected text and |selected_range| is the
// character range.
func (d *RenderHandler) OnTextSelectionChanged(browser *Browser, selected_text string, selected_range *Range) {
	lookupRenderHandlerProxy(d.Base()).OnTextSelectionChanged(d, browser, selected_text, selected_range)
}

//nolint:gocritic
//export gocef_render_handler_on_text_selection_changed
func gocef_render_handler_on_text_selection_changed(self *C.cef_render_handler_t, browser *C.cef_browser_t, selected_text *C.cef_string_t, selected_range *C.cef_range_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	selected_text_ := cefstrToString(selected_text)
	selected_range_ := selected_range.toGo()
	proxy__.OnTextSelectionChanged(me__, (*Browser)(browser), selected_text_, selected_range_)
}

// OnVirtualKeyboardRequested (on_virtual_keyboard_requested)
// Called when an on-screen keyboard should be shown or hidden for the
// specified |browser|. |input_mode| specifies what kind of keyboard should be
// opened. If |input_mode| is CEF_TEXT_INPUT_MODE_NONE, any existing keyboard
// for this browser should be hidden.
func (d *RenderHandler) OnVirtualKeyboardRequested(browser *Browser, input_mode TextInputMode) {
	lookupRenderHandlerProxy(d.Base()).OnVirtualKeyboardRequested(d, browser, input_mode)
}

//nolint:gocritic
//export gocef_render_handler_on_virtual_keyboard_requested
func gocef_render_handler_on_virtual_keyboard_requested(self *C.cef_render_handler_t, browser *C.cef_browser_t, input_mode C.cef_text_input_mode_t) {
	me__ := (*RenderHandler)(self)
	proxy__ := lookupRenderHandlerProxy(me__.Base())
	proxy__.OnVirtualKeyboardRequested(me__, (*Browser)(browser), TextInputMode(input_mode))
}
