// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "FileDialogCallback_gen.h"
	"C"
)

// FileDialogCallbackProxy defines methods required for using FileDialogCallback.
type FileDialogCallbackProxy interface {
	Cont(self *FileDialogCallback, selected_accept_filter int32, file_paths StringList)
	Cancel(self *FileDialogCallback)
}

// FileDialogCallback (cef_file_dialog_callback_t from include/capi/cef_dialog_handler_capi.h)
// Callback structure for asynchronous continuation of file dialog requests.
type FileDialogCallback C.cef_file_dialog_callback_t

// NewFileDialogCallback creates a new FileDialogCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewFileDialogCallback(proxy FileDialogCallbackProxy) *FileDialogCallback {
	result := (*FileDialogCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_file_dialog_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_file_dialog_callback_proxy(result.toNative())
	}
	return result
}

func (d *FileDialogCallback) toNative() *C.cef_file_dialog_callback_t {
	return (*C.cef_file_dialog_callback_t)(d)
}

func lookupFileDialogCallbackProxy(obj *BaseRefCounted) FileDialogCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(FileDialogCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type FileDialogCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *FileDialogCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the file selection. |selected_accept_filter| should be the 0-based
// index of the value selected from the accept filters array passed to
// cef_dialog_handler_t::OnFileDialog. |file_paths| should be a single value
// or a list of values depending on the dialog mode. An NULL |file_paths|
// value is treated the same as calling cancel().
func (d *FileDialogCallback) Cont(selected_accept_filter int32, file_paths StringList) {
	lookupFileDialogCallbackProxy(d.Base()).Cont(d, selected_accept_filter, file_paths)
}

//nolint:gocritic
//export gocef_file_dialog_callback_cont
func gocef_file_dialog_callback_cont(self *C.cef_file_dialog_callback_t, selected_accept_filter C.int, file_paths C.cef_string_list_t) {
	me__ := (*FileDialogCallback)(self)
	proxy__ := lookupFileDialogCallbackProxy(me__.Base())
	proxy__.Cont(me__, int32(selected_accept_filter), StringList(file_paths))
}

// Cancel (cancel)
// Cancel the file selection.
func (d *FileDialogCallback) Cancel() {
	lookupFileDialogCallbackProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_file_dialog_callback_cancel
func gocef_file_dialog_callback_cancel(self *C.cef_file_dialog_callback_t) {
	me__ := (*FileDialogCallback)(self)
	proxy__ := lookupFileDialogCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
