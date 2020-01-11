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
	// int gocef_scheme_registrar_add_custom_scheme(cef_scheme_registrar_t * self, cef_string_t * schemeName, int options, int (CEF_CALLBACK *callback__)(cef_scheme_registrar_t *, cef_string_t *, int)) { return callback__(self, schemeName, options); }
	"C"
)

// SchemeRegistrar (cef_scheme_registrar_t from include/capi/cef_scheme_capi.h)
// Structure that manages custom scheme registrations.
type SchemeRegistrar C.cef_scheme_registrar_t

func (d *SchemeRegistrar) toNative() *C.cef_scheme_registrar_t {
	return (*C.cef_scheme_registrar_t)(d)
}

// Base (base)
// Base structure.
func (d *SchemeRegistrar) Base() *BaseScoped {
	return (*BaseScoped)(&d.base)
}

// AddCustomScheme (add_custom_scheme)
// Register a custom scheme. This function should not be called for the built-
// in HTTP, HTTPS, FILE, FTP, ABOUT and DATA schemes.
//
// See cef_scheme_options_t for possible values for |options|.
//
// This function may be called on any thread. It should only be called once
// per unique |scheme_name| value. If |scheme_name| is already registered or
// if an error occurs this function will return false (0).
func (d *SchemeRegistrar) AddCustomScheme(schemeName string, options int32) int32 {
	schemeName_ := C.cef_string_userfree_alloc()
	setCEFStr(schemeName, schemeName_)
	defer func() {
		C.cef_string_userfree_free(schemeName_)
	}()
	return int32(C.gocef_scheme_registrar_add_custom_scheme(d.toNative(), (*C.cef_string_t)(schemeName_), C.int(options), d.add_custom_scheme))
}
