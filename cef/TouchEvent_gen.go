// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// TouchEvent (cef_touch_event_t from include/internal/cef_types.h)
// Structure representing touch event information.
type TouchEvent struct {
	// ID (id)
	// Id of a touch point. Must be unique per touch, can be any number except -1.
	// Note that a maximum of 16 concurrent touches will be tracked; touches
	// beyond that will be ignored.
	ID int32
	// X (x)
	// X coordinate relative to the left side of the view.
	X float32
	// Y (y)
	// Y coordinate relative to the top side of the view.
	Y float32
	// RadiusX (radius_x)
	// X radius in pixels. Set to 0 if not applicable.
	RadiusX float32
	// RadiusY (radius_y)
	// Y radius in pixels. Set to 0 if not applicable.
	RadiusY float32
	// RotationAngle (rotation_angle)
	// Rotation angle in radians. Set to 0 if not applicable.
	RotationAngle float32
	// Pressure (pressure)
	// The normalized pressure of the pointer input in the range of [0,1].
	// Set to 0 if not applicable.
	Pressure float32
	// Type (_type)
	// The state of the touch point. Touches begin with one CEF_TET_PRESSED event
	// followed by zero or more CEF_TET_MOVED events and finally one
	// CEF_TET_RELEASED or CEF_TET_CANCELLED event. Events not respecting this
	// order will be ignored.
	Type TouchEventType
	// Modifiers (modifiers)
	// Bit flags describing any pressed modifier keys. See
	// cef_event_flags_t for values.
	Modifiers uint32
	// PointerType (pointer_type)
	// The device type that caused the event.
	PointerType PointerType
}

// NewTouchEvent creates a new TouchEvent.
func NewTouchEvent() *TouchEvent {
	return &TouchEvent{}
}

func (d *TouchEvent) toNative(native *C.cef_touch_event_t) *C.cef_touch_event_t {
	if d == nil {
		return nil
	}
	native.id = C.int(d.ID)
	native.x = C.float(d.X)
	native.y = C.float(d.Y)
	native.radius_x = C.float(d.RadiusX)
	native.radius_y = C.float(d.RadiusY)
	native.rotation_angle = C.float(d.RotationAngle)
	native.pressure = C.float(d.Pressure)
	native._type = C.cef_touch_event_type_t(d.Type)
	native.modifiers = C.uint32(d.Modifiers)
	native.pointer_type = C.cef_pointer_type_t(d.PointerType)
	return native
}

func (n *C.cef_touch_event_t) toGo() *TouchEvent {
	if n == nil {
		return nil
	}
	var d TouchEvent
	n.intoGo(&d)
	return &d
}

func (n *C.cef_touch_event_t) intoGo(d *TouchEvent) {
	d.ID = int32(n.id)
	d.X = float32(n.x)
	d.Y = float32(n.y)
	d.RadiusX = float32(n.radius_x)
	d.RadiusY = float32(n.radius_y)
	d.RotationAngle = float32(n.rotation_angle)
	d.Pressure = float32(n.pressure)
	d.Type = TouchEventType(n._type)
	d.Modifiers = uint32(n.modifiers)
	d.PointerType = PointerType(n.pointer_type)
}
