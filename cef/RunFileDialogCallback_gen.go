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
	// #include "RunFileDialogCallback_gen.h"
	"C"
)

// RunFileDialogCallbackProxy defines methods required for using RunFileDialogCallback.
type RunFileDialogCallbackProxy interface {
	OnFileDialogDismissed(self *RunFileDialogCallback, selectedAcceptFilter int32, filePaths StringList)
}

// RunFileDialogCallback (cef_run_file_dialog_callback_t from include/capi/cef_browser_capi.h)
// Callback structure for cef_browser_host_t::RunFileDialog. The functions of
// this structure will be called on the browser process UI thread.
type RunFileDialogCallback C.cef_run_file_dialog_callback_t

// NewRunFileDialogCallback creates a new RunFileDialogCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRunFileDialogCallback(proxy RunFileDialogCallbackProxy) *RunFileDialogCallback {
	result := (*RunFileDialogCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_run_file_dialog_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_run_file_dialog_callback_proxy(result.toNative())
	}
	return result
}

func (d *RunFileDialogCallback) toNative() *C.cef_run_file_dialog_callback_t {
	return (*C.cef_run_file_dialog_callback_t)(d)
}

func lookupRunFileDialogCallbackProxy(obj *BaseRefCounted) RunFileDialogCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RunFileDialogCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RunFileDialogCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RunFileDialogCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnFileDialogDismissed (on_file_dialog_dismissed)
// Called asynchronously after the file dialog is dismissed.
// |selected_accept_filter| is the 0-based index of the value selected from
// the accept filters array passed to cef_browser_host_t::RunFileDialog.
// |file_paths| will be a single value or a list of values depending on the
// dialog mode. If the selection was cancelled |file_paths| will be NULL.
func (d *RunFileDialogCallback) OnFileDialogDismissed(selectedAcceptFilter int32, filePaths StringList) {
	lookupRunFileDialogCallbackProxy(d.Base()).OnFileDialogDismissed(d, selectedAcceptFilter, filePaths)
}

//nolint:gocritic
//export gocef_run_file_dialog_callback_on_file_dialog_dismissed
func gocef_run_file_dialog_callback_on_file_dialog_dismissed(self *C.cef_run_file_dialog_callback_t, selectedAcceptFilter C.int, filePaths C.cef_string_list_t) {
	me__ := (*RunFileDialogCallback)(self)
	proxy__ := lookupRunFileDialogCallbackProxy(me__.Base())
	proxy__.OnFileDialogDismissed(me__, int32(selectedAcceptFilter), StringList(filePaths))
}
