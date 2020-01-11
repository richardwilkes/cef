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
	// #include "JsdialogCallback_gen.h"
	"C"
)

// JsdialogCallbackProxy defines methods required for using JsdialogCallback.
type JsdialogCallbackProxy interface {
	Cont(self *JsdialogCallback, success int32, userInput string)
}

// JsdialogCallback (cef_jsdialog_callback_t from include/capi/cef_jsdialog_handler_capi.h)
// Callback structure used for asynchronous continuation of JavaScript dialog
// requests.
type JsdialogCallback C.cef_jsdialog_callback_t

// NewJsdialogCallback creates a new JsdialogCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewJsdialogCallback(proxy JsdialogCallbackProxy) *JsdialogCallback {
	result := (*JsdialogCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_jsdialog_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_jsdialog_callback_proxy(result.toNative())
	}
	return result
}

func (d *JsdialogCallback) toNative() *C.cef_jsdialog_callback_t {
	return (*C.cef_jsdialog_callback_t)(d)
}

func lookupJsdialogCallbackProxy(obj *BaseRefCounted) JsdialogCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(JsdialogCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type JsdialogCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *JsdialogCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the JS dialog request. Set |success| to true (1) if the OK button
// was pressed. The |user_input| value should be specified for prompt dialogs.
func (d *JsdialogCallback) Cont(success int32, userInput string) {
	lookupJsdialogCallbackProxy(d.Base()).Cont(d, success, userInput)
}

//nolint:gocritic
//export gocef_jsdialog_callback_cont
func gocef_jsdialog_callback_cont(self *C.cef_jsdialog_callback_t, success C.int, userInput *C.cef_string_t) {
	me__ := (*JsdialogCallback)(self)
	proxy__ := lookupJsdialogCallbackProxy(me__.Base())
	userInput_ := cefstrToString(userInput)
	proxy__.Cont(me__, int32(success), userInput_)
}
