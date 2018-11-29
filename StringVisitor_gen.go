// Code generated - DO NOT EDIT.

package cef

import (
	// #include "StringVisitor_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// StringVisitorProxy defines methods required for using StringVisitor.
type StringVisitorProxy interface {
	Visit(self *StringVisitor, string string)
}

// StringVisitor (cef_string_visitor_t from include/capi/cef_string_visitor_capi.h)
// Implement this structure to receive string values asynchronously.
type StringVisitor C.cef_string_visitor_t

// NewStringVisitor creates a new StringVisitor with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewStringVisitor(proxy StringVisitorProxy) *StringVisitor {
	result := (*StringVisitor)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_string_visitor_t, proxy)))
	if proxy != nil {
		C.gocef_set_string_visitor_proxy(result.toNative())
	}
	return result
}

func (d *StringVisitor) toNative() *C.cef_string_visitor_t {
	return (*C.cef_string_visitor_t)(d)
}

func lookupStringVisitorProxy(obj *BaseRefCounted) StringVisitorProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(StringVisitorProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type StringVisitorProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *StringVisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method that will be executed.
func (d *StringVisitor) Visit(string string) {
	lookupStringVisitorProxy(d.Base()).Visit(d, string)
}

//export gocef_string_visitor_visit
func gocef_string_visitor_visit(self *C.cef_string_visitor_t, string *C.cef_string_t) {
	me__ := (*StringVisitor)(self)
	proxy__ := lookupStringVisitorProxy(me__.Base())
	proxy__.Visit(me__, cefstrToString(string))
}
