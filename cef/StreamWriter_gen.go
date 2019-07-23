// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// size_t gocef_stream_writer_write(cef_stream_writer_t * self, void * ptr, size_t size, size_t n, size_t (CEF_CALLBACK *callback__)(cef_stream_writer_t *, void *, size_t, size_t)) { return callback__(self, ptr, size, n); }
	// int gocef_stream_writer_seek(cef_stream_writer_t * self, int64 offset, int whence, int (CEF_CALLBACK *callback__)(cef_stream_writer_t *, int64, int)) { return callback__(self, offset, whence); }
	// int64 gocef_stream_writer_tell(cef_stream_writer_t * self, int64 (CEF_CALLBACK *callback__)(cef_stream_writer_t *)) { return callback__(self); }
	// int gocef_stream_writer_flush(cef_stream_writer_t * self, int (CEF_CALLBACK *callback__)(cef_stream_writer_t *)) { return callback__(self); }
	// int gocef_stream_writer_may_block(cef_stream_writer_t * self, int (CEF_CALLBACK *callback__)(cef_stream_writer_t *)) { return callback__(self); }
	"C"
)

// StreamWriter (cef_stream_writer_t from include/capi/cef_stream_capi.h)
// Structure used to write data to a stream. The functions of this structure may
// be called on any thread.
type StreamWriter C.cef_stream_writer_t

func (d *StreamWriter) toNative() *C.cef_stream_writer_t {
	return (*C.cef_stream_writer_t)(d)
}

// Base (base)
// Base structure.
func (d *StreamWriter) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Write (write)
// Write raw binary data.
func (d *StreamWriter) Write(ptr unsafe.Pointer, size, n uint64) uint64 {
	return uint64(C.gocef_stream_writer_write(d.toNative(), ptr, C.size_t(size), C.size_t(n), d.write))
}

// Seek (seek)
// Seek to the specified offset position. |whence| may be any one of SEEK_CUR,
// SEEK_END or SEEK_SET. Returns zero on success and non-zero on failure.
func (d *StreamWriter) Seek(offset int64, whence int32) int32 {
	return int32(C.gocef_stream_writer_seek(d.toNative(), C.int64(offset), C.int(whence), d.seek))
}

// Tell (tell)
// Return the current offset position.
func (d *StreamWriter) Tell() int64 {
	return int64(C.gocef_stream_writer_tell(d.toNative(), d.tell))
}

// Flush (flush)
// Flush the stream.
func (d *StreamWriter) Flush() int32 {
	return int32(C.gocef_stream_writer_flush(d.toNative(), d.flush))
}

// MayBlock (may_block)
// Returns true (1) if this writer performs work like accessing the file
// system which may block. Used as a hint for determining the thread to access
// the writer from.
func (d *StreamWriter) MayBlock() int32 {
	return int32(C.gocef_stream_writer_may_block(d.toNative(), d.may_block))
}
