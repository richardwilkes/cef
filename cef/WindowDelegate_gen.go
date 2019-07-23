// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "WindowDelegate_gen.h"
	"C"
)

// WindowDelegateProxy defines methods required for using WindowDelegate.
type WindowDelegateProxy interface {
	OnWindowCreated(self *WindowDelegate, window *Window)
	OnWindowDestroyed(self *WindowDelegate, window *Window)
	GetParentWindow(self *WindowDelegate, window *Window, is_menu, can_activate_menu *int32) *Window
	IsFrameless(self *WindowDelegate, window *Window) int32
	CanResize(self *WindowDelegate, window *Window) int32
	CanMaximize(self *WindowDelegate, window *Window) int32
	CanMinimize(self *WindowDelegate, window *Window) int32
	CanClose(self *WindowDelegate, window *Window) int32
	OnAccelerator(self *WindowDelegate, window *Window, command_id int32) int32
	OnKeyEvent(self *WindowDelegate, window *Window, event *KeyEvent) int32
}

// WindowDelegate (cef_window_delegate_t from include/capi/views/cef_window_delegate_capi.h)
// Implement this structure to handle window events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type WindowDelegate C.cef_window_delegate_t

// NewWindowDelegate creates a new WindowDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewWindowDelegate(proxy WindowDelegateProxy) *WindowDelegate {
	result := (*WindowDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_window_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_window_delegate_proxy(result.toNative())
	}
	return result
}

func (d *WindowDelegate) toNative() *C.cef_window_delegate_t {
	return (*C.cef_window_delegate_t)(d)
}

func lookupWindowDelegateProxy(obj *BaseRefCounted) WindowDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(WindowDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type WindowDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *WindowDelegate) Base() *PanelDelegate {
	return (&d.base).toGo()
}

// OnWindowCreated (on_window_created)
// Called when |window| is created.
func (d *WindowDelegate) OnWindowCreated(window *Window) {
	lookupWindowDelegateProxy(d.Base().Base.Base()).OnWindowCreated(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_on_window_created
func gocef_window_delegate_on_window_created(self *C.cef_window_delegate_t, window *C.cef_window_t) {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	proxy__.OnWindowCreated(me__, (*Window)(window))
}

// OnWindowDestroyed (on_window_destroyed)
// Called when |window| is destroyed. Release all references to |window| and
// do not attempt to execute any functions on |window| after this callback
// returns.
func (d *WindowDelegate) OnWindowDestroyed(window *Window) {
	lookupWindowDelegateProxy(d.Base().Base.Base()).OnWindowDestroyed(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_on_window_destroyed
func gocef_window_delegate_on_window_destroyed(self *C.cef_window_delegate_t, window *C.cef_window_t) {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	proxy__.OnWindowDestroyed(me__, (*Window)(window))
}

// GetParentWindow (get_parent_window)
// Return the parent for |window| or NULL if the |window| does not have a
// parent. Windows with parents will not get a taskbar button. Set |is_menu|
// to true (1) if |window| will be displayed as a menu, in which case it will
// not be clipped to the parent window bounds. Set |can_activate_menu| to
// false (0) if |is_menu| is true (1) and |window| should not be activated
// (given keyboard focus) when displayed.
func (d *WindowDelegate) GetParentWindow(window *Window, is_menu, can_activate_menu *int32) *Window {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).GetParentWindow(d, window, is_menu, can_activate_menu)
}

//nolint:gocritic
//export gocef_window_delegate_get_parent_window
func gocef_window_delegate_get_parent_window(self *C.cef_window_delegate_t, window *C.cef_window_t, is_menu *C.int, can_activate_menu *C.int) *C.cef_window_t {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return (proxy__.GetParentWindow(me__, (*Window)(window), (*int32)(is_menu), (*int32)(can_activate_menu))).toNative()
}

// IsFrameless (is_frameless)
// Return true (1) if |window| should be created without a frame or title bar.
// The window will be resizable if can_resize() returns true (1). Use
// cef_window_t::set_draggable_regions() to specify draggable regions.
func (d *WindowDelegate) IsFrameless(window *Window) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).IsFrameless(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_is_frameless
func gocef_window_delegate_is_frameless(self *C.cef_window_delegate_t, window *C.cef_window_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.IsFrameless(me__, (*Window)(window)))
}

// CanResize (can_resize)
// Return true (1) if |window| can be resized.
func (d *WindowDelegate) CanResize(window *Window) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).CanResize(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_can_resize
func gocef_window_delegate_can_resize(self *C.cef_window_delegate_t, window *C.cef_window_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.CanResize(me__, (*Window)(window)))
}

// CanMaximize (can_maximize)
// Return true (1) if |window| can be maximized.
func (d *WindowDelegate) CanMaximize(window *Window) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).CanMaximize(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_can_maximize
func gocef_window_delegate_can_maximize(self *C.cef_window_delegate_t, window *C.cef_window_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.CanMaximize(me__, (*Window)(window)))
}

// CanMinimize (can_minimize)
// Return true (1) if |window| can be minimized.
func (d *WindowDelegate) CanMinimize(window *Window) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).CanMinimize(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_can_minimize
func gocef_window_delegate_can_minimize(self *C.cef_window_delegate_t, window *C.cef_window_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.CanMinimize(me__, (*Window)(window)))
}

// CanClose (can_close)
// Return true (1) if |window| can be closed. This will be called for user-
// initiated window close actions and when cef_window_t::close() is called.
func (d *WindowDelegate) CanClose(window *Window) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).CanClose(d, window)
}

//nolint:gocritic
//export gocef_window_delegate_can_close
func gocef_window_delegate_can_close(self *C.cef_window_delegate_t, window *C.cef_window_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.CanClose(me__, (*Window)(window)))
}

// OnAccelerator (on_accelerator)
// Called when a keyboard accelerator registered with
// cef_window_t::SetAccelerator is triggered. Return true (1) if the
// accelerator was handled or false (0) otherwise.
func (d *WindowDelegate) OnAccelerator(window *Window, command_id int32) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).OnAccelerator(d, window, command_id)
}

//nolint:gocritic
//export gocef_window_delegate_on_accelerator
func gocef_window_delegate_on_accelerator(self *C.cef_window_delegate_t, window *C.cef_window_t, command_id C.int) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	return C.int(proxy__.OnAccelerator(me__, (*Window)(window), int32(command_id)))
}

// OnKeyEvent (on_key_event)
// Called after all other controls in the window have had a chance to handle
// the event. |event| contains information about the keyboard event. Return
// true (1) if the keyboard event was handled or false (0) otherwise.
func (d *WindowDelegate) OnKeyEvent(window *Window, event *KeyEvent) int32 {
	return lookupWindowDelegateProxy(d.Base().Base.Base()).OnKeyEvent(d, window, event)
}

//nolint:gocritic
//export gocef_window_delegate_on_key_event
func gocef_window_delegate_on_key_event(self *C.cef_window_delegate_t, window *C.cef_window_t, event *C.cef_key_event_t) C.int {
	me__ := (*WindowDelegate)(self)
	proxy__ := lookupWindowDelegateProxy(me__.Base().Base.Base())
	event_ := event.toGo()
	return C.int(proxy__.OnKeyEvent(me__, (*Window)(window), event_))
}
