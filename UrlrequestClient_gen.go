package cef

import (
	// #include "UrlrequestClient_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// UrlrequestClientProxy defines methods required for using UrlrequestClient.
type UrlrequestClientProxy interface {
	OnRequestComplete(self *UrlrequestClient, request *Urlrequest)
	OnUploadProgress(self *UrlrequestClient, request *Urlrequest, current int64, total int64)
	OnDownloadProgress(self *UrlrequestClient, request *Urlrequest, current int64, total int64)
	OnDownloadData(self *UrlrequestClient, request *Urlrequest, data unsafe.Pointer, data_length uint64)
	GetAuthCredentials(self *UrlrequestClient, isProxy int32, host string, port int32, realm string, scheme string, callback *AuthCallback) int32
}

// UrlrequestClient (cef_urlrequest_client_t from include/capi/cef_urlrequest_capi.h)
// Structure that should be implemented by the cef_urlrequest_t client. The
// functions of this structure will be called on the same thread that created
// the request unless otherwise documented.
type UrlrequestClient C.cef_urlrequest_client_t

// NewUrlrequestClient creates a new UrlrequestClient with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewUrlrequestClient(proxy UrlrequestClientProxy) *UrlrequestClient {
	result := (*UrlrequestClient)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_urlrequest_client_t, proxy)))
	if proxy != nil {
		C.gocef_set_urlrequest_client_proxy(result.toNative())
	}
	return result
}

func (d *UrlrequestClient) toNative() *C.cef_urlrequest_client_t {
	return (*C.cef_urlrequest_client_t)(d)
}

func lookupUrlrequestClientProxy(obj *BaseRefCounted) UrlrequestClientProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(UrlrequestClientProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type UrlrequestClientProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *UrlrequestClient) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnRequestComplete (on_request_complete)
// Notifies the client that the request has completed. Use the
// cef_urlrequest_t::GetRequestStatus function to determine if the request was
// successful or not.
func (d *UrlrequestClient) OnRequestComplete(request *Urlrequest) {
	lookupUrlrequestClientProxy(d.Base()).OnRequestComplete(d, request)
}

//export gocef_urlrequest_client_on_request_complete
func gocef_urlrequest_client_on_request_complete(self *C.cef_urlrequest_client_t, request *C.cef_urlrequest_t) {
	me__ := (*UrlrequestClient)(self)
	proxy__ := lookupUrlrequestClientProxy(me__.Base())
	proxy__.OnRequestComplete(me__, (*Urlrequest)(request))
}

// OnUploadProgress (on_upload_progress)
// Notifies the client of upload progress. |current| denotes the number of
// bytes sent so far and |total| is the total size of uploading data (or -1 if
// chunked upload is enabled). This function will only be called if the
// UR_FLAG_REPORT_UPLOAD_PROGRESS flag is set on the request.
func (d *UrlrequestClient) OnUploadProgress(request *Urlrequest, current, total int64) {
	lookupUrlrequestClientProxy(d.Base()).OnUploadProgress(d, request, current, total)
}

//export gocef_urlrequest_client_on_upload_progress
func gocef_urlrequest_client_on_upload_progress(self *C.cef_urlrequest_client_t, request *C.cef_urlrequest_t, current C.int64, total C.int64) {
	me__ := (*UrlrequestClient)(self)
	proxy__ := lookupUrlrequestClientProxy(me__.Base())
	proxy__.OnUploadProgress(me__, (*Urlrequest)(request), int64(current), int64(total))
}

// OnDownloadProgress (on_download_progress)
// Notifies the client of download progress. |current| denotes the number of
// bytes received up to the call and |total| is the expected total size of the
// response (or -1 if not determined).
func (d *UrlrequestClient) OnDownloadProgress(request *Urlrequest, current, total int64) {
	lookupUrlrequestClientProxy(d.Base()).OnDownloadProgress(d, request, current, total)
}

//export gocef_urlrequest_client_on_download_progress
func gocef_urlrequest_client_on_download_progress(self *C.cef_urlrequest_client_t, request *C.cef_urlrequest_t, current C.int64, total C.int64) {
	me__ := (*UrlrequestClient)(self)
	proxy__ := lookupUrlrequestClientProxy(me__.Base())
	proxy__.OnDownloadProgress(me__, (*Urlrequest)(request), int64(current), int64(total))
}

// OnDownloadData (on_download_data)
// Called when some part of the response is read. |data| contains the current
// bytes received since the last call. This function will not be called if the
// UR_FLAG_NO_DOWNLOAD_DATA flag is set on the request.
func (d *UrlrequestClient) OnDownloadData(request *Urlrequest, data unsafe.Pointer, data_length uint64) {
	lookupUrlrequestClientProxy(d.Base()).OnDownloadData(d, request, data, data_length)
}

//export gocef_urlrequest_client_on_download_data
func gocef_urlrequest_client_on_download_data(self *C.cef_urlrequest_client_t, request *C.cef_urlrequest_t, data unsafe.Pointer, data_length C.size_t) {
	me__ := (*UrlrequestClient)(self)
	proxy__ := lookupUrlrequestClientProxy(me__.Base())
	proxy__.OnDownloadData(me__, (*Urlrequest)(request), data, uint64(data_length))
}

// GetAuthCredentials (get_auth_credentials)
// Called on the IO thread when the browser needs credentials from the user.
// |isProxy| indicates whether the host is a proxy server. |host| contains the
// hostname and |port| contains the port number. Return true (1) to continue
// the request and call cef_auth_callback_t::cont() when the authentication
// information is available. Return false (0) to cancel the request. This
// function will only be called for requests initiated from the browser
// process.
func (d *UrlrequestClient) GetAuthCredentials(isProxy int32, host string, port int32, realm, scheme string, callback *AuthCallback) int32 {
	return lookupUrlrequestClientProxy(d.Base()).GetAuthCredentials(d, isProxy, host, port, realm, scheme, callback)
}

//export gocef_urlrequest_client_get_auth_credentials
func gocef_urlrequest_client_get_auth_credentials(self *C.cef_urlrequest_client_t, isProxy C.int, host *C.cef_string_t, port C.int, realm *C.cef_string_t, scheme *C.cef_string_t, callback *C.cef_auth_callback_t) C.int {
	me__ := (*UrlrequestClient)(self)
	proxy__ := lookupUrlrequestClientProxy(me__.Base())
	return C.int(proxy__.GetAuthCredentials(me__, int32(isProxy), cefstrToString(host), int32(port), cefstrToString(realm), cefstrToString(scheme), (*AuthCallback)(callback)))
}
