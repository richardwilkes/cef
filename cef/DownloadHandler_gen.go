// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "DownloadHandler_gen.h"
	"C"
)

// DownloadHandlerProxy defines methods required for using DownloadHandler.
type DownloadHandlerProxy interface {
	OnBeforeDownload(self *DownloadHandler, browser *Browser, download_item *DownloadItem, suggested_name string, callback *BeforeDownloadCallback)
	OnDownloadUpdated(self *DownloadHandler, browser *Browser, download_item *DownloadItem, callback *DownloadItemCallback)
}

// DownloadHandler (cef_download_handler_t from include/capi/cef_download_handler_capi.h)
// Structure used to handle file downloads. The functions of this structure will
// called on the browser process UI thread.
type DownloadHandler C.cef_download_handler_t

// NewDownloadHandler creates a new DownloadHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDownloadHandler(proxy DownloadHandlerProxy) *DownloadHandler {
	result := (*DownloadHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_download_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_download_handler_proxy(result.toNative())
	}
	return result
}

func (d *DownloadHandler) toNative() *C.cef_download_handler_t {
	return (*C.cef_download_handler_t)(d)
}

func lookupDownloadHandlerProxy(obj *BaseRefCounted) DownloadHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DownloadHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DownloadHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DownloadHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnBeforeDownload (on_before_download)
// Called before a download begins. |suggested_name| is the suggested name for
// the download file. By default the download will be canceled. Execute
// |callback| either asynchronously or in this function to continue the
// download if desired. Do not keep a reference to |download_item| outside of
// this function.
func (d *DownloadHandler) OnBeforeDownload(browser *Browser, download_item *DownloadItem, suggested_name string, callback *BeforeDownloadCallback) {
	lookupDownloadHandlerProxy(d.Base()).OnBeforeDownload(d, browser, download_item, suggested_name, callback)
}

//nolint:gocritic
//export gocef_download_handler_on_before_download
func gocef_download_handler_on_before_download(self *C.cef_download_handler_t, browser *C.cef_browser_t, download_item *C.cef_download_item_t, suggested_name *C.cef_string_t, callback *C.cef_before_download_callback_t) {
	me__ := (*DownloadHandler)(self)
	proxy__ := lookupDownloadHandlerProxy(me__.Base())
	suggested_name_ := cefstrToString(suggested_name)
	proxy__.OnBeforeDownload(me__, (*Browser)(browser), (*DownloadItem)(download_item), suggested_name_, (*BeforeDownloadCallback)(callback))
}

// OnDownloadUpdated (on_download_updated)
// Called when a download's status or progress information has been updated.
// This may be called multiple times before and after on_before_download().
// Execute |callback| either asynchronously or in this function to cancel the
// download if desired. Do not keep a reference to |download_item| outside of
// this function.
func (d *DownloadHandler) OnDownloadUpdated(browser *Browser, download_item *DownloadItem, callback *DownloadItemCallback) {
	lookupDownloadHandlerProxy(d.Base()).OnDownloadUpdated(d, browser, download_item, callback)
}

//nolint:gocritic
//export gocef_download_handler_on_download_updated
func gocef_download_handler_on_download_updated(self *C.cef_download_handler_t, browser *C.cef_browser_t, download_item *C.cef_download_item_t, callback *C.cef_download_item_callback_t) {
	me__ := (*DownloadHandler)(self)
	proxy__ := lookupDownloadHandlerProxy(me__.Base())
	proxy__.OnDownloadUpdated(me__, (*Browser)(browser), (*DownloadItem)(download_item), (*DownloadItemCallback)(callback))
}
