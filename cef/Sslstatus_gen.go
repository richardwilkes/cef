// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// int gocef_sslstatus_is_secure_connection(cef_sslstatus_t * self, int (CEF_CALLBACK *callback__)(cef_sslstatus_t *)) { return callback__(self); }
	// cef_cert_status_t gocef_sslstatus_get_cert_status(cef_sslstatus_t * self, cef_cert_status_t (CEF_CALLBACK *callback__)(cef_sslstatus_t *)) { return callback__(self); }
	// cef_ssl_version_t gocef_sslstatus_get_sslversion(cef_sslstatus_t * self, cef_ssl_version_t (CEF_CALLBACK *callback__)(cef_sslstatus_t *)) { return callback__(self); }
	// cef_ssl_content_status_t gocef_sslstatus_get_content_status(cef_sslstatus_t * self, cef_ssl_content_status_t (CEF_CALLBACK *callback__)(cef_sslstatus_t *)) { return callback__(self); }
	// cef_x509certificate_t * gocef_sslstatus_get_x509certificate(cef_sslstatus_t * self, cef_x509certificate_t * (CEF_CALLBACK *callback__)(cef_sslstatus_t *)) { return callback__(self); }
	"C"
)

// Sslstatus (cef_sslstatus_t from include/capi/cef_ssl_status_capi.h)
// Structure representing the SSL information for a navigation entry.
type Sslstatus C.cef_sslstatus_t

func (d *Sslstatus) toNative() *C.cef_sslstatus_t {
	return (*C.cef_sslstatus_t)(d)
}

// Base (base)
// Base structure.
func (d *Sslstatus) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsSecureConnection (is_secure_connection)
// Returns true (1) if the status is related to a secure SSL/TLS connection.
func (d *Sslstatus) IsSecureConnection() int32 {
	return int32(C.gocef_sslstatus_is_secure_connection(d.toNative(), d.is_secure_connection))
}

// GetCertStatus (get_cert_status)
// Returns a bitmask containing any and all problems verifying the server
// certificate.
func (d *Sslstatus) GetCertStatus() CertStatus {
	return CertStatus(C.gocef_sslstatus_get_cert_status(d.toNative(), d.get_cert_status))
}

// GetSslversion (get_sslversion)
// Returns the SSL version used for the SSL connection.
func (d *Sslstatus) GetSslversion() SSLVersion {
	return SSLVersion(C.gocef_sslstatus_get_sslversion(d.toNative(), d.get_sslversion))
}

// GetContentStatus (get_content_status)
// Returns a bitmask containing the page security content status.
func (d *Sslstatus) GetContentStatus() SSLContentStatus {
	return SSLContentStatus(C.gocef_sslstatus_get_content_status(d.toNative(), d.get_content_status))
}

// GetX509certificate (get_x509certificate)
// Returns the X.509 certificate.
func (d *Sslstatus) GetX509certificate() *X509certificate {
	return (*X509certificate)(C.gocef_sslstatus_get_x509certificate(d.toNative(), d.get_x509certificate))
}
