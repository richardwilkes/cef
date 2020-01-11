// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"
)

import (
	// #include "capi_gen.h"
	"C"
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
	if d == nil {
		return nil
	}
	native.base = *(*C.cef_view_delegate_t)(unsafe.Pointer(d.Base))
	return native
}

func (n *C.cef_panel_delegate_t) toGo() *PanelDelegate {
	if n == nil {
		return nil
	}
	var d PanelDelegate
	n.intoGo(&d)
	return &d
}

func (n *C.cef_panel_delegate_t) intoGo(d *PanelDelegate) {
	d.Base = (*ViewDelegate)(&n.base)
}
