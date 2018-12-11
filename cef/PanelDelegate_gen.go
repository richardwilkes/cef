// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
	"unsafe"
)

// PanelDelegate (cef_panel_delegate_t from include/capi/views/cef_panel_delegate_capi.h)
// Implement this structure to handle Panel events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type PanelDelegate struct {
	// Base (base)
	// Base structure.
	Base *ViewDelegate
}

// NewPanelDelegate creates a new PanelDelegate.
func NewPanelDelegate() *PanelDelegate {
	return &PanelDelegate{}
}

func (d *PanelDelegate) toNative(native *C.cef_panel_delegate_t) *C.cef_panel_delegate_t {
	native.base = *(*C.cef_view_delegate_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_panel_delegate_t) toGo() *PanelDelegate {
	var d PanelDelegate
	n.intoGo(&d)
	return &d
}

func (n *C.cef_panel_delegate_t) intoGo(d *PanelDelegate) {
	d.Base = (*ViewDelegate)(&n.base)
}
