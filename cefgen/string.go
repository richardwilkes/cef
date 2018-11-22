package cef

import (
	// #include <stdlib.h>
	// #include <string.h>
	// #include "include/capi/cef_base_capi.h"
	// #cgo CFLAGS: -I ${SRCDIR}/../cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/../cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/../cef/Release -lcef -Wl,--subsystem,windows
	"C"
	"unsafe"
)

func newCEFStr(str string) *C.cef_string_t {
	utf8str := C.CString(str)
	cefstr := (*C.cef_string_t)(C.calloc(1, C.size_t(unsafe.Sizeof(C.cef_string_t{}))))
	C.cef_string_from_utf8(utf8str, C.strlen(utf8str), cefstr)
	C.free(unsafe.Pointer(utf8str))
	return cefstr
}

func setCEFStr(str string, cefStr *C.cef_string_t) {
	utf8str := C.CString(str)
	C.cef_string_from_utf8(utf8str, C.strlen(utf8str), cefStr)
	C.free(unsafe.Pointer(utf8str))
}

func cefstrToString(cefstr *C.cef_string_t) string {
	if cefstr == nil {
		return ""
	}
	utf8str := C.cef_string_userfree_utf8_alloc()
	C.cef_string_to_utf8(cefstr.str, cefstr.length, utf8str)
	str := C.GoString(utf8str.str)
	C.cef_string_userfree_utf8_free(utf8str)
	return str
}

func cefuserfreestrToString(cefstr C.cef_string_userfree_t) string {
	if cefstr == nil {
		return ""
	}
	utf8str := C.cef_string_userfree_utf8_alloc()
	C.cef_string_to_utf8(cefstr.str, cefstr.length, utf8str)
	str := C.GoString(utf8str.str)
	C.cef_string_userfree_utf8_free(utf8str)
	C.cef_string_userfree_free(cefstr)
	return str
}
