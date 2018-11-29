package cef

import (
	// #include "DeleteCookiesCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// DeleteCookiesCallbackProxy defines methods required for using DeleteCookiesCallback.
type DeleteCookiesCallbackProxy interface {
	OnComplete(self *DeleteCookiesCallback, num_deleted int32)
}

// DeleteCookiesCallback (cef_delete_cookies_callback_t from include/capi/cef_cookie_capi.h)
// Structure to implement to be notified of asynchronous completion via
// cef_cookie_manager_t::delete_cookies().
type DeleteCookiesCallback C.cef_delete_cookies_callback_t

// NewDeleteCookiesCallback creates a new DeleteCookiesCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDeleteCookiesCallback(proxy DeleteCookiesCallbackProxy) *DeleteCookiesCallback {
	result := (*DeleteCookiesCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_delete_cookies_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_delete_cookies_callback_proxy(result.toNative())
	}
	return result
}

func (d *DeleteCookiesCallback) toNative() *C.cef_delete_cookies_callback_t {
	return (*C.cef_delete_cookies_callback_t)(d)
}

func lookupDeleteCookiesCallbackProxy(obj *BaseRefCounted) DeleteCookiesCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DeleteCookiesCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DeleteCookiesCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DeleteCookiesCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnComplete (on_complete)
// Method that will be called upon completion. |num_deleted| will be the
// number of cookies that were deleted or -1 if unknown.
func (d *DeleteCookiesCallback) OnComplete(num_deleted int32) {
	lookupDeleteCookiesCallbackProxy(d.Base()).OnComplete(d, num_deleted)
}

//export gocef_delete_cookies_callback_on_complete
func gocef_delete_cookies_callback_on_complete(self *C.cef_delete_cookies_callback_t, num_deleted C.int) {
	me__ := (*DeleteCookiesCallback)(self)
	proxy__ := lookupDeleteCookiesCallbackProxy(me__.Base())
	proxy__.OnComplete(me__, int32(num_deleted))
}
