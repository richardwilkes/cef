// Code generated - DO NOT EDIT.

package cef

import (
	// #include "DownloadItemCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// DownloadItemCallbackProxy defines methods required for using DownloadItemCallback.
type DownloadItemCallbackProxy interface {
	Cancel(self *DownloadItemCallback)
	Pause(self *DownloadItemCallback)
	Resume(self *DownloadItemCallback)
}

// DownloadItemCallback (cef_download_item_callback_t from include/capi/cef_download_handler_capi.h)
// Callback structure used to asynchronously cancel a download.
type DownloadItemCallback C.cef_download_item_callback_t

// NewDownloadItemCallback creates a new DownloadItemCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDownloadItemCallback(proxy DownloadItemCallbackProxy) *DownloadItemCallback {
	result := (*DownloadItemCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_download_item_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_download_item_callback_proxy(result.toNative())
	}
	return result
}

func (d *DownloadItemCallback) toNative() *C.cef_download_item_callback_t {
	return (*C.cef_download_item_callback_t)(d)
}

func lookupDownloadItemCallbackProxy(obj *BaseRefCounted) DownloadItemCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DownloadItemCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DownloadItemCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DownloadItemCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cancel (cancel)
// Call to cancel the download.
func (d *DownloadItemCallback) Cancel() {
	lookupDownloadItemCallbackProxy(d.Base()).Cancel(d)
}

//export gocef_download_item_callback_cancel
func gocef_download_item_callback_cancel(self *C.cef_download_item_callback_t) {
	me__ := (*DownloadItemCallback)(self)
	proxy__ := lookupDownloadItemCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}

// Pause (pause)
// Call to pause the download.
func (d *DownloadItemCallback) Pause() {
	lookupDownloadItemCallbackProxy(d.Base()).Pause(d)
}

//export gocef_download_item_callback_pause
func gocef_download_item_callback_pause(self *C.cef_download_item_callback_t) {
	me__ := (*DownloadItemCallback)(self)
	proxy__ := lookupDownloadItemCallbackProxy(me__.Base())
	proxy__.Pause(me__)
}

// Resume (resume)
// Call to resume the download.
func (d *DownloadItemCallback) Resume() {
	lookupDownloadItemCallbackProxy(d.Base()).Resume(d)
}

//export gocef_download_item_callback_resume
func gocef_download_item_callback_resume(self *C.cef_download_item_callback_t) {
	me__ := (*DownloadItemCallback)(self)
	proxy__ := lookupDownloadItemCallbackProxy(me__.Base())
	proxy__.Resume(me__)
}
