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
	// #include "DialogHandler_gen.h"
	"C"
)

// DialogHandlerProxy defines methods required for using DialogHandler.
type DialogHandlerProxy interface {
	OnFileDialog(self *DialogHandler, browser *Browser, mode FileDialogMode, title, defaultFilePath string, acceptFilters StringList, selectedAcceptFilter int32, callback *FileDialogCallback) int32
}

// DialogHandler (cef_dialog_handler_t from include/capi/cef_dialog_handler_capi.h)
// Implement this structure to handle dialog events. The functions of this
// structure will be called on the browser process UI thread.
type DialogHandler C.cef_dialog_handler_t

// NewDialogHandler creates a new DialogHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDialogHandler(proxy DialogHandlerProxy) *DialogHandler {
	result := (*DialogHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_dialog_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_dialog_handler_proxy(result.toNative())
	}
	return result
}

func (d *DialogHandler) toNative() *C.cef_dialog_handler_t {
	return (*C.cef_dialog_handler_t)(d)
}

func lookupDialogHandlerProxy(obj *BaseRefCounted) DialogHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DialogHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DialogHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DialogHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnFileDialog (on_file_dialog)
// Called to run a file chooser dialog. |mode| represents the type of dialog
// to display. |title| to the title to be used for the dialog and may be NULL
// to show the default title ("Open" or "Save" depending on the mode).
// |default_file_path| is the path with optional directory and/or file name
// component that should be initially selected in the dialog. |accept_filters|
// are used to restrict the selectable file types and may any combination of
// (a) valid lower-cased MIME types (e.g. "text/*" or "image/*"), (b)
// individual file extensions (e.g. ".txt" or ".png"), or (c) combined
// description and file extension delimited using "|" and ";" (e.g. "Image
// Types|.png;.gif;.jpg"). |selected_accept_filter| is the 0-based index of
// the filter that should be selected by default. To display a custom dialog
// return true (1) and execute |callback| either inline or at a later time. To
// display the default dialog return false (0).
func (d *DialogHandler) OnFileDialog(browser *Browser, mode FileDialogMode, title, defaultFilePath string, acceptFilters StringList, selectedAcceptFilter int32, callback *FileDialogCallback) int32 {
	return lookupDialogHandlerProxy(d.Base()).OnFileDialog(d, browser, mode, title, defaultFilePath, acceptFilters, selectedAcceptFilter, callback)
}

//nolint:gocritic
//export gocef_dialog_handler_on_file_dialog
func gocef_dialog_handler_on_file_dialog(self *C.cef_dialog_handler_t, browser *C.cef_browser_t, mode C.cef_file_dialog_mode_t, title *C.cef_string_t, defaultFilePath *C.cef_string_t, acceptFilters C.cef_string_list_t, selectedAcceptFilter C.int, callback *C.cef_file_dialog_callback_t) C.int {
	me__ := (*DialogHandler)(self)
	proxy__ := lookupDialogHandlerProxy(me__.Base())
	title_ := cefstrToString(title)
	defaultFilePath_ := cefstrToString(defaultFilePath)
	return C.int(proxy__.OnFileDialog(me__, (*Browser)(browser), FileDialogMode(mode), title_, defaultFilePath_, StringList(acceptFilters), int32(selectedAcceptFilter), (*FileDialogCallback)(callback)))
}
