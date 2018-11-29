// Code generated - DO NOT EDIT.

package cef

import (
	// #include "LifeSpanHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// LifeSpanHandlerProxy defines methods required for using LifeSpanHandler.
type LifeSpanHandlerProxy interface {
	OnBeforePopup(self *LifeSpanHandler, browser *Browser, frame *Frame, target_url string, target_frame_name string, target_disposition WindowOpenDisposition, user_gesture int32, popupFeatures *PopupFeatures, windowInfo *WindowInfo, client **Client, settings *BrowserSettings, no_javascript_access *int32) int32
	OnAfterCreated(self *LifeSpanHandler, browser *Browser)
	DoClose(self *LifeSpanHandler, browser *Browser) int32
	OnBeforeClose(self *LifeSpanHandler, browser *Browser)
}

// LifeSpanHandler (cef_life_span_handler_t from include/capi/cef_life_span_handler_capi.h)
// Implement this structure to handle events related to browser life span. The
// functions of this structure will be called on the UI thread unless otherwise
// indicated.
type LifeSpanHandler C.cef_life_span_handler_t

// NewLifeSpanHandler creates a new LifeSpanHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewLifeSpanHandler(proxy LifeSpanHandlerProxy) *LifeSpanHandler {
	result := (*LifeSpanHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_life_span_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_life_span_handler_proxy(result.toNative())
	}
	return result
}

func (d *LifeSpanHandler) toNative() *C.cef_life_span_handler_t {
	return (*C.cef_life_span_handler_t)(d)
}

func lookupLifeSpanHandlerProxy(obj *BaseRefCounted) LifeSpanHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(LifeSpanHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type LifeSpanHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *LifeSpanHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnBeforePopup (on_before_popup)
// Called on the UI thread before a new popup browser is created. The
// |browser| and |frame| values represent the source of the popup request. The
// |target_url| and |target_frame_name| values indicate where the popup
// browser should navigate and may be NULL if not specified with the request.
// The |target_disposition| value indicates where the user intended to open
// the popup (e.g. current tab, new tab, etc). The |user_gesture| value will
// be true (1) if the popup was opened via explicit user gesture (e.g.
// clicking a link) or false (0) if the popup opened automatically (e.g. via
// the DomContentLoaded event). The |popupFeatures| structure contains
// additional information about the requested popup window. To allow creation
// of the popup browser optionally modify |windowInfo|, |client|, |settings|
// and |no_javascript_access| and return false (0). To cancel creation of the
// popup browser return true (1). The |client| and |settings| values will
// default to the source browser's values. If the |no_javascript_access| value
// is set to false (0) the new browser will not be scriptable and may not be
// hosted in the same renderer process as the source browser. Any
// modifications to |windowInfo| will be ignored if the parent browser is
// wrapped in a cef_browser_view_t. Popup browser creation will be canceled if
// the parent browser is destroyed before the popup browser creation completes
// (indicated by a call to OnAfterCreated for the popup browser).
func (d *LifeSpanHandler) OnBeforePopup(browser *Browser, frame *Frame, target_url, target_frame_name string, target_disposition WindowOpenDisposition, user_gesture int32, popupFeatures *PopupFeatures, windowInfo *WindowInfo, client **Client, settings *BrowserSettings, no_javascript_access *int32) int32 {
	return lookupLifeSpanHandlerProxy(d.Base()).OnBeforePopup(d, browser, frame, target_url, target_frame_name, target_disposition, user_gesture, popupFeatures, windowInfo, client, settings, no_javascript_access)
}

//export gocef_life_span_handler_on_before_popup
func gocef_life_span_handler_on_before_popup(self *C.cef_life_span_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, target_url *C.cef_string_t, target_frame_name *C.cef_string_t, target_disposition C.cef_window_open_disposition_t, user_gesture C.int, popupFeatures *C.cef_popup_features_t, windowInfo *C.cef_window_info_t, client **C.cef_client_t, settings *C.cef_browser_settings_t, no_javascript_access *C.int) C.int {
	me__ := (*LifeSpanHandler)(self)
	proxy__ := lookupLifeSpanHandlerProxy(me__.Base())
	var vpopupFeatures PopupFeatures
	var vwindowInfo WindowInfo
	vclient := (*Client)(*client)
	var vsettings BrowserSettings
	return C.int(proxy__.OnBeforePopup(me__, (*Browser)(browser), (*Frame)(frame), cefstrToString(target_url), cefstrToString(target_frame_name), WindowOpenDisposition(target_disposition), int32(user_gesture), vpopupFeatures.fromNative(popupFeatures), vwindowInfo.fromNative(windowInfo), &vclient, vsettings.fromNative(settings), (*int32)(no_javascript_access)))
}

// OnAfterCreated (on_after_created)
// Called after a new browser is created. This callback will be the first
// notification that references |browser|.
func (d *LifeSpanHandler) OnAfterCreated(browser *Browser) {
	lookupLifeSpanHandlerProxy(d.Base()).OnAfterCreated(d, browser)
}

//export gocef_life_span_handler_on_after_created
func gocef_life_span_handler_on_after_created(self *C.cef_life_span_handler_t, browser *C.cef_browser_t) {
	me__ := (*LifeSpanHandler)(self)
	proxy__ := lookupLifeSpanHandlerProxy(me__.Base())
	proxy__.OnAfterCreated(me__, (*Browser)(browser))
}

// DoClose (do_close)
// Called when a browser has recieved a request to close. This may result
// directly from a call to cef_browser_host_t::*close_browser() or indirectly
// if the browser is parented to a top-level window created by CEF and the
// user attempts to close that window (by clicking the 'X', for example). The
// do_close() function will be called after the JavaScript 'onunload' event
// has been fired.
//
// An application should handle top-level owner window close notifications by
// calling cef_browser_host_t::try_close_browser() or
// cef_browser_host_t::CloseBrowser(false (0)) instead of allowing the window
// to close immediately (see the examples below). This gives CEF an
// opportunity to process the 'onbeforeunload' event and optionally cancel the
// close before do_close() is called.
//
// When windowed rendering is enabled CEF will internally create a window or
// view to host the browser. In that case returning false (0) from do_close()
// will send the standard close notification to the browser's top-level owner
// window (e.g. WM_CLOSE on Windows, performClose: on OS X, "delete_event" on
// Linux or cef_window_delegate_t::can_close() callback from Views). If the
// browser's host window/view has already been destroyed (via view hierarchy
// tear-down, for example) then do_close() will not be called for that browser
// since is no longer possible to cancel the close.
//
// When windowed rendering is disabled returning false (0) from do_close()
// will cause the browser object to be destroyed immediately.
//
// If the browser's top-level owner window requires a non-standard close
// notification then send that notification from do_close() and return true
// (1).
//
// The cef_life_span_handler_t::on_before_close() function will be called
// after do_close() (if do_close() is called) and immediately before the
// browser object is destroyed. The application should only exit after
// on_before_close() has been called for all existing browsers.
//
// The below examples describe what should happen during window close when the
// browser is parented to an application-provided top-level window.
//
// Example 1: Using cef_browser_host_t::try_close_browser(). This is
// recommended for clients using standard close handling and windows created
// on the browser process UI thread. 1.  User clicks the window close button
// which sends a close notification to
//     the application's top-level window.
// 2.  Application's top-level window receives the close notification and
//     calls TryCloseBrowser() (which internally calls CloseBrowser(false)).
//     TryCloseBrowser() returns false so the client cancels the window close.
// 3.  JavaScript 'onbeforeunload' handler executes and shows the close
//     confirmation dialog (which can be overridden via
//     CefJSDialogHandler::OnBeforeUnloadDialog()).
// 4.  User approves the close. 5.  JavaScript 'onunload' handler executes. 6.
// CEF sends a close notification to the application's top-level window
//     (because DoClose() returned false by default).
// 7.  Application's top-level window receives the close notification and
//     calls TryCloseBrowser(). TryCloseBrowser() returns true so the client
//     allows the window close.
// 8.  Application's top-level window is destroyed. 9.  Application's
// on_before_close() handler is called and the browser object
//     is destroyed.
// 10. Application exits by calling cef_quit_message_loop() if no other
// browsers
//     exist.
//
// Example 2: Using cef_browser_host_t::CloseBrowser(false (0)) and
// implementing the do_close() callback. This is recommended for clients using
// non-standard close handling or windows that were not created on the browser
// process UI thread. 1.  User clicks the window close button which sends a
// close notification to
//     the application's top-level window.
// 2.  Application's top-level window receives the close notification and:
//     A. Calls CefBrowserHost::CloseBrowser(false).
//     B. Cancels the window close.
// 3.  JavaScript 'onbeforeunload' handler executes and shows the close
//     confirmation dialog (which can be overridden via
//     CefJSDialogHandler::OnBeforeUnloadDialog()).
// 4.  User approves the close. 5.  JavaScript 'onunload' handler executes. 6.
// Application's do_close() handler is called. Application will:
//     A. Set a flag to indicate that the next close attempt will be allowed.
//     B. Return false.
// 7.  CEF sends an close notification to the application's top-level window.
// 8.  Application's top-level window receives the close notification and
//     allows the window to close based on the flag from #6B.
// 9.  Application's top-level window is destroyed. 10. Application's
// on_before_close() handler is called and the browser object
//     is destroyed.
// 11. Application exits by calling cef_quit_message_loop() if no other
// browsers
//     exist.
func (d *LifeSpanHandler) DoClose(browser *Browser) int32 {
	return lookupLifeSpanHandlerProxy(d.Base()).DoClose(d, browser)
}

//export gocef_life_span_handler_do_close
func gocef_life_span_handler_do_close(self *C.cef_life_span_handler_t, browser *C.cef_browser_t) C.int {
	me__ := (*LifeSpanHandler)(self)
	proxy__ := lookupLifeSpanHandlerProxy(me__.Base())
	return C.int(proxy__.DoClose(me__, (*Browser)(browser)))
}

// OnBeforeClose (on_before_close)
// Called just before a browser is destroyed. Release all references to the
// browser object and do not attempt to execute any functions on the browser
// object after this callback returns. This callback will be the last
// notification that references |browser|. See do_close() documentation for
// additional usage information.
func (d *LifeSpanHandler) OnBeforeClose(browser *Browser) {
	lookupLifeSpanHandlerProxy(d.Base()).OnBeforeClose(d, browser)
}

//export gocef_life_span_handler_on_before_close
func gocef_life_span_handler_on_before_close(self *C.cef_life_span_handler_t, browser *C.cef_browser_t) {
	me__ := (*LifeSpanHandler)(self)
	proxy__ := lookupLifeSpanHandlerProxy(me__.Base())
	proxy__.OnBeforeClose(me__, (*Browser)(browser))
}
