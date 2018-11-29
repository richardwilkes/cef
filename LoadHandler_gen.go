// Code generated - DO NOT EDIT.

package cef

import (
	// #include "LoadHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// LoadHandlerProxy defines methods required for using LoadHandler.
type LoadHandlerProxy interface {
	OnLoadingStateChange(self *LoadHandler, browser *Browser, isLoading int32, canGoBack int32, canGoForward int32)
	OnLoadStart(self *LoadHandler, browser *Browser, frame *Frame, transition_type TransitionType)
	OnLoadEnd(self *LoadHandler, browser *Browser, frame *Frame, httpStatusCode int32)
	OnLoadError(self *LoadHandler, browser *Browser, frame *Frame, errorCode Errorcode, errorText string, failedUrl string)
}

// LoadHandler (cef_load_handler_t from include/capi/cef_load_handler_capi.h)
// Implement this structure to handle events related to browser load status. The
// functions of this structure will be called on the browser process UI thread
// or render process main thread (TID_RENDERER).
type LoadHandler C.cef_load_handler_t

// NewLoadHandler creates a new LoadHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewLoadHandler(proxy LoadHandlerProxy) *LoadHandler {
	result := (*LoadHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_load_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_load_handler_proxy(result.toNative())
	}
	return result
}

func (d *LoadHandler) toNative() *C.cef_load_handler_t {
	return (*C.cef_load_handler_t)(d)
}

func lookupLoadHandlerProxy(obj *BaseRefCounted) LoadHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(LoadHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type LoadHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *LoadHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnLoadingStateChange (on_loading_state_change)
// Called when the loading state has changed. This callback will be executed
// twice -- once when loading is initiated either programmatically or by user
// action, and once when loading is terminated due to completion, cancellation
// of failure. It will be called before any calls to OnLoadStart and after all
// calls to OnLoadError and/or OnLoadEnd.
func (d *LoadHandler) OnLoadingStateChange(browser *Browser, isLoading, canGoBack, canGoForward int32) {
	lookupLoadHandlerProxy(d.Base()).OnLoadingStateChange(d, browser, isLoading, canGoBack, canGoForward)
}

//export gocef_load_handler_on_loading_state_change
func gocef_load_handler_on_loading_state_change(self *C.cef_load_handler_t, browser *C.cef_browser_t, isLoading C.int, canGoBack C.int, canGoForward C.int) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadingStateChange(me__, (*Browser)(browser), int32(isLoading), int32(canGoBack), int32(canGoForward))
}

// OnLoadStart (on_load_start)
// Called after a navigation has been committed and before the browser begins
// loading contents in the frame. The |frame| value will never be NULL -- call
// the is_main() function to check if this frame is the main frame.
// |transition_type| provides information about the source of the navigation
// and an accurate value is only available in the browser process. Multiple
// frames may be loading at the same time. Sub-frames may start or continue
// loading after the main frame load has ended. This function will not be
// called for same page navigations (fragments, history state, etc.) or for
// navigations that fail or are canceled before commit. For notification of
// overall browser load status use OnLoadingStateChange instead.
func (d *LoadHandler) OnLoadStart(browser *Browser, frame *Frame, transition_type TransitionType) {
	lookupLoadHandlerProxy(d.Base()).OnLoadStart(d, browser, frame, transition_type)
}

//export gocef_load_handler_on_load_start
func gocef_load_handler_on_load_start(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, transition_type C.cef_transition_type_t) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadStart(me__, (*Browser)(browser), (*Frame)(frame), TransitionType(transition_type))
}

// OnLoadEnd (on_load_end)
// Called when the browser is done loading a frame. The |frame| value will
// never be NULL -- call the is_main() function to check if this frame is the
// main frame. Multiple frames may be loading at the same time. Sub-frames may
// start or continue loading after the main frame load has ended. This
// function will not be called for same page navigations (fragments, history
// state, etc.) or for navigations that fail or are canceled before commit.
// For notification of overall browser load status use OnLoadingStateChange
// instead.
func (d *LoadHandler) OnLoadEnd(browser *Browser, frame *Frame, httpStatusCode int32) {
	lookupLoadHandlerProxy(d.Base()).OnLoadEnd(d, browser, frame, httpStatusCode)
}

//export gocef_load_handler_on_load_end
func gocef_load_handler_on_load_end(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, httpStatusCode C.int) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadEnd(me__, (*Browser)(browser), (*Frame)(frame), int32(httpStatusCode))
}

// OnLoadError (on_load_error)
// Called when a navigation fails or is canceled. This function may be called
// by itself if before commit or in combination with OnLoadStart/OnLoadEnd if
// after commit. |errorCode| is the error code number, |errorText| is the
// error text and |failedUrl| is the URL that failed to load. See
// net\base\net_error_list.h for complete descriptions of the error codes.
func (d *LoadHandler) OnLoadError(browser *Browser, frame *Frame, errorCode Errorcode, errorText, failedUrl string) {
	lookupLoadHandlerProxy(d.Base()).OnLoadError(d, browser, frame, errorCode, errorText, failedUrl)
}

//export gocef_load_handler_on_load_error
func gocef_load_handler_on_load_error(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, errorCode C.cef_errorcode_t, errorText *C.cef_string_t, failedUrl *C.cef_string_t) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadError(me__, (*Browser)(browser), (*Frame)(frame), Errorcode(errorCode), cefstrToString(errorText), cefstrToString(failedUrl))
}
