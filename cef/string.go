package cef

import "unsafe"

import (
	// #include <stdlib.h>
	// #include <string.h>
	// #include "include/capi/cef_base_capi.h"
	"C"
)

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
