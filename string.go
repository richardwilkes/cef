package cef

import (
	// #include <stdlib.h>
	// #include <string.h>
	// #include "include/capi/cef_base_capi.h"
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
