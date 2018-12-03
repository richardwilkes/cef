// Code generated - DO NOT EDIT.

package cef

import (
	// #include "AccessibilityHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// AccessibilityHandlerProxy defines methods required for using AccessibilityHandler.
type AccessibilityHandlerProxy interface {
	OnAccessibilityTreeChange(self *AccessibilityHandler, value *Value)
	OnAccessibilityLocationChange(self *AccessibilityHandler, value *Value)
}

// AccessibilityHandler (cef_accessibility_handler_t from include/capi/cef_accessibility_handler_capi.h)
// Implement this structure to receive accessibility notification when
// accessibility events have been registered. The functions of this structure
// will be called on the UI thread.
type AccessibilityHandler C.cef_accessibility_handler_t

// NewAccessibilityHandler creates a new AccessibilityHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewAccessibilityHandler(proxy AccessibilityHandlerProxy) *AccessibilityHandler {
	result := (*AccessibilityHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_accessibility_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_accessibility_handler_proxy(result.toNative())
	}
	return result
}

func (d *AccessibilityHandler) toNative() *C.cef_accessibility_handler_t {
	return (*C.cef_accessibility_handler_t)(d)
}

func lookupAccessibilityHandlerProxy(obj *BaseRefCounted) AccessibilityHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(AccessibilityHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type AccessibilityHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *AccessibilityHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnAccessibilityTreeChange (on_accessibility_tree_change)
// Called after renderer process sends accessibility tree changes to the
// browser process.
func (d *AccessibilityHandler) OnAccessibilityTreeChange(value *Value) {
	lookupAccessibilityHandlerProxy(d.Base()).OnAccessibilityTreeChange(d, value)
}

//export gocef_accessibility_handler_on_accessibility_tree_change
func gocef_accessibility_handler_on_accessibility_tree_change(self *C.cef_accessibility_handler_t, value *C.cef_value_t) {
	me__ := (*AccessibilityHandler)(self)
	proxy__ := lookupAccessibilityHandlerProxy(me__.Base())
	proxy__.OnAccessibilityTreeChange(me__, (*Value)(value))
}

// OnAccessibilityLocationChange (on_accessibility_location_change)
// Called after renderer process sends accessibility location changes to the
// browser process.
func (d *AccessibilityHandler) OnAccessibilityLocationChange(value *Value) {
	lookupAccessibilityHandlerProxy(d.Base()).OnAccessibilityLocationChange(d, value)
}

//export gocef_accessibility_handler_on_accessibility_location_change
func gocef_accessibility_handler_on_accessibility_location_change(self *C.cef_accessibility_handler_t, value *C.cef_value_t) {
	me__ := (*AccessibilityHandler)(self)
	proxy__ := lookupAccessibilityHandlerProxy(me__.Base())
	proxy__.OnAccessibilityLocationChange(me__, (*Value)(value))
}
