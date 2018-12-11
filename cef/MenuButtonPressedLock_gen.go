// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
	"unsafe"
)

// MenuButtonPressedLock (cef_menu_button_pressed_lock_t from include/capi/views/cef_menu_button_delegate_capi.h)
// MenuButton pressed lock is released when this object is destroyed.
type MenuButtonPressedLock struct {
	// Base (base)
	// Base structure.
	Base *BaseRefCounted
}

// NewMenuButtonPressedLock creates a new MenuButtonPressedLock.
func NewMenuButtonPressedLock() *MenuButtonPressedLock {
	return &MenuButtonPressedLock{}
}

func (d *MenuButtonPressedLock) toNative(native *C.cef_menu_button_pressed_lock_t) *C.cef_menu_button_pressed_lock_t {
	native.base = *(*C.cef_base_ref_counted_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_menu_button_pressed_lock_t) toGo() *MenuButtonPressedLock {
	var d MenuButtonPressedLock
	n.intoGo(&d)
	return &d
}

func (n *C.cef_menu_button_pressed_lock_t) intoGo(d *MenuButtonPressedLock) {
	d.Base = (*BaseRefCounted)(&n.base)
}
