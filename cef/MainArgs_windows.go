// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cef

import (
	"syscall"
	"unsafe"
)

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

func (d *MainArgs) platformInit() {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetModuleHandleW")
	if h, _, _ := proc.Call(0); h != 0 {
		d.native.instance = C.HINSTANCE(unsafe.Pointer(h))
	}
}
