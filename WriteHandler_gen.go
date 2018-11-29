// Code generated - DO NOT EDIT.

package cef

import (
	// #include "WriteHandler_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// WriteHandlerProxy defines methods required for using WriteHandler.
type WriteHandlerProxy interface {
	Write(self *WriteHandler, ptr unsafe.Pointer, size uint64, n uint64) uint64
	Seek(self *WriteHandler, offset int64, whence int32) int32
	Tell(self *WriteHandler) int64
	Flush(self *WriteHandler) int32
	MayBlock(self *WriteHandler) int32
}

// WriteHandler (cef_write_handler_t from include/capi/cef_stream_capi.h)
// Structure the client can implement to provide a custom stream writer. The
// functions of this structure may be called on any thread.
type WriteHandler C.cef_write_handler_t

// NewWriteHandler creates a new WriteHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewWriteHandler(proxy WriteHandlerProxy) *WriteHandler {
	result := (*WriteHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_write_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_write_handler_proxy(result.toNative())
	}
	return result
}

func (d *WriteHandler) toNative() *C.cef_write_handler_t {
	return (*C.cef_write_handler_t)(d)
}

func lookupWriteHandlerProxy(obj *BaseRefCounted) WriteHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(WriteHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type WriteHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *WriteHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Write (write)
// Write raw binary data.
func (d *WriteHandler) Write(ptr unsafe.Pointer, size, n uint64) uint64 {
	return lookupWriteHandlerProxy(d.Base()).Write(d, ptr, size, n)
}

//export gocef_write_handler_write
func gocef_write_handler_write(self *C.cef_write_handler_t, ptr unsafe.Pointer, size C.size_t, n C.size_t) C.size_t {
	me__ := (*WriteHandler)(self)
	proxy__ := lookupWriteHandlerProxy(me__.Base())
	return C.size_t(proxy__.Write(me__, ptr, uint64(size), uint64(n)))
}

// Seek (seek)
// Seek to the specified offset position. |whence| may be any one of SEEK_CUR,
// SEEK_END or SEEK_SET. Return zero on success and non-zero on failure.
func (d *WriteHandler) Seek(offset int64, whence int32) int32 {
	return lookupWriteHandlerProxy(d.Base()).Seek(d, offset, whence)
}

//export gocef_write_handler_seek
func gocef_write_handler_seek(self *C.cef_write_handler_t, offset C.int64, whence C.int) C.int {
	me__ := (*WriteHandler)(self)
	proxy__ := lookupWriteHandlerProxy(me__.Base())
	return C.int(proxy__.Seek(me__, int64(offset), int32(whence)))
}

// Tell (tell)
// Return the current offset position.
func (d *WriteHandler) Tell() int64 {
	return lookupWriteHandlerProxy(d.Base()).Tell(d)
}

//export gocef_write_handler_tell
func gocef_write_handler_tell(self *C.cef_write_handler_t) C.int64 {
	me__ := (*WriteHandler)(self)
	proxy__ := lookupWriteHandlerProxy(me__.Base())
	return C.int64(proxy__.Tell(me__))
}

// Flush (flush)
// Flush the stream.
func (d *WriteHandler) Flush() int32 {
	return lookupWriteHandlerProxy(d.Base()).Flush(d)
}

//export gocef_write_handler_flush
func gocef_write_handler_flush(self *C.cef_write_handler_t) C.int {
	me__ := (*WriteHandler)(self)
	proxy__ := lookupWriteHandlerProxy(me__.Base())
	return C.int(proxy__.Flush(me__))
}

// MayBlock (may_block)
// Return true (1) if this handler performs work like accessing the file
// system which may block. Used as a hint for determining the thread to access
// the handler from.
func (d *WriteHandler) MayBlock() int32 {
	return lookupWriteHandlerProxy(d.Base()).MayBlock(d)
}

//export gocef_write_handler_may_block
func gocef_write_handler_may_block(self *C.cef_write_handler_t) C.int {
	me__ := (*WriteHandler)(self)
	proxy__ := lookupWriteHandlerProxy(me__.Base())
	return C.int(proxy__.MayBlock(me__))
}
