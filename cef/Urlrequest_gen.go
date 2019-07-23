// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_request_t * gocef_urlrequest_get_request(cef_urlrequest_t * self, cef_request_t * (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// cef_urlrequest_client_t * gocef_urlrequest_get_client(cef_urlrequest_t * self, cef_urlrequest_client_t * (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// cef_urlrequest_status_t gocef_urlrequest_get_request_status(cef_urlrequest_t * self, cef_urlrequest_status_t (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// cef_errorcode_t gocef_urlrequest_get_request_error(cef_urlrequest_t * self, cef_errorcode_t (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// cef_response_t * gocef_urlrequest_get_response(cef_urlrequest_t * self, cef_response_t * (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// int gocef_urlrequest_response_was_cached(cef_urlrequest_t * self, int (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	// void gocef_urlrequest_cancel(cef_urlrequest_t * self, void (CEF_CALLBACK *callback__)(cef_urlrequest_t *)) { return callback__(self); }
	"C"
)

// Urlrequest (cef_urlrequest_t from include/capi/cef_urlrequest_capi.h)
// Structure used to make a URL request. URL requests are not associated with a
// browser instance so no cef_client_t callbacks will be executed. URL requests
// can be created on any valid CEF thread in either the browser or render
// process. Once created the functions of the URL request object must be
// accessed on the same thread that created it.
type Urlrequest C.cef_urlrequest_t

func (d *Urlrequest) toNative() *C.cef_urlrequest_t {
	return (*C.cef_urlrequest_t)(d)
}

// Base (base)
// Base structure.
func (d *Urlrequest) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetRequest (get_request)
// Returns the request object used to create this URL request. The returned
// object is read-only and should not be modified.
func (d *Urlrequest) GetRequest() *Request {
	return (*Request)(C.gocef_urlrequest_get_request(d.toNative(), d.get_request))
}

// GetClient (get_client)
// Returns the client.
func (d *Urlrequest) GetClient() *UrlrequestClient {
	return (*UrlrequestClient)(C.gocef_urlrequest_get_client(d.toNative(), d.get_client))
}

// GetRequestStatus (get_request_status)
// Returns the request status.
func (d *Urlrequest) GetRequestStatus() UrlrequestStatus {
	return UrlrequestStatus(C.gocef_urlrequest_get_request_status(d.toNative(), d.get_request_status))
}

// GetRequestError (get_request_error)
// Returns the request error if status is UR_CANCELED or UR_FAILED, or 0
// otherwise.
func (d *Urlrequest) GetRequestError() Errorcode {
	return Errorcode(C.gocef_urlrequest_get_request_error(d.toNative(), d.get_request_error))
}

// GetResponse (get_response)
// Returns the response, or NULL if no response information is available.
// Response information will only be available after the upload has completed.
// The returned object is read-only and should not be modified.
func (d *Urlrequest) GetResponse() *Response {
	return (*Response)(C.gocef_urlrequest_get_response(d.toNative(), d.get_response))
}

// ResponseWasCached (response_was_cached)
// Returns true (1) if the response body was served from the cache. This
// includes responses for which revalidation was required.
func (d *Urlrequest) ResponseWasCached() int32 {
	return int32(C.gocef_urlrequest_response_was_cached(d.toNative(), d.response_was_cached))
}

// Cancel (cancel)
// Cancel the request.
func (d *Urlrequest) Cancel() {
	C.gocef_urlrequest_cancel(d.toNative(), d.cancel)
}
