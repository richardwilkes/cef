package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// RequestContextSettings holds context initialization settings.
//
// Defined in include/internal/cef_types.h
type RequestContextSettings struct {
	native *C.cef_request_context_settings_t
}

// NewRequestContextSettings creates a new default RequestContextSettings
// instance.
func NewRequestContextSettings() *RequestContextSettings {
	settings := &RequestContextSettings{native: (*C.cef_request_context_settings_t)(C.calloc(1, C.sizeof_struct__cef_request_context_settings_t))}
	settings.native.size = C.sizeof_struct__cef_request_context_settings_t
	return settings
}

// SetCachePath sets the location where cache data will be stored on disk. If
// empty then browsers will be created in "incognito mode" where in-memory
// caches are used for storage and no data is persisted to disk. HTML5
// databases such as localStorage will only persist across sessions if a
// cache path is specified. To share the global browser cache and related
// configuration set this value to match the value set in
// Settings.SetCachePath.
func (s *RequestContextSettings) SetCachePath(path string) {
	setCEFStr(path, &s.native.cache_path)
}

// EnablePersistentSessionCookies to persist session cookies (cookies without
// an expiry date or validity interval) by default when using the global
// cookie manager. Session cookies are generally intended to be transient and
// most Web browsers do not persist them. The cache path (via SetCachePath)
// must also be specified to enable this feature. Can be set globally using
// Settings.EnablePersistentSessionCookies. This value will be ignored if
// SetCachePath is not called on this object or if the path that was set
// matches the one set via Settings.SetCachePath.
func (s *RequestContextSettings) EnablePersistentSessionCookies() {
	s.native.persist_session_cookies = 1
}

// EnablePersistentUserPreferences to persist user preferences as a JSON file
// in the cache path directory. The cache path (via SetCachePath) must also be
// specified to enable this feature. Can be set globally using
// Settings.EnablePersistentUserPreferences. This value will be ignored if
// SetCachePath is not called on this object or if the path that was set
// matches the one set via Settings.SetCachePath.
func (s *RequestContextSettings) EnablePersistentUserPreferences() {
	s.native.persist_user_preferences = 1
}

// IgnoreCertificateErrors causes errors related to invalid SSL certificates
// to be ignored. Enabling this setting can lead to potential security
// vulnerabilities like "man in the middle" attacks. Applications that load
// content from the internet should not enable this setting. Can be set
// globally using Settings.IgnoreCertificateErrors. This value will be ignored
// if SetCachePath is not called on this object or if the path that was set
// matches the one set via Settings.SetCachePath.
func (s *RequestContextSettings) IgnoreCertificateErrors() {
	s.native.ignore_certificate_errors = 1
}

// EnableNetSecurityExpiration enables date-based expiration of built in
// network security information (i.e. certificate transparency logs, HSTS
// preloading and pinning information). Enabling this option improves network
// security but may cause HTTPS load failures when using CEF binaries built
// more than 10 weeks in the past. See
// https://www.certificate-transparency.org/ and https://www.chromium.org/hsts
// for details. Can be set globally using
// Settings.EnableNetSecurityExpiration.
func (s *RequestContextSettings) EnableNetSecurityExpiration() {
	s.native.enable_net_security_expiration = 1
}

// SetAcceptLanguageList sets a comma delimited ordered list of language codes
// without any whitespace that will be used in the "Accept-Language" HTTP
// header. Can be set globally using Settings.SetAcceptLanguageList. This
// value will be ignored if SetCachePath is not called on this object or if
// the path that was set matches the one set via Settings.SetCachePath.
func (s *RequestContextSettings) SetAcceptLanguageList(list string) {
	setCEFStr(list, &s.native.accept_language_list)
}
