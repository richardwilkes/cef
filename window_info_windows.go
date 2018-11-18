package cef

import (
	// #include "include/internal/cef_types.h"
	"C"
	"unsafe"
)

func (w *WindowInfo) platformInit(parent WindowHandle) {
	w.native.style = C.WS_CHILD | C.WS_CLIPCHILDREN | C.WS_CLIPSIBLINGS | C.WS_TABSTOP | C.WS_VISIBLE
	w.native.parent_window = C.cef_window_handle_t(unsafe.Pointer(parent))
}
