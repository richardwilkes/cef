package cef

import (
	// #include "include/capi/cef_client_capi.h"
	// #include "ref.h"
	"C"
	"unsafe"
)

type Client *C.cef_client_t

// NewClient creates a new default Client instance.
func NewClient() Client {
	return Client(unsafe.Pointer(C.gocef_refcnt_alloc(C.sizeof_struct__cef_client_t)))
}
