// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "CookieVisitor_gen.h"
	"C"
)

// CookieVisitorProxy defines methods required for using CookieVisitor.
type CookieVisitorProxy interface {
	Visit(self *CookieVisitor, cookie *Cookie, count, total int32, deleteCookie *int32) int32
}

// CookieVisitor (cef_cookie_visitor_t from include/capi/cef_cookie_capi.h)
// Structure to implement for visiting cookie values. The functions of this
// structure will always be called on the UI thread.
type CookieVisitor C.cef_cookie_visitor_t

// NewCookieVisitor creates a new CookieVisitor with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewCookieVisitor(proxy CookieVisitorProxy) *CookieVisitor {
	result := (*CookieVisitor)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_cookie_visitor_t, proxy)))
	if proxy != nil {
		C.gocef_set_cookie_visitor_proxy(result.toNative())
	}
	return result
}

func (d *CookieVisitor) toNative() *C.cef_cookie_visitor_t {
	return (*C.cef_cookie_visitor_t)(d)
}

func lookupCookieVisitorProxy(obj *BaseRefCounted) CookieVisitorProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(CookieVisitorProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type CookieVisitorProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *CookieVisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method that will be called once for each cookie. |count| is the 0-based
// index for the current cookie. |total| is the total number of cookies. Set
// |deleteCookie| to true (1) to delete the cookie currently being visited.
// Return false (0) to stop visiting cookies. This function may never be
// called if no cookies are found.
func (d *CookieVisitor) Visit(cookie *Cookie, count, total int32, deleteCookie *int32) int32 {
	return lookupCookieVisitorProxy(d.Base()).Visit(d, cookie, count, total, deleteCookie)
}

//nolint:gocritic
//export gocef_cookie_visitor_visit
func gocef_cookie_visitor_visit(self *C.cef_cookie_visitor_t, cookie *C.cef_cookie_t, count C.int, total C.int, deleteCookie *C.int) C.int {
	me__ := (*CookieVisitor)(self)
	proxy__ := lookupCookieVisitorProxy(me__.Base())
	cookie_ := cookie.toGo()
	return C.int(proxy__.Visit(me__, cookie_, int32(count), int32(total), (*int32)(deleteCookie)))
}
