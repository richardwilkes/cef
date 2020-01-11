// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "LoadHandler_gen.h"
	"C"
)

// LoadHandlerProxy defines methods required for using LoadHandler.
type LoadHandlerProxy interface {
	OnLoadingStateChange(self *LoadHandler, browser *Browser, isLoading, canGoBack, canGoForward int32)
	OnLoadStart(self *LoadHandler, browser *Browser, frame *Frame, transitionType TransitionType)
	OnLoadEnd(self *LoadHandler, browser *Browser, frame *Frame, hTTPStatusCode int32)
	OnLoadError(self *LoadHandler, browser *Browser, frame *Frame, errorCode Errorcode, errorText, failedURL string)
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

//nolint:gocritic
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
func (d *LoadHandler) OnLoadStart(browser *Browser, frame *Frame, transitionType TransitionType) {
	lookupLoadHandlerProxy(d.Base()).OnLoadStart(d, browser, frame, transitionType)
}

//nolint:gocritic
//export gocef_load_handler_on_load_start
func gocef_load_handler_on_load_start(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, transitionType C.cef_transition_type_t) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadStart(me__, (*Browser)(browser), (*Frame)(frame), TransitionType(transitionType))
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
func (d *LoadHandler) OnLoadEnd(browser *Browser, frame *Frame, hTTPStatusCode int32) {
	lookupLoadHandlerProxy(d.Base()).OnLoadEnd(d, browser, frame, hTTPStatusCode)
}

//nolint:gocritic
//export gocef_load_handler_on_load_end
func gocef_load_handler_on_load_end(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, hTTPStatusCode C.int) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	proxy__.OnLoadEnd(me__, (*Browser)(browser), (*Frame)(frame), int32(hTTPStatusCode))
}

// OnLoadError (on_load_error)
// Called when a navigation fails or is canceled. This function may be called
// by itself if before commit or in combination with OnLoadStart/OnLoadEnd if
// after commit. |errorCode| is the error code number, |errorText| is the
// error text and |failedUrl| is the URL that failed to load. See
// net\base\net_error_list.h for complete descriptions of the error codes.
func (d *LoadHandler) OnLoadError(browser *Browser, frame *Frame, errorCode Errorcode, errorText, failedURL string) {
	lookupLoadHandlerProxy(d.Base()).OnLoadError(d, browser, frame, errorCode, errorText, failedURL)
}

//nolint:gocritic
//export gocef_load_handler_on_load_error
func gocef_load_handler_on_load_error(self *C.cef_load_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, errorCode C.cef_errorcode_t, errorText *C.cef_string_t, failedURL *C.cef_string_t) {
	me__ := (*LoadHandler)(self)
	proxy__ := lookupLoadHandlerProxy(me__.Base())
	errorText_ := cefstrToString(errorText)
	failedURL_ := cefstrToString(failedURL)
	proxy__.OnLoadError(me__, (*Browser)(browser), (*Frame)(frame), Errorcode(errorCode), errorText_, failedURL_)
}
