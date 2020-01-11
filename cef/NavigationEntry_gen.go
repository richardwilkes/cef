// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_navigation_entry_is_valid(cef_navigation_entry_t * self, int (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_navigation_entry_get_url(cef_navigation_entry_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_navigation_entry_get_display_url(cef_navigation_entry_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_navigation_entry_get_original_url(cef_navigation_entry_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_navigation_entry_get_title(cef_navigation_entry_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_transition_type_t gocef_navigation_entry_get_transition_type(cef_navigation_entry_t * self, cef_transition_type_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// int gocef_navigation_entry_has_post_data(cef_navigation_entry_t * self, int (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_time_t gocef_navigation_entry_get_completion_time(cef_navigation_entry_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// int gocef_navigation_entry_get_http_status_code(cef_navigation_entry_t * self, int (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	// cef_sslstatus_t * gocef_navigation_entry_get_sslstatus(cef_navigation_entry_t * self, cef_sslstatus_t * (CEF_CALLBACK *callback__)(cef_navigation_entry_t *)) { return callback__(self); }
	"C"
)

// NavigationEntry (cef_navigation_entry_t from include/capi/cef_navigation_entry_capi.h)
// Structure used to represent an entry in navigation history.
type NavigationEntry C.cef_navigation_entry_t

func (d *NavigationEntry) toNative() *C.cef_navigation_entry_t {
	return (*C.cef_navigation_entry_t)(d)
}

// Base (base)
// Base structure.
func (d *NavigationEntry) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsValid (is_valid)
// Returns true (1) if this object is valid. Do not call any other functions
// if this function returns false (0).
func (d *NavigationEntry) IsValid() int32 {
	return int32(C.gocef_navigation_entry_is_valid(d.toNative(), d.is_valid))
}

// GetURL (get_url)
// Returns the actual URL of the page. For some pages this may be data: URL or
// similar. Use get_display_url() to return a display-friendly version.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *NavigationEntry) GetURL() string {
	return cefuserfreestrToString(C.gocef_navigation_entry_get_url(d.toNative(), d.get_url))
}

// GetDisplayURL (get_display_url)
// Returns a display-friendly version of the URL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *NavigationEntry) GetDisplayURL() string {
	return cefuserfreestrToString(C.gocef_navigation_entry_get_display_url(d.toNative(), d.get_display_url))
}

// GetOriginalURL (get_original_url)
// Returns the original URL that was entered by the user before any redirects.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *NavigationEntry) GetOriginalURL() string {
	return cefuserfreestrToString(C.gocef_navigation_entry_get_original_url(d.toNative(), d.get_original_url))
}

// GetTitle (get_title)
// Returns the title set by the page. This value may be NULL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *NavigationEntry) GetTitle() string {
	return cefuserfreestrToString(C.gocef_navigation_entry_get_title(d.toNative(), d.get_title))
}

// GetTransitionType (get_transition_type)
// Returns the transition type which indicates what the user did to move to
// this page from the previous page.
func (d *NavigationEntry) GetTransitionType() TransitionType {
	return TransitionType(C.gocef_navigation_entry_get_transition_type(d.toNative(), d.get_transition_type))
}

// HasPostData (has_post_data)
// Returns true (1) if this navigation includes post data.
func (d *NavigationEntry) HasPostData() int32 {
	return int32(C.gocef_navigation_entry_has_post_data(d.toNative(), d.has_post_data))
}

// GetCompletionTime (get_completion_time)
// Returns the time for the last known successful navigation completion. A
// navigation may be completed more than once if the page is reloaded. May be
// 0 if the navigation has not yet completed.
func (d *NavigationEntry) GetCompletionTime() Time {
	cresult__ := C.gocef_navigation_entry_get_completion_time(d.toNative(), d.get_completion_time)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetHTTPStatusCode (get_http_status_code)
// Returns the HTTP status code for the last known successful navigation
// response. May be 0 if the response has not yet been received or if the
// navigation has not yet completed.
func (d *NavigationEntry) GetHTTPStatusCode() int32 {
	return int32(C.gocef_navigation_entry_get_http_status_code(d.toNative(), d.get_http_status_code))
}

// GetSslstatus (get_sslstatus)
// Returns the SSL information for this navigation entry.
func (d *NavigationEntry) GetSslstatus() *Sslstatus {
	return (*Sslstatus)(C.gocef_navigation_entry_get_sslstatus(d.toNative(), d.get_sslstatus))
}
