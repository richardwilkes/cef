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
	// #include "StringVisitor_gen.h"
	"C"
)

// StringVisitorProxy defines methods required for using StringVisitor.
type StringVisitorProxy interface {
	Visit(self *StringVisitor, string_r string)
}

// StringVisitor (cef_string_visitor_t from include/capi/cef_string_visitor_capi.h)
// Implement this structure to receive string values asynchronously.
type StringVisitor C.cef_string_visitor_t

// NewStringVisitor creates a new StringVisitor with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewStringVisitor(proxy StringVisitorProxy) *StringVisitor {
	result := (*StringVisitor)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_string_visitor_t, proxy)))
	if proxy != nil {
		C.gocef_set_string_visitor_proxy(result.toNative())
	}
	return result
}

func (d *StringVisitor) toNative() *C.cef_string_visitor_t {
	return (*C.cef_string_visitor_t)(d)
}

func lookupStringVisitorProxy(obj *BaseRefCounted) StringVisitorProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(StringVisitorProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type StringVisitorProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *StringVisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method that will be executed.
func (d *StringVisitor) Visit(string_r string) {
	lookupStringVisitorProxy(d.Base()).Visit(d, string_r)
}

//nolint:gocritic
//export gocef_string_visitor_visit
func gocef_string_visitor_visit(self *C.cef_string_visitor_t, string_r *C.cef_string_t) {
	me__ := (*StringVisitor)(self)
	proxy__ := lookupStringVisitorProxy(me__.Base())
	string_r_ := cefstrToString(string_r)
	proxy__.Visit(me__, string_r_)
}
