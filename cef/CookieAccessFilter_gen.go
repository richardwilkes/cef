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
	// #include "CookieAccessFilter_gen.h"
	"C"
)

// CookieAccessFilterProxy defines methods required for using CookieAccessFilter.
type CookieAccessFilterProxy interface {
	CanSendCookie(self *CookieAccessFilter, browser *Browser, frame *Frame, request *Request, cookie *Cookie) int32
	CanSaveCookie(self *CookieAccessFilter, browser *Browser, frame *Frame, request *Request, response *Response, cookie *Cookie) int32
}

// CookieAccessFilter (cef_cookie_access_filter_t from include/capi/cef_resource_request_handler_capi.h)
// Implement this structure to filter cookies that may be sent or received from
// resource requests. The functions of this structure will be called on the IO
// thread unless otherwise indicated.
type CookieAccessFilter C.cef_cookie_access_filter_t

// NewCookieAccessFilter creates a new CookieAccessFilter with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewCookieAccessFilter(proxy CookieAccessFilterProxy) *CookieAccessFilter {
	result := (*CookieAccessFilter)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_cookie_access_filter_t, proxy)))
	if proxy != nil {
		C.gocef_set_cookie_access_filter_proxy(result.toNative())
	}
	return result
}

func (d *CookieAccessFilter) toNative() *C.cef_cookie_access_filter_t {
	return (*C.cef_cookie_access_filter_t)(d)
}

func lookupCookieAccessFilterProxy(obj *BaseRefCounted) CookieAccessFilterProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(CookieAccessFilterProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type CookieAccessFilterProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *CookieAccessFilter) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// CanSendCookie (can_send_cookie)
// Called on the IO thread before a resource request is sent. The |browser|
// and |frame| values represent the source of the request, and may be NULL for
// requests originating from service workers or cef_urlrequest_t. |request|
// cannot be modified in this callback. Return true (1) if the specified
// cookie can be sent with the request or false (0) otherwise.
func (d *CookieAccessFilter) CanSendCookie(browser *Browser, frame *Frame, request *Request, cookie *Cookie) int32 {
	return lookupCookieAccessFilterProxy(d.Base()).CanSendCookie(d, browser, frame, request, cookie)
}

//nolint:gocritic
//export gocef_cookie_access_filter_can_send_cookie
func gocef_cookie_access_filter_can_send_cookie(self *C.cef_cookie_access_filter_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, cookie *C.cef_cookie_t) C.int {
	me__ := (*CookieAccessFilter)(self)
	proxy__ := lookupCookieAccessFilterProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.CanSendCookie(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), cookie_))
}

// CanSaveCookie (can_save_cookie)
// Called on the IO thread after a resource response is received. The
// |browser| and |frame| values represent the source of the request, and may
// be NULL for requests originating from service workers or cef_urlrequest_t.
// |request| cannot be modified in this callback. Return true (1) if the
// specified cookie returned with the response can be saved or false (0)
// otherwise.
func (d *CookieAccessFilter) CanSaveCookie(browser *Browser, frame *Frame, request *Request, response *Response, cookie *Cookie) int32 {
	return lookupCookieAccessFilterProxy(d.Base()).CanSaveCookie(d, browser, frame, request, response, cookie)
}

//nolint:gocritic
//export gocef_cookie_access_filter_can_save_cookie
func gocef_cookie_access_filter_can_save_cookie(self *C.cef_cookie_access_filter_t, browser *C.cef_browser_t, frame *C.cef_frame_t, request *C.cef_request_t, response *C.cef_response_t, cookie *C.cef_cookie_t) C.int {
	me__ := (*CookieAccessFilter)(self)
	proxy__ := lookupCookieAccessFilterProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.CanSaveCookie(me__, (*Browser)(browser), (*Frame)(frame), (*Request)(request), (*Response)(response), cookie_))
}
