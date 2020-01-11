// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_request_is_read_only(cef_request_t * self, int (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_request_get_url(cef_request_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// void gocef_request_set_url(cef_request_t * self, cef_string_t * uRL, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *)) { return callback__(self, uRL); }
	// cef_string_userfree_t gocef_request_get_method(cef_request_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// void gocef_request_set_method(cef_request_t * self, cef_string_t * method, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *)) { return callback__(self, method); }
	// void gocef_request_set_referrer(cef_request_t * self, cef_string_t * referrerURL, cef_referrer_policy_t policy, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *, cef_referrer_policy_t)) { return callback__(self, referrerURL, policy); }
	// cef_string_userfree_t gocef_request_get_referrer_url(cef_request_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// cef_referrer_policy_t gocef_request_get_referrer_policy(cef_request_t * self, cef_referrer_policy_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// cef_post_data_t * gocef_request_get_post_data(cef_request_t * self, cef_post_data_t * (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// void gocef_request_set_post_data(cef_request_t * self, cef_post_data_t * postData, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_post_data_t *)) { return callback__(self, postData); }
	// void gocef_request_get_header_map(cef_request_t * self, cef_string_multimap_t headerMap, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_multimap_t)) { return callback__(self, headerMap); }
	// void gocef_request_set_header_map(cef_request_t * self, cef_string_multimap_t headerMap, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_multimap_t)) { return callback__(self, headerMap); }
	// cef_string_userfree_t gocef_request_get_header_by_name(cef_request_t * self, cef_string_t * name, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *)) { return callback__(self, name); }
	// void gocef_request_set_header_by_name(cef_request_t * self, cef_string_t * name, cef_string_t * value, int overwrite, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *, cef_string_t *, int)) { return callback__(self, name, value, overwrite); }
	// void gocef_request_set(cef_request_t * self, cef_string_t * uRL, cef_string_t * method, cef_post_data_t * postData, cef_string_multimap_t headerMap, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *, cef_string_t *, cef_post_data_t *, cef_string_multimap_t)) { return callback__(self, uRL, method, postData, headerMap); }
	// int gocef_request_get_flags(cef_request_t * self, int (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// void gocef_request_set_flags(cef_request_t * self, int flags, void (CEF_CALLBACK *callback__)(cef_request_t *, int)) { return callback__(self, flags); }
	// cef_string_userfree_t gocef_request_get_first_party_for_cookies(cef_request_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// void gocef_request_set_first_party_for_cookies(cef_request_t * self, cef_string_t * uRL, void (CEF_CALLBACK *callback__)(cef_request_t *, cef_string_t *)) { return callback__(self, uRL); }
	// cef_resource_type_t gocef_request_get_resource_type(cef_request_t * self, cef_resource_type_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// cef_transition_type_t gocef_request_get_transition_type(cef_request_t * self, cef_transition_type_t (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	// uint64 gocef_request_get_identifier(cef_request_t * self, uint64 (CEF_CALLBACK *callback__)(cef_request_t *)) { return callback__(self); }
	"C"
)

// Request (cef_request_t from include/capi/cef_request_capi.h)
// Structure used to represent a web request. The functions of this structure
// may be called on any thread.
type Request C.cef_request_t

func (d *Request) toNative() *C.cef_request_t {
	return (*C.cef_request_t)(d)
}

// Base (base)
// Base structure.
func (d *Request) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsReadOnly (is_read_only)
// Returns true (1) if this object is read-only.
func (d *Request) IsReadOnly() int32 {
	return int32(C.gocef_request_is_read_only(d.toNative(), d.is_read_only))
}

// GetURL (get_url)
// Get the fully qualified URL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Request) GetURL() string {
	return cefuserfreestrToString(C.gocef_request_get_url(d.toNative(), d.get_url))
}

// SetURL (set_url)
// Set the fully qualified URL.
func (d *Request) SetURL(uRL string) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	C.gocef_request_set_url(d.toNative(), (*C.cef_string_t)(uRL_), d.set_url)
}

// GetMethod (get_method)
// Get the request function type. The value will default to POST if post data
// is provided and GET otherwise.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Request) GetMethod() string {
	return cefuserfreestrToString(C.gocef_request_get_method(d.toNative(), d.get_method))
}

// SetMethod (set_method)
// Set the request function type.
func (d *Request) SetMethod(method string) {
	method_ := C.cef_string_userfree_alloc()
	setCEFStr(method, method_)
	defer func() {
		C.cef_string_userfree_free(method_)
	}()
	C.gocef_request_set_method(d.toNative(), (*C.cef_string_t)(method_), d.set_method)
}

// SetReferrer (set_referrer)
// Set the referrer URL and policy. If non-NULL the referrer URL must be fully
// qualified with an HTTP or HTTPS scheme component. Any username, password or
// ref component will be removed.
func (d *Request) SetReferrer(referrerURL string, policy ReferrerPolicy) {
	referrerURL_ := C.cef_string_userfree_alloc()
	setCEFStr(referrerURL, referrerURL_)
	defer func() {
		C.cef_string_userfree_free(referrerURL_)
	}()
	C.gocef_request_set_referrer(d.toNative(), (*C.cef_string_t)(referrerURL_), C.cef_referrer_policy_t(policy), d.set_referrer)
}

// GetReferrerURL (get_referrer_url)
// Get the referrer URL.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Request) GetReferrerURL() string {
	return cefuserfreestrToString(C.gocef_request_get_referrer_url(d.toNative(), d.get_referrer_url))
}

// GetReferrerPolicy (get_referrer_policy)
// Get the referrer policy.
func (d *Request) GetReferrerPolicy() ReferrerPolicy {
	return ReferrerPolicy(C.gocef_request_get_referrer_policy(d.toNative(), d.get_referrer_policy))
}

// GetPostData (get_post_data)
// Get the post data.
func (d *Request) GetPostData() *PostData {
	return (*PostData)(C.gocef_request_get_post_data(d.toNative(), d.get_post_data))
}

// SetPostData (set_post_data)
// Set the post data.
func (d *Request) SetPostData(postData *PostData) {
	C.gocef_request_set_post_data(d.toNative(), postData.toNative(), d.set_post_data)
}

// GetHeaderMap (get_header_map)
// Get the header values. Will not include the Referer value if any.
func (d *Request) GetHeaderMap(headerMap StringMultimap) {
	C.gocef_request_get_header_map(d.toNative(), C.cef_string_multimap_t(headerMap), d.get_header_map)
}

// SetHeaderMap (set_header_map)
// Set the header values. If a Referer value exists in the header map it will
// be removed and ignored.
func (d *Request) SetHeaderMap(headerMap StringMultimap) {
	C.gocef_request_set_header_map(d.toNative(), C.cef_string_multimap_t(headerMap), d.set_header_map)
}

// GetHeaderByName (get_header_by_name)
// Returns the first header value for |name| or an NULL string if not found.
// Will not return the Referer value if any. Use GetHeaderMap instead if
// |name| might have multiple values.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Request) GetHeaderByName(name string) string {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return cefuserfreestrToString(C.gocef_request_get_header_by_name(d.toNative(), (*C.cef_string_t)(name_), d.get_header_by_name))
}

// SetHeaderByName (set_header_by_name)
// Set the header |name| to |value|. If |overwrite| is true (1) any existing
// values will be replaced with the new value. If |overwrite| is false (0) any
// existing values will not be overwritten. The Referer value cannot be set
// using this function.
func (d *Request) SetHeaderByName(name, value string, overwrite int32) {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	value_ := C.cef_string_userfree_alloc()
	setCEFStr(value, value_)
	defer func() {
		C.cef_string_userfree_free(value_)
	}()
	C.gocef_request_set_header_by_name(d.toNative(), (*C.cef_string_t)(name_), (*C.cef_string_t)(value_), C.int(overwrite), d.set_header_by_name)
}

// Set (set)
// Set all values at one time.
func (d *Request) Set(uRL, method string, postData *PostData, headerMap StringMultimap) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	method_ := C.cef_string_userfree_alloc()
	setCEFStr(method, method_)
	defer func() {
		C.cef_string_userfree_free(method_)
	}()
	C.gocef_request_set(d.toNative(), (*C.cef_string_t)(uRL_), (*C.cef_string_t)(method_), postData.toNative(), C.cef_string_multimap_t(headerMap), d.set)
}

// GetFlags (get_flags)
// Get the flags used in combination with cef_urlrequest_t. See
// cef_urlrequest_flags_t for supported values.
func (d *Request) GetFlags() int32 {
	return int32(C.gocef_request_get_flags(d.toNative(), d.get_flags))
}

// SetFlags (set_flags)
// Set the flags used in combination with cef_urlrequest_t.  See
// cef_urlrequest_flags_t for supported values.
func (d *Request) SetFlags(flags int32) {
	C.gocef_request_set_flags(d.toNative(), C.int(flags), d.set_flags)
}

// GetFirstPartyForCookies (get_first_party_for_cookies)
// Get the URL to the first party for cookies used in combination with
// cef_urlrequest_t.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Request) GetFirstPartyForCookies() string {
	return cefuserfreestrToString(C.gocef_request_get_first_party_for_cookies(d.toNative(), d.get_first_party_for_cookies))
}

// SetFirstPartyForCookies (set_first_party_for_cookies)
// Set the URL to the first party for cookies used in combination with
// cef_urlrequest_t.
func (d *Request) SetFirstPartyForCookies(uRL string) {
	uRL_ := C.cef_string_userfree_alloc()
	setCEFStr(uRL, uRL_)
	defer func() {
		C.cef_string_userfree_free(uRL_)
	}()
	C.gocef_request_set_first_party_for_cookies(d.toNative(), (*C.cef_string_t)(uRL_), d.set_first_party_for_cookies)
}

// GetResourceType (get_resource_type)
// Get the resource type for this request. Only available in the browser
// process.
func (d *Request) GetResourceType() ResourceType {
	return ResourceType(C.gocef_request_get_resource_type(d.toNative(), d.get_resource_type))
}

// GetTransitionType (get_transition_type)
// Get the transition type for this request. Only available in the browser
// process and only applies to requests that represent a main frame or sub-
// frame navigation.
func (d *Request) GetTransitionType() TransitionType {
	return TransitionType(C.gocef_request_get_transition_type(d.toNative(), d.get_transition_type))
}

// GetIdentifier (get_identifier)
// Returns the globally unique identifier for this request or 0 if not
// specified. Can be used by cef_resource_request_handler_t implementations in
// the browser process to track a single request across multiple callbacks.
func (d *Request) GetIdentifier() uint64 {
	return uint64(C.gocef_request_get_identifier(d.toNative(), d.get_identifier))
}
