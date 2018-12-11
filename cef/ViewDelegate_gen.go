// Code generated - DO NOT EDIT.

package cef

import (
	// #include "ViewDelegate_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// ViewDelegateProxy defines methods required for using ViewDelegate.
type ViewDelegateProxy interface {
	GetPreferredSize(self *ViewDelegate, view *View) Size
	GetMinimumSize(self *ViewDelegate, view *View) Size
	GetMaximumSize(self *ViewDelegate, view *View) Size
	GetHeightForWidth(self *ViewDelegate, view *View, width int32) int32
	OnParentViewChanged(self *ViewDelegate, view *View, added int32, parent *View)
	OnChildViewChanged(self *ViewDelegate, view *View, added int32, child *View)
	OnFocus(self *ViewDelegate, view *View)
	OnBlur(self *ViewDelegate, view *View)
}

// ViewDelegate (cef_view_delegate_t from include/capi/views/cef_view_delegate_capi.h)
// Implement this structure to handle view events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type ViewDelegate C.cef_view_delegate_t

// NewViewDelegate creates a new ViewDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewViewDelegate(proxy ViewDelegateProxy) *ViewDelegate {
	result := (*ViewDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_view_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_view_delegate_proxy(result.toNative())
	}
	return result
}

func (d *ViewDelegate) toNative() *C.cef_view_delegate_t {
	return (*C.cef_view_delegate_t)(d)
}

func lookupViewDelegateProxy(obj *BaseRefCounted) ViewDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ViewDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ViewDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ViewDelegate) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetPreferredSize (get_preferred_size)
// Return the preferred size for |view|. The Layout will use this information
// to determine the display size.
func (d *ViewDelegate) GetPreferredSize(view *View) Size {
	return lookupViewDelegateProxy(d.Base()).GetPreferredSize(d, view)
}

//export gocef_view_delegate_get_preferred_size
func gocef_view_delegate_get_preferred_size(self *C.cef_view_delegate_t, view *C.cef_view_t) C.cef_size_t {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	call__ := proxy__.GetPreferredSize(me__, (*View)(view))
	var result__ C.cef_size_t
	call__.toNative(&result__)
	return result__
}

// GetMinimumSize (get_minimum_size)
// Return the minimum size for |view|.
func (d *ViewDelegate) GetMinimumSize(view *View) Size {
	return lookupViewDelegateProxy(d.Base()).GetMinimumSize(d, view)
}

//export gocef_view_delegate_get_minimum_size
func gocef_view_delegate_get_minimum_size(self *C.cef_view_delegate_t, view *C.cef_view_t) C.cef_size_t {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	call__ := proxy__.GetMinimumSize(me__, (*View)(view))
	var result__ C.cef_size_t
	call__.toNative(&result__)
	return result__
}

// GetMaximumSize (get_maximum_size)
// Return the maximum size for |view|.
func (d *ViewDelegate) GetMaximumSize(view *View) Size {
	return lookupViewDelegateProxy(d.Base()).GetMaximumSize(d, view)
}

//export gocef_view_delegate_get_maximum_size
func gocef_view_delegate_get_maximum_size(self *C.cef_view_delegate_t, view *C.cef_view_t) C.cef_size_t {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	call__ := proxy__.GetMaximumSize(me__, (*View)(view))
	var result__ C.cef_size_t
	call__.toNative(&result__)
	return result__
}

// GetHeightForWidth (get_height_for_width)
// Return the height necessary to display |view| with the provided |width|. If
// not specified the result of get_preferred_size().height will be used by
// default. Override if |view|'s preferred height depends upon the width (for
// example, with Labels).
func (d *ViewDelegate) GetHeightForWidth(view *View, width int32) int32 {
	return lookupViewDelegateProxy(d.Base()).GetHeightForWidth(d, view, width)
}

//export gocef_view_delegate_get_height_for_width
func gocef_view_delegate_get_height_for_width(self *C.cef_view_delegate_t, view *C.cef_view_t, width C.int) C.int {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	return C.int(proxy__.GetHeightForWidth(me__, (*View)(view), int32(width)))
}

// OnParentViewChanged (on_parent_view_changed)
// Called when the parent of |view| has changed. If |view| is being added to
// |parent| then |added| will be true (1). If |view| is being removed from
// |parent| then |added| will be false (0). If |view| is being reparented the
// remove notification will be sent before the add notification. Do not modify
// the view hierarchy in this callback.
func (d *ViewDelegate) OnParentViewChanged(view *View, added int32, parent *View) {
	lookupViewDelegateProxy(d.Base()).OnParentViewChanged(d, view, added, parent)
}

//export gocef_view_delegate_on_parent_view_changed
func gocef_view_delegate_on_parent_view_changed(self *C.cef_view_delegate_t, view *C.cef_view_t, added C.int, parent *C.cef_view_t) {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	proxy__.OnParentViewChanged(me__, (*View)(view), int32(added), (*View)(parent))
}

// OnChildViewChanged (on_child_view_changed)
// Called when a child of |view| has changed. If |child| is being added to
// |view| then |added| will be true (1). If |child| is being removed from
// |view| then |added| will be false (0). If |child| is being reparented the
// remove notification will be sent to the old parent before the add
// notification is sent to the new parent. Do not modify the view hierarchy in
// this callback.
func (d *ViewDelegate) OnChildViewChanged(view *View, added int32, child *View) {
	lookupViewDelegateProxy(d.Base()).OnChildViewChanged(d, view, added, child)
}

//export gocef_view_delegate_on_child_view_changed
func gocef_view_delegate_on_child_view_changed(self *C.cef_view_delegate_t, view *C.cef_view_t, added C.int, child *C.cef_view_t) {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	proxy__.OnChildViewChanged(me__, (*View)(view), int32(added), (*View)(child))
}

// OnFocus (on_focus)
// Called when |view| gains focus.
func (d *ViewDelegate) OnFocus(view *View) {
	lookupViewDelegateProxy(d.Base()).OnFocus(d, view)
}

//export gocef_view_delegate_on_focus
func gocef_view_delegate_on_focus(self *C.cef_view_delegate_t, view *C.cef_view_t) {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	proxy__.OnFocus(me__, (*View)(view))
}

// OnBlur (on_blur)
// Called when |view| loses focus.
func (d *ViewDelegate) OnBlur(view *View) {
	lookupViewDelegateProxy(d.Base()).OnBlur(d, view)
}

//export gocef_view_delegate_on_blur
func gocef_view_delegate_on_blur(self *C.cef_view_delegate_t, view *C.cef_view_t) {
	me__ := (*ViewDelegate)(self)
	proxy__ := lookupViewDelegateProxy(me__.Base())
	proxy__.OnBlur(me__, (*View)(view))
}
