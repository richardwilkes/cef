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
	// #include "RequestCallback_gen.h"
	"C"
)

// RequestCallbackProxy defines methods required for using RequestCallback.
type RequestCallbackProxy interface {
	Cont(self *RequestCallback, allow int32)
	Cancel(self *RequestCallback)
}

// RequestCallback (cef_request_callback_t from include/capi/cef_request_callback_capi.h)
// Callback structure used for asynchronous continuation of url requests.
type RequestCallback C.cef_request_callback_t

// NewRequestCallback creates a new RequestCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRequestCallback(proxy RequestCallbackProxy) *RequestCallback {
	result := (*RequestCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_request_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_request_callback_proxy(result.toNative())
	}
	return result
}

func (d *RequestCallback) toNative() *C.cef_request_callback_t {
	return (*C.cef_request_callback_t)(d)
}

func lookupRequestCallbackProxy(obj *BaseRefCounted) RequestCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RequestCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RequestCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RequestCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the url request. If |allow| is true (1) the request will be
// continued. Otherwise, the request will be canceled.
func (d *RequestCallback) Cont(allow int32) {
	lookupRequestCallbackProxy(d.Base()).Cont(d, allow)
}

//nolint:gocritic
//export gocef_request_callback_cont
func gocef_request_callback_cont(self *C.cef_request_callback_t, allow C.int) {
	me__ := (*RequestCallback)(self)
	proxy__ := lookupRequestCallbackProxy(me__.Base())
	proxy__.Cont(me__, int32(allow))
}

// Cancel (cancel)
// Cancel the url request.
func (d *RequestCallback) Cancel() {
	lookupRequestCallbackProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_request_callback_cancel
func gocef_request_callback_cancel(self *C.cef_request_callback_t) {
	me__ := (*RequestCallback)(self)
	proxy__ := lookupRequestCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
