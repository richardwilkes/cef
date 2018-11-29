// Code generated - DO NOT EDIT.

package cef

import (
	// #include "NavigationEntryVisitor_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// NavigationEntryVisitorProxy defines methods required for using NavigationEntryVisitor.
type NavigationEntryVisitorProxy interface {
	Visit(self *NavigationEntryVisitor, entry *NavigationEntry, current int32, index int32, total int32) int32
}

// NavigationEntryVisitor (cef_navigation_entry_visitor_t from include/capi/cef_browser_capi.h)
// Callback structure for cef_browser_host_t::GetNavigationEntries. The
// functions of this structure will be called on the browser process UI thread.
type NavigationEntryVisitor C.cef_navigation_entry_visitor_t

// NewNavigationEntryVisitor creates a new NavigationEntryVisitor with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewNavigationEntryVisitor(proxy NavigationEntryVisitorProxy) *NavigationEntryVisitor {
	result := (*NavigationEntryVisitor)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_navigation_entry_visitor_t, proxy)))
	if proxy != nil {
		C.gocef_set_navigation_entry_visitor_proxy(result.toNative())
	}
	return result
}

func (d *NavigationEntryVisitor) toNative() *C.cef_navigation_entry_visitor_t {
	return (*C.cef_navigation_entry_visitor_t)(d)
}

func lookupNavigationEntryVisitorProxy(obj *BaseRefCounted) NavigationEntryVisitorProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(NavigationEntryVisitorProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type NavigationEntryVisitorProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *NavigationEntryVisitor) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// Visit (visit)
// Method that will be executed. Do not keep a reference to |entry| outside of
// this callback. Return true (1) to continue visiting entries or false (0) to
// stop. |current| is true (1) if this entry is the currently loaded
// navigation entry. |index| is the 0-based index of this entry and |total| is
// the total number of entries.
func (d *NavigationEntryVisitor) Visit(entry *NavigationEntry, current, index, total int32) int32 {
	return lookupNavigationEntryVisitorProxy(d.Base()).Visit(d, entry, current, index, total)
}

//export gocef_navigation_entry_visitor_visit
func gocef_navigation_entry_visitor_visit(self *C.cef_navigation_entry_visitor_t, entry *C.cef_navigation_entry_t, current C.int, index C.int, total C.int) C.int {
	me__ := (*NavigationEntryVisitor)(self)
	proxy__ := lookupNavigationEntryVisitorProxy(me__.Base())
	return C.int(proxy__.Visit(me__, (*NavigationEntry)(entry), int32(current), int32(index), int32(total)))
}
