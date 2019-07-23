// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"
)

import (
	// #include "capi_gen.h"
	"C"
)

// FillLayout (cef_fill_layout_t from include/capi/views/cef_fill_layout_capi.h)
// A simple Layout that causes the associated Panel's one child to be sized to
// match the bounds of its parent. Methods must be called on the browser process
// UI thread unless otherwise indicated.
type FillLayout struct {
	// Base (base)
	// Base structure.
	Base *Layout
}

// NewFillLayout creates a new FillLayout.
func NewFillLayout() *FillLayout {
	return &FillLayout{}
}

func (d *FillLayout) toNative(native *C.cef_fill_layout_t) *C.cef_fill_layout_t {
	if d == nil {
		return nil
	}
	native.base = *(*C.cef_layout_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_fill_layout_t) toGo() *FillLayout {
	if n == nil {
		return nil
	}
	var d FillLayout
	n.intoGo(&d)
	return &d
}

func (n *C.cef_fill_layout_t) intoGo(d *FillLayout) {
	d.Base = (*Layout)(&n.base)
}
