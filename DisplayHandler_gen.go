// Code generated - DO NOT EDIT.

package cef

import (
	// #include "DisplayHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// DisplayHandlerProxy defines methods required for using DisplayHandler.
type DisplayHandlerProxy interface {
	OnAddressChange(self *DisplayHandler, browser *Browser, frame *Frame, url string)
	OnTitleChange(self *DisplayHandler, browser *Browser, title string)
	OnFaviconUrlchange(self *DisplayHandler, browser *Browser, icon_urls StringList)
	OnFullscreenModeChange(self *DisplayHandler, browser *Browser, fullscreen int32)
	OnTooltip(self *DisplayHandler, browser *Browser, text string) int32
	OnStatusMessage(self *DisplayHandler, browser *Browser, value string)
	OnConsoleMessage(self *DisplayHandler, browser *Browser, level LogSeverity, message, source string, line int32) int32
	OnAutoResize(self *DisplayHandler, browser *Browser, new_size *Size) int32
	OnLoadingProgressChange(self *DisplayHandler, browser *Browser, progress float64)
}

// DisplayHandler (cef_display_handler_t from include/capi/cef_display_handler_capi.h)
// Implement this structure to handle events related to browser display state.
// The functions of this structure will be called on the UI thread.
type DisplayHandler C.cef_display_handler_t

// NewDisplayHandler creates a new DisplayHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDisplayHandler(proxy DisplayHandlerProxy) *DisplayHandler {
	result := (*DisplayHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_display_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_display_handler_proxy(result.toNative())
	}
	return result
}

func (d *DisplayHandler) toNative() *C.cef_display_handler_t {
	return (*C.cef_display_handler_t)(d)
}

func lookupDisplayHandlerProxy(obj *BaseRefCounted) DisplayHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DisplayHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DisplayHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DisplayHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnAddressChange (on_address_change)
// Called when a frame's address has changed.
func (d *DisplayHandler) OnAddressChange(browser *Browser, frame *Frame, url string) {
	lookupDisplayHandlerProxy(d.Base()).OnAddressChange(d, browser, frame, url)
}

//export gocef_display_handler_on_address_change
func gocef_display_handler_on_address_change(self *C.cef_display_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, url *C.cef_string_t) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnAddressChange(me__, (*Browser)(browser), (*Frame)(frame), cefstrToString(url))
}

// OnTitleChange (on_title_change)
// Called when the page title changes.
func (d *DisplayHandler) OnTitleChange(browser *Browser, title string) {
	lookupDisplayHandlerProxy(d.Base()).OnTitleChange(d, browser, title)
}

//export gocef_display_handler_on_title_change
func gocef_display_handler_on_title_change(self *C.cef_display_handler_t, browser *C.cef_browser_t, title *C.cef_string_t) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnTitleChange(me__, (*Browser)(browser), cefstrToString(title))
}

// OnFaviconUrlchange (on_favicon_urlchange)
// Called when the page icon changes.
func (d *DisplayHandler) OnFaviconUrlchange(browser *Browser, icon_urls StringList) {
	lookupDisplayHandlerProxy(d.Base()).OnFaviconUrlchange(d, browser, icon_urls)
}

//export gocef_display_handler_on_favicon_urlchange
func gocef_display_handler_on_favicon_urlchange(self *C.cef_display_handler_t, browser *C.cef_browser_t, icon_urls C.cef_string_list_t) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnFaviconUrlchange(me__, (*Browser)(browser), StringList(icon_urls))
}

// OnFullscreenModeChange (on_fullscreen_mode_change)
// Called when web content in the page has toggled fullscreen mode. If
// |fullscreen| is true (1) the content will automatically be sized to fill
// the browser content area. If |fullscreen| is false (0) the content will
// automatically return to its original size and position. The client is
// responsible for resizing the browser if desired.
func (d *DisplayHandler) OnFullscreenModeChange(browser *Browser, fullscreen int32) {
	lookupDisplayHandlerProxy(d.Base()).OnFullscreenModeChange(d, browser, fullscreen)
}

//export gocef_display_handler_on_fullscreen_mode_change
func gocef_display_handler_on_fullscreen_mode_change(self *C.cef_display_handler_t, browser *C.cef_browser_t, fullscreen C.int) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnFullscreenModeChange(me__, (*Browser)(browser), int32(fullscreen))
}

// OnTooltip (on_tooltip)
// Called when the browser is about to display a tooltip. |text| contains the
// text that will be displayed in the tooltip. To handle the display of the
// tooltip yourself return true (1). Otherwise, you can optionally modify
// |text| and then return false (0) to allow the browser to display the
// tooltip. When window rendering is disabled the application is responsible
// for drawing tooltips and the return value is ignored.
func (d *DisplayHandler) OnTooltip(browser *Browser, text string) int32 {
	return lookupDisplayHandlerProxy(d.Base()).OnTooltip(d, browser, text)
}

//export gocef_display_handler_on_tooltip
func gocef_display_handler_on_tooltip(self *C.cef_display_handler_t, browser *C.cef_browser_t, text *C.cef_string_t) C.int {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	return C.int(proxy__.OnTooltip(me__, (*Browser)(browser), cefstrToString(text)))
}

// OnStatusMessage (on_status_message)
// Called when the browser receives a status message. |value| contains the
// text that will be displayed in the status message.
func (d *DisplayHandler) OnStatusMessage(browser *Browser, value string) {
	lookupDisplayHandlerProxy(d.Base()).OnStatusMessage(d, browser, value)
}

//export gocef_display_handler_on_status_message
func gocef_display_handler_on_status_message(self *C.cef_display_handler_t, browser *C.cef_browser_t, value *C.cef_string_t) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnStatusMessage(me__, (*Browser)(browser), cefstrToString(value))
}

// OnConsoleMessage (on_console_message)
// Called to display a console message. Return true (1) to stop the message
// from being output to the console.
func (d *DisplayHandler) OnConsoleMessage(browser *Browser, level LogSeverity, message, source string, line int32) int32 {
	return lookupDisplayHandlerProxy(d.Base()).OnConsoleMessage(d, browser, level, message, source, line)
}

//export gocef_display_handler_on_console_message
func gocef_display_handler_on_console_message(self *C.cef_display_handler_t, browser *C.cef_browser_t, level C.cef_log_severity_t, message *C.cef_string_t, source *C.cef_string_t, line C.int) C.int {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	return C.int(proxy__.OnConsoleMessage(me__, (*Browser)(browser), LogSeverity(level), cefstrToString(message), cefstrToString(source), int32(line)))
}

// OnAutoResize (on_auto_resize)
// Called when auto-resize is enabled via
// cef_browser_host_t::SetAutoResizeEnabled and the contents have auto-
// resized. |new_size| will be the desired size in view coordinates. Return
// true (1) if the resize was handled or false (0) for default handling.
func (d *DisplayHandler) OnAutoResize(browser *Browser, new_size *Size) int32 {
	return lookupDisplayHandlerProxy(d.Base()).OnAutoResize(d, browser, new_size)
}

//export gocef_display_handler_on_auto_resize
func gocef_display_handler_on_auto_resize(self *C.cef_display_handler_t, browser *C.cef_browser_t, new_size *C.cef_size_t) C.int {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	var vnew_size Size
	return C.int(proxy__.OnAutoResize(me__, (*Browser)(browser), vnew_size.fromNative(new_size)))
}

// OnLoadingProgressChange (on_loading_progress_change)
// Called when the overall page loading progress has changed. |progress|
// ranges from 0.0 to 1.0.
func (d *DisplayHandler) OnLoadingProgressChange(browser *Browser, progress float64) {
	lookupDisplayHandlerProxy(d.Base()).OnLoadingProgressChange(d, browser, progress)
}

//export gocef_display_handler_on_loading_progress_change
func gocef_display_handler_on_loading_progress_change(self *C.cef_display_handler_t, browser *C.cef_browser_t, progress C.double) {
	me__ := (*DisplayHandler)(self)
	proxy__ := lookupDisplayHandlerProxy(me__.Base())
	proxy__.OnLoadingProgressChange(me__, (*Browser)(browser), float64(progress))
}
