// Code generated - DO NOT EDIT.

package cef

import (
	// #include "BeforeDownloadCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// BeforeDownloadCallbackProxy defines methods required for using BeforeDownloadCallback.
type BeforeDownloadCallbackProxy interface {
	Cont(self *BeforeDownloadCallback, download_path string, show_dialog int32)
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
func (d *BeforeDownloadCallback) Cont(download_path string, show_dialog int32) {
	lookupBeforeDownloadCallbackProxy(d.Base()).Cont(d, download_path, show_dialog)
}

//export gocef_before_download_callback_cont
func gocef_before_download_callback_cont(self *C.cef_before_download_callback_t, download_path *C.cef_string_t, show_dialog C.int) {
	me__ := (*BeforeDownloadCallback)(self)
	proxy__ := lookupBeforeDownloadCallbackProxy(me__.Base())
	proxy__.Cont(me__, cefstrToString(download_path), int32(show_dialog))
}
