// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// cef_browser_t * gocef_browser_host_get_browser(cef_browser_host_t * self, cef_browser_t * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_close_browser(cef_browser_host_t * self, int forceClose, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, forceClose); }
	// int gocef_browser_host_try_close_browser(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_set_focus(cef_browser_host_t * self, int focus, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, focus); }
	// void * gocef_browser_host_get_window_handle(cef_browser_host_t * self, void * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void * gocef_browser_host_get_opener_window_handle(cef_browser_host_t * self, void * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// int gocef_browser_host_has_view(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// cef_client_t * gocef_browser_host_get_client(cef_browser_host_t * self, cef_client_t * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// cef_request_context_t * gocef_browser_host_get_request_context(cef_browser_host_t * self, cef_request_context_t * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// double gocef_browser_host_get_zoom_level(cef_browser_host_t * self, double (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_set_zoom_level(cef_browser_host_t * self, double zoomLevel, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, double)) { return callback__(self, zoomLevel); }
	// void gocef_browser_host_run_file_dialog(cef_browser_host_t * self, cef_file_dialog_mode_t mode, cef_string_t * title, cef_string_t * defaultFilePath, cef_string_list_t acceptFilters, int selectedAcceptFilter, cef_run_file_dialog_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_file_dialog_mode_t, cef_string_t *, cef_string_t *, cef_string_list_t, int, cef_run_file_dialog_callback_t *)) { return callback__(self, mode, title, defaultFilePath, acceptFilters, selectedAcceptFilter, callback); }
	// void gocef_browser_host_start_download(cef_browser_host_t * self, cef_string_t * uRL, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *)) { return callback__(self, uRL); }
	// void gocef_browser_host_download_image(cef_browser_host_t * self, cef_string_t * imageURL, int isFavicon, uint32 maxImageSize, int bypassCache, cef_download_image_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *, int, uint32, int, cef_download_image_callback_t *)) { return callback__(self, imageURL, isFavicon, maxImageSize, bypassCache, callback); }
	// void gocef_browser_host_print(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_print_to_pdf(cef_browser_host_t * self, cef_string_t * path, cef_pdf_print_settings_t * settings, cef_pdf_print_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *, cef_pdf_print_settings_t *, cef_pdf_print_callback_t *)) { return callback__(self, path, settings, callback); }
	// void gocef_browser_host_find(cef_browser_host_t * self, int identifier, cef_string_t * searchText, int forward, int matchCase, int findNext, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int, cef_string_t *, int, int, int)) { return callback__(self, identifier, searchText, forward, matchCase, findNext); }
	// void gocef_browser_host_stop_finding(cef_browser_host_t * self, int clearSelection, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, clearSelection); }
	// void gocef_browser_host_show_dev_tools(cef_browser_host_t * self, cef_window_info_t * windowInfo, cef_client_t * client, cef_browser_settings_t * settings, cef_point_t * inspectElementAt, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_window_info_t *, cef_client_t *, cef_browser_settings_t *, cef_point_t *)) { return callback__(self, windowInfo, client, settings, inspectElementAt); }
	// void gocef_browser_host_close_dev_tools(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// int gocef_browser_host_has_dev_tools(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_get_navigation_entries(cef_browser_host_t * self, cef_navigation_entry_visitor_t * visitor, int currentOnly, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_navigation_entry_visitor_t *, int)) { return callback__(self, visitor, currentOnly); }
	// void gocef_browser_host_set_mouse_cursor_change_disabled(cef_browser_host_t * self, int disabled, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, disabled); }
	// int gocef_browser_host_is_mouse_cursor_change_disabled(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_replace_misspelling(cef_browser_host_t * self, cef_string_t * word, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *)) { return callback__(self, word); }
	// void gocef_browser_host_add_word_to_dictionary(cef_browser_host_t * self, cef_string_t * word, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *)) { return callback__(self, word); }
	// int gocef_browser_host_is_window_rendering_disabled(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_was_resized(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_was_hidden(cef_browser_host_t * self, int hidden, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, hidden); }
	// void gocef_browser_host_notify_screen_info_changed(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_invalidate(cef_browser_host_t * self, cef_paint_element_type_t type_r, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_paint_element_type_t)) { return callback__(self, type_r); }
	// void gocef_browser_host_send_external_begin_frame(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_send_key_event(cef_browser_host_t * self, cef_key_event_t * event, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_key_event_t *)) { return callback__(self, event); }
	// void gocef_browser_host_send_mouse_click_event(cef_browser_host_t * self, cef_mouse_event_t * event, cef_mouse_button_type_t type_r, int mouseUp, int clickCount, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_mouse_event_t *, cef_mouse_button_type_t, int, int)) { return callback__(self, event, type_r, mouseUp, clickCount); }
	// void gocef_browser_host_send_mouse_move_event(cef_browser_host_t * self, cef_mouse_event_t * event, int mouseLeave, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_mouse_event_t *, int)) { return callback__(self, event, mouseLeave); }
	// void gocef_browser_host_send_mouse_wheel_event(cef_browser_host_t * self, cef_mouse_event_t * event, int deltaX, int deltaY, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_mouse_event_t *, int, int)) { return callback__(self, event, deltaX, deltaY); }
	// void gocef_browser_host_send_touch_event(cef_browser_host_t * self, cef_touch_event_t * event, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_touch_event_t *)) { return callback__(self, event); }
	// void gocef_browser_host_send_focus_event(cef_browser_host_t * self, int setFocus, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, setFocus); }
	// void gocef_browser_host_send_capture_lost_event(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_notify_move_or_resize_started(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// int gocef_browser_host_get_windowless_frame_rate(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_set_windowless_frame_rate(cef_browser_host_t * self, int frameRate, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, frameRate); }
	// void gocef_browser_host_ime_set_composition(cef_browser_host_t * self, cef_string_t * text, size_t underlinesCount, cef_composition_underline_t * underlines, cef_range_t * replacementRange, cef_range_t * selectionRange, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *, size_t, cef_composition_underline_t *, cef_range_t *, cef_range_t *)) { return callback__(self, text, underlinesCount, underlines, replacementRange, selectionRange); }
	// void gocef_browser_host_ime_commit_text(cef_browser_host_t * self, cef_string_t * text, cef_range_t * replacementRange, int relativeCursorPos, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_string_t *, cef_range_t *, int)) { return callback__(self, text, replacementRange, relativeCursorPos); }
	// void gocef_browser_host_ime_finish_composing_text(cef_browser_host_t * self, int keepSelection, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, keepSelection); }
	// void gocef_browser_host_ime_cancel_composition(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_drag_target_drag_enter(cef_browser_host_t * self, cef_drag_data_t * dragData, cef_mouse_event_t * event, cef_drag_operations_mask_t allowedOps, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_drag_data_t *, cef_mouse_event_t *, cef_drag_operations_mask_t)) { return callback__(self, dragData, event, allowedOps); }
	// void gocef_browser_host_drag_target_drag_over(cef_browser_host_t * self, cef_mouse_event_t * event, cef_drag_operations_mask_t allowedOps, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_mouse_event_t *, cef_drag_operations_mask_t)) { return callback__(self, event, allowedOps); }
	// void gocef_browser_host_drag_target_drag_leave(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_drag_target_drop(cef_browser_host_t * self, cef_mouse_event_t * event, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_mouse_event_t *)) { return callback__(self, event); }
	// void gocef_browser_host_drag_source_ended_at(cef_browser_host_t * self, int x, int y, cef_drag_operations_mask_t op, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int, int, cef_drag_operations_mask_t)) { return callback__(self, x, y, op); }
	// void gocef_browser_host_drag_source_system_drag_ended(cef_browser_host_t * self, void (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// cef_navigation_entry_t * gocef_browser_host_get_visible_navigation_entry(cef_browser_host_t * self, cef_navigation_entry_t * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_set_accessibility_state(cef_browser_host_t * self, cef_state_t accessibilityState, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, cef_state_t)) { return callback__(self, accessibilityState); }
	// void gocef_browser_host_set_auto_resize_enabled(cef_browser_host_t * self, int enabled, cef_size_t * minSize, cef_size_t * maxSize, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int, cef_size_t *, cef_size_t *)) { return callback__(self, enabled, minSize, maxSize); }
	// cef_extension_t * gocef_browser_host_get_extension(cef_browser_host_t * self, cef_extension_t * (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// int gocef_browser_host_is_background_host(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	// void gocef_browser_host_set_audio_muted(cef_browser_host_t * self, int mute, void (CEF_CALLBACK *callback__)(cef_browser_host_t *, int)) { return callback__(self, mute); }
	// int gocef_browser_host_is_audio_muted(cef_browser_host_t * self, int (CEF_CALLBACK *callback__)(cef_browser_host_t *)) { return callback__(self); }
	"C"
)

// BrowserHost (cef_browser_host_t from include/capi/cef_browser_capi.h)
// Structure used to represent the browser process aspects of a browser window.
// The functions of this structure can only be called in the browser process.
// They may be called on any thread in that process unless otherwise indicated
// in the comments.
type BrowserHost C.cef_browser_host_t

func (d *BrowserHost) toNative() *C.cef_browser_host_t {
	return (*C.cef_browser_host_t)(d)
}

// Base (base)
// Base structure.
func (d *BrowserHost) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetBrowser (get_browser)
// Returns the hosted browser object.
func (d *BrowserHost) GetBrowser() *Browser {
	return (*Browser)(C.gocef_browser_host_get_browser(d.toNative(), d.get_browser))
}

// CloseBrowser (close_browser)
// Request that the browser close. The JavaScript 'onbeforeunload' event will
// be fired. If |force_close| is false (0) the event handler, if any, will be
// allowed to prompt the user and the user can optionally cancel the close. If
// |force_close| is true (1) the prompt will not be displayed and the close
// will proceed. Results in a call to cef_life_span_handler_t::do_close() if
// the event handler allows the close or if |force_close| is true (1). See
// cef_life_span_handler_t::do_close() documentation for additional usage
// information.
func (d *BrowserHost) CloseBrowser(forceClose int32) {
	C.gocef_browser_host_close_browser(d.toNative(), C.int(forceClose), d.close_browser)
}

// TryCloseBrowser (try_close_browser)
// Helper for closing a browser. Call this function from the top-level window
// close handler. Internally this calls CloseBrowser(false (0)) if the close
// has not yet been initiated. This function returns false (0) while the close
// is pending and true (1) after the close has completed. See close_browser()
// and cef_life_span_handler_t::do_close() documentation for additional usage
// information. This function must be called on the browser process UI thread.
func (d *BrowserHost) TryCloseBrowser() int32 {
	return int32(C.gocef_browser_host_try_close_browser(d.toNative(), d.try_close_browser))
}

// SetFocus (set_focus)
// Set whether the browser is focused.
func (d *BrowserHost) SetFocus(focus int32) {
	C.gocef_browser_host_set_focus(d.toNative(), C.int(focus), d.set_focus)
}

// GetWindowHandle (get_window_handle)
// Retrieve the window handle for this browser. If this browser is wrapped in
// a cef_browser_view_t this function should be called on the browser process
// UI thread and it will return the handle for the top-level native window.
func (d *BrowserHost) GetWindowHandle() unsafe.Pointer {
	return C.gocef_browser_host_get_window_handle(d.toNative(), d.get_window_handle)
}

// GetOpenerWindowHandle (get_opener_window_handle)
// Retrieve the window handle of the browser that opened this browser. Will
// return NULL for non-popup windows or if this browser is wrapped in a
// cef_browser_view_t. This function can be used in combination with custom
// handling of modal windows.
func (d *BrowserHost) GetOpenerWindowHandle() unsafe.Pointer {
	return C.gocef_browser_host_get_opener_window_handle(d.toNative(), d.get_opener_window_handle)
}

// HasView (has_view)
// Returns true (1) if this browser is wrapped in a cef_browser_view_t.
func (d *BrowserHost) HasView() int32 {
	return int32(C.gocef_browser_host_has_view(d.toNative(), d.has_view))
}

// GetClient (get_client)
// Returns the client for this browser.
func (d *BrowserHost) GetClient() *Client {
	return (*Client)(C.gocef_browser_host_get_client(d.toNative(), d.get_client))
}

// GetRequestContext (get_request_context)
// Returns the request context for this browser.
func (d *BrowserHost) GetRequestContext() *RequestContext {
	return (*RequestContext)(C.gocef_browser_host_get_request_context(d.toNative(), d.get_request_context))
}

// GetZoomLevel (get_zoom_level)
// Get the current zoom level. The default zoom level is 0.0. This function
// can only be called on the UI thread.
func (d *BrowserHost) GetZoomLevel() float64 {
	return float64(C.gocef_browser_host_get_zoom_level(d.toNative(), d.get_zoom_level))
}

// SetZoomLevel (set_zoom_level)
// Change the zoom level to the specified value. Specify 0.0 to reset the zoom
// level. If called on the UI thread the change will be applied immediately.
// Otherwise, the change will be applied asynchronously on the UI thread.
func (d *BrowserHost) SetZoomLevel(zoomLevel float64) {
	C.gocef_browser_host_set_zoom_level(d.toNative(), C.double(zoomLevel), d.set_zoom_level)
}

// RunFileDialog (run_file_dialog)
// Call to run a file chooser dialog. Only a single file chooser dialog may be
// pending at any given time. |mode| represents the type of dialog to display.
// |title| to the title to be used for the dialog and may be NULL to show the
// default title ("Open" or "Save" depending on the mode). |default_file_path|
// is the path with optional directory and/or file name component that will be
// initially selected in the dialog. |accept_filters| are used to restrict the
// selectable file types and may any combination of (a) valid lower-cased MIME
// types (e.g. "text/*" or "image/*"), (b) individual file extensions (e.g.
// ".txt" or ".png"), or (c) combined description and file extension delimited
// using "|" and ";" (e.g. "Image Types|.png;.gif;.jpg").
// |selected_accept_filter| is the 0-based index of the filter that will be
// selected by default. |callback| will be executed after the dialog is
// dismissed or immediately if another dialog is already pending. The dialog
// will be initiated asynchronously on the UI thread.
func (d *BrowserHost) RunFileDialog(mode FileDialogMode, title, defaultFilePath string, acceptFilters StringList, selectedAcceptFilter int32, callback *RunFileDialogCallback) {
	title_ := C.cef_string_userfree_alloc()
	setCEFStr(title, title_)
	defer func() {
		C.cef_string_userfree_free(title_)
	}()
	defaultFilePath_ := C.cef_string_userfree_alloc()
	setCEFStr(defaultFilePath, defaultFilePath_)
	defer func() {
		C.cef_string_userfree_free(defaultFilePath_)
	}()
	C.gocef_browser_host_run_file_dialog(d.toNative(), C.cef_file_dialog_mode_t(mode), (*C.cef_string_t)(title_), (*C.cef_string_t)(defaultFilePath_), C.cef_string_list_t(acceptFilters), C.int(selectedAcceptFilter), callback.toNative(), d.run_file_dialog)
}

// StartDownload (start_download)
// Download the file at |url| using cef_download_handler_t.
func (d *BrowserHost) StartDownload(uRL string) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	C.gocef_browser_host_start_download(d.toNative(), (*C.cef_string_t)(uRL_), d.start_download)
}

// DownloadImage (download_image)
// Download |image_url| and execute |callback| on completion with the images
// received from the renderer. If |is_favicon| is true (1) then cookies are
// not sent and not accepted during download. Images with density independent
// pixel (DIP) sizes larger than |max_image_size| are filtered out from the
// image results. Versions of the image at different scale factors may be
// downloaded up to the maximum scale factor supported by the system. If there
// are no image results <= |max_image_size| then the smallest image is resized
// to |max_image_size| and is the only result. A |max_image_size| of 0 means
// unlimited. If |bypass_cache| is true (1) then |image_url| is requested from
// the server even if it is present in the browser cache.
func (d *BrowserHost) DownloadImage(imageURL string, isFavicon int32, maxImageSize uint32, bypassCache int32, callback *DownloadImageCallback) {
	imageURL_ := C.cef_string_userfree_alloc()
	setCEFStr(imageURL, imageURL_)
	defer func() {
		C.cef_string_userfree_free(imageURL_)
	}()
	C.gocef_browser_host_download_image(d.toNative(), (*C.cef_string_t)(imageURL_), C.int(isFavicon), C.uint32(maxImageSize), C.int(bypassCache), callback.toNative(), d.download_image)
}

// Print (print)
// Print the current browser contents.
func (d *BrowserHost) Print() {
	C.gocef_browser_host_print(d.toNative(), d.print)
}

// PrintToPDF (print_to_pdf)
// Print the current browser contents to the PDF file specified by |path| and
// execute |callback| on completion. The caller is responsible for deleting
// |path| when done. For PDF printing to work on Linux you must implement the
// cef_print_handler_t::GetPdfPaperSize function.
func (d *BrowserHost) PrintToPDF(path string, settings *PDFPrintSettings, callback *PDFPrintCallback) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	C.gocef_browser_host_print_to_pdf(d.toNative(), (*C.cef_string_t)(path_), settings.toNative(&C.cef_pdf_print_settings_t{}), callback.toNative(), d.print_to_pdf)
}

// Find (find)
// Search for |searchText|. |identifier| must be a unique ID and these IDs
// must strictly increase so that newer requests always have greater IDs than
// older requests. If |identifier| is zero or less than the previous ID value
// then it will be automatically assigned a new valid ID. |forward| indicates
// whether to search forward or backward within the page. |matchCase|
// indicates whether the search should be case-sensitive. |findNext| indicates
// whether this is the first request or a follow-up. The cef_find_handler_t
// instance, if any, returned via cef_client_t::GetFindHandler will be called
// to report find results.
func (d *BrowserHost) Find(identifier int32, searchText string, forward, matchCase, findNext int32) {
	searchText_ := C.cef_string_userfree_alloc()
	setCEFStr(searchText, searchText_)
	defer func() {
		C.cef_string_userfree_free(searchText_)
	}()
	C.gocef_browser_host_find(d.toNative(), C.int(identifier), (*C.cef_string_t)(searchText_), C.int(forward), C.int(matchCase), C.int(findNext), d.find)
}

// StopFinding (stop_finding)
// Cancel all searches that are currently going on.
func (d *BrowserHost) StopFinding(clearSelection int32) {
	C.gocef_browser_host_stop_finding(d.toNative(), C.int(clearSelection), d.stop_finding)
}

// ShowDevTools (show_dev_tools)
// Open developer tools (DevTools) in its own browser. The DevTools browser
// will remain associated with this browser. If the DevTools browser is
// already open then it will be focused, in which case the |windowInfo|,
// |client| and |settings| parameters will be ignored. If |inspect_element_at|
// is non-NULL then the element at the specified (x,y) location will be
// inspected. The |windowInfo| parameter will be ignored if this browser is
// wrapped in a cef_browser_view_t.
func (d *BrowserHost) ShowDevTools(windowInfo *WindowInfo, client *Client, settings *BrowserSettings, inspectElementAt *Point) {
	C.gocef_browser_host_show_dev_tools(d.toNative(), windowInfo.toNative(&C.cef_window_info_t{}), client.toNative(), settings.toNative(&C.cef_browser_settings_t{}), inspectElementAt.toNative(&C.cef_point_t{}), d.show_dev_tools)
}

// CloseDevTools (close_dev_tools)
// Explicitly close the associated DevTools browser, if any.
func (d *BrowserHost) CloseDevTools() {
	C.gocef_browser_host_close_dev_tools(d.toNative(), d.close_dev_tools)
}

// HasDevTools (has_dev_tools)
// Returns true (1) if this browser currently has an associated DevTools
// browser. Must be called on the browser process UI thread.
func (d *BrowserHost) HasDevTools() int32 {
	return int32(C.gocef_browser_host_has_dev_tools(d.toNative(), d.has_dev_tools))
}

// GetNavigationEntries (get_navigation_entries)
// Retrieve a snapshot of current navigation entries as values sent to the
// specified visitor. If |current_only| is true (1) only the current
// navigation entry will be sent, otherwise all navigation entries will be
// sent.
func (d *BrowserHost) GetNavigationEntries(visitor *NavigationEntryVisitor, currentOnly int32) {
	C.gocef_browser_host_get_navigation_entries(d.toNative(), visitor.toNative(), C.int(currentOnly), d.get_navigation_entries)
}

// SetMouseCursorChangeDisabled (set_mouse_cursor_change_disabled)
// Set whether mouse cursor change is disabled.
func (d *BrowserHost) SetMouseCursorChangeDisabled(disabled int32) {
	C.gocef_browser_host_set_mouse_cursor_change_disabled(d.toNative(), C.int(disabled), d.set_mouse_cursor_change_disabled)
}

// IsMouseCursorChangeDisabled (is_mouse_cursor_change_disabled)
// Returns true (1) if mouse cursor change is disabled.
func (d *BrowserHost) IsMouseCursorChangeDisabled() int32 {
	return int32(C.gocef_browser_host_is_mouse_cursor_change_disabled(d.toNative(), d.is_mouse_cursor_change_disabled))
}

// ReplaceMisspelling (replace_misspelling)
// If a misspelled word is currently selected in an editable node calling this
// function will replace it with the specified |word|.
func (d *BrowserHost) ReplaceMisspelling(word string) {
	word_ := C.cef_string_userfree_alloc()
	setCEFStr(word, word_)
	defer func() {
		C.cef_string_userfree_free(word_)
	}()
	C.gocef_browser_host_replace_misspelling(d.toNative(), (*C.cef_string_t)(word_), d.replace_misspelling)
}

// AddWordToDictionary (add_word_to_dictionary)
// Add the specified |word| to the spelling dictionary.
func (d *BrowserHost) AddWordToDictionary(word string) {
	word_ := C.cef_string_userfree_alloc()
	setCEFStr(word, word_)
	defer func() {
		C.cef_string_userfree_free(word_)
	}()
	C.gocef_browser_host_add_word_to_dictionary(d.toNative(), (*C.cef_string_t)(word_), d.add_word_to_dictionary)
}

// IsWindowRenderingDisabled (is_window_rendering_disabled)
// Returns true (1) if window rendering is disabled.
func (d *BrowserHost) IsWindowRenderingDisabled() int32 {
	return int32(C.gocef_browser_host_is_window_rendering_disabled(d.toNative(), d.is_window_rendering_disabled))
}

// WasResized (was_resized)
// Notify the browser that the widget has been resized. The browser will first
// call cef_render_handler_t::GetViewRect to get the new size and then call
// cef_render_handler_t::OnPaint asynchronously with the updated regions. This
// function is only used when window rendering is disabled.
func (d *BrowserHost) WasResized() {
	C.gocef_browser_host_was_resized(d.toNative(), d.was_resized)
}

// WasHidden (was_hidden)
// Notify the browser that it has been hidden or shown. Layouting and
// cef_render_handler_t::OnPaint notification will stop when the browser is
// hidden. This function is only used when window rendering is disabled.
func (d *BrowserHost) WasHidden(hidden int32) {
	C.gocef_browser_host_was_hidden(d.toNative(), C.int(hidden), d.was_hidden)
}

// NotifyScreenInfoChanged (notify_screen_info_changed)
// Send a notification to the browser that the screen info has changed. The
// browser will then call cef_render_handler_t::GetScreenInfo to update the
// screen information with the new values. This simulates moving the webview
// window from one display to another, or changing the properties of the
// current display. This function is only used when window rendering is
// disabled.
func (d *BrowserHost) NotifyScreenInfoChanged() {
	C.gocef_browser_host_notify_screen_info_changed(d.toNative(), d.notify_screen_info_changed)
}

// Invalidate (invalidate)
// Invalidate the view. The browser will call cef_render_handler_t::OnPaint
// asynchronously. This function is only used when window rendering is
// disabled.
func (d *BrowserHost) Invalidate(type_r PaintElementType) {
	C.gocef_browser_host_invalidate(d.toNative(), C.cef_paint_element_type_t(type_r), d.invalidate)
}

// SendExternalBeginFrame (send_external_begin_frame)
// Issue a BeginFrame request to Chromium.  Only valid when
// cef_window_tInfo::external_begin_frame_enabled is set to true (1).
func (d *BrowserHost) SendExternalBeginFrame() {
	C.gocef_browser_host_send_external_begin_frame(d.toNative(), d.send_external_begin_frame)
}

// SendKeyEvent (send_key_event)
// Send a key event to the browser.
func (d *BrowserHost) SendKeyEvent(event *KeyEvent) {
	C.gocef_browser_host_send_key_event(d.toNative(), event.toNative(&C.cef_key_event_t{}), d.send_key_event)
}

// SendMouseClickEvent (send_mouse_click_event)
// Send a mouse click event to the browser. The |x| and |y| coordinates are
// relative to the upper-left corner of the view.
func (d *BrowserHost) SendMouseClickEvent(event *MouseEvent, type_r MouseButtonType, mouseUp, clickCount int32) {
	C.gocef_browser_host_send_mouse_click_event(d.toNative(), event.toNative(&C.cef_mouse_event_t{}), C.cef_mouse_button_type_t(type_r), C.int(mouseUp), C.int(clickCount), d.send_mouse_click_event)
}

// SendMouseMoveEvent (send_mouse_move_event)
// Send a mouse move event to the browser. The |x| and |y| coordinates are
// relative to the upper-left corner of the view.
func (d *BrowserHost) SendMouseMoveEvent(event *MouseEvent, mouseLeave int32) {
	C.gocef_browser_host_send_mouse_move_event(d.toNative(), event.toNative(&C.cef_mouse_event_t{}), C.int(mouseLeave), d.send_mouse_move_event)
}

// SendMouseWheelEvent (send_mouse_wheel_event)
// Send a mouse wheel event to the browser. The |x| and |y| coordinates are
// relative to the upper-left corner of the view. The |deltaX| and |deltaY|
// values represent the movement delta in the X and Y directions respectively.
// In order to scroll inside select popups with window rendering disabled
// cef_render_handler_t::GetScreenPoint should be implemented properly.
func (d *BrowserHost) SendMouseWheelEvent(event *MouseEvent, deltaX, deltaY int32) {
	C.gocef_browser_host_send_mouse_wheel_event(d.toNative(), event.toNative(&C.cef_mouse_event_t{}), C.int(deltaX), C.int(deltaY), d.send_mouse_wheel_event)
}

// SendTouchEvent (send_touch_event)
// Send a touch event to the browser for a windowless browser.
func (d *BrowserHost) SendTouchEvent(event *TouchEvent) {
	C.gocef_browser_host_send_touch_event(d.toNative(), event.toNative(&C.cef_touch_event_t{}), d.send_touch_event)
}

// SendFocusEvent (send_focus_event)
// Send a focus event to the browser.
func (d *BrowserHost) SendFocusEvent(setFocus int32) {
	C.gocef_browser_host_send_focus_event(d.toNative(), C.int(setFocus), d.send_focus_event)
}

// SendCaptureLostEvent (send_capture_lost_event)
// Send a capture lost event to the browser.
func (d *BrowserHost) SendCaptureLostEvent() {
	C.gocef_browser_host_send_capture_lost_event(d.toNative(), d.send_capture_lost_event)
}

// NotifyMoveOrResizeStarted (notify_move_or_resize_started)
// Notify the browser that the window hosting it is about to be moved or
// resized. This function is only used on Windows and Linux.
func (d *BrowserHost) NotifyMoveOrResizeStarted() {
	C.gocef_browser_host_notify_move_or_resize_started(d.toNative(), d.notify_move_or_resize_started)
}

// GetWindowlessFrameRate (get_windowless_frame_rate)
// Returns the maximum rate in frames per second (fps) that
// cef_render_handler_t:: OnPaint will be called for a windowless browser. The
// actual fps may be lower if the browser cannot generate frames at the
// requested rate. The minimum value is 1 and the maximum value is 60 (default
// 30). This function can only be called on the UI thread.
func (d *BrowserHost) GetWindowlessFrameRate() int32 {
	return int32(C.gocef_browser_host_get_windowless_frame_rate(d.toNative(), d.get_windowless_frame_rate))
}

// SetWindowlessFrameRate (set_windowless_frame_rate)
// Set the maximum rate in frames per second (fps) that cef_render_handler_t::
// OnPaint will be called for a windowless browser. The actual fps may be
// lower if the browser cannot generate frames at the requested rate. The
// minimum value is 1 and the maximum value is 60 (default 30). Can also be
// set at browser creation via cef_browser_tSettings.windowless_frame_rate.
func (d *BrowserHost) SetWindowlessFrameRate(frameRate int32) {
	C.gocef_browser_host_set_windowless_frame_rate(d.toNative(), C.int(frameRate), d.set_windowless_frame_rate)
}

// ImeSetComposition (ime_set_composition)
// Begins a new composition or updates the existing composition. Blink has a
// special node (a composition node) that allows the input function to change
// text without affecting other DOM nodes. |text| is the optional text that
// will be inserted into the composition node. |underlines| is an optional set
// of ranges that will be underlined in the resulting text.
// |replacement_range| is an optional range of the existing text that will be
// replaced. |selection_range| is an optional range of the resulting text that
// will be selected after insertion or replacement. The |replacement_range|
// value is only used on OS X.
//
// This function may be called multiple times as the composition changes. When
// the client is done making changes the composition should either be canceled
// or completed. To cancel the composition call ImeCancelComposition. To
// complete the composition call either ImeCommitText or
// ImeFinishComposingText. Completion is usually signaled when:
//   A. The client receives a WM_IME_COMPOSITION message with a GCS_RESULTSTR
//      flag (on Windows), or;
//   B. The client receives a "commit" signal of GtkIMContext (on Linux), or;
//   C. insertText of NSTextInput is called (on Mac).
//
// This function is only used when window rendering is disabled.
func (d *BrowserHost) ImeSetComposition(text string, underlinesCount uint64, underlines *CompositionUnderline, replacementRange, selectionRange *Range) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_browser_host_ime_set_composition(d.toNative(), (*C.cef_string_t)(text_), C.size_t(underlinesCount), underlines.toNative(&C.cef_composition_underline_t{}), replacementRange.toNative(&C.cef_range_t{}), selectionRange.toNative(&C.cef_range_t{}), d.ime_set_composition)
}

// ImeCommitText (ime_commit_text)
// Completes the existing composition by optionally inserting the specified
// |text| into the composition node. |replacement_range| is an optional range
// of the existing text that will be replaced. |relative_cursor_pos| is where
// the cursor will be positioned relative to the current cursor position. See
// comments on ImeSetComposition for usage. The |replacement_range| and
// |relative_cursor_pos| values are only used on OS X. This function is only
// used when window rendering is disabled.
func (d *BrowserHost) ImeCommitText(text string, replacementRange *Range, relativeCursorPos int32) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_browser_host_ime_commit_text(d.toNative(), (*C.cef_string_t)(text_), replacementRange.toNative(&C.cef_range_t{}), C.int(relativeCursorPos), d.ime_commit_text)
}

// ImeFinishComposingText (ime_finish_composing_text)
// Completes the existing composition by applying the current composition node
// contents. If |keep_selection| is false (0) the current selection, if any,
// will be discarded. See comments on ImeSetComposition for usage. This
// function is only used when window rendering is disabled.
func (d *BrowserHost) ImeFinishComposingText(keepSelection int32) {
	C.gocef_browser_host_ime_finish_composing_text(d.toNative(), C.int(keepSelection), d.ime_finish_composing_text)
}

// ImeCancelComposition (ime_cancel_composition)
// Cancels the existing composition and discards the composition node contents
// without applying them. See comments on ImeSetComposition for usage. This
// function is only used when window rendering is disabled.
func (d *BrowserHost) ImeCancelComposition() {
	C.gocef_browser_host_ime_cancel_composition(d.toNative(), d.ime_cancel_composition)
}

// DragTargetDragEnter (drag_target_drag_enter)
// Call this function when the user drags the mouse into the web view (before
// calling DragTargetDragOver/DragTargetLeave/DragTargetDrop). |drag_data|
// should not contain file contents as this type of data is not allowed to be
// dragged into the web view. File contents can be removed using
// cef_drag_data_t::ResetFileContents (for example, if |drag_data| comes from
// cef_render_handler_t::StartDragging). This function is only used when
// window rendering is disabled.
func (d *BrowserHost) DragTargetDragEnter(dragData *DragData, event *MouseEvent, allowedOps DragOperationsMask) {
	C.gocef_browser_host_drag_target_drag_enter(d.toNative(), dragData.toNative(), event.toNative(&C.cef_mouse_event_t{}), C.cef_drag_operations_mask_t(allowedOps), d.drag_target_drag_enter)
}

// DragTargetDragOver (drag_target_drag_over)
// Call this function each time the mouse is moved across the web view during
// a drag operation (after calling DragTargetDragEnter and before calling
// DragTargetDragLeave/DragTargetDrop). This function is only used when window
// rendering is disabled.
func (d *BrowserHost) DragTargetDragOver(event *MouseEvent, allowedOps DragOperationsMask) {
	C.gocef_browser_host_drag_target_drag_over(d.toNative(), event.toNative(&C.cef_mouse_event_t{}), C.cef_drag_operations_mask_t(allowedOps), d.drag_target_drag_over)
}

// DragTargetDragLeave (drag_target_drag_leave)
// Call this function when the user drags the mouse out of the web view (after
// calling DragTargetDragEnter). This function is only used when window
// rendering is disabled.
func (d *BrowserHost) DragTargetDragLeave() {
	C.gocef_browser_host_drag_target_drag_leave(d.toNative(), d.drag_target_drag_leave)
}

// DragTargetDrop (drag_target_drop)
// Call this function when the user completes the drag operation by dropping
// the object onto the web view (after calling DragTargetDragEnter). The
// object being dropped is |drag_data|, given as an argument to the previous
// DragTargetDragEnter call. This function is only used when window rendering
// is disabled.
func (d *BrowserHost) DragTargetDrop(event *MouseEvent) {
	C.gocef_browser_host_drag_target_drop(d.toNative(), event.toNative(&C.cef_mouse_event_t{}), d.drag_target_drop)
}

// DragSourceEndedAt (drag_source_ended_at)
// Call this function when the drag operation started by a
// cef_render_handler_t::StartDragging call has ended either in a drop or by
// being cancelled. |x| and |y| are mouse coordinates relative to the upper-
// left corner of the view. If the web view is both the drag source and the
// drag target then all DragTarget* functions should be called before
// DragSource* mthods. This function is only used when window rendering is
// disabled.
func (d *BrowserHost) DragSourceEndedAt(x, y int32, op DragOperationsMask) {
	C.gocef_browser_host_drag_source_ended_at(d.toNative(), C.int(x), C.int(y), C.cef_drag_operations_mask_t(op), d.drag_source_ended_at)
}

// DragSourceSystemDragEnded (drag_source_system_drag_ended)
// Call this function when the drag operation started by a
// cef_render_handler_t::StartDragging call has completed. This function may
// be called immediately without first calling DragSourceEndedAt to cancel a
// drag operation. If the web view is both the drag source and the drag target
// then all DragTarget* functions should be called before DragSource* mthods.
// This function is only used when window rendering is disabled.
func (d *BrowserHost) DragSourceSystemDragEnded() {
	C.gocef_browser_host_drag_source_system_drag_ended(d.toNative(), d.drag_source_system_drag_ended)
}

// GetVisibleNavigationEntry (get_visible_navigation_entry)
// Returns the current visible navigation entry for this browser. This
// function can only be called on the UI thread.
func (d *BrowserHost) GetVisibleNavigationEntry() *NavigationEntry {
	return (*NavigationEntry)(C.gocef_browser_host_get_visible_navigation_entry(d.toNative(), d.get_visible_navigation_entry))
}

// SetAccessibilityState (set_accessibility_state)
// Set accessibility state for all frames. |accessibility_state| may be
// default, enabled or disabled. If |accessibility_state| is STATE_DEFAULT
// then accessibility will be disabled by default and the state may be further
// controlled with the "force-renderer-accessibility" and "disable-renderer-
// accessibility" command-line switches. If |accessibility_state| is
// STATE_ENABLED then accessibility will be enabled. If |accessibility_state|
// is STATE_DISABLED then accessibility will be completely disabled.
//
// For windowed browsers accessibility will be enabled in Complete mode (which
// corresponds to kAccessibilityModeComplete in Chromium). In this mode all
// platform accessibility objects will be created and managed by Chromium's
// internal implementation. The client needs only to detect the screen reader
// and call this function appropriately. For example, on macOS the client can
// handle the @"AXEnhancedUserStructure" accessibility attribute to detect
// VoiceOver state changes and on Windows the client can handle WM_GETOBJECT
// with OBJID_CLIENT to detect accessibility readers.
//
// For windowless browsers accessibility will be enabled in TreeOnly mode
// (which corresponds to kAccessibilityModeWebContentsOnly in Chromium). In
// this mode renderer accessibility is enabled, the full tree is computed, and
// events are passed to CefAccessibiltyHandler, but platform accessibility
// objects are not created. The client may implement platform accessibility
// objects using CefAccessibiltyHandler callbacks if desired.
func (d *BrowserHost) SetAccessibilityState(accessibilityState State) {
	C.gocef_browser_host_set_accessibility_state(d.toNative(), C.cef_state_t(accessibilityState), d.set_accessibility_state)
}

// SetAutoResizeEnabled (set_auto_resize_enabled)
// Enable notifications of auto resize via
// cef_display_handler_t::OnAutoResize. Notifications are disabled by default.
// |min_size| and |max_size| define the range of allowed sizes.
func (d *BrowserHost) SetAutoResizeEnabled(enabled int32, minSize, maxSize *Size) {
	C.gocef_browser_host_set_auto_resize_enabled(d.toNative(), C.int(enabled), minSize.toNative(&C.cef_size_t{}), maxSize.toNative(&C.cef_size_t{}), d.set_auto_resize_enabled)
}

// GetExtension (get_extension)
// Returns the extension hosted in this browser or NULL if no extension is
// hosted. See cef_request_tContext::LoadExtension for details.
func (d *BrowserHost) GetExtension() *Extension {
	return (*Extension)(C.gocef_browser_host_get_extension(d.toNative(), d.get_extension))
}

// IsBackgroundHost (is_background_host)
// Returns true (1) if this browser is hosting an extension background script.
// Background hosts do not have a window and are not displayable. See
// cef_request_tContext::LoadExtension for details.
func (d *BrowserHost) IsBackgroundHost() int32 {
	return int32(C.gocef_browser_host_is_background_host(d.toNative(), d.is_background_host))
}

// SetAudioMuted (set_audio_muted)
//  Set whether the browser's audio is muted.
func (d *BrowserHost) SetAudioMuted(mute int32) {
	C.gocef_browser_host_set_audio_muted(d.toNative(), C.int(mute), d.set_audio_muted)
}

// IsAudioMuted (is_audio_muted)
// Returns true (1) if the browser's audio is muted.  This function can only
// be called on the UI thread.
func (d *BrowserHost) IsAudioMuted() int32 {
	return int32(C.gocef_browser_host_is_audio_muted(d.toNative(), d.is_audio_muted))
}
