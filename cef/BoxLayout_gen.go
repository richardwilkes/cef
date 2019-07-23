// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_box_layout_set_flex_for_view(cef_box_layout_t * self, cef_view_t * view, int flex, void (CEF_CALLBACK *callback__)(cef_box_layout_t *, cef_view_t *, int)) { return callback__(self, view, flex); }
	// void gocef_box_layout_clear_flex_for_view(cef_box_layout_t * self, cef_view_t * view, void (CEF_CALLBACK *callback__)(cef_box_layout_t *, cef_view_t *)) { return callback__(self, view); }
	"C"
)

// BoxLayout (cef_box_layout_t from include/capi/views/cef_box_layout_capi.h)
// A Layout manager that arranges child views vertically or horizontally in a
// side-by-side fashion with spacing around and between the child views. The
// child views are always sized according to their preferred size. If the host's
// bounds provide insufficient space, child views will be clamped. Excess space
// will not be distributed. Methods must be called on the browser process UI
// thread unless otherwise indicated.
type BoxLayout C.cef_box_layout_t

func (d *BoxLayout) toNative() *C.cef_box_layout_t {
	return (*C.cef_box_layout_t)(d)
}

// Base (base)
// Base structure.
func (d *BoxLayout) Base() *Layout {
	return (*Layout)(&d.base)
}

// SetFlexForView (set_flex_for_view)
// Set the flex weight for the given |view|. Using the preferred size as the
// basis, free space along the main axis is distributed to views in the ratio
// of their flex weights. Similarly, if the views will overflow the parent,
// space is subtracted in these ratios. A flex of 0 means this view is not
// resized. Flex values must not be negative.
func (d *BoxLayout) SetFlexForView(view *View, flex int32) {
	C.gocef_box_layout_set_flex_for_view(d.toNative(), view.toNative(), C.int(flex), d.set_flex_for_view)
}

// ClearFlexForView (clear_flex_for_view)
// Clears the flex for the given |view|, causing it to use the default flex
// specified via cef_box_layout_tSettings.default_flex.
func (d *BoxLayout) ClearFlexForView(view *View) {
	C.gocef_box_layout_clear_flex_for_view(d.toNative(), view.toNative(), d.clear_flex_for_view)
}
