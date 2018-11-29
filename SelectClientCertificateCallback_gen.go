package cef

import (
	// #include "SelectClientCertificateCallback_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// SelectClientCertificateCallbackProxy defines methods required for using SelectClientCertificateCallback.
type SelectClientCertificateCallbackProxy interface {
	Select(self *SelectClientCertificateCallback, cert *X509certificate)
}

// SelectClientCertificateCallback (cef_select_client_certificate_callback_t from include/capi/cef_request_handler_capi.h)
// Callback structure used to select a client certificate for authentication.
type SelectClientCertificateCallback C.cef_select_client_certificate_callback_t

// NewSelectClientCertificateCallback creates a new SelectClientCertificateCallback with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewSelectClientCertificateCallback(proxy SelectClientCertificateCallbackProxy) *SelectClientCertificateCallback {
	result := (*SelectClientCertificateCallback)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_select_client_certificate_callback_t, proxy)))
	if proxy != nil {
		C.gocef_set_select_client_certificate_callback_proxy(result.toNative())
	}
	return result
}

func (d *SelectClientCertificateCallback) toNative() *C.cef_select_client_certificate_callback_t {
	return (*C.cef_select_client_certificate_callback_t)(d)
}

func lookupSelectClientCertificateCallbackProxy(obj *BaseRefCounted) SelectClientCertificateCallbackProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(SelectClientCertificateCallbackProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type SelectClientCertificateCallbackProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *SelectClientCertificateCallback) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Select (_select)
// Chooses the specified certificate for client certificate authentication.
// NULL value means that no client certificate should be used.
func (d *SelectClientCertificateCallback) Select(cert *X509certificate) {
	lookupSelectClientCertificateCallbackProxy(d.Base()).Select(d, cert)
}

//export gocef_select_client_certificate_callback__select
func gocef_select_client_certificate_callback__select(self *C.cef_select_client_certificate_callback_t, cert *C.cef_x509certificate_t) {
	me__ := (*SelectClientCertificateCallback)(self)
	proxy__ := lookupSelectClientCertificateCallbackProxy(me__.Base())
	proxy__.Select(me__, (*X509certificate)(cert))
}
