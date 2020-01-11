// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "RunContextMenuCallback_gen.h"
	"C"
)

// RunContextMenuCallbackProxy defines methods required for using RunContextMenuCallback.
type RunContextMenuCallbackProxy interface {
	Cont(self *RunContextMenuCallback, commandID int32, eventFlags EventFlags)
	Cancel(self *RunContextMenuCallback)
}

// RunContextMenuCallback (cef_run_context_menu_callback_t from include/capi/cef_context_menu_handler_capi.h)
// Callback structure used for continuation of custom context menu display.
type RunContextMenuCallback C.cef_run_context_menu_callback_t

// NewRunContextMenuCallback creates a new RunContextMenuCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewRunContextMenuCallback(proxy RunContextMenuCallbackProxy) *RunContextMenuCallback {
	result := (*RunContextMenuCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_run_context_menu_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_run_context_menu_callback_proxy(result.toNative())
	}
	return result
}

func (d *RunContextMenuCallback) toNative() *C.cef_run_context_menu_callback_t {
	return (*C.cef_run_context_menu_callback_t)(d)
}

func lookupRunContextMenuCallbackProxy(obj *BaseRefCounted) RunContextMenuCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(RunContextMenuCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type RunContextMenuCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *RunContextMenuCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Cont (cont)
// Complete context menu display by selecting the specified |command_id| and
// |event_flags|.
func (d *RunContextMenuCallback) Cont(commandID int32, eventFlags EventFlags) {
	lookupRunContextMenuCallbackProxy(d.Base()).Cont(d, commandID, eventFlags)
}

//nolint:gocritic
//export gocef_run_context_menu_callback_cont
func gocef_run_context_menu_callback_cont(self *C.cef_run_context_menu_callback_t, commandID C.int, eventFlags C.cef_event_flags_t) {
	me__ := (*RunContextMenuCallback)(self)
	proxy__ := lookupRunContextMenuCallbackProxy(me__.Base())
	proxy__.Cont(me__, int32(commandID), EventFlags(eventFlags))
}

// Cancel (cancel)
// Cancel context menu display.
func (d *RunContextMenuCallback) Cancel() {
	lookupRunContextMenuCallbackProxy(d.Base()).Cancel(d)
}

//nolint:gocritic
//export gocef_run_context_menu_callback_cancel
func gocef_run_context_menu_callback_cancel(self *C.cef_run_context_menu_callback_t) {
	me__ := (*RunContextMenuCallback)(self)
	proxy__ := lookupRunContextMenuCallbackProxy(me__.Base())
	proxy__.Cancel(me__)
}
