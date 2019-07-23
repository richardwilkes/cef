// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_cert_status_t gocef_sslinfo_get_cert_status(cef_sslinfo_t * self, cef_cert_status_t (CEF_CALLBACK *callback__)(cef_sslinfo_t *)) { return callback__(self); }
	// cef_x509certificate_t * gocef_sslinfo_get_x509certificate(cef_sslinfo_t * self, cef_x509certificate_t * (CEF_CALLBACK *callback__)(cef_sslinfo_t *)) { return callback__(self); }
	"C"
)

// Sslinfo (cef_sslinfo_t from include/capi/cef_ssl_info_capi.h)
// Structure representing SSL information.
type Sslinfo C.cef_sslinfo_t

func (d *Sslinfo) toNative() *C.cef_sslinfo_t {
	return (*C.cef_sslinfo_t)(d)
}

// Base (base)
// Base structure.
func (d *Sslinfo) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetCertStatus (get_cert_status)
// Returns a bitmask containing any and all problems verifying the server
// certificate.
func (d *Sslinfo) GetCertStatus() CertStatus {
	return CertStatus(C.gocef_sslinfo_get_cert_status(d.toNative(), d.get_cert_status))
}

// GetX509certificate (get_x509certificate)
// Returns the X.509 certificate.
func (d *Sslinfo) GetX509certificate() *X509certificate {
	return (*X509certificate)(C.gocef_sslinfo_get_x509certificate(d.toNative(), d.get_x509certificate))
}
