// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "JsdialogCallback_gen.h"
	"C"
)

// JsdialogCallbackProxy defines methods required for using JsdialogCallback.
type JsdialogCallbackProxy interface {
	Cont(self *JsdialogCallback, success int32, user_input string)
}

// JsdialogCallback (cef_jsdialog_callback_t from include/capi/cef_jsdialog_handler_capi.h)
// Callback structure used for asynchronous continuation of JavaScript dialog
// requests.
type JsdialogCallback C.cef_jsdialog_callback_t

// NewJsdialogCallback creates a new JsdialogCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewJsdialogCallback(proxy JsdialogCallbackProxy) *JsdialogCallback {
	result := (*JsdialogCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_jsdialog_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_jsdialog_callback_proxy(result.toNative())
	}
	return result
}

func (d *JsdialogCallback) toNative() *C.cef_jsdialog_callback_t {
	return (*C.cef_jsdialog_callback_t)(d)
}

func lookupJsdialogCallbackProxy(obj *BaseRefCounted) JsdialogCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(JsdialogCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type JsdialogCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *JsdialogCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Continue the JS dialog request. Set |success| to true (1) if the OK button
// was pressed. The |user_input| value should be specified for prompt dialogs.
func (d *JsdialogCallback) Cont(success int32, user_input string) {
	lookupJsdialogCallbackProxy(d.Base()).Cont(d, success, user_input)
}

//nolint:gocritic
//export gocef_jsdialog_callback_cont
func gocef_jsdialog_callback_cont(self *C.cef_jsdialog_callback_t, success C.int, user_input *C.cef_string_t) {
	me__ := (*JsdialogCallback)(self)
	proxy__ := lookupJsdialogCallbackProxy(me__.Base())
	user_input_ := cefstrToString(user_input)
	proxy__.Cont(me__, int32(success), user_input_)
}
