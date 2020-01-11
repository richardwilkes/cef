// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// Code created from "typedef.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "include/internal/cef_string.h"
	// #include "include/internal/cef_string_list.h"
	// #include "include/internal/cef_string_map.h"
	// #include "include/internal/cef_string_multimap.h"
	// #include "include/internal/cef_types.h"
	"C"
)

// Char (cef_char_t from include/internal/cef_string.h)
type Char uint16

// Color (cef_color_t from include/internal/cef_types.h)
// 32-bit ARGB color value, not premultiplied. The color components are always
// in a known order. Equivalent to the SkColor type.
type Color uint32

// StringList (cef_string_list_t from include/internal/cef_string_list.h)
// CEF string maps are a set of key/value string pairs.
type StringList unsafe.Pointer

// StringMap (cef_string_map_t from include/internal/cef_string_map.h)
// CEF string maps are a set of key/value string pairs.
type StringMap unsafe.Pointer

// StringMultimap (cef_string_multimap_t from include/internal/cef_string_multimap.h)
// CEF string multimaps are a set of key/value string pairs.
// More than one value can be assigned to a single key.
type StringMultimap unsafe.Pointer
