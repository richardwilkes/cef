// Code created from "struct.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// RequestContextSettings (cef_request_context_settings_t from include/internal/cef_types.h)
// Request context initialization settings. Specify NULL or 0 to get the
// recommended default values.
type RequestContextSettings struct {
	// Size (size)
	// Size of this structure.
	Size uint64
	// CachePath (cache_path)
	// The location where cache data for this request context will be stored on
	// disk. If non-empty this must be either equal to or a child directory of
	// CefSettings.root_cache_path. If empty then browsers will be created in
	// "incognito mode" where in-memory caches are used for storage and no data is
	// persisted to disk. HTML5 databases such as localStorage will only persist
	// across sessions if a cache path is specified. To share the global browser
	// cache and related configuration set this value to match the
	// CefSettings.cache_path value.
	CachePath string
	// PersistSessionCookies (persist_session_cookies)
	// To persist session cookies (cookies without an expiry date or validity
	// interval) by default when using the global cookie manager set this value to
	// true (1). Session cookies are generally intended to be transient and most
	// Web browsers do not persist them. Can be set globally using the
	// CefSettings.persist_session_cookies value. This value will be ignored if
	// |cache_path| is empty or if it matches the CefSettings.cache_path value.
	PersistSessionCookies int32
	// PersistUserPreferences (persist_user_preferences)
	// To persist user preferences as a JSON file in the cache path directory set
	// this value to true (1). Can be set globally using the
	// CefSettings.persist_user_preferences value. This value will be ignored if
	// |cache_path| is empty or if it matches the CefSettings.cache_path value.
	PersistUserPreferences int32
	// IgnoreCertificateErrors (ignore_certificate_errors)
	// Set to true (1) to ignore errors related to invalid SSL certificates.
	// Enabling this setting can lead to potential security vulnerabilities like
	// "man in the middle" attacks. Applications that load content from the
	// internet should not enable this setting. Can be set globally using the
	// CefSettings.ignore_certificate_errors value. This value will be ignored if
	// |cache_path| matches the CefSettings.cache_path value.
	IgnoreCertificateErrors int32
	// EnableNetSecurityExpiration (enable_net_security_expiration)
	// Set to true (1) to enable date-based expiration of built in network
	// security information (i.e. certificate transparency logs, HSTS preloading
	// and pinning information). Enabling this option improves network security
	// but may cause HTTPS load failures when using CEF binaries built more than
	// 10 weeks in the past. See https://www.certificate-transparency.org/ and
	// https://www.chromium.org/hsts for details. Can be set globally using the
	// CefSettings.enable_net_security_expiration value.
	EnableNetSecurityExpiration int32
	// AcceptLanguageList (accept_language_list)
	// Comma delimited ordered list of language codes without any whitespace that
	// will be used in the "Accept-Language" HTTP header. Can be set globally
	// using the CefSettings.accept_language_list value or overridden on a per-
	// browser basis using the CefBrowserSettings.accept_language_list value. If
	// all values are empty then "en-US,en" will be used. This value will be
	// ignored if |cache_path| matches the CefSettings.cache_path value.
	AcceptLanguageList string
}

// NewRequestContextSettings creates a new RequestContextSettings.
func NewRequestContextSettings() *RequestContextSettings {
	return &RequestContextSettings{
		Size: C.sizeof_struct__cef_request_context_settings_t,
	}
}

func (d *RequestContextSettings) toNative(native *C.cef_request_context_settings_t) *C.cef_request_context_settings_t {
	if d == nil {
		return nil
	}
	native.size = C.size_t(d.Size)
	setCEFStr(d.CachePath, &native.cache_path)
	native.persist_session_cookies = C.int(d.PersistSessionCookies)
	native.persist_user_preferences = C.int(d.PersistUserPreferences)
	native.ignore_certificate_errors = C.int(d.IgnoreCertificateErrors)
	native.enable_net_security_expiration = C.int(d.EnableNetSecurityExpiration)
	setCEFStr(d.AcceptLanguageList, &native.accept_language_list)
	return native
}

func (n *C.cef_request_context_settings_t) toGo() *RequestContextSettings {
	if n == nil {
		return nil
	}
	var d RequestContextSettings
	n.intoGo(&d)
	return &d
}

func (n *C.cef_request_context_settings_t) intoGo(d *RequestContextSettings) {
	d.Size = uint64(n.size)
	d.CachePath = cefstrToString(&n.cache_path)
	d.PersistSessionCookies = int32(n.persist_session_cookies)
	d.PersistUserPreferences = int32(n.persist_user_preferences)
	d.IgnoreCertificateErrors = int32(n.ignore_certificate_errors)
	d.EnableNetSecurityExpiration = int32(n.enable_net_security_expiration)
	d.AcceptLanguageList = cefstrToString(&n.accept_language_list)
}
