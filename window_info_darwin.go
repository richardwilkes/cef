package cef

import (
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"
)

func (w *WindowInfo) platformInit(parent WindowHandle) {
	w.native.parent_view = C.cef_window_handle_t(unsafe.Pointer(parent))
}
