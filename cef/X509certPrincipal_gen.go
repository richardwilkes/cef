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
	// cef_string_userfree_t gocef_x509cert_principal_get_display_name(cef_x509cert_principal_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_x509cert_principal_get_common_name(cef_x509cert_principal_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_x509cert_principal_get_locality_name(cef_x509cert_principal_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_x509cert_principal_get_state_or_province_name(cef_x509cert_principal_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_x509cert_principal_get_country_name(cef_x509cert_principal_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *)) { return callback__(self); }
	// void gocef_x509cert_principal_get_street_addresses(cef_x509cert_principal_t * self, cef_string_list_t addresses, void (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *, cef_string_list_t)) { return callback__(self, addresses); }
	// void gocef_x509cert_principal_get_organization_names(cef_x509cert_principal_t * self, cef_string_list_t names, void (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *, cef_string_list_t)) { return callback__(self, names); }
	// void gocef_x509cert_principal_get_organization_unit_names(cef_x509cert_principal_t * self, cef_string_list_t names, void (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *, cef_string_list_t)) { return callback__(self, names); }
	// void gocef_x509cert_principal_get_domain_components(cef_x509cert_principal_t * self, cef_string_list_t components, void (CEF_CALLBACK *callback__)(cef_x509cert_principal_t *, cef_string_list_t)) { return callback__(self, components); }
	"C"
)

// X509certPrincipal (cef_x509cert_principal_t from include/capi/cef_x509_certificate_capi.h)
// Structure representing the issuer or subject field of an X.509 certificate.
type X509certPrincipal C.cef_x509cert_principal_t

func (d *X509certPrincipal) toNative() *C.cef_x509cert_principal_t {
	return (*C.cef_x509cert_principal_t)(d)
}

// Base (base)
// Base structure.
func (d *X509certPrincipal) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetDisplayName (get_display_name)
// Returns a name that can be used to represent the issuer. It tries in this
// order: Common Name (CN), Organization Name (O) and Organizational Unit Name
// (OU) and returns the first non-NULL one found.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *X509certPrincipal) GetDisplayName() string {
	return cefuserfreestrToString(C.gocef_x509cert_principal_get_display_name(d.toNative(), d.get_display_name))
}

// GetCommonName (get_common_name)
// Returns the common name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *X509certPrincipal) GetCommonName() string {
	return cefuserfreestrToString(C.gocef_x509cert_principal_get_common_name(d.toNative(), d.get_common_name))
}

// GetLocalityName (get_locality_name)
// Returns the locality name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *X509certPrincipal) GetLocalityName() string {
	return cefuserfreestrToString(C.gocef_x509cert_principal_get_locality_name(d.toNative(), d.get_locality_name))
}

// GetStateOrProvinceName (get_state_or_province_name)
// Returns the state or province name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *X509certPrincipal) GetStateOrProvinceName() string {
	return cefuserfreestrToString(C.gocef_x509cert_principal_get_state_or_province_name(d.toNative(), d.get_state_or_province_name))
}

// GetCountryName (get_country_name)
// Returns the country name.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *X509certPrincipal) GetCountryName() string {
	return cefuserfreestrToString(C.gocef_x509cert_principal_get_country_name(d.toNative(), d.get_country_name))
}

// GetStreetAddresses (get_street_addresses)
// Retrieve the list of street addresses.
func (d *X509certPrincipal) GetStreetAddresses(addresses StringList) {
	C.gocef_x509cert_principal_get_street_addresses(d.toNative(), C.cef_string_list_t(addresses), d.get_street_addresses)
}

// GetOrganizationNames (get_organization_names)
// Retrieve the list of organization names.
func (d *X509certPrincipal) GetOrganizationNames(names StringList) {
	C.gocef_x509cert_principal_get_organization_names(d.toNative(), C.cef_string_list_t(names), d.get_organization_names)
}

// GetOrganizationUnitNames (get_organization_unit_names)
// Retrieve the list of organization unit names.
func (d *X509certPrincipal) GetOrganizationUnitNames(names StringList) {
	C.gocef_x509cert_principal_get_organization_unit_names(d.toNative(), C.cef_string_list_t(names), d.get_organization_unit_names)
}

// GetDomainComponents (get_domain_components)
// Retrieve the list of domain components.
func (d *X509certPrincipal) GetDomainComponents(components StringList) {
	C.gocef_x509cert_principal_get_domain_components(d.toNative(), C.cef_string_list_t(components), d.get_domain_components)
}
