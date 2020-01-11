// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "KeyboardHandler_gen.h"
	"C"
)

// KeyboardHandlerProxy defines methods required for using KeyboardHandler.
type KeyboardHandlerProxy interface {
	OnPreKeyEvent(self *KeyboardHandler, browser *Browser, event *KeyEvent, osEvent unsafe.Pointer, isKeyboardShortcut *int32) int32
	OnKeyEvent(self *KeyboardHandler, browser *Browser, event *KeyEvent, osEvent unsafe.Pointer) int32
}

// KeyboardHandler (cef_keyboard_handler_t from include/capi/cef_keyboard_handler_capi.h)
// Implement this structure to handle events related to keyboard input. The
// functions of this structure will be called on the UI thread.
type KeyboardHandler C.cef_keyboard_handler_t

// NewKeyboardHandler creates a new KeyboardHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewKeyboardHandler(proxy KeyboardHandlerProxy) *KeyboardHandler {
	result := (*KeyboardHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_keyboard_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_keyboard_handler_proxy(result.toNative())
	}
	return result
}

func (d *KeyboardHandler) toNative() *C.cef_keyboard_handler_t {
	return (*C.cef_keyboard_handler_t)(d)
}

func lookupKeyboardHandlerProxy(obj *BaseRefCounted) KeyboardHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(KeyboardHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type KeyboardHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *KeyboardHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnPreKeyEvent (on_pre_key_event)
// Called before a keyboard event is sent to the renderer. |event| contains
// information about the keyboard event. |os_event| is the operating system
// event message, if any. Return true (1) if the event was handled or false
// (0) otherwise. If the event will be handled in on_key_event() as a keyboard
// shortcut set |is_keyboard_shortcut| to true (1) and return false (0).
func (d *KeyboardHandler) OnPreKeyEvent(browser *Browser, event *KeyEvent, osEvent unsafe.Pointer, isKeyboardShortcut *int32) int32 {
	return lookupKeyboardHandlerProxy(d.Base()).OnPreKeyEvent(d, browser, event, osEvent, isKeyboardShortcut)
}

//nolint:gocritic
//export gocef_keyboard_handler_on_pre_key_event
func gocef_keyboard_handler_on_pre_key_event(self *C.cef_keyboard_handler_t, browser *C.cef_browser_t, event *C.cef_key_event_t, osEvent unsafe.Pointer, isKeyboardShortcut *C.int) C.int {
	me__ := (*KeyboardHandler)(self)
	proxy__ := lookupKeyboardHandlerProxy(me__.Base())
	event_ := event.toGo()
	return C.int(proxy__.OnPreKeyEvent(me__, (*Browser)(browser), event_, osEvent, (*int32)(isKeyboardShortcut)))
}

// OnKeyEvent (on_key_event)
// Called after the renderer and JavaScript in the page has had a chance to
// handle the event. |event| contains information about the keyboard event.
// |os_event| is the operating system event message, if any. Return true (1)
// if the keyboard event was handled or false (0) otherwise.
func (d *KeyboardHandler) OnKeyEvent(browser *Browser, event *KeyEvent, osEvent unsafe.Pointer) int32 {
	return lookupKeyboardHandlerProxy(d.Base()).OnKeyEvent(d, browser, event, osEvent)
}

//nolint:gocritic
//export gocef_keyboard_handler_on_key_event
func gocef_keyboard_handler_on_key_event(self *C.cef_keyboard_handler_t, browser *C.cef_browser_t, event *C.cef_key_event_t, osEvent unsafe.Pointer) C.int {
	me__ := (*KeyboardHandler)(self)
	proxy__ := lookupKeyboardHandlerProxy(me__.Base())
	event_ := event.toGo()
	return C.int(proxy__.OnKeyEvent(me__, (*Browser)(browser), event_, osEvent))
}
