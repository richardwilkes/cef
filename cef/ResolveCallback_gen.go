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
	// #include "ResolveCallback_gen.h"
	"C"
)

// ResolveCallbackProxy defines methods required for using ResolveCallback.
type ResolveCallbackProxy interface {
	OnResolveCompleted(self *ResolveCallback, result Errorcode, resolvedIps StringList)
}

// ResolveCallback (cef_resolve_callback_t from include/capi/cef_request_context_capi.h)
// Callback structure for cef_request_tContext::ResolveHost.
type ResolveCallback C.cef_resolve_callback_t

// NewResolveCallback creates a new ResolveCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResolveCallback(proxy ResolveCallbackProxy) *ResolveCallback {
	result := (*ResolveCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resolve_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_resolve_callback_proxy(result.toNative())
	}
	return result
}

func (d *ResolveCallback) toNative() *C.cef_resolve_callback_t {
	return (*C.cef_resolve_callback_t)(d)
}

func lookupResolveCallbackProxy(obj *BaseRefCounted) ResolveCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResolveCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResolveCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResolveCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnResolveCompleted (on_resolve_completed)
// Called on the UI thread after the ResolveHost request has completed.
// |result| will be the result code. |resolved_ips| will be the list of
// resolved IP addresses or NULL if the resolution failed.
func (d *ResolveCallback) OnResolveCompleted(result Errorcode, resolvedIps StringList) {
	lookupResolveCallbackProxy(d.Base()).OnResolveCompleted(d, result, resolvedIps)
}

//nolint:gocritic
//export gocef_resolve_callback_on_resolve_completed
func gocef_resolve_callback_on_resolve_completed(self *C.cef_resolve_callback_t, result C.cef_errorcode_t, resolvedIps C.cef_string_list_t) {
	me__ := (*ResolveCallback)(self)
	proxy__ := lookupResolveCallbackProxy(me__.Base())
	proxy__.OnResolveCompleted(me__, Errorcode(result), StringList(resolvedIps))
}
