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
	// #include "ResourceReadCallback_gen.h"
	"C"
)

// ResourceReadCallbackProxy defines methods required for using ResourceReadCallback.
type ResourceReadCallbackProxy interface {
	Cont(self *ResourceReadCallback, bytesRead int32)
}

// ResourceReadCallback (cef_resource_read_callback_t from include/capi/cef_resource_handler_capi.h)
// Callback for asynchronous continuation of cef_resource_handler_t::read().
type ResourceReadCallback C.cef_resource_read_callback_t

// NewResourceReadCallback creates a new ResourceReadCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResourceReadCallback(proxy ResourceReadCallbackProxy) *ResourceReadCallback {
	result := (*ResourceReadCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resource_read_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_resource_read_callback_proxy(result.toNative())
	}
	return result
}

func (d *ResourceReadCallback) toNative() *C.cef_resource_read_callback_t {
	return (*C.cef_resource_read_callback_t)(d)
}

func lookupResourceReadCallbackProxy(obj *BaseRefCounted) ResourceReadCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResourceReadCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResourceReadCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResourceReadCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Callback for asynchronous continuation of read(). If |bytes_read| == 0 the
// response will be considered complete. If |bytes_read| > 0 then read() will
// be called again until the request is complete (based on either the result
// or the expected content length). If |bytes_read| < 0 then the request will
// fail and the |bytes_read| value will be treated as the error code.
func (d *ResourceReadCallback) Cont(bytesRead int32) {
	lookupResourceReadCallbackProxy(d.Base()).Cont(d, bytesRead)
}

//nolint:gocritic
//export gocef_resource_read_callback_cont
func gocef_resource_read_callback_cont(self *C.cef_resource_read_callback_t, bytesRead C.int) {
	me__ := (*ResourceReadCallback)(self)
	proxy__ := lookupResourceReadCallbackProxy(me__.Base())
	proxy__.Cont(me__, int32(bytesRead))
}
