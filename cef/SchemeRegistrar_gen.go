// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	// int gocef_scheme_registrar_add_custom_scheme(cef_scheme_registrar_t * self, cef_string_t * scheme_name, int is_standard, int is_local, int is_display_isolated, int is_secure, int is_cors_enabled, int is_csp_bypassing, int (CEF_CALLBACK *callback__)(cef_scheme_registrar_t *, cef_string_t *, int, int, int, int, int, int)) { return callback__(self, scheme_name, is_standard, is_local, is_display_isolated, is_secure, is_cors_enabled, is_csp_bypassing); }
	"C"
)

// SchemeRegistrar (cef_scheme_registrar_t from include/capi/cef_scheme_capi.h)
// Structure that manages custom scheme registrations.
type SchemeRegistrar C.cef_scheme_registrar_t

func (d *SchemeRegistrar) toNative() *C.cef_scheme_registrar_t {
	return (*C.cef_scheme_registrar_t)(d)
}

// Base (base)
// Base structure.
func (d *SchemeRegistrar) Base() *BaseScoped {
	return (*BaseScoped)(&d.base)
}

// AddCustomScheme (add_custom_scheme)
// Register a custom scheme. This function should not be called for the built-
// in HTTP, HTTPS, FILE, FTP, ABOUT and DATA schemes.
//
// If |is_standard| is true (1) the scheme will be treated as a standard
// scheme. Standard schemes are subject to URL canonicalization and parsing
// rules as defined in the Common Internet Scheme Syntax RFC 1738 Section 3.1
// available at http://www.ietf.org/rfc/rfc1738.txt
//
// In particular, the syntax for standard scheme URLs must be of the form:
// <pre>
//  [scheme]://[username]:[password]@[host]:[port]/[url-path]
// </pre> Standard scheme URLs must have a host component that is a fully
// qualified domain name as defined in Section 3.5 of RFC 1034 [13] and
// Section 2.1 of RFC 1123. These URLs will be canonicalized to
// "scheme://host/path" in the simplest case and
// "scheme://username:password@host:port/path" in the most explicit case. For
// example, "scheme:host/path" and "scheme:///host/path" will both be
// canonicalized to "scheme://host/path". The origin of a standard scheme URL
// is the combination of scheme, host and port (i.e., "scheme://host:port" in
// the most explicit case).
//
// For non-standard scheme URLs only the "scheme:" component is parsed and
// canonicalized. The remainder of the URL will be passed to the handler as-
// is. For example, "scheme:///some%20text" will remain the same. Non-standard
// scheme URLs cannot be used as a target for form submission.
//
// If |is_local| is true (1) the scheme will be treated with the same security
// rules as those applied to "file" URLs. Normal pages cannot link to or
// access local URLs. Also, by default, local URLs can only perform
// XMLHttpRequest calls to the same URL (origin + path) that originated the
// request. To allow XMLHttpRequest calls from a local URL to other URLs with
// the same origin set the CefSettings.file_access_from_file_urls_allowed
// value to true (1). To allow XMLHttpRequest calls from a local URL to all
// origins set the CefSettings.universal_access_from_file_urls_allowed value
// to true (1).
//
// If |is_display_isolated| is true (1) the scheme can only be displayed from
// other content hosted with the same scheme. For example, pages in other
// origins cannot create iframes or hyperlinks to URLs with the scheme. For
// schemes that must be accessible from other schemes set this value to false
// (0), set |is_cors_enabled| to true (1), and use CORS "Access-Control-Allow-
// Origin" headers to further restrict access.
//
// If |is_secure| is true (1) the scheme will be treated with the same
// security rules as those applied to "https" URLs. For example, loading this
// scheme from other secure schemes will not trigger mixed content warnings.
//
// If |is_cors_enabled| is true (1) the scheme can be sent CORS requests. This
// value should be true (1) in most cases where |is_standard| is true (1).
//
// If |is_csp_bypassing| is true (1) the scheme can bypass Content-Security-
// Policy (CSP) checks. This value should be false (0) in most cases where
// |is_standard| is true (1).
//
// This function may be called on any thread. It should only be called once
// per unique |scheme_name| value. If |scheme_name| is already registered or
// if an error occurs this function will return false (0).
func (d *SchemeRegistrar) AddCustomScheme(scheme_name string, is_standard, is_local, is_display_isolated, is_secure, is_cors_enabled, is_csp_bypassing int32) int32 {
	scheme_name_ := C.cef_string_userfree_alloc()
	setCEFStr(scheme_name, scheme_name_)
	defer func() {
		C.cef_string_userfree_free(scheme_name_)
	}()
	return int32(C.gocef_scheme_registrar_add_custom_scheme(d.toNative(), (*C.cef_string_t)(scheme_name_), C.int(is_standard), C.int(is_local), C.int(is_display_isolated), C.int(is_secure), C.int(is_cors_enabled), C.int(is_csp_bypassing), d.add_custom_scheme))
}
