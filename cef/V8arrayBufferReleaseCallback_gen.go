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
	// #include "V8arrayBufferReleaseCallback_gen.h"
	"C"
)

// V8arrayBufferReleaseCallbackProxy defines methods required for using V8arrayBufferReleaseCallback.
type V8arrayBufferReleaseCallbackProxy interface {
	ReleaseBuffer(self *V8arrayBufferReleaseCallback, buffer unsafe.Pointer)
}

// V8arrayBufferReleaseCallback (cef_v8array_buffer_release_callback_t from include/capi/cef_v8_capi.h)
// Callback structure that is passed to cef_v8value_t::CreateArrayBuffer.
type V8arrayBufferReleaseCallback C.cef_v8array_buffer_release_callback_t

// NewV8arrayBufferReleaseCallback creates a new V8arrayBufferReleaseCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewV8arrayBufferReleaseCallback(proxy V8arrayBufferReleaseCallbackProxy) *V8arrayBufferReleaseCallback {
	result := (*V8arrayBufferReleaseCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_v8array_buffer_release_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_v8array_buffer_release_callback_proxy(result.toNative())
	}
	return result
}

func (d *V8arrayBufferReleaseCallback) toNative() *C.cef_v8array_buffer_release_callback_t {
	return (*C.cef_v8array_buffer_release_callback_t)(d)
}

func lookupV8arrayBufferReleaseCallbackProxy(obj *BaseRefCounted) V8arrayBufferReleaseCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(V8arrayBufferReleaseCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type V8arrayBufferReleaseCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *V8arrayBufferReleaseCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// ReleaseBuffer (release_buffer)
// Called to release |buffer| when the ArrayBuffer JS object is garbage
// collected. |buffer| is the value that was passed to CreateArrayBuffer along
// with this object.
func (d *V8arrayBufferReleaseCallback) ReleaseBuffer(buffer unsafe.Pointer) {
	lookupV8arrayBufferReleaseCallbackProxy(d.Base()).ReleaseBuffer(d, buffer)
}

//nolint:gocritic
//export gocef_v8array_buffer_release_callback_release_buffer
func gocef_v8array_buffer_release_callback_release_buffer(self *C.cef_v8array_buffer_release_callback_t, buffer unsafe.Pointer) {
	me__ := (*V8arrayBufferReleaseCallback)(self)
	proxy__ := lookupV8arrayBufferReleaseCallbackProxy(me__.Base())
	proxy__.ReleaseBuffer(me__, buffer)
}
