// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ResponseFilter_gen.h"
	"C"
)

// ResponseFilterProxy defines methods required for using ResponseFilter.
type ResponseFilterProxy interface {
	InitFilter(self *ResponseFilter) int32
	Filter(self *ResponseFilter, dataIn unsafe.Pointer, dataInSize uint64, dataInRead *uint64, dataOut unsafe.Pointer, dataOutSize uint64, dataOutWritten *uint64) ResponseFilterStatus
}

// ResponseFilter (cef_response_filter_t from include/capi/cef_response_filter_capi.h)
// Implement this structure to filter resource response content. The functions
// of this structure will be called on the browser process IO thread.
type ResponseFilter C.cef_response_filter_t

// NewResponseFilter creates a new ResponseFilter with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResponseFilter(proxy ResponseFilterProxy) *ResponseFilter {
	result := (*ResponseFilter)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_response_filter_t, proxy)))
	if proxy != nil {
		C.gocef_set_response_filter_proxy(result.toNative())
	}
	return result
}

func (d *ResponseFilter) toNative() *C.cef_response_filter_t {
	return (*C.cef_response_filter_t)(d)
}

func lookupResponseFilterProxy(obj *BaseRefCounted) ResponseFilterProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResponseFilterProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResponseFilterProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResponseFilter) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// InitFilter (init_filter)
// Initialize the response filter. Will only be called a single time. The
// filter will not be installed if this function returns false (0).
func (d *ResponseFilter) InitFilter() int32 {
	return lookupResponseFilterProxy(d.Base()).InitFilter(d)
}

//nolint:gocritic
//export gocef_response_filter_init_filter
func gocef_response_filter_init_filter(self *C.cef_response_filter_t) C.int {
	me__ := (*ResponseFilter)(self)
	proxy__ := lookupResponseFilterProxy(me__.Base())
	return C.int(proxy__.InitFilter(me__))
}

// Filter (filter)
// Called to filter a chunk of data. Expected usage is as follows:
//
//  A. Read input data from |data_in| and set |data_in_read| to the number of
//     bytes that were read up to a maximum of |data_in_size|. |data_in| will
//     be NULL if |data_in_size| is zero.
//  B. Write filtered output data to |data_out| and set |data_out_written| to
//     the number of bytes that were written up to a maximum of
//     |data_out_size|. If no output data was written then all data must be
//     read from |data_in| (user must set |data_in_read| = |data_in_size|).
//  C. Return RESPONSE_FILTER_DONE if all output data was written or
//     RESPONSE_FILTER_NEED_MORE_DATA if output data is still pending.
//
// This function will be called repeatedly until the input buffer has been
// fully read (user sets |data_in_read| = |data_in_size|) and there is no more
// input data to filter (the resource response is complete). This function may
// then be called an additional time with an NULL input buffer if the user
// filled the output buffer (set |data_out_written| = |data_out_size|) and
// returned RESPONSE_FILTER_NEED_MORE_DATA to indicate that output data is
// still pending.
//
// Calls to this function will stop when one of the following conditions is
// met:
//
//  A. There is no more input data to filter (the resource response is
//     complete) and the user sets |data_out_written| = 0 or returns
//     RESPONSE_FILTER_DONE to indicate that all data has been written, or;
//  B. The user returns RESPONSE_FILTER_ERROR to indicate an error.
//
// Do not keep a reference to the buffers passed to this function.
func (d *ResponseFilter) Filter(dataIn unsafe.Pointer, dataInSize uint64, dataInRead *uint64, dataOut unsafe.Pointer, dataOutSize uint64, dataOutWritten *uint64) ResponseFilterStatus {
	return lookupResponseFilterProxy(d.Base()).Filter(d, dataIn, dataInSize, dataInRead, dataOut, dataOutSize, dataOutWritten)
}

//nolint:gocritic
//export gocef_response_filter_filter
func gocef_response_filter_filter(self *C.cef_response_filter_t, dataIn unsafe.Pointer, dataInSize C.size_t, dataInRead *C.size_t, dataOut unsafe.Pointer, dataOutSize C.size_t, dataOutWritten *C.size_t) C.cef_response_filter_status_t {
	me__ := (*ResponseFilter)(self)
	proxy__ := lookupResponseFilterProxy(me__.Base())
	return C.cef_response_filter_status_t(proxy__.Filter(me__, dataIn, uint64(dataInSize), (*uint64)(dataInRead), dataOut, uint64(dataOutSize), (*uint64)(dataOutWritten)))
}
