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
	// #include "FindHandler_gen.h"
	"C"
)

// FindHandlerProxy defines methods required for using FindHandler.
type FindHandlerProxy interface {
	OnFindResult(self *FindHandler, browser *Browser, identifier, count int32, selectionRect *Rect, activeMatchOrdinal, finalUpdate int32)
}

// FindHandler (cef_find_handler_t from include/capi/cef_find_handler_capi.h)
// Implement this structure to handle events related to find results. The
// functions of this structure will be called on the UI thread.
type FindHandler C.cef_find_handler_t

// NewFindHandler creates a new FindHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewFindHandler(proxy FindHandlerProxy) *FindHandler {
	result := (*FindHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_find_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_find_handler_proxy(result.toNative())
	}
	return result
}

func (d *FindHandler) toNative() *C.cef_find_handler_t {
	return (*C.cef_find_handler_t)(d)
}

func lookupFindHandlerProxy(obj *BaseRefCounted) FindHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(FindHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type FindHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *FindHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnFindResult (on_find_result)
// Called to report find results returned by cef_browser_host_t::find().
// |identifer| is the identifier passed to find(), |count| is the number of
// matches currently identified, |selectionRect| is the location of where the
// match was found (in window coordinates), |activeMatchOrdinal| is the
// current position in the search results, and |finalUpdate| is true (1) if
// this is the last find notification.
func (d *FindHandler) OnFindResult(browser *Browser, identifier, count int32, selectionRect *Rect, activeMatchOrdinal, finalUpdate int32) {
	lookupFindHandlerProxy(d.Base()).OnFindResult(d, browser, identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
}

//nolint:gocritic
//export gocef_find_handler_on_find_result
func gocef_find_handler_on_find_result(self *C.cef_find_handler_t, browser *C.cef_browser_t, identifier C.int, count C.int, selectionRect *C.cef_rect_t, activeMatchOrdinal C.int, finalUpdate C.int) {
	me__ := (*FindHandler)(self)
	proxy__ := lookupFindHandlerProxy(me__.Base())
	selectionRect_ := selectionRect.toGo()
	proxy__.OnFindResult(me__, (*Browser)(browser), int32(identifier), int32(count), selectionRect_, int32(activeMatchOrdinal), int32(finalUpdate))
}
