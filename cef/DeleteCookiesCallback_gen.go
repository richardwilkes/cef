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
	// #include "DeleteCookiesCallback_gen.h"
	"C"
)

// DeleteCookiesCallbackProxy defines methods required for using DeleteCookiesCallback.
type DeleteCookiesCallbackProxy interface {
	OnComplete(self *DeleteCookiesCallback, numDeleted int32)
}

// DeleteCookiesCallback (cef_delete_cookies_callback_t from include/capi/cef_cookie_capi.h)
// Structure to implement to be notified of asynchronous completion via
// cef_cookie_manager_t::delete_cookies().
type DeleteCookiesCallback C.cef_delete_cookies_callback_t

// NewDeleteCookiesCallback creates a new DeleteCookiesCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDeleteCookiesCallback(proxy DeleteCookiesCallbackProxy) *DeleteCookiesCallback {
	result := (*DeleteCookiesCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_delete_cookies_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_delete_cookies_callback_proxy(result.toNative())
	}
	return result
}

func (d *DeleteCookiesCallback) toNative() *C.cef_delete_cookies_callback_t {
	return (*C.cef_delete_cookies_callback_t)(d)
}

func lookupDeleteCookiesCallbackProxy(obj *BaseRefCounted) DeleteCookiesCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DeleteCookiesCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DeleteCookiesCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DeleteCookiesCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnComplete (on_complete)
// Method that will be called upon completion. |num_deleted| will be the
// number of cookies that were deleted.
func (d *DeleteCookiesCallback) OnComplete(numDeleted int32) {
	lookupDeleteCookiesCallbackProxy(d.Base()).OnComplete(d, numDeleted)
}

//nolint:gocritic
//export gocef_delete_cookies_callback_on_complete
func gocef_delete_cookies_callback_on_complete(self *C.cef_delete_cookies_callback_t, numDeleted C.int) {
	me__ := (*DeleteCookiesCallback)(self)
	proxy__ := lookupDeleteCookiesCallbackProxy(me__.Base())
	proxy__.OnComplete(me__, int32(numDeleted))
}
