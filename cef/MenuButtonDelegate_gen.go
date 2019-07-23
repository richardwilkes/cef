// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "MenuButtonDelegate_gen.h"
	"C"
)

// MenuButtonDelegateProxy defines methods required for using MenuButtonDelegate.
type MenuButtonDelegateProxy interface {
	OnMenuButtonPressed(self *MenuButtonDelegate, menu_button *MenuButton, screen_point *Point, button_pressed_lock *MenuButtonPressedLock)
}

// MenuButtonDelegate (cef_menu_button_delegate_t from include/capi/views/cef_menu_button_delegate_capi.h)
// Implement this structure to handle MenuButton events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type MenuButtonDelegate C.cef_menu_button_delegate_t

// NewMenuButtonDelegate creates a new MenuButtonDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewMenuButtonDelegate(proxy MenuButtonDelegateProxy) *MenuButtonDelegate {
	result := (*MenuButtonDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_menu_button_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_menu_button_delegate_proxy(result.toNative())
	}
	return result
}

func (d *MenuButtonDelegate) toNative() *C.cef_menu_button_delegate_t {
	return (*C.cef_menu_button_delegate_t)(d)
}

func lookupMenuButtonDelegateProxy(obj *BaseRefCounted) MenuButtonDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(MenuButtonDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type MenuButtonDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *MenuButtonDelegate) Base() *ButtonDelegate {
	return (*ButtonDelegate)(&d.base)
}

// OnMenuButtonPressed (on_menu_button_pressed)
// Called when |button| is pressed. Call cef_menu_button_t::show_menu() to
// show a popup menu at |screen_point|. When showing a custom popup such as a
// window keep a reference to |button_pressed_lock| until the popup is hidden
// to maintain the pressed button state.
func (d *MenuButtonDelegate) OnMenuButtonPressed(menu_button *MenuButton, screen_point *Point, button_pressed_lock *MenuButtonPressedLock) {
	lookupMenuButtonDelegateProxy(d.Base().Base().Base()).OnMenuButtonPressed(d, menu_button, screen_point, button_pressed_lock)
}

//nolint:gocritic
//export gocef_menu_button_delegate_on_menu_button_pressed
func gocef_menu_button_delegate_on_menu_button_pressed(self *C.cef_menu_button_delegate_t, menu_button *C.cef_menu_button_t, screen_point *C.cef_point_t, button_pressed_lock *C.cef_menu_button_pressed_lock_t) {
	me__ := (*MenuButtonDelegate)(self)
	proxy__ := lookupMenuButtonDelegateProxy(me__.Base().Base().Base())
	screen_point_ := screen_point.toGo()
	button_pressed_lock_ := button_pressed_lock.toGo()
	proxy__.OnMenuButtonPressed(me__, (*MenuButton)(menu_button), screen_point_, button_pressed_lock_)
}
