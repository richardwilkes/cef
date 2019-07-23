// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// void gocef_cookie_manager_set_supported_schemes(cef_cookie_manager_t * self, cef_string_list_t schemes, cef_completion_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_list_t, cef_completion_callback_t *)) { return callback__(self, schemes, callback); }
	// int gocef_cookie_manager_visit_all_cookies(cef_cookie_manager_t * self, cef_cookie_visitor_t * visitor, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_cookie_visitor_t *)) { return callback__(self, visitor); }
	// int gocef_cookie_manager_visit_url_cookies(cef_cookie_manager_t * self, cef_string_t * url, int includeHttpOnly, cef_cookie_visitor_t * visitor, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, int, cef_cookie_visitor_t *)) { return callback__(self, url, includeHttpOnly, visitor); }
	// int gocef_cookie_manager_set_cookie(cef_cookie_manager_t * self, cef_string_t * url, cef_cookie_t * cookie, cef_set_cookie_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, cef_cookie_t *, cef_set_cookie_callback_t *)) { return callback__(self, url, cookie, callback); }
	// int gocef_cookie_manager_delete_cookies(cef_cookie_manager_t * self, cef_string_t * url, cef_string_t * cookie_name, cef_delete_cookies_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, cef_string_t *, cef_delete_cookies_callback_t *)) { return callback__(self, url, cookie_name, callback); }
	// int gocef_cookie_manager_set_storage_path(cef_cookie_manager_t * self, cef_string_t * path, int persist_session_cookies, cef_completion_callback_t * callback, int (CEF_CALLBACK *callback__)(cef_cookie_manager_t *, cef_string_t *, int, cef_completion_callback_t *)) { return callback__(self, path, persist_session_cookies, callback); }
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
// Set the schemes supported by this manager. The default schemes ("http",
// "https", "ws" and "wss") will always be supported. If |callback| is non-
// NULL it will be executed asnychronously on the IO thread after the change
// has been applied. Must be called before any cookies are accessed.
func (d *CookieManager) SetSupportedSchemes(schemes StringList, callback *CompletionCallback) {
	C.gocef_cookie_manager_set_supported_schemes(d.toNative(), C.cef_string_list_t(schemes), callback.toNative(), d.set_supported_schemes)
}

// VisitAllCookies (visit_all_cookies)
// Visit all cookies on the IO thread. The returned cookies are ordered by
// longest path, then by earliest creation date. Returns false (0) if cookies
// cannot be accessed.
func (d *CookieManager) VisitAllCookies(visitor *CookieVisitor) int32 {
	return int32(C.gocef_cookie_manager_visit_all_cookies(d.toNative(), visitor.toNative(), d.visit_all_cookies))
}

// VisitUrlCookies (visit_url_cookies)
// Visit a subset of cookies on the IO thread. The results are filtered by the
// given url scheme, host, domain and path. If |includeHttpOnly| is true (1)
// HTTP-only cookies will also be included in the results. The returned
// cookies are ordered by longest path, then by earliest creation date.
// Returns false (0) if cookies cannot be accessed.
func (d *CookieManager) VisitUrlCookies(url string, includeHttpOnly int32, visitor *CookieVisitor) int32 {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	return int32(C.gocef_cookie_manager_visit_url_cookies(d.toNative(), (*C.cef_string_t)(url_), C.int(includeHttpOnly), visitor.toNative(), d.visit_url_cookies))
}

// SetCookie (set_cookie)
// Sets a cookie given a valid URL and explicit user-provided cookie
// attributes. This function expects each attribute to be well-formed. It will
// check for disallowed characters (e.g. the ';' character is disallowed
// within the cookie value attribute) and fail without setting the cookie if
// such characters are found. If |callback| is non-NULL it will be executed
// asnychronously on the IO thread after the cookie has been set. Returns
// false (0) if an invalid URL is specified or if cookies cannot be accessed.
func (d *CookieManager) SetCookie(url string, cookie *Cookie, callback *SetCookieCallback) int32 {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	return int32(C.gocef_cookie_manager_set_cookie(d.toNative(), (*C.cef_string_t)(url_), cookie.toNative(&C.cef_cookie_t{}), callback.toNative(), d.set_cookie))
}

// DeleteCookies (delete_cookies)
// Delete all cookies that match the specified parameters. If both |url| and
// |cookie_name| values are specified all host and domain cookies matching
// both will be deleted. If only |url| is specified all host cookies (but not
// domain cookies) irrespective of path will be deleted. If |url| is NULL all
// cookies for all hosts and domains will be deleted. If |callback| is non-
// NULL it will be executed asnychronously on the IO thread after the cookies
// have been deleted. Returns false (0) if a non-NULL invalid URL is specified
// or if cookies cannot be accessed. Cookies can alternately be deleted using
// the Visit*Cookies() functions.
func (d *CookieManager) DeleteCookies(url, cookie_name string, callback *DeleteCookiesCallback) int32 {
	url_ := C.cef_string_userfree_alloc()
	setCEFStr(url, url_)
	defer func() {
		C.cef_string_userfree_free(url_)
	}()
	cookie_name_ := C.cef_string_userfree_alloc()
	setCEFStr(cookie_name, cookie_name_)
	defer func() {
		C.cef_string_userfree_free(cookie_name_)
	}()
	return int32(C.gocef_cookie_manager_delete_cookies(d.toNative(), (*C.cef_string_t)(url_), (*C.cef_string_t)(cookie_name_), callback.toNative(), d.delete_cookies))
}

// SetStoragePath (set_storage_path)
// Sets the directory path that will be used for storing cookie data. If
// |path| is NULL data will be stored in memory only. Otherwise, data will be
// stored at the specified |path|. To persist session cookies (cookies without
// an expiry date or validity interval) set |persist_session_cookies| to true
// (1). Session cookies are generally intended to be transient and most Web
// browsers do not persist them. If |callback| is non-NULL it will be executed
// asnychronously on the IO thread after the manager's storage has been
// initialized. Returns false (0) if cookies cannot be accessed.
func (d *CookieManager) SetStoragePath(path string, persist_session_cookies int32, callback *CompletionCallback) int32 {
	path_ := C.cef_string_userfree_alloc()
	setCEFStr(path, path_)
	defer func() {
		C.cef_string_userfree_free(path_)
	}()
	return int32(C.gocef_cookie_manager_set_storage_path(d.toNative(), (*C.cef_string_t)(path_), C.int(persist_session_cookies), callback.toNative(), d.set_storage_path))
}

// FlushStore (flush_store)
// Flush the backing store (if any) to disk. If |callback| is non-NULL it will
// be executed asnychronously on the IO thread after the flush is complete.
// Returns false (0) if cookies cannot be accessed.
func (d *CookieManager) FlushStore(callback *CompletionCallback) int32 {
	return int32(C.gocef_cookie_manager_flush_store(d.toNative(), callback.toNative(), d.flush_store))
}
