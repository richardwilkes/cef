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
	// #include "BeforeDownloadCallback_gen.h"
	"C"
)

// BeforeDownloadCallbackProxy defines methods required for using BeforeDownloadCallback.
type BeforeDownloadCallbackProxy interface {
	Cont(self *BeforeDownloadCallback, downloadPath string, showDialog int32)
}

// BeforeDownloadCallback (cef_before_download_callback_t from include/capi/cef_download_handler_capi.h)
// Callback structure used to asynchronously continue a download.
type BeforeDownloadCallback C.cef_before_download_callback_t

// NewBeforeDownloadCallback creates a new BeforeDownloadCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewBeforeDownloadCallback(proxy BeforeDownloadCallbackProxy) *BeforeDownloadCallback {
	result := (*BeforeDownloadCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_before_download_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_before_download_callback_proxy(result.toNative())
	}
	return result
}

func (d *BeforeDownloadCallback) toNative() *C.cef_before_download_callback_t {
	return (*C.cef_before_download_callback_t)(d)
}

func lookupBeforeDownloadCallbackProxy(obj *BaseRefCounted) BeforeDownloadCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(BeforeDownloadCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type BeforeDownloadCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *BeforeDownloadCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Call to continue the download. Set |download_path| to the full file path
// for the download including the file name or leave blank to use the
// suggested name and the default temp directory. Set |show_dialog| to true
// (1) if you do wish to show the default "Save As" dialog.
func (d *BeforeDownloadCallback) Cont(downloadPath string, showDialog int32) {
	lookupBeforeDownloadCallbackProxy(d.Base()).Cont(d, downloadPath, showDialog)
}

//nolint:gocritic
//export gocef_before_download_callback_cont
func gocef_before_download_callback_cont(self *C.cef_before_download_callback_t, downloadPath *C.cef_string_t, showDialog C.int) {
	me__ := (*BeforeDownloadCallback)(self)
	proxy__ := lookupBeforeDownloadCallbackProxy(me__.Base())
	downloadPath_ := cefstrToString(downloadPath)
	proxy__.Cont(me__, downloadPath_, int32(showDialog))
}
