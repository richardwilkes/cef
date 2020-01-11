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
	// #include "AuthCallback_gen.h"
	"C"
)

// AuthCallbackProxy defines methods required for using AuthCallback.
type AuthCallbackProxy interface {
	Cont(self *AuthCallback, username, password string)
	Cancel(self *AuthCallback)
}

// AuthCallback (cef_auth_callback_t from include/capi/cef_auth_callback_capi.h)
// Callback structure used for asynchronous continuation of authentication
// requests.
type AuthCallback C.cef_auth_callback_t

// NewAuthCallback creates a new AuthCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewAuthCallback(proxy AuthCallbackProxy) *AuthCallback {
	result := (*AuthCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_auth_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_auth_callback_proxy(result.toNative())
	}
	return result
}

func (d *AuthCallback) toNative() *C.cef_auth_callback_t {
	return (*C.cef_auth_callback_t)(d)
}

func lookupAuthCallbackProxy(obj *BaseRefCounted) AuthCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(AuthCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type AuthCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *AuthCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the authentication request.
func (d *AuthCallback) Cont(username, password string) {
	lookupAuthCallbackProxy(d.Base()).Cont(d, username, password)
}

//nolint:gocritic
//export gocef_auth_callback_cont
func gocef_auth_callback_cont(self *C.cef_auth_callback_t, username *C.cef_string_t, password *C.cef_string_t) {
	me__ := (*AuthCallback)(self)
	proxy__ := lookupAuthCallbackProxy(me__.Base())
	username_ := cefstrToString(username)
	password_ := cefstrToString(password)
	proxy__.Cont(me__, username_, password_)
}

// Cancel (cancel)
// Cancel the authentication request.
func (d *AuthCallback) Cancel() {
	lookupAuthCallbackProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_auth_callback_cancel
func gocef_auth_callback_cancel(self *C.cef_auth_callback_t) {
	me__ := (*AuthCallback)(self)
	proxy__ := lookupAuthCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
