// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// size_t gocef_stream_reader_read(cef_stream_reader_t * self, void * ptr, size_t size, size_t n, size_t (CEF_CALLBACK *callback__)(cef_stream_reader_t *, void *, size_t, size_t)) { return callback__(self, ptr, size, n); }
	// int gocef_stream_reader_seek(cef_stream_reader_t * self, int64 offset, int whence, int (CEF_CALLBACK *callback__)(cef_stream_reader_t *, int64, int)) { return callback__(self, offset, whence); }
	// int64 gocef_stream_reader_tell(cef_stream_reader_t * self, int64 (CEF_CALLBACK *callback__)(cef_stream_reader_t *)) { return callback__(self); }
	// int gocef_stream_reader_eof(cef_stream_reader_t * self, int (CEF_CALLBACK *callback__)(cef_stream_reader_t *)) { return callback__(self); }
	// int gocef_stream_reader_may_block(cef_stream_reader_t * self, int (CEF_CALLBACK *callback__)(cef_stream_reader_t *)) { return callback__(self); }
	"C"
)

// StreamReader (cef_stream_reader_t from include/capi/cef_stream_capi.h)
// Structure used to read data from a stream. The functions of this structure
// may be called on any thread.
type StreamReader C.cef_stream_reader_t

func (d *StreamReader) toNative() *C.cef_stream_reader_t {
	return (*C.cef_stream_reader_t)(d)
}

// Base (base)
// Base structure.
func (d *StreamReader) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Read (read)
// Read raw binary data.
func (d *StreamReader) Read(ptr unsafe.Pointer, size, n uint64) uint64 {
	return uint64(C.gocef_stream_reader_read(d.toNative(), ptr, C.size_t(size), C.size_t(n), d.read))
}

// Seek (seek)
// Seek to the specified offset position. |whence| may be any one of SEEK_CUR,
// SEEK_END or SEEK_SET. Returns zero on success and non-zero on failure.
func (d *StreamReader) Seek(offset int64, whence int32) int32 {
	return int32(C.gocef_stream_reader_seek(d.toNative(), C.int64(offset), C.int(whence), d.seek))
}

// Tell (tell)
// Return the current offset position.
func (d *StreamReader) Tell() int64 {
	return int64(C.gocef_stream_reader_tell(d.toNative(), d.tell))
}

// EOF (eof)
// Return non-zero if at end of file.
func (d *StreamReader) EOF() int32 {
	return int32(C.gocef_stream_reader_eof(d.toNative(), d.eof))
}

// MayBlock (may_block)
// Returns true (1) if this reader performs work like accessing the file
// system which may block. Used as a hint for determining the thread to access
// the reader from.
func (d *StreamReader) MayBlock() int32 {
	return int32(C.gocef_stream_reader_may_block(d.toNative(), d.may_block))
}
