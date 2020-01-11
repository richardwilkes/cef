// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_label_button_t * gocef_button_as_label_button(cef_button_t * self, cef_label_button_t * (CEF_CALLBACK *callback__)(cef_button_t *)) { return callback__(self); }
	// void gocef_button_set_state(cef_button_t * self, cef_button_state_t state, void (CEF_CALLBACK *callback__)(cef_button_t *, cef_button_state_t)) { return callback__(self, state); }
	// cef_button_state_t gocef_button_get_state(cef_button_t * self, cef_button_state_t (CEF_CALLBACK *callback__)(cef_button_t *)) { return callback__(self); }
	// void gocef_button_set_ink_drop_enabled(cef_button_t * self, int enabled, void (CEF_CALLBACK *callback__)(cef_button_t *, int)) { return callback__(self, enabled); }
	// void gocef_button_set_tooltip_text(cef_button_t * self, cef_string_t * tooltipText, void (CEF_CALLBACK *callback__)(cef_button_t *, cef_string_t *)) { return callback__(self, tooltipText); }
	// void gocef_button_set_accessible_name(cef_button_t * self, cef_string_t * name, void (CEF_CALLBACK *callback__)(cef_button_t *, cef_string_t *)) { return callback__(self, name); }
	"C"
)

// Button (cef_button_t from include/capi/views/cef_button_capi.h)
// A View representing a button. Depending on the specific type, the button
// could be implemented by a native control or custom rendered. Methods must be
// called on the browser process UI thread unless otherwise indicated.
type Button C.cef_button_t

func (d *Button) toNative() *C.cef_button_t {
	return (*C.cef_button_t)(d)
}

// Base (base)
// Base structure.
func (d *Button) Base() *View {
	return (*View)(&d.base)
}

// AsLabelButton (as_label_button)
// Returns this Button as a LabelButton or NULL if this is not a LabelButton.
func (d *Button) AsLabelButton() *LabelButton {
	return (*LabelButton)(C.gocef_button_as_label_button(d.toNative(), d.as_label_button))
}

// SetState (set_state)
// Sets the current display state of the Button.
func (d *Button) SetState(state ButtonState) {
	C.gocef_button_set_state(d.toNative(), C.cef_button_state_t(state), d.set_state)
}

// GetState (get_state)
// Returns the current display state of the Button.
func (d *Button) GetState() ButtonState {
	return ButtonState(C.gocef_button_get_state(d.toNative(), d.get_state))
}

// SetInkDropEnabled (set_ink_drop_enabled)
// Sets the Button will use an ink drop effect for displaying state changes.
func (d *Button) SetInkDropEnabled(enabled int32) {
	C.gocef_button_set_ink_drop_enabled(d.toNative(), C.int(enabled), d.set_ink_drop_enabled)
}

// SetTooltipText (set_tooltip_text)
// Sets the tooltip text that will be displayed when the user hovers the mouse
// cursor over the Button.
func (d *Button) SetTooltipText(tooltipText string) {
	tooltipText_ := C.cef_string_userfree_alloc()
	setCEFStr(tooltipText, tooltipText_)
	defer func() {
		C.cef_string_userfree_free(tooltipText_)
	}()
	C.gocef_button_set_tooltip_text(d.toNative(), (*C.cef_string_t)(tooltipText_), d.set_tooltip_text)
}

// SetAccessibleName (set_accessible_name)
// Sets the accessible name that will be exposed to assistive technology (AT).
func (d *Button) SetAccessibleName(name string) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	C.gocef_button_set_accessible_name(d.toNative(), (*C.cef_string_t)(name_), d.set_accessible_name)
}
