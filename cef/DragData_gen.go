// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_drag_data_t * gocef_drag_data_clone(cef_drag_data_t * self, cef_drag_data_t * (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// int gocef_drag_data_is_read_only(cef_drag_data_t * self, int (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// int gocef_drag_data_is_link(cef_drag_data_t * self, int (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// int gocef_drag_data_is_fragment(cef_drag_data_t * self, int (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// int gocef_drag_data_is_file(cef_drag_data_t * self, int (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_link_url(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_link_title(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_link_metadata(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_fragment_text(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_fragment_html(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_fragment_base_url(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_drag_data_get_file_name(cef_drag_data_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// size_t gocef_drag_data_get_file_contents(cef_drag_data_t * self, cef_stream_writer_t * writer, size_t (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_stream_writer_t *)) { return callback__(self, writer); }
	// int gocef_drag_data_get_file_names(cef_drag_data_t * self, cef_string_list_t names, int (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_list_t)) { return callback__(self, names); }
	// void gocef_drag_data_set_link_url(cef_drag_data_t * self, cef_string_t * uRL, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, uRL); }
	// void gocef_drag_data_set_link_title(cef_drag_data_t * self, cef_string_t * title, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, title); }
	// void gocef_drag_data_set_link_metadata(cef_drag_data_t * self, cef_string_t * data, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, data); }
	// void gocef_drag_data_set_fragment_text(cef_drag_data_t * self, cef_string_t * text, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, text); }
	// void gocef_drag_data_set_fragment_html(cef_drag_data_t * self, cef_string_t * hTML, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, hTML); }
	// void gocef_drag_data_set_fragment_base_url(cef_drag_data_t * self, cef_string_t * baseURL, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *)) { return callback__(self, baseURL); }
	// void gocef_drag_data_reset_file_contents(cef_drag_data_t * self, void (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// void gocef_drag_data_add_file(cef_drag_data_t * self, cef_string_t * path, cef_string_t * displayName, void (CEF_CALLBACK *callback__)(cef_drag_data_t *, cef_string_t *, cef_string_t *)) { return callback__(self, path, displayName); }
	// cef_image_t * gocef_drag_data_get_image(cef_drag_data_t * self, cef_image_t * (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// cef_point_t gocef_drag_data_get_image_hotspot(cef_drag_data_t * self, cef_point_t (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	// int gocef_drag_data_has_image(cef_drag_data_t * self, int (CEF_CALLBACK *callback__)(cef_drag_data_t *)) { return callback__(self); }
	"C"
)

// DragData (cef_drag_data_t from include/capi/cef_drag_data_capi.h)
// Structure used to represent drag data. The functions of this structure may be
// called on any thread.
type DragData C.cef_drag_data_t

func (d *DragData) toNative() *C.cef_drag_data_t {
	return (*C.cef_drag_data_t)(d)
}

// Base (base)
// Base structure.
func (d *DragData) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Clone (clone)
// Returns a copy of the current object.
func (d *DragData) Clone() *DragData {
	return (*DragData)(C.gocef_drag_data_clone(d.toNative(), d.clone))
}

// IsReadOnly (is_read_only)
// Returns true (1) if this object is read-only.
func (d *DragData) IsReadOnly() int32 {
	return int32(C.gocef_drag_data_is_read_only(d.toNative(), d.is_read_only))
}

// IsLink (is_link)
// Returns true (1) if the drag data is a link.
func (d *DragData) IsLink() int32 {
	return int32(C.gocef_drag_data_is_link(d.toNative(), d.is_link))
}

// IsFragment (is_fragment)
// Returns true (1) if the drag data is a text or html fragment.
func (d *DragData) IsFragment() int32 {
	return int32(C.gocef_drag_data_is_fragment(d.toNative(), d.is_fragment))
}

// IsFile (is_file)
// Returns true (1) if the drag data is a file.
func (d *DragData) IsFile() int32 {
	return int32(C.gocef_drag_data_is_file(d.toNative(), d.is_file))
}

// GetLinkURL (get_link_url)
// Return the link URL that is being dragged.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetLinkURL() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_link_url(d.toNative(), d.get_link_url))
}

// GetLinkTitle (get_link_title)
// Return the title associated with the link being dragged.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetLinkTitle() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_link_title(d.toNative(), d.get_link_title))
}

// GetLinkMetadata (get_link_metadata)
// Return the metadata, if any, associated with the link being dragged.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetLinkMetadata() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_link_metadata(d.toNative(), d.get_link_metadata))
}

// GetFragmentText (get_fragment_text)
// Return the plain text fragment that is being dragged.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetFragmentText() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_fragment_text(d.toNative(), d.get_fragment_text))
}

// GetFragmentHTML (get_fragment_html)
// Return the text/html fragment that is being dragged.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetFragmentHTML() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_fragment_html(d.toNative(), d.get_fragment_html))
}

// GetFragmentBaseURL (get_fragment_base_url)
// Return the base URL that the fragment came from. This value is used for
// resolving relative URLs and may be NULL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetFragmentBaseURL() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_fragment_base_url(d.toNative(), d.get_fragment_base_url))
}

// GetFileName (get_file_name)
// Return the name of the file being dragged out of the browser window.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *DragData) GetFileName() string {
	return cefuserfreestrToString(C.gocef_drag_data_get_file_name(d.toNative(), d.get_file_name))
}

// GetFileContents (get_file_contents)
// Write the contents of the file being dragged out of the web view into
// |writer|. Returns the number of bytes sent to |writer|. If |writer| is NULL
// this function will return the size of the file contents in bytes. Call
// get_file_name() to get a suggested name for the file.
func (d *DragData) GetFileContents(writer *StreamWriter) uint64 {
	return uint64(C.gocef_drag_data_get_file_contents(d.toNative(), writer.toNative(), d.get_file_contents))
}

// GetFileNames (get_file_names)
// Retrieve the list of file names that are being dragged into the browser
// window.
func (d *DragData) GetFileNames(names StringList) int32 {
	return int32(C.gocef_drag_data_get_file_names(d.toNative(), C.cef_string_list_t(names), d.get_file_names))
}

// SetLinkURL (set_link_url)
// Set the link URL that is being dragged.
func (d *DragData) SetLinkURL(uRL string) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	C.gocef_drag_data_set_link_url(d.toNative(), (*C.cef_string_t)(uRL_), d.set_link_url)
}

// SetLinkTitle (set_link_title)
// Set the title associated with the link being dragged.
func (d *DragData) SetLinkTitle(title string) {
	title_ := C.cef_string_userfree_alloc()
	setCEFStr(title, title_)
	defer func() {
		C.cef_string_userfree_free(title_)
	}()
	C.gocef_drag_data_set_link_title(d.toNative(), (*C.cef_string_t)(title_), d.set_link_title)
}

// SetLinkMetadata (set_link_metadata)
// Set the metadata associated with the link being dragged.
func (d *DragData) SetLinkMetadata(data string) {
	data_ := C.cef_string_userfree_alloc()
	setCEFStr(data, data_)
	defer func() {
		C.cef_string_userfree_free(data_)
	}()
	C.gocef_drag_data_set_link_metadata(d.toNative(), (*C.cef_string_t)(data_), d.set_link_metadata)
}

// SetFragmentText (set_fragment_text)
// Set the plain text fragment that is being dragged.
func (d *DragData) SetFragmentText(text string) {
	text_ := C.cef_string_userfree_alloc()
	setCEFStr(text, text_)
	defer func() {
		C.cef_string_userfree_free(text_)
	}()
	C.gocef_drag_data_set_fragment_text(d.toNative(), (*C.cef_string_t)(text_), d.set_fragment_text)
}

// SetFragmentHTML (set_fragment_html)
// Set the text/html fragment that is being dragged.
func (d *DragData) SetFragmentHTML(hTML string) {
	hTML_ := C.cef_string_userfree_alloc()
	setCEFStr(hTML, hTML_)
	defer func() {
		C.cef_string_userfree_free(hTML_)
	}()
	C.gocef_drag_data_set_fragment_html(d.toNative(), (*C.cef_string_t)(hTML_), d.set_fragment_html)
}

// SetFragmentBaseURL (set_fragment_base_url)
// Set the base URL that the fragment came from.
func (d *DragData) SetFragmentBaseURL(baseURL string) {
	baseURL_ := C.cef_string_userfree_alloc()
	setCEFStr(baseURL, baseURL_)
	defer func() {
		C.cef_string_userfree_free(baseURL_)
	}()
	C.gocef_drag_data_set_fragment_base_url(d.toNative(), (*C.cef_string_t)(baseURL_), d.set_fragment_base_url)
}

// ResetFileContents (reset_file_contents)
// Reset the file contents. You should do this before calling
// cef_browser_host_t::DragTargetDragEnter as the web view does not allow us
// to drag in this kind of data.
func (d *DragData) ResetFileContents() {
	C.gocef_drag_data_reset_file_contents(d.toNative(), d.reset_file_contents)
}

// AddFile (add_file)
// Add a file that is being dragged into the webview.
func (d *DragData) AddFile(path, displayName string) {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	displayName_ := C.cef_string_userfree_alloc()
	setCEFStr(displayName, displayName_)
	defer func() {
		C.cef_string_userfree_free(displayName_)
	}()
	C.gocef_drag_data_add_file(d.toNative(), (*C.cef_string_t)(path_), (*C.cef_string_t)(displayName_), d.add_file)
}

// GetImage (get_image)
// Get the image representation of drag data. May return NULL if no image
// representation is available.
func (d *DragData) GetImage() *Image {
	return (*Image)(C.gocef_drag_data_get_image(d.toNative(), d.get_image))
}

// GetImageHotspot (get_image_hotspot)
// Get the image hotspot (drag start location relative to image dimensions).
func (d *DragData) GetImageHotspot() Point {
	cresult__ := C.gocef_drag_data_get_image_hotspot(d.toNative(), d.get_image_hotspot)
	var result__ Point
	(&cresult__).intoGo(&result__)
	return result__
}

// HasImage (has_image)
// Returns true (1) if an image representation of drag data is available.
func (d *DragData) HasImage() int32 {
	return int32(C.gocef_drag_data_has_image(d.toNative(), d.has_image))
}
