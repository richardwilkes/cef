// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "DownloadImageCallback_gen.h"
	"C"
)

// DownloadImageCallbackProxy defines methods required for using DownloadImageCallback.
type DownloadImageCallbackProxy interface {
	OnDownloadImageFinished(self *DownloadImageCallback, imageURL string, hTTPStatusCode int32, image *Image)
}

// DownloadImageCallback (cef_download_image_callback_t from include/capi/cef_browser_capi.h)
// Callback structure for cef_browser_host_t::DownloadImage. The functions of
// this structure will be called on the browser process UI thread.
type DownloadImageCallback C.cef_download_image_callback_t

// NewDownloadImageCallback creates a new DownloadImageCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDownloadImageCallback(proxy DownloadImageCallbackProxy) *DownloadImageCallback {
	result := (*DownloadImageCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_download_image_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_download_image_callback_proxy(result.toNative())
	}
	return result
}

func (d *DownloadImageCallback) toNative() *C.cef_download_image_callback_t {
	return (*C.cef_download_image_callback_t)(d)
}

func lookupDownloadImageCallbackProxy(obj *BaseRefCounted) DownloadImageCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DownloadImageCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DownloadImageCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DownloadImageCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnDownloadImageFinished (on_download_image_finished)
// Method that will be executed when the image download has completed.
// |image_url| is the URL that was downloaded and |http_status_code| is the
// resulting HTTP status code. |image| is the resulting image, possibly at
// multiple scale factors, or NULL if the download failed.
func (d *DownloadImageCallback) OnDownloadImageFinished(imageURL string, hTTPStatusCode int32, image *Image) {
	lookupDownloadImageCallbackProxy(d.Base()).OnDownloadImageFinished(d, imageURL, hTTPStatusCode, image)
}

//nolint:gocritic
//export gocef_download_image_callback_on_download_image_finished
func gocef_download_image_callback_on_download_image_finished(self *C.cef_download_image_callback_t, imageURL *C.cef_string_t, hTTPStatusCode C.int, image *C.cef_image_t) {
	me__ := (*DownloadImageCallback)(self)
	proxy__ := lookupDownloadImageCallbackProxy(me__.Base())
	imageURL_ := cefstrToString(imageURL)
	proxy__.OnDownloadImageFinished(me__, imageURL_, int32(hTTPStatusCode), (*Image)(image))
}
