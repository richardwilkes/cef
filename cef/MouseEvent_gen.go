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
	// #include "capi_gen.h"
	"C"
)

// MouseEvent (cef_mouse_event_t from include/internal/cef_types.h)
// Structure representing mouse event information.
type MouseEvent struct {
	// X (x)
	// X coordinate relative to the left side of the view.
	X int32
	// Y (y)
	// Y coordinate relative to the top side of the view.
	Y int32
	// Modifiers (modifiers)
	// Bit flags describing any pressed modifier keys. See
	// cef_event_flags_t for values.
	Modifiers uint32
}

// NewMouseEvent creates a new MouseEvent.
func NewMouseEvent() *MouseEvent {
	return &MouseEvent{}
}

func (d *MouseEvent) toNative(native *C.cef_mouse_event_t) *C.cef_mouse_event_t {
	if d == nil {
		return nil
	}
	native.x = C.int(d.X)
	native.y = C.int(d.Y)
	native.modifiers = C.uint32(d.Modifiers)
	return native
}

func (n *C.cef_mouse_event_t) toGo() *MouseEvent {
	if n == nil {
		return nil
	}
	var d MouseEvent
	n.intoGo(&d)
	return &d
}

func (n *C.cef_mouse_event_t) intoGo(d *MouseEvent) {
	d.X = int32(n.x)
	d.Y = int32(n.y)
	d.Modifiers = uint32(n.modifiers)
}
