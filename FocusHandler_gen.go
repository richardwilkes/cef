package cef

import (
	// #include "FocusHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// FocusHandlerProxy defines methods required for using FocusHandler.
type FocusHandlerProxy interface {
	OnTakeFocus(self *FocusHandler, browser *Browser, next int32)
	OnSetFocus(self *FocusHandler, browser *Browser, source FocusSource) int32
	OnGotFocus(self *FocusHandler, browser *Browser)
}

// FocusHandler (cef_focus_handler_t from include/capi/cef_focus_handler_capi.h)
// Implement this structure to handle events related to focus. The functions of
// this structure will be called on the UI thread.
type FocusHandler C.cef_focus_handler_t

// NewFocusHandler creates a new FocusHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewFocusHandler(proxy FocusHandlerProxy) *FocusHandler {
	result := (*FocusHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_focus_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_focus_handler_proxy(result.toNative())
	}
	return result
}

func (d *FocusHandler) toNative() *C.cef_focus_handler_t {
	return (*C.cef_focus_handler_t)(d)
}

func lookupFocusHandlerProxy(obj *BaseRefCounted) FocusHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(FocusHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type FocusHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *FocusHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnTakeFocus (on_take_focus)
// Called when the browser component is about to loose focus. For instance, if
// focus was on the last HTML element and the user pressed the TAB key. |next|
// will be true (1) if the browser is giving focus to the next component and
// false (0) if the browser is giving focus to the previous component.
func (d *FocusHandler) OnTakeFocus(browser *Browser, next int32) {
	lookupFocusHandlerProxy(d.Base()).OnTakeFocus(d, browser, next)
}

//export gocef_focus_handler_on_take_focus
func gocef_focus_handler_on_take_focus(self *C.cef_focus_handler_t, browser *C.cef_browser_t, next C.int) {
	me__ := (*FocusHandler)(self)
	proxy__ := lookupFocusHandlerProxy(me__.Base())
	proxy__.OnTakeFocus(me__, (*Browser)(browser), int32(next))
}

// OnSetFocus (on_set_focus)
// Called when the browser component is requesting focus. |source| indicates
// where the focus request is originating from. Return false (0) to allow the
// focus to be set or true (1) to cancel setting the focus.
func (d *FocusHandler) OnSetFocus(browser *Browser, source FocusSource) int32 {
	return lookupFocusHandlerProxy(d.Base()).OnSetFocus(d, browser, source)
}

//export gocef_focus_handler_on_set_focus
func gocef_focus_handler_on_set_focus(self *C.cef_focus_handler_t, browser *C.cef_browser_t, source C.cef_focus_source_t) C.int {
	me__ := (*FocusHandler)(self)
	proxy__ := lookupFocusHandlerProxy(me__.Base())
	return C.int(proxy__.OnSetFocus(me__, (*Browser)(browser), FocusSource(source)))
}

// OnGotFocus (on_got_focus)
// Called when the browser component has received focus.
func (d *FocusHandler) OnGotFocus(browser *Browser) {
	lookupFocusHandlerProxy(d.Base()).OnGotFocus(d, browser)
}

//export gocef_focus_handler_on_got_focus
func gocef_focus_handler_on_got_focus(self *C.cef_focus_handler_t, browser *C.cef_browser_t) {
	me__ := (*FocusHandler)(self)
	proxy__ := lookupFocusHandlerProxy(me__.Base())
	proxy__.OnGotFocus(me__, (*Browser)(browser))
}
