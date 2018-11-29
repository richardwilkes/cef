// Code generated - DO NOT EDIT.

package cef

import (
	// #include "ReadHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// ReadHandlerProxy defines methods required for using ReadHandler.
type ReadHandlerProxy interface {
	Read(self *ReadHandler, ptr unsafe.Pointer, size uint64, n uint64) uint64
	Seek(self *ReadHandler, offset int64, whence int32) int32
	Tell(self *ReadHandler) int64
	Eof(self *ReadHandler) int32
	MayBlock(self *ReadHandler) int32
}

// ReadHandler (cef_read_handler_t from include/capi/cef_stream_capi.h)
// Structure the client can implement to provide a custom stream reader. The
// functions of this structure may be called on any thread.
type ReadHandler C.cef_read_handler_t

// NewReadHandler creates a new ReadHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewReadHandler(proxy ReadHandlerProxy) *ReadHandler {
	result := (*ReadHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_read_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_read_handler_proxy(result.toNative())
	}
	return result
}

func (d *ReadHandler) toNative() *C.cef_read_handler_t {
	return (*C.cef_read_handler_t)(d)
}

func lookupReadHandlerProxy(obj *BaseRefCounted) ReadHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ReadHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ReadHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ReadHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Read (read)
// Read raw binary data.
func (d *ReadHandler) Read(ptr unsafe.Pointer, size, n uint64) uint64 {
	return lookupReadHandlerProxy(d.Base()).Read(d, ptr, size, n)
}

//export gocef_read_handler_read
func gocef_read_handler_read(self *C.cef_read_handler_t, ptr unsafe.Pointer, size C.size_t, n C.size_t) C.size_t {
	me__ := (*ReadHandler)(self)
	proxy__ := lookupReadHandlerProxy(me__.Base())
	return C.size_t(proxy__.Read(me__, ptr, uint64(size), uint64(n)))
}

// Seek (seek)
// Seek to the specified offset position. |whence| may be any one of SEEK_CUR,
// SEEK_END or SEEK_SET. Return zero on success and non-zero on failure.
func (d *ReadHandler) Seek(offset int64, whence int32) int32 {
	return lookupReadHandlerProxy(d.Base()).Seek(d, offset, whence)
}

//export gocef_read_handler_seek
func gocef_read_handler_seek(self *C.cef_read_handler_t, offset C.int64, whence C.int) C.int {
	me__ := (*ReadHandler)(self)
	proxy__ := lookupReadHandlerProxy(me__.Base())
	return C.int(proxy__.Seek(me__, int64(offset), int32(whence)))
}

// Tell (tell)
// Return the current offset position.
func (d *ReadHandler) Tell() int64 {
	return lookupReadHandlerProxy(d.Base()).Tell(d)
}

//export gocef_read_handler_tell
func gocef_read_handler_tell(self *C.cef_read_handler_t) C.int64 {
	me__ := (*ReadHandler)(self)
	proxy__ := lookupReadHandlerProxy(me__.Base())
	return C.int64(proxy__.Tell(me__))
}

// Eof (eof)
// Return non-zero if at end of file.
func (d *ReadHandler) Eof() int32 {
	return lookupReadHandlerProxy(d.Base()).Eof(d)
}

//export gocef_read_handler_eof
func gocef_read_handler_eof(self *C.cef_read_handler_t) C.int {
	me__ := (*ReadHandler)(self)
	proxy__ := lookupReadHandlerProxy(me__.Base())
	return C.int(proxy__.Eof(me__))
}

// MayBlock (may_block)
// Return true (1) if this handler performs work like accessing the file
// system which may block. Used as a hint for determining the thread to access
// the handler from.
func (d *ReadHandler) MayBlock() int32 {
	return lookupReadHandlerProxy(d.Base()).MayBlock(d)
}

//export gocef_read_handler_may_block
func gocef_read_handler_may_block(self *C.cef_read_handler_t) C.int {
	me__ := (*ReadHandler)(self)
	proxy__ := lookupReadHandlerProxy(me__.Base())
	return C.int(proxy__.MayBlock(me__))
}
