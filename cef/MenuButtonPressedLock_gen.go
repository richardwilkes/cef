// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"
)

import (
	// #include "capi_gen.h"
	"C"
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
	if d == nil {
		return nil
	}
	native.base = *(*C.cef_base_ref_counted_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_menu_button_pressed_lock_t) toGo() *MenuButtonPressedLock {
	if n == nil {
		return nil
	}
	var d MenuButtonPressedLock
	n.intoGo(&d)
	return &d
}

func (n *C.cef_menu_button_pressed_lock_t) intoGo(d *MenuButtonPressedLock) {
	d.Base = (*BaseRefCounted)(&n.base)
}
