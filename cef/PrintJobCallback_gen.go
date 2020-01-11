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
	// #include "PrintJobCallback_gen.h"
	"C"
)

// PrintJobCallbackProxy defines methods required for using PrintJobCallback.
type PrintJobCallbackProxy interface {
	Cont(self *PrintJobCallback)
}

// PrintJobCallback (cef_print_job_callback_t from include/capi/cef_print_handler_capi.h)
// Callback structure for asynchronous continuation of print job requests.
type PrintJobCallback C.cef_print_job_callback_t

// NewPrintJobCallback creates a new PrintJobCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewPrintJobCallback(proxy PrintJobCallbackProxy) *PrintJobCallback {
	result := (*PrintJobCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_print_job_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_print_job_callback_proxy(result.toNative())
	}
	return result
}

func (d *PrintJobCallback) toNative() *C.cef_print_job_callback_t {
	return (*C.cef_print_job_callback_t)(d)
}

func lookupPrintJobCallbackProxy(obj *BaseRefCounted) PrintJobCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(PrintJobCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type PrintJobCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *PrintJobCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Indicate completion of the print job.
func (d *PrintJobCallback) Cont() {
	lookupPrintJobCallbackProxy(d.Base()).Cont(d)
}

//nolint:gocritic
//export gocef_print_job_callback_cont
func gocef_print_job_callback_cont(self *C.cef_print_job_callback_t) {
	me__ := (*PrintJobCallback)(self)
	proxy__ := lookupPrintJobCallbackProxy(me__.Base())
	proxy__.Cont(me__)
}
