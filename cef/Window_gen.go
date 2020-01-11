// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// void gocef_window_show(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_hide(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_center_window(cef_window_t * self, cef_size_t * size, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_size_t *)) { return callback__(self, size); }
	// void gocef_window_close(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// int gocef_window_is_closed(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_activate(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_deactivate(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// int gocef_window_is_active(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_bring_to_top(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_always_on_top(cef_window_t * self, int onTop, void (CEF_CALLBACK *callback__)(cef_window_t *, int)) { return callback__(self, onTop); }
	// int gocef_window_is_always_on_top(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_maximize(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_minimize(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_restore(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_fullscreen(cef_window_t * self, int fullscreen, void (CEF_CALLBACK *callback__)(cef_window_t *, int)) { return callback__(self, fullscreen); }
	// int gocef_window_is_maximized(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// int gocef_window_is_minimized(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// int gocef_window_is_fullscreen(cef_window_t * self, int (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_title(cef_window_t * self, cef_string_t * title, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_string_t *)) { return callback__(self, title); }
	// cef_string_userfree_t gocef_window_get_title(cef_window_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_window_icon(cef_window_t * self, cef_image_t * image, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_image_t *)) { return callback__(self, image); }
	// cef_image_t * gocef_window_get_window_icon(cef_window_t * self, cef_image_t * (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_window_app_icon(cef_window_t * self, cef_image_t * image, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_image_t *)) { return callback__(self, image); }
	// cef_image_t * gocef_window_get_window_app_icon(cef_window_t * self, cef_image_t * (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_show_menu(cef_window_t * self, cef_menu_model_t * menuModel, cef_point_t * screenPoint, cef_menu_anchor_position_t anchorPosition, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_menu_model_t *, cef_point_t *, cef_menu_anchor_position_t)) { return callback__(self, menuModel, screenPoint, anchorPosition); }
	// void gocef_window_cancel_menu(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// cef_display_t * gocef_window_get_display(cef_window_t * self, cef_display_t * (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// cef_rect_t gocef_window_get_client_area_bounds_in_screen(cef_window_t * self, cef_rect_t (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_set_draggable_regions(cef_window_t * self, size_t regionsCount, cef_draggable_region_t * regions, void (CEF_CALLBACK *callback__)(cef_window_t *, size_t, cef_draggable_region_t *)) { return callback__(self, regionsCount, regions); }
	// void * gocef_window_get_window_handle(cef_window_t * self, void * (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	// void gocef_window_send_key_press(cef_window_t * self, int keyCode, uint32 eventFlags, void (CEF_CALLBACK *callback__)(cef_window_t *, int, uint32)) { return callback__(self, keyCode, eventFlags); }
	// void gocef_window_send_mouse_move(cef_window_t * self, int screenX, int screenY, void (CEF_CALLBACK *callback__)(cef_window_t *, int, int)) { return callback__(self, screenX, screenY); }
	// void gocef_window_send_mouse_events(cef_window_t * self, cef_mouse_button_type_t button, int mouseDown, int mouseUp, void (CEF_CALLBACK *callback__)(cef_window_t *, cef_mouse_button_type_t, int, int)) { return callback__(self, button, mouseDown, mouseUp); }
	// void gocef_window_set_accelerator(cef_window_t * self, int commandID, int keyCode, int shiftPressed, int ctrlPressed, int altPressed, void (CEF_CALLBACK *callback__)(cef_window_t *, int, int, int, int, int)) { return callback__(self, commandID, keyCode, shiftPressed, ctrlPressed, altPressed); }
	// void gocef_window_remove_accelerator(cef_window_t * self, int commandID, void (CEF_CALLBACK *callback__)(cef_window_t *, int)) { return callback__(self, commandID); }
	// void gocef_window_remove_all_accelerators(cef_window_t * self, void (CEF_CALLBACK *callback__)(cef_window_t *)) { return callback__(self); }
	"C"
)

// Window (cef_window_t from include/capi/views/cef_window_capi.h)
// A Window is a top-level Window/widget in the Views hierarchy. By default it
// will have a non-client area with title bar, icon and buttons that supports
// moving and resizing. All size and position values are in density independent
// pixels (DIP) unless otherwise indicated. Methods must be called on the
// browser process UI thread unless otherwise indicated.
type Window C.cef_window_t

func (d *Window) toNative() *C.cef_window_t {
	return (*C.cef_window_t)(d)
}

// Base (base)
// Base structure.
func (d *Window) Base() *Panel {
	return (*Panel)(&d.base)
}

// Show (show)
// Show the Window.
func (d *Window) Show() {
	C.gocef_window_show(d.toNative(), d.show)
}

// Hide (hide)
// Hide the Window.
func (d *Window) Hide() {
	C.gocef_window_hide(d.toNative(), d.hide)
}

// CenterWindow (center_window)
// Sizes the Window to |size| and centers it in the current display.
func (d *Window) CenterWindow(size *Size) {
	C.gocef_window_center_window(d.toNative(), size.toNative(&C.cef_size_t{}), d.center_window)
}

// Close (close)
// Close the Window.
func (d *Window) Close() {
	C.gocef_window_close(d.toNative(), d.close)
}

// IsClosed (is_closed)
// Returns true (1) if the Window has been closed.
func (d *Window) IsClosed() int32 {
	return int32(C.gocef_window_is_closed(d.toNative(), d.is_closed))
}

// Activate (activate)
// Activate the Window, assuming it already exists and is visible.
func (d *Window) Activate() {
	C.gocef_window_activate(d.toNative(), d.activate)
}

// Deactivate (deactivate)
// Deactivate the Window, making the next Window in the Z order the active
// Window.
func (d *Window) Deactivate() {
	C.gocef_window_deactivate(d.toNative(), d.deactivate)
}

// IsActive (is_active)
// Returns whether the Window is the currently active Window.
func (d *Window) IsActive() int32 {
	return int32(C.gocef_window_is_active(d.toNative(), d.is_active))
}

// BringToTop (bring_to_top)
// Bring this Window to the top of other Windows in the Windowing system.
func (d *Window) BringToTop() {
	C.gocef_window_bring_to_top(d.toNative(), d.bring_to_top)
}

// SetAlwaysOnTop (set_always_on_top)
// Set the Window to be on top of other Windows in the Windowing system.
func (d *Window) SetAlwaysOnTop(onTop int32) {
	C.gocef_window_set_always_on_top(d.toNative(), C.int(onTop), d.set_always_on_top)
}

// IsAlwaysOnTop (is_always_on_top)
// Returns whether the Window has been set to be on top of other Windows in
// the Windowing system.
func (d *Window) IsAlwaysOnTop() int32 {
	return int32(C.gocef_window_is_always_on_top(d.toNative(), d.is_always_on_top))
}

// Maximize (maximize)
// Maximize the Window.
func (d *Window) Maximize() {
	C.gocef_window_maximize(d.toNative(), d.maximize)
}

// Minimize (minimize)
// Minimize the Window.
func (d *Window) Minimize() {
	C.gocef_window_minimize(d.toNative(), d.minimize)
}

// Restore (restore)
// Restore the Window.
func (d *Window) Restore() {
	C.gocef_window_restore(d.toNative(), d.restore)
}

// SetFullscreen (set_fullscreen)
// Set fullscreen Window state.
func (d *Window) SetFullscreen(fullscreen int32) {
	C.gocef_window_set_fullscreen(d.toNative(), C.int(fullscreen), d.set_fullscreen)
}

// IsMaximized (is_maximized)
// Returns true (1) if the Window is maximized.
func (d *Window) IsMaximized() int32 {
	return int32(C.gocef_window_is_maximized(d.toNative(), d.is_maximized))
}

// IsMinimized (is_minimized)
// Returns true (1) if the Window is minimized.
func (d *Window) IsMinimized() int32 {
	return int32(C.gocef_window_is_minimized(d.toNative(), d.is_minimized))
}

// IsFullscreen (is_fullscreen)
// Returns true (1) if the Window is fullscreen.
func (d *Window) IsFullscreen() int32 {
	return int32(C.gocef_window_is_fullscreen(d.toNative(), d.is_fullscreen))
}

// SetTitle (set_title)
// Set the Window title.
func (d *Window) SetTitle(title string) {
	title_ := C.cef_string_userfree_alloc()
	setCEFStr(title, title_)
	defer func() {
		C.cef_string_userfree_free(title_)
	}()
	C.gocef_window_set_title(d.toNative(), (*C.cef_string_t)(title_), d.set_title)
}

// GetTitle (get_title)
// Get the Window title.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Window) GetTitle() string {
	return cefuserfreestrToString(C.gocef_window_get_title(d.toNative(), d.get_title))
}

// SetWindowIcon (set_window_icon)
// Set the Window icon. This should be a 16x16 icon suitable for use in the
// Windows's title bar.
func (d *Window) SetWindowIcon(image *Image) {
	C.gocef_window_set_window_icon(d.toNative(), image.toNative(), d.set_window_icon)
}

// GetWindowIcon (get_window_icon)
// Get the Window icon.
func (d *Window) GetWindowIcon() *Image {
	return (*Image)(C.gocef_window_get_window_icon(d.toNative(), d.get_window_icon))
}

// SetWindowAppIcon (set_window_app_icon)
// Set the Window App icon. This should be a larger icon for use in the host
// environment app switching UI. On Windows, this is the ICON_BIG used in Alt-
// Tab list and Windows taskbar. The Window icon will be used by default if no
// Window App icon is specified.
func (d *Window) SetWindowAppIcon(image *Image) {
	C.gocef_window_set_window_app_icon(d.toNative(), image.toNative(), d.set_window_app_icon)
}

// GetWindowAppIcon (get_window_app_icon)
// Get the Window App icon.
func (d *Window) GetWindowAppIcon() *Image {
	return (*Image)(C.gocef_window_get_window_app_icon(d.toNative(), d.get_window_app_icon))
}

// ShowMenu (show_menu)
// Show a menu with contents |menu_model|. |screen_point| specifies the menu
// position in screen coordinates. |anchor_position| specifies how the menu
// will be anchored relative to |screen_point|.
func (d *Window) ShowMenu(menuModel *MenuModel, screenPoint *Point, anchorPosition MenuAnchorPosition) {
	C.gocef_window_show_menu(d.toNative(), menuModel.toNative(), screenPoint.toNative(&C.cef_point_t{}), C.cef_menu_anchor_position_t(anchorPosition), d.show_menu)
}

// CancelMenu (cancel_menu)
// Cancel the menu that is currently showing, if any.
func (d *Window) CancelMenu() {
	C.gocef_window_cancel_menu(d.toNative(), d.cancel_menu)
}

// GetDisplay (get_display)
// Returns the Display that most closely intersects the bounds of this Window.
// May return NULL if this Window is not currently displayed.
func (d *Window) GetDisplay() *Display {
	return (*Display)(C.gocef_window_get_display(d.toNative(), d.get_display))
}

// GetClientAreaBoundsInScreen (get_client_area_bounds_in_screen)
// Returns the bounds (size and position) of this Window's client area.
// Position is in screen coordinates.
func (d *Window) GetClientAreaBoundsInScreen() Rect {
	cresult__ := C.gocef_window_get_client_area_bounds_in_screen(d.toNative(), d.get_client_area_bounds_in_screen)
	var result__ Rect
	(&cresult__).intoGo(&result__)
	return result__
}

// SetDraggableRegions (set_draggable_regions)
// Set the regions where mouse events will be intercepted by this Window to
// support drag operations. Call this function with an NULL vector to clear
// the draggable regions. The draggable region bounds should be in window
// coordinates.
func (d *Window) SetDraggableRegions(regionsCount uint64, regions *DraggableRegion) {
	C.gocef_window_set_draggable_regions(d.toNative(), C.size_t(regionsCount), regions.toNative(&C.cef_draggable_region_t{}), d.set_draggable_regions)
}

// GetWindowHandle (get_window_handle)
// Retrieve the platform window handle for this Window.
func (d *Window) GetWindowHandle() unsafe.Pointer {
	return C.gocef_window_get_window_handle(d.toNative(), d.get_window_handle)
}

// SendKeyPress (send_key_press)
// Simulate a key press. |key_code| is the VKEY_* value from Chromium's
// ui/events/keycodes/keyboard_codes.h header (VK_* values on Windows).
// |event_flags| is some combination of EVENTFLAG_SHIFT_DOWN,
// EVENTFLAG_CONTROL_DOWN and/or EVENTFLAG_ALT_DOWN. This function is exposed
// primarily for testing purposes.
func (d *Window) SendKeyPress(keyCode int32, eventFlags uint32) {
	C.gocef_window_send_key_press(d.toNative(), C.int(keyCode), C.uint32(eventFlags), d.send_key_press)
}

// SendMouseMove (send_mouse_move)
// Simulate a mouse move. The mouse cursor will be moved to the specified
// (screen_x, screen_y) position. This function is exposed primarily for
// testing purposes.
func (d *Window) SendMouseMove(screenX, screenY int32) {
	C.gocef_window_send_mouse_move(d.toNative(), C.int(screenX), C.int(screenY), d.send_mouse_move)
}

// SendMouseEvents (send_mouse_events)
// Simulate mouse down and/or mouse up events. |button| is the mouse button
// type. If |mouse_down| is true (1) a mouse down event will be sent. If
// |mouse_up| is true (1) a mouse up event will be sent. If both are true (1)
// a mouse down event will be sent followed by a mouse up event (equivalent to
// clicking the mouse button). The events will be sent using the current
// cursor position so make sure to call send_mouse_move() first to position
// the mouse. This function is exposed primarily for testing purposes.
func (d *Window) SendMouseEvents(button MouseButtonType, mouseDown, mouseUp int32) {
	C.gocef_window_send_mouse_events(d.toNative(), C.cef_mouse_button_type_t(button), C.int(mouseDown), C.int(mouseUp), d.send_mouse_events)
}

// SetAccelerator (set_accelerator)
// Set the keyboard accelerator for the specified |command_id|. |key_code| can
// be any virtual key or character value. cef_window_delegate_t::OnAccelerator
// will be called if the keyboard combination is triggered while this window
// has focus.
func (d *Window) SetAccelerator(commandID, keyCode, shiftPressed, ctrlPressed, altPressed int32) {
	C.gocef_window_set_accelerator(d.toNative(), C.int(commandID), C.int(keyCode), C.int(shiftPressed), C.int(ctrlPressed), C.int(altPressed), d.set_accelerator)
}

// RemoveAccelerator (remove_accelerator)
// Remove the keyboard accelerator for the specified |command_id|.
func (d *Window) RemoveAccelerator(commandID int32) {
	C.gocef_window_remove_accelerator(d.toNative(), C.int(commandID), d.remove_accelerator)
}

// RemoveAllAccelerators (remove_all_accelerators)
// Remove all keyboard accelerators.
func (d *Window) RemoveAllAccelerators() {
	C.gocef_window_remove_all_accelerators(d.toNative(), d.remove_all_accelerators)
}
