// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

// +build !windows

package cef

import (
	"os"
	"unsafe"
)

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

func (d *MainArgs) platformInit() {
	if len(os.Args) > 0 {
		cp := C.calloc(C.size_t(len(os.Args)), C.size_t(unsafe.Sizeof(uintptr(0))))
		p := (*[1<<30 - 1]*C.char)(cp)
		for i, arg := range os.Args {
			p[i] = C.CString(arg)
		}
		d.native.argc = C.int(len(os.Args))
		d.native.argv = (**C.char)(cp)
	}
}
