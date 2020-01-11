// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_response_is_read_only(cef_response_t * self, int (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// cef_errorcode_t gocef_response_get_error(cef_response_t * self, cef_errorcode_t (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_error(cef_response_t * self, cef_errorcode_t error_r, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_errorcode_t)) { return callback__(self, error_r); }
	// int gocef_response_get_status(cef_response_t * self, int (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_status(cef_response_t * self, int status, void (CEF_CALLBACK *callback__)(cef_response_t *, int)) { return callback__(self, status); }
	// cef_string_userfree_t gocef_response_get_status_text(cef_response_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_status_text(cef_response_t * self, cef_string_t * statusText, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_t *)) { return callback__(self, statusText); }
	// cef_string_userfree_t gocef_response_get_mime_type(cef_response_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_mime_type(cef_response_t * self, cef_string_t * mimeType, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_t *)) { return callback__(self, mimeType); }
	// cef_string_userfree_t gocef_response_get_charset(cef_response_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_charset(cef_response_t * self, cef_string_t * charset, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_t *)) { return callback__(self, charset); }
	// cef_string_userfree_t gocef_response_get_header(cef_response_t * self, cef_string_t * name, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_t *)) { return callback__(self, name); }
	// void gocef_response_get_header_map(cef_response_t * self, cef_string_multimap_t headerMap, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_multimap_t)) { return callback__(self, headerMap); }
	// void gocef_response_set_header_map(cef_response_t * self, cef_string_multimap_t headerMap, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_multimap_t)) { return callback__(self, headerMap); }
	// cef_string_userfree_t gocef_response_get_url(cef_response_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_response_t *)) { return callback__(self); }
	// void gocef_response_set_url(cef_response_t * self, cef_string_t * uRL, void (CEF_CALLBACK *callback__)(cef_response_t *, cef_string_t *)) { return callback__(self, uRL); }
	"C"
)

// Response (cef_response_t from include/capi/cef_response_capi.h)
// Structure used to represent a web response. The functions of this structure
// may be called on any thread.
type Response C.cef_response_t

func (d *Response) toNative() *C.cef_response_t {
	return (*C.cef_response_t)(d)
}

// Base (base)
// Base structure.
func (d *Response) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsReadOnly (is_read_only)
// Returns true (1) if this object is read-only.
func (d *Response) IsReadOnly() int32 {
	return int32(C.gocef_response_is_read_only(d.toNative(), d.is_read_only))
}

// GetError (get_error)
// Get the response error code. Returns ERR_NONE if there was no error.
func (d *Response) GetError() Errorcode {
	return Errorcode(C.gocef_response_get_error(d.toNative(), d.get_error))
}

// SetError (set_error)
// Set the response error code. This can be used by custom scheme handlers to
// return errors during initial request processing.
func (d *Response) SetError(error_r Errorcode) {
	C.gocef_response_set_error(d.toNative(), C.cef_errorcode_t(error_r), d.set_error)
}

// GetStatus (get_status)
// Get the response status code.
func (d *Response) GetStatus() int32 {
	return int32(C.gocef_response_get_status(d.toNative(), d.get_status))
}

// SetStatus (set_status)
// Set the response status code.
func (d *Response) SetStatus(status int32) {
	C.gocef_response_set_status(d.toNative(), C.int(status), d.set_status)
}

// GetStatusText (get_status_text)
// Get the response status text.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Response) GetStatusText() string {
	return cefuserfreestrToString(C.gocef_response_get_status_text(d.toNative(), d.get_status_text))
}

// SetStatusText (set_status_text)
// Set the response status text.
func (d *Response) SetStatusText(statusText string) {
	statusText_ := C.cef_string_userfree_alloc()
	setCEFStr(statusText, statusText_)
	defer func() {
		C.cef_string_userfree_free(statusText_)
	}()
	C.gocef_response_set_status_text(d.toNative(), (*C.cef_string_t)(statusText_), d.set_status_text)
}

// GetMimeType (get_mime_type)
// Get the response mime type.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Response) GetMimeType() string {
	return cefuserfreestrToString(C.gocef_response_get_mime_type(d.toNative(), d.get_mime_type))
}

// SetMimeType (set_mime_type)
// Set the response mime type.
func (d *Response) SetMimeType(mimeType string) {
	mimeType_ := C.cef_string_userfree_alloc()
	setCEFStr(mimeType, mimeType_)
	defer func() {
		C.cef_string_userfree_free(mimeType_)
	}()
	C.gocef_response_set_mime_type(d.toNative(), (*C.cef_string_t)(mimeType_), d.set_mime_type)
}

// GetCharset (get_charset)
// Get the response charset.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Response) GetCharset() string {
	return cefuserfreestrToString(C.gocef_response_get_charset(d.toNative(), d.get_charset))
}

// SetCharset (set_charset)
// Set the response charset.
func (d *Response) SetCharset(charset string) {
	charset_ := C.cef_string_userfree_alloc()
	setCEFStr(charset, charset_)
	defer func() {
		C.cef_string_userfree_free(charset_)
	}()
	C.gocef_response_set_charset(d.toNative(), (*C.cef_string_t)(charset_), d.set_charset)
}

// GetHeader (get_header)
// Get the value for the specified response header field.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Response) GetHeader(name string) string {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return cefuserfreestrToString(C.gocef_response_get_header(d.toNative(), (*C.cef_string_t)(name_), d.get_header))
}

// GetHeaderMap (get_header_map)
// Get all response header fields.
func (d *Response) GetHeaderMap(headerMap StringMultimap) {
	C.gocef_response_get_header_map(d.toNative(), C.cef_string_multimap_t(headerMap), d.get_header_map)
}

// SetHeaderMap (set_header_map)
// Set all response header fields.
func (d *Response) SetHeaderMap(headerMap StringMultimap) {
	C.gocef_response_set_header_map(d.toNative(), C.cef_string_multimap_t(headerMap), d.set_header_map)
}

// GetURL (get_url)
// Get the resolved URL after redirects or changed as a result of HSTS.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Response) GetURL() string {
	return cefuserfreestrToString(C.gocef_response_get_url(d.toNative(), d.get_url))
}

// SetURL (set_url)
// Set the resolved URL after redirects or changed as a result of HSTS.
func (d *Response) SetURL(uRL string) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	C.gocef_response_set_url(d.toNative(), (*C.cef_string_t)(uRL_), d.set_url)
}
