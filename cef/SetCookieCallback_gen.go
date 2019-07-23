// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "SetCookieCallback_gen.h"
	"C"
)

// SetCookieCallbackProxy defines methods required for using SetCookieCallback.
type SetCookieCallbackProxy interface {
	OnComplete(self *SetCookieCallback, success int32)
}

// SetCookieCallback (cef_set_cookie_callback_t from include/capi/cef_cookie_capi.h)
// Structure to implement to be notified of asynchronous completion via
// cef_cookie_manager_t::set_cookie().
type SetCookieCallback C.cef_set_cookie_callback_t

// NewSetCookieCallback creates a new SetCookieCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewSetCookieCallback(proxy SetCookieCallbackProxy) *SetCookieCallback {
	result := (*SetCookieCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_set_cookie_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_set_cookie_callback_proxy(result.toNative())
	}
	return result
}

func (d *SetCookieCallback) toNative() *C.cef_set_cookie_callback_t {
	return (*C.cef_set_cookie_callback_t)(d)
}

func lookupSetCookieCallbackProxy(obj *BaseRefCounted) SetCookieCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(SetCookieCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type SetCookieCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *SetCookieCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnComplete (on_complete)
// Method that will be called upon completion. |success| will be true (1) if
// the cookie was set successfully.
func (d *SetCookieCallback) OnComplete(success int32) {
	lookupSetCookieCallbackProxy(d.Base()).OnComplete(d, success)
}

//nolint:gocritic
//export gocef_set_cookie_callback_on_complete
func gocef_set_cookie_callback_on_complete(self *C.cef_set_cookie_callback_t, success C.int) {
	me__ := (*SetCookieCallback)(self)
	proxy__ := lookupSetCookieCallbackProxy(me__.Base())
	proxy__.OnComplete(me__, int32(success))
}
