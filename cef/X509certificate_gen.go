// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_x509cert_principal_t * gocef_x509certificate_get_subject(cef_x509certificate_t * self, cef_x509cert_principal_t * (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_x509cert_principal_t * gocef_x509certificate_get_issuer(cef_x509certificate_t * self, cef_x509cert_principal_t * (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_binary_value_t * gocef_x509certificate_get_serial_number(cef_x509certificate_t * self, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_time_t gocef_x509certificate_get_valid_start(cef_x509certificate_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_time_t gocef_x509certificate_get_valid_expiry(cef_x509certificate_t * self, cef_time_t (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_binary_value_t * gocef_x509certificate_get_derencoded(cef_x509certificate_t * self, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// cef_binary_value_t * gocef_x509certificate_get_pemencoded(cef_x509certificate_t * self, cef_binary_value_t * (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// size_t gocef_x509certificate_get_issuer_chain_size(cef_x509certificate_t * self, size_t (CEF_CALLBACK *callback__)(cef_x509certificate_t *)) { return callback__(self); }
	// void gocef_x509certificate_get_derencoded_issuer_chain(cef_x509certificate_t * self, size_t * chainCount, cef_binary_value_t ** chain, void (CEF_CALLBACK *callback__)(cef_x509certificate_t *, size_t *, cef_binary_value_t **)) { return callback__(self, chainCount, chain); }
	// void gocef_x509certificate_get_pemencoded_issuer_chain(cef_x509certificate_t * self, size_t * chainCount, cef_binary_value_t ** chain, void (CEF_CALLBACK *callback__)(cef_x509certificate_t *, size_t *, cef_binary_value_t **)) { return callback__(self, chainCount, chain); }
	"C"
)

// X509certificate (cef_x509certificate_t from include/capi/cef_x509_certificate_capi.h)
// Structure representing a X.509 certificate.
type X509certificate C.cef_x509certificate_t

func (d *X509certificate) toNative() *C.cef_x509certificate_t {
	return (*C.cef_x509certificate_t)(d)
}

// Base (base)
// Base structure.
func (d *X509certificate) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetSubject (get_subject)
// Returns the subject of the X.509 certificate. For HTTPS server certificates
// this represents the web server.  The common name of the subject should
// match the host name of the web server.
func (d *X509certificate) GetSubject() *X509certPrincipal {
	return (*X509certPrincipal)(C.gocef_x509certificate_get_subject(d.toNative(), d.get_subject))
}

// GetIssuer (get_issuer)
// Returns the issuer of the X.509 certificate.
func (d *X509certificate) GetIssuer() *X509certPrincipal {
	return (*X509certPrincipal)(C.gocef_x509certificate_get_issuer(d.toNative(), d.get_issuer))
}

// GetSerialNumber (get_serial_number)
// Returns the DER encoded serial number for the X.509 certificate. The value
// possibly includes a leading 00 byte.
func (d *X509certificate) GetSerialNumber() *BinaryValue {
	return (*BinaryValue)(C.gocef_x509certificate_get_serial_number(d.toNative(), d.get_serial_number))
}

// GetValidStart (get_valid_start)
// Returns the date before which the X.509 certificate is invalid.
// CefTime.GetTimeT() will return 0 if no date was specified.
func (d *X509certificate) GetValidStart() Time {
	cresult__ := C.gocef_x509certificate_get_valid_start(d.toNative(), d.get_valid_start)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetValidExpiry (get_valid_expiry)
// Returns the date after which the X.509 certificate is invalid.
// CefTime.GetTimeT() will return 0 if no date was specified.
func (d *X509certificate) GetValidExpiry() Time {
	cresult__ := C.gocef_x509certificate_get_valid_expiry(d.toNative(), d.get_valid_expiry)
	var result__ Time
	(&cresult__).intoGo(&result__)
	return result__
}

// GetDerencoded (get_derencoded)
// Returns the DER encoded data for the X.509 certificate.
func (d *X509certificate) GetDerencoded() *BinaryValue {
	return (*BinaryValue)(C.gocef_x509certificate_get_derencoded(d.toNative(), d.get_derencoded))
}

// GetPemencoded (get_pemencoded)
// Returns the PEM encoded data for the X.509 certificate.
func (d *X509certificate) GetPemencoded() *BinaryValue {
	return (*BinaryValue)(C.gocef_x509certificate_get_pemencoded(d.toNative(), d.get_pemencoded))
}

// GetIssuerChainSize (get_issuer_chain_size)
// Returns the number of certificates in the issuer chain. If 0, the
// certificate is self-signed.
func (d *X509certificate) GetIssuerChainSize() uint64 {
	return uint64(C.gocef_x509certificate_get_issuer_chain_size(d.toNative(), d.get_issuer_chain_size))
}

// GetDerencodedIssuerChain (get_derencoded_issuer_chain)
// Returns the DER encoded data for the certificate issuer chain. If we failed
// to encode a certificate in the chain it is still present in the array but
// is an NULL string.
func (d *X509certificate) GetDerencodedIssuerChain(chainCount *uint64, chain **BinaryValue) {
	chain_ := (*chain).toNative()
	C.gocef_x509certificate_get_derencoded_issuer_chain(d.toNative(), (*C.size_t)(chainCount), &chain_, d.get_derencoded_issuer_chain)
}

// GetPemencodedIssuerChain (get_pemencoded_issuer_chain)
// Returns the PEM encoded data for the certificate issuer chain. If we failed
// to encode a certificate in the chain it is still present in the array but
// is an NULL string.
func (d *X509certificate) GetPemencodedIssuerChain(chainCount *uint64, chain **BinaryValue) {
	chain_ := (*chain).toNative()
	C.gocef_x509certificate_get_pemencoded_issuer_chain(d.toNative(), (*C.size_t)(chainCount), &chain_, d.get_pemencoded_issuer_chain)
}
