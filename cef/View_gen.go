// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_browser_view_t * gocef_view_as_browser_view(cef_view_t * self, cef_browser_view_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_button_t * gocef_view_as_button(cef_view_t * self, cef_button_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_panel_t * gocef_view_as_panel(cef_view_t * self, cef_panel_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_scroll_view_t * gocef_view_as_scroll_view(cef_view_t * self, cef_scroll_view_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_textfield_t * gocef_view_as_textfield(cef_view_t * self, cef_textfield_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_view_get_type_string(cef_view_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_view_to_string(cef_view_t * self, int includeChildren, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, includeChildren); }
	// int gocef_view_is_valid(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_is_attached(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_is_same(cef_view_t * self, cef_view_t * that, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_view_t *)) { return callback__(self, that); }
	// cef_view_delegate_t * gocef_view_get_delegate(cef_view_t * self, cef_view_delegate_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_window_t * gocef_view_get_window(cef_view_t * self, cef_window_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_get_id(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_id(cef_view_t * self, int iD, void (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, iD); }
	// int gocef_view_get_group_id(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_group_id(cef_view_t * self, int groupID, void (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, groupID); }
	// cef_view_t * gocef_view_get_parent_view(cef_view_t * self, cef_view_t * (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_view_t * gocef_view_get_view_for_id(cef_view_t * self, int iD, cef_view_t * (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, iD); }
	// void gocef_view_set_bounds(cef_view_t * self, cef_rect_t * bounds, void (CEF_CALLBACK *callback__)(cef_view_t *, cef_rect_t *)) { return callback__(self, bounds); }
	// cef_rect_t gocef_view_get_bounds(cef_view_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_rect_t gocef_view_get_bounds_in_screen(cef_view_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_size(cef_view_t * self, cef_size_t * size, void (CEF_CALLBACK *callback__)(cef_view_t *, cef_size_t *)) { return callback__(self, size); }
	// cef_size_t gocef_view_get_size(cef_view_t * self, cef_size_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_position(cef_view_t * self, cef_point_t * position, void (CEF_CALLBACK *callback__)(cef_view_t *, cef_point_t *)) { return callback__(self, position); }
	// cef_point_t gocef_view_get_position(cef_view_t * self, cef_point_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_size_t gocef_view_get_preferred_size(cef_view_t * self, cef_size_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_size_to_preferred_size(cef_view_t * self, void (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_size_t gocef_view_get_minimum_size(cef_view_t * self, cef_size_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// cef_size_t gocef_view_get_maximum_size(cef_view_t * self, cef_size_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_get_height_for_width(cef_view_t * self, int width, int (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, width); }
	// void gocef_view_invalidate_layout(cef_view_t * self, void (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_visible(cef_view_t * self, int visible, void (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, visible); }
	// int gocef_view_is_visible(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_is_drawn(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_enabled(cef_view_t * self, int enabled, void (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, enabled); }
	// int gocef_view_is_enabled(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_focusable(cef_view_t * self, int focusable, void (CEF_CALLBACK *callback__)(cef_view_t *, int)) { return callback__(self, focusable); }
	// int gocef_view_is_focusable(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_is_accessibility_focusable(cef_view_t * self, int (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_request_focus(cef_view_t * self, void (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// void gocef_view_set_background_color(cef_view_t * self, cef_color_t color, void (CEF_CALLBACK *callback__)(cef_view_t *, cef_color_t)) { return callback__(self, color); }
	// cef_color_t gocef_view_get_background_color(cef_view_t * self, cef_color_t (CEF_CALLBACK *callback__)(cef_view_t *)) { return callback__(self); }
	// int gocef_view_convert_point_to_screen(cef_view_t * self, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_point_t *)) { return callback__(self, point); }
	// int gocef_view_convert_point_from_screen(cef_view_t * self, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_point_t *)) { return callback__(self, point); }
	// int gocef_view_convert_point_to_window(cef_view_t * self, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_point_t *)) { return callback__(self, point); }
	// int gocef_view_convert_point_from_window(cef_view_t * self, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_point_t *)) { return callback__(self, point); }
	// int gocef_view_convert_point_to_view(cef_view_t * self, cef_view_t * view, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_view_t *, cef_point_t *)) { return callback__(self, view, point); }
	// int gocef_view_convert_point_from_view(cef_view_t * self, cef_view_t * view, cef_point_t * point, int (CEF_CALLBACK *callback__)(cef_view_t *, cef_view_t *, cef_point_t *)) { return callback__(self, view, point); }
	"C"
)

// View (cef_view_t from include/capi/views/cef_view_capi.h)
// A View is a rectangle within the views View hierarchy. It is the base
// structure for all Views. All size and position values are in density
// independent pixels (DIP) unless otherwise indicated. Methods must be called
// on the browser process UI thread unless otherwise indicated.
type View C.cef_view_t

func (d *View) toNative() *C.cef_view_t {
	return (*C.cef_view_t)(d)
}

// Base (base)
// Base structure.
func (d *View) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// AsBrowserView (as_browser_view)
// Returns this View as a BrowserView or NULL if this is not a BrowserView.
func (d *View) AsBrowserView() *BrowserView {
	return (*BrowserView)(C.gocef_view_as_browser_view(d.toNative(), d.as_browser_view))
}

// AsButton (as_button)
// Returns this View as a Button or NULL if this is not a Button.
func (d *View) AsButton() *Button {
	return (*Button)(C.gocef_view_as_button(d.toNative(), d.as_button))
}

// AsPanel (as_panel)
// Returns this View as a Panel or NULL if this is not a Panel.
func (d *View) AsPanel() *Panel {
	return (*Panel)(C.gocef_view_as_panel(d.toNative(), d.as_panel))
}

// AsScrollView (as_scroll_view)
// Returns this View as a ScrollView or NULL if this is not a ScrollView.
func (d *View) AsScrollView() *ScrollView {
	return (*ScrollView)(C.gocef_view_as_scroll_view(d.toNative(), d.as_scroll_view))
}

// AsTextfield (as_textfield)
// Returns this View as a Textfield or NULL if this is not a Textfield.
func (d *View) AsTextfield() *Textfield {
	return (*Textfield)(C.gocef_view_as_textfield(d.toNative(), d.as_textfield))
}

// GetTypeString (get_type_string)
// Returns the type of this View as a string. Used primarily for testing
// purposes.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *View) GetTypeString() string {
	return cefuserfreestrToString(C.gocef_view_get_type_string(d.toNative(), d.get_type_string))
}

// ToString (to_string)
// Returns a string representation of this View which includes the type and
// various type-specific identifying attributes. If |include_children| is true
// (1) any child Views will also be included. Used primarily for testing
// purposes.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *View) ToString(includeChildren int32) string {
	return cefuserfreestrToString(C.gocef_view_to_string(d.toNative(), C.int(includeChildren), d.to_string))
}

// IsValid (is_valid)
// Returns true (1) if this View is valid.
func (d *View) IsValid() int32 {
	return int32(C.gocef_view_is_valid(d.toNative(), d.is_valid))
}

// IsAttached (is_attached)
// Returns true (1) if this View is currently attached to another View. A View
// can only be attached to one View at a time.
func (d *View) IsAttached() int32 {
	return int32(C.gocef_view_is_attached(d.toNative(), d.is_attached))
}

// IsSame (is_same)
// Returns true (1) if this View is the same as |that| View.
func (d *View) IsSame(that *View) int32 {
	return int32(C.gocef_view_is_same(d.toNative(), that.toNative(), d.is_same))
}

// GetDelegate (get_delegate)
// Returns the delegate associated with this View, if any.
func (d *View) GetDelegate() *ViewDelegate {
	return (*ViewDelegate)(C.gocef_view_get_delegate(d.toNative(), d.get_delegate))
}

// GetWindow (get_window)
// Returns the top-level Window hosting this View, if any.
func (d *View) GetWindow() *Window {
	return (*Window)(C.gocef_view_get_window(d.toNative(), d.get_window))
}

// GetID (get_id)
// Returns the ID for this View.
func (d *View) GetID() int32 {
	return int32(C.gocef_view_get_id(d.toNative(), d.get_id))
}

// SetID (set_id)
// Sets the ID for this View. ID should be unique within the subtree that you
// intend to search for it. 0 is the default ID for views.
func (d *View) SetID(iD int32) {
	C.gocef_view_set_id(d.toNative(), C.int(iD), d.set_id)
}

// GetGroupID (get_group_id)
// Returns the group id of this View, or -1 if not set.
func (d *View) GetGroupID() int32 {
	return int32(C.gocef_view_get_group_id(d.toNative(), d.get_group_id))
}

// SetGroupID (set_group_id)
// A group id is used to tag Views which are part of the same logical group.
// Focus can be moved between views with the same group using the arrow keys.
// The group id is immutable once it's set.
func (d *View) SetGroupID(groupID int32) {
	C.gocef_view_set_group_id(d.toNative(), C.int(groupID), d.set_group_id)
}

// GetParentView (get_parent_view)
// Returns the View that contains this View, if any.
func (d *View) GetParentView() *View {
	return (*View)(C.gocef_view_get_parent_view(d.toNative(), d.get_parent_view))
}

// GetViewForID (get_view_for_id)
// Recursively descends the view tree starting at this View, and returns the
// first child that it encounters with the given ID. Returns NULL if no
// matching child view is found.
func (d *View) GetViewForID(iD int32) *View {
	return (*View)(C.gocef_view_get_view_for_id(d.toNative(), C.int(iD), d.get_view_for_id))
}

// SetBounds (set_bounds)
// Sets the bounds (size and position) of this View. Position is in parent
// coordinates.
func (d *View) SetBounds(bounds *Rect) {
	C.gocef_view_set_bounds(d.toNative(), bounds.toNative(&C.cef_rect_t{}), d.set_bounds)
}

// GetBounds (get_bounds)
// Returns the bounds (size and position) of this View. Position is in parent
// coordinates.
func (d *View) GetBounds() Rect {
	cresult__ := C.gocef_view_get_bounds(d.toNative(), d.get_bounds)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// GetBoundsInScreen (get_bounds_in_screen)
// Returns the bounds (size and position) of this View. Position is in screen
// coordinates.
func (d *View) GetBoundsInScreen() Rect {
	cresult__ := C.gocef_view_get_bounds_in_screen(d.toNative(), d.get_bounds_in_screen)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// SetSize (set_size)
// Sets the size of this View without changing the position.
func (d *View) SetSize(size *Size) {
	C.gocef_view_set_size(d.toNative(), size.toNative(&C.cef_size_t{}), d.set_size)
}

// GetSize (get_size)
// Returns the size of this View.
func (d *View) GetSize() Size {
	cresult__ := C.gocef_view_get_size(d.toNative(), d.get_size)
	var result__ Size
	(&cresult__).intoGo(&result__)
	return result__
}

// SetPosition (set_position)
// Sets the position of this View without changing the size. |position| is in
// parent coordinates.
func (d *View) SetPosition(position *Point) {
	C.gocef_view_set_position(d.toNative(), position.toNative(&C.cef_point_t{}), d.set_position)
}

// GetPosition (get_position)
// Returns the position of this View. Position is in parent coordinates.
func (d *View) GetPosition() Point {
	cresult__ := C.gocef_view_get_position(d.toNative(), d.get_position)
	var result__ Point
	(&cresult__).intoGo(&result__)
	return result__
}

// GetPreferredSize (get_preferred_size)
// Returns the size this View would like to be if enough space is available.
func (d *View) GetPreferredSize() Size {
	cresult__ := C.gocef_view_get_preferred_size(d.toNative(), d.get_preferred_size)
	var result__ Size
	(&cresult__).intoGo(&result__)
	return result__
}

// SizeToPreferredSize (size_to_preferred_size)
// Size this View to its preferred size.
func (d *View) SizeToPreferredSize() {
	C.gocef_view_size_to_preferred_size(d.toNative(), d.size_to_preferred_size)
}

// GetMinimumSize (get_minimum_size)
// Returns the minimum size for this View.
func (d *View) GetMinimumSize() Size {
	cresult__ := C.gocef_view_get_minimum_size(d.toNative(), d.get_minimum_size)
	var result__ Size
	(&cresult__).intoGo(&result__)
	return result__
}

// GetMaximumSize (get_maximum_size)
// Returns the maximum size for this View.
func (d *View) GetMaximumSize() Size {
	cresult__ := C.gocef_view_get_maximum_size(d.toNative(), d.get_maximum_size)
	var result__ Size
	(&cresult__).intoGo(&result__)
	return result__
}

// GetHeightForWidth (get_height_for_width)
// Returns the height necessary to display this View with the provided width.
func (d *View) GetHeightForWidth(width int32) int32 {
	return int32(C.gocef_view_get_height_for_width(d.toNative(), C.int(width), d.get_height_for_width))
}

// InvalidateLayout (invalidate_layout)
// Indicate that this View and all parent Views require a re-layout. This
// ensures the next call to layout() will propagate to this View even if the
// bounds of parent Views do not change.
func (d *View) InvalidateLayout() {
	C.gocef_view_invalidate_layout(d.toNative(), d.invalidate_layout)
}

// SetVisible (set_visible)
// Sets whether this View is visible. Windows are hidden by default and other
// views are visible by default. This View and any parent views must be set as
// visible for this View to be drawn in a Window. If this View is set as
// hidden then it and any child views will not be drawn and, if any of those
// views currently have focus, then focus will also be cleared. Painting is
// scheduled as needed. If this View is a Window then calling this function is
// equivalent to calling the Window show() and hide() functions.
func (d *View) SetVisible(visible int32) {
	C.gocef_view_set_visible(d.toNative(), C.int(visible), d.set_visible)
}

// IsVisible (is_visible)
// Returns whether this View is visible. A view may be visible but still not
// drawn in a Window if any parent views are hidden. If this View is a Window
// then a return value of true (1) indicates that this Window is currently
// visible to the user on-screen. If this View is not a Window then call
// is_drawn() to determine whether this View and all parent views are visible
// and will be drawn.
func (d *View) IsVisible() int32 {
	return int32(C.gocef_view_is_visible(d.toNative(), d.is_visible))
}

// IsDrawn (is_drawn)
// Returns whether this View is visible and drawn in a Window. A view is drawn
// if it and all parent views are visible. If this View is a Window then
// calling this function is equivalent to calling is_visible(). Otherwise, to
// determine if the containing Window is visible to the user on-screen call
// is_visible() on the Window.
func (d *View) IsDrawn() int32 {
	return int32(C.gocef_view_is_drawn(d.toNative(), d.is_drawn))
}

// SetEnabled (set_enabled)
// Set whether this View is enabled. A disabled View does not receive keyboard
// or mouse inputs. If |enabled| differs from the current value the View will
// be repainted. Also, clears focus if the focused View is disabled.
func (d *View) SetEnabled(enabled int32) {
	C.gocef_view_set_enabled(d.toNative(), C.int(enabled), d.set_enabled)
}

// IsEnabled (is_enabled)
// Returns whether this View is enabled.
func (d *View) IsEnabled() int32 {
	return int32(C.gocef_view_is_enabled(d.toNative(), d.is_enabled))
}

// SetFocusable (set_focusable)
// Sets whether this View is capable of taking focus. It will clear focus if
// the focused View is set to be non-focusable. This is false (0) by default
// so that a View used as a container does not get the focus.
func (d *View) SetFocusable(focusable int32) {
	C.gocef_view_set_focusable(d.toNative(), C.int(focusable), d.set_focusable)
}

// IsFocusable (is_focusable)
// Returns true (1) if this View is focusable, enabled and drawn.
func (d *View) IsFocusable() int32 {
	return int32(C.gocef_view_is_focusable(d.toNative(), d.is_focusable))
}

// IsAccessibilityFocusable (is_accessibility_focusable)
// Return whether this View is focusable when the user requires full keyboard
// access, even though it may not be normally focusable.
func (d *View) IsAccessibilityFocusable() int32 {
	return int32(C.gocef_view_is_accessibility_focusable(d.toNative(), d.is_accessibility_focusable))
}

// RequestFocus (request_focus)
// Request keyboard focus. If this View is focusable it will become the
// focused View.
func (d *View) RequestFocus() {
	C.gocef_view_request_focus(d.toNative(), d.request_focus)
}

// SetBackgroundColor (set_background_color)
// Sets the background color for this View.
func (d *View) SetBackgroundColor(color Color) {
	C.gocef_view_set_background_color(d.toNative(), C.cef_color_t(color), d.set_background_color)
}

// GetBackgroundColor (get_background_color)
// Returns the background color for this View.
func (d *View) GetBackgroundColor() Color {
	return Color(C.gocef_view_get_background_color(d.toNative(), d.get_background_color))
}

// ConvertPointToScreen (convert_point_to_screen)
// Convert |point| from this View's coordinate system to that of the screen.
// This View must belong to a Window when calling this function. Returns true
// (1) if the conversion is successful or false (0) otherwise. Use
// cef_display_t::convert_point_to_pixels() after calling this function if
// further conversion to display-specific pixel coordinates is desired.
func (d *View) ConvertPointToScreen(point *Point) int32 {
	return int32(C.gocef_view_convert_point_to_screen(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_to_screen))
}

// ConvertPointFromScreen (convert_point_from_screen)
// Convert |point| to this View's coordinate system from that of the screen.
// This View must belong to a Window when calling this function. Returns true
// (1) if the conversion is successful or false (0) otherwise. Use
// cef_display_t::convert_point_from_pixels() before calling this function if
// conversion from display-specific pixel coordinates is necessary.
func (d *View) ConvertPointFromScreen(point *Point) int32 {
	return int32(C.gocef_view_convert_point_from_screen(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_from_screen))
}

// ConvertPointToWindow (convert_point_to_window)
// Convert |point| from this View's coordinate system to that of the Window.
// This View must belong to a Window when calling this function. Returns true
// (1) if the conversion is successful or false (0) otherwise.
func (d *View) ConvertPointToWindow(point *Point) int32 {
	return int32(C.gocef_view_convert_point_to_window(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_to_window))
}

// ConvertPointFromWindow (convert_point_from_window)
// Convert |point| to this View's coordinate system from that of the Window.
// This View must belong to a Window when calling this function. Returns true
// (1) if the conversion is successful or false (0) otherwise.
func (d *View) ConvertPointFromWindow(point *Point) int32 {
	return int32(C.gocef_view_convert_point_from_window(d.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_from_window))
}

// ConvertPointToView (convert_point_to_view)
// Convert |point| from this View's coordinate system to that of |view|.
// |view| needs to be in the same Window but not necessarily the same view
// hierarchy. Returns true (1) if the conversion is successful or false (0)
// otherwise.
func (d *View) ConvertPointToView(view *View, point *Point) int32 {
	return int32(C.gocef_view_convert_point_to_view(d.toNative(), view.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_to_view))
}

// ConvertPointFromView (convert_point_from_view)
// Convert |point| to this View's coordinate system from that |view|. |view|
// needs to be in the same Window but not necessarily the same view hierarchy.
// Returns true (1) if the conversion is successful or false (0) otherwise.
func (d *View) ConvertPointFromView(view *View, point *Point) int32 {
	return int32(C.gocef_view_convert_point_from_view(d.toNative(), view.toNative(), point.toNative(&C.cef_point_t{}), d.convert_point_from_view))
}
