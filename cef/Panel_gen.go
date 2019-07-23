// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_window_t * gocef_panel_as_window(cef_panel_t * self, cef_window_t * (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// cef_fill_layout_t * gocef_panel_set_to_fill_layout(cef_panel_t * self, cef_fill_layout_t * (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// cef_box_layout_t * gocef_panel_set_to_box_layout(cef_panel_t * self, cef_box_layout_settings_t * settings, cef_box_layout_t * (CEF_CALLBACK *callback__)(cef_panel_t *, cef_box_layout_settings_t *)) { return callback__(self, settings); }
	// cef_layout_t * gocef_panel_get_layout(cef_panel_t * self, cef_layout_t * (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// void gocef_panel_layout(cef_panel_t * self, void (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// void gocef_panel_add_child_view(cef_panel_t * self, cef_view_t * view, void (CEF_CALLBACK *callback__)(cef_panel_t *, cef_view_t *)) { return callback__(self, view); }
	// void gocef_panel_add_child_view_at(cef_panel_t * self, cef_view_t * view, int index, void (CEF_CALLBACK *callback__)(cef_panel_t *, cef_view_t *, int)) { return callback__(self, view, index); }
	// void gocef_panel_reorder_child_view(cef_panel_t * self, cef_view_t * view, int index, void (CEF_CALLBACK *callback__)(cef_panel_t *, cef_view_t *, int)) { return callback__(self, view, index); }
	// void gocef_panel_remove_child_view(cef_panel_t * self, cef_view_t * view, void (CEF_CALLBACK *callback__)(cef_panel_t *, cef_view_t *)) { return callback__(self, view); }
	// void gocef_panel_remove_all_child_views(cef_panel_t * self, void (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// size_t gocef_panel_get_child_view_count(cef_panel_t * self, size_t (CEF_CALLBACK *callback__)(cef_panel_t *)) { return callback__(self); }
	// cef_view_t * gocef_panel_get_child_view_at(cef_panel_t * self, int index, cef_view_t * (CEF_CALLBACK *callback__)(cef_panel_t *, int)) { return callback__(self, index); }
	"C"
)

// Panel (cef_panel_t from include/capi/views/cef_panel_capi.h)
// A Panel is a container in the views hierarchy that can contain other Views as
// children. Methods must be called on the browser process UI thread unless
// otherwise indicated.
type Panel C.cef_panel_t

func (d *Panel) toNative() *C.cef_panel_t {
	return (*C.cef_panel_t)(d)
}

// Base (base)
// Base structure.
func (d *Panel) Base() *View {
	return (*View)(&d.base)
}

// AsWindow (as_window)
// Returns this Panel as a Window or NULL if this is not a Window.
func (d *Panel) AsWindow() *Window {
	return (*Window)(C.gocef_panel_as_window(d.toNative(), d.as_window))
}

// SetToFillLayout (set_to_fill_layout)
// Set this Panel's Layout to FillLayout and return the FillLayout object.
func (d *Panel) SetToFillLayout() *FillLayout {
	return (C.gocef_panel_set_to_fill_layout(d.toNative(), d.set_to_fill_layout)).toGo()
}

// SetToBoxLayout (set_to_box_layout)
// Set this Panel's Layout to BoxLayout and return the BoxLayout object.
func (d *Panel) SetToBoxLayout(settings *BoxLayoutSettings) *BoxLayout {
	return (*BoxLayout)(C.gocef_panel_set_to_box_layout(d.toNative(), settings.toNative(&C.cef_box_layout_settings_t{}), d.set_to_box_layout))
}

// GetLayout (get_layout)
// Get the Layout.
func (d *Panel) GetLayout() *Layout {
	return (*Layout)(C.gocef_panel_get_layout(d.toNative(), d.get_layout))
}

// Layout (layout)
// Lay out the child Views (set their bounds based on sizing heuristics
// specific to the current Layout).
func (d *Panel) Layout() {
	C.gocef_panel_layout(d.toNative(), d.layout)
}

// AddChildView (add_child_view)
// Add a child View.
func (d *Panel) AddChildView(view *View) {
	C.gocef_panel_add_child_view(d.toNative(), view.toNative(), d.add_child_view)
}

// AddChildViewAt (add_child_view_at)
// Add a child View at the specified |index|. If |index| matches the result of
// GetChildCount() then the View will be added at the end.
func (d *Panel) AddChildViewAt(view *View, index int32) {
	C.gocef_panel_add_child_view_at(d.toNative(), view.toNative(), C.int(index), d.add_child_view_at)
}

// ReorderChildView (reorder_child_view)
// Move the child View to the specified |index|. A negative value for |index|
// will move the View to the end.
func (d *Panel) ReorderChildView(view *View, index int32) {
	C.gocef_panel_reorder_child_view(d.toNative(), view.toNative(), C.int(index), d.reorder_child_view)
}

// RemoveChildView (remove_child_view)
// Remove a child View. The View can then be added to another Panel.
func (d *Panel) RemoveChildView(view *View) {
	C.gocef_panel_remove_child_view(d.toNative(), view.toNative(), d.remove_child_view)
}

// RemoveAllChildViews (remove_all_child_views)
// Remove all child Views. The removed Views will be deleted if the client
// holds no references to them.
func (d *Panel) RemoveAllChildViews() {
	C.gocef_panel_remove_all_child_views(d.toNative(), d.remove_all_child_views)
}

// GetChildViewCount (get_child_view_count)
// Returns the number of child Views.
func (d *Panel) GetChildViewCount() uint64 {
	return uint64(C.gocef_panel_get_child_view_count(d.toNative(), d.get_child_view_count))
}

// GetChildViewAt (get_child_view_at)
// Returns the child View at the specified |index|.
func (d *Panel) GetChildViewAt(index int32) *View {
	return (*View)(C.gocef_panel_get_child_view_at(d.toNative(), C.int(index), d.get_child_view_at))
}
