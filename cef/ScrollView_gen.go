// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_scroll_view_set_content_view(cef_scroll_view_t * self, cef_view_t * view, void (CEF_CALLBACK *callback__)(cef_scroll_view_t *, cef_view_t *)) { return callback__(self, view); }
	// cef_view_t * gocef_scroll_view_get_content_view(cef_scroll_view_t * self, cef_view_t * (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	// cef_rect_t gocef_scroll_view_get_visible_content_rect(cef_scroll_view_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	// int gocef_scroll_view_has_horizontal_scrollbar(cef_scroll_view_t * self, int (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	// int gocef_scroll_view_get_horizontal_scrollbar_height(cef_scroll_view_t * self, int (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	// int gocef_scroll_view_has_vertical_scrollbar(cef_scroll_view_t * self, int (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	// int gocef_scroll_view_get_vertical_scrollbar_width(cef_scroll_view_t * self, int (CEF_CALLBACK *callback__)(cef_scroll_view_t *)) { return callback__(self); }
	"C"
)

// ScrollView (cef_scroll_view_t from include/capi/views/cef_scroll_view_capi.h)
// A ScrollView will show horizontal and/or vertical scrollbars when necessary
// based on the size of the attached content view. Methods must be called on the
// browser process UI thread unless otherwise indicated.
type ScrollView C.cef_scroll_view_t

func (d *ScrollView) toNative() *C.cef_scroll_view_t {
	return (*C.cef_scroll_view_t)(d)
}

// Base (base)
// Base structure.
func (d *ScrollView) Base() *View {
	return (*View)(&d.base)
}

// SetContentView (set_content_view)
// Set the content View. The content View must have a specified size (e.g. via
// cef_view_t::SetBounds or cef_view_tDelegate::GetPreferredSize).
func (d *ScrollView) SetContentView(view *View) {
	C.gocef_scroll_view_set_content_view(d.toNative(), view.toNative(), d.set_content_view)
}

// GetContentView (get_content_view)
// Returns the content View.
func (d *ScrollView) GetContentView() *View {
	return (*View)(C.gocef_scroll_view_get_content_view(d.toNative(), d.get_content_view))
}

// GetVisibleContentRect (get_visible_content_rect)
// Returns the visible region of the content View.
func (d *ScrollView) GetVisibleContentRect() Rect {
	cresult__ := C.gocef_scroll_view_get_visible_content_rect(d.toNative(), d.get_visible_content_rect)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// HasHorizontalScrollbar (has_horizontal_scrollbar)
// Returns true (1) if the horizontal scrollbar is currently showing.
func (d *ScrollView) HasHorizontalScrollbar() int32 {
	return int32(C.gocef_scroll_view_has_horizontal_scrollbar(d.toNative(), d.has_horizontal_scrollbar))
}

// GetHorizontalScrollbarHeight (get_horizontal_scrollbar_height)
// Returns the height of the horizontal scrollbar.
func (d *ScrollView) GetHorizontalScrollbarHeight() int32 {
	return int32(C.gocef_scroll_view_get_horizontal_scrollbar_height(d.toNative(), d.get_horizontal_scrollbar_height))
}

// HasVerticalScrollbar (has_vertical_scrollbar)
// Returns true (1) if the vertical scrollbar is currently showing.
func (d *ScrollView) HasVerticalScrollbar() int32 {
	return int32(C.gocef_scroll_view_has_vertical_scrollbar(d.toNative(), d.has_vertical_scrollbar))
}

// GetVerticalScrollbarWidth (get_vertical_scrollbar_width)
// Returns the width of the vertical scrollbar.
func (d *ScrollView) GetVerticalScrollbarWidth() int32 {
	return int32(C.gocef_scroll_view_get_vertical_scrollbar_width(d.toNative(), d.get_vertical_scrollbar_width))
}
