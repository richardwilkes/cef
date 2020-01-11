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
	// #include "GetExtensionResourceCallback_gen.h"
	"C"
)

// GetExtensionResourceCallbackProxy defines methods required for using GetExtensionResourceCallback.
type GetExtensionResourceCallbackProxy interface {
	Cont(self *GetExtensionResourceCallback, stream *StreamReader)
	Cancel(self *GetExtensionResourceCallback)
}

// GetExtensionResourceCallback (cef_get_extension_resource_callback_t from include/capi/cef_extension_handler_capi.h)
// Callback structure used for asynchronous continuation of
// cef_extension_tHandler::GetExtensionResource.
type GetExtensionResourceCallback C.cef_get_extension_resource_callback_t

// NewGetExtensionResourceCallback creates a new GetExtensionResourceCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewGetExtensionResourceCallback(proxy GetExtensionResourceCallbackProxy) *GetExtensionResourceCallback {
	result := (*GetExtensionResourceCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_get_extension_resource_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_get_extension_resource_callback_proxy(result.toNative())
	}
	return result
}

func (d *GetExtensionResourceCallback) toNative() *C.cef_get_extension_resource_callback_t {
	return (*C.cef_get_extension_resource_callback_t)(d)
}

func lookupGetExtensionResourceCallbackProxy(obj *BaseRefCounted) GetExtensionResourceCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(GetExtensionResourceCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type GetExtensionResourceCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *GetExtensionResourceCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the request. Read the resource contents from |stream|.
func (d *GetExtensionResourceCallback) Cont(stream *StreamReader) {
	lookupGetExtensionResourceCallbackProxy(d.Base()).Cont(d, stream)
}

//nolint:gocritic
//export gocef_get_extension_resource_callback_cont
func gocef_get_extension_resource_callback_cont(self *C.cef_get_extension_resource_callback_t, stream *C.cef_stream_reader_t) {
	me__ := (*GetExtensionResourceCallback)(self)
	proxy__ := lookupGetExtensionResourceCallbackProxy(me__.Base())
	proxy__.Cont(me__, (*StreamReader)(stream))
}

// Cancel (cancel)
// Cancel the request.
func (d *GetExtensionResourceCallback) Cancel() {
	lookupGetExtensionResourceCallbackProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_get_extension_resource_callback_cancel
func gocef_get_extension_resource_callback_cancel(self *C.cef_get_extension_resource_callback_t) {
	me__ := (*GetExtensionResourceCallback)(self)
	proxy__ := lookupGetExtensionResourceCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
