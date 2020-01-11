// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_cookie_manager_set_supported_schemes(cef_cookie_manager_t * self, cef_string_list_t schemes, int includeDefaults, cef_completion_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_list_t, int, cef_completion_callback_t *)) { return callback__(self, schemes, includeDefaults, callback); }
	// int gocef_cookie_manager_visit_all_cookies(cef_cookie_manager_t * self, cef_cookie_visitor_t * visitor, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_cookie_visitor_t *)) { return callback__(self, visitor); }
	// int gocef_cookie_manager_visit_url_cookies(cef_cookie_manager_t * self, cef_string_t * uRL, int includeHTTPOnly, cef_cookie_visitor_t * visitor, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, int, cef_cookie_visitor_t *)) { return callback__(self, uRL, includeHTTPOnly, visitor); }
	// int gocef_cookie_manager_set_cookie(cef_cookie_manager_t * self, cef_string_t * uRL, cef_cookie_t * cookie, cef_set_cookie_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, cef_cookie_t *, cef_set_cookie_callback_t *)) { return callback__(self, uRL, cookie, callback); }
	// int gocef_cookie_manager_delete_cookies(cef_cookie_manager_t * self, cef_string_t * uRL, cef_string_t * cookieName, cef_delete_cookies_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, cef_string_t *, cef_delete_cookies_callback_t *)) { return callback__(self, uRL, cookieName, callback); }
	// int gocef_cookie_manager_flush_store(cef_cookie_manager_t * self, cef_completion_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_completion_callback_t *)) { return callback__(self, callback); }
	"C"
)

// CookieManager (cef_cookie_manager_t from include/capi/cef_cookie_capi.h)
// Structure used for managing cookies. The functions of this structure may be
// called on any thread unless otherwise indicated.
type CookieManager C.cef_cookie_manager_t

func (d *CookieManager) toNative() *C.cef_cookie_manager_t {
	return (*C.cef_cookie_manager_t)(d)
}

// Base (base)
// Base structure.
func (d *CookieManager) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// SetSupportedSchemes (set_supported_schemes)
// Set the schemes supported by this manager. If |include_defaults| is true
// (1) the default schemes ("http", "https", "ws" and "wss") will also be
// supported. Calling this function with an NULL |schemes| value and
// |include_defaults| set to false (0) will disable all loading and saving of
// cookies for this manager. If |callback| is non-NULL it will be executed
// asnychronously on the UI thread after the change has been applied. Must be
// called before any cookies are accessed.
func (d *CookieManager) SetSupportedSchemes(schemes StringList, includeDefaults int32, callback *CompletionCallback) {
	C.gocef_cookie_manager_set_supported_schemes(d.toNative(), C.cef_string_list_t(schemes), C.int(includeDefaults), callback.toNative(), d.set_supported_schemes)
}

// VisitAllCookies (visit_all_cookies)
// Visit all cookies on the UI thread. The returned cookies are ordered by
// longest path, then by earliest creation date. Returns false (0) if cookies
// cannot be accessed.
func (d *CookieManager) VisitAllCookies(visitor *CookieVisitor) int32 {
	return int32(C.gocef_cookie_manager_visit_all_cookies(d.toNative(), visitor.toNative(), d.visit_all_cookies))
}

// VisitURLCookies (visit_url_cookies)
// Visit a subset of cookies on the UI thread. The results are filtered by the
// given url scheme, host, domain and path. If |includeHttpOnly| is true (1)
// HTTP-only cookies will also be included in the results. The returned
// cookies are ordered by longest path, then by earliest creation date.
// Returns false (0) if cookies cannot be accessed.
func (d *CookieManager) VisitURLCookies(uRL string, includeHTTPOnly int32, visitor *CookieVisitor) int32 {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	return int32(C.gocef_cookie_manager_visit_url_cookies(d.toNative(), (*C.cef_string_t)(uRL_), C.int(includeHTTPOnly), visitor.toNative(), d.visit_url_cookies))
}

// SetCookie (set_cookie)
// Sets a cookie given a valid URL and explicit user-provided cookie
// attributes. This function expects each attribute to be well-formed. It will
// check for disallowed characters (e.g. the ';' character is disallowed
// within the cookie value attribute) and fail without setting the cookie if
// such characters are found. If |callback| is non-NULL it will be executed
// asnychronously on the UI thread after the cookie has been set. Returns
// false (0) if an invalid URL is specified or if cookies cannot be accessed.
func (d *CookieManager) SetCookie(uRL string, cookie *Cookie, callback *SetCookieCallback) int32 {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	return int32(C.gocef_cookie_manager_set_cookie(d.toNative(), (*C.cef_string_t)(uRL_), cookie.toNative(&C.cef_cookie_t{}), callback.toNative(), d.set_cookie))
}

// DeleteCookies (delete_cookies)
// Delete all cookies that match the specified parameters. If both |url| and
// |cookie_name| values are specified all host and domain cookies matching
// both will be deleted. If only |url| is specified all host cookies (but not
// domain cookies) irrespective of path will be deleted. If |url| is NULL all
// cookies for all hosts and domains will be deleted. If |callback| is non-
// NULL it will be executed asnychronously on the UI thread after the cookies
// have been deleted. Returns false (0) if a non-NULL invalid URL is specified
// or if cookies cannot be accessed. Cookies can alternately be deleted using
// the Visit*Cookies() functions.
func (d *CookieManager) DeleteCookies(uRL, cookieName string, callback *DeleteCookiesCallback) int32 {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	cookieName_ := C.cef_string_userfree_alloc()
	setCEFStr(cookieName, cookieName_)
	defer func() {
		C.cef_string_userfree_free(cookieName_)
	}()
	return int32(C.gocef_cookie_manager_delete_cookies(d.toNative(), (*C.cef_string_t)(uRL_), (*C.cef_string_t)(cookieName_), callback.toNative(), d.delete_cookies))
}

// FlushStore (flush_store)
// Flush the backing store (if any) to disk. If |callback| is non-NULL it will
// be executed asnychronously on the UI thread after the flush is complete.
// Returns false (0) if cookies cannot be accessed.
func (d *CookieManager) FlushStore(callback *CompletionCallback) int32 {
	return int32(C.gocef_cookie_manager_flush_store(d.toNative(), callback.toNative(), d.flush_store))
}
