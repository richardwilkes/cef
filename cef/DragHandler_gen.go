// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "DragHandler_gen.h"
	"C"
)

// DragHandlerProxy defines methods required for using DragHandler.
type DragHandlerProxy interface {
	OnDragEnter(self *DragHandler, browser *Browser, dragData *DragData, mask DragOperationsMask) int32
	OnDraggableRegionsChanged(self *DragHandler, browser *Browser, frame *Frame, regionsCount uint64, regions *DraggableRegion)
}

// DragHandler (cef_drag_handler_t from include/capi/cef_drag_handler_capi.h)
// Implement this structure to handle events related to dragging. The functions
// of this structure will be called on the UI thread.
type DragHandler C.cef_drag_handler_t

// NewDragHandler creates a new DragHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewDragHandler(proxy DragHandlerProxy) *DragHandler {
	result := (*DragHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_drag_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_drag_handler_proxy(result.toNative())
	}
	return result
}

func (d *DragHandler) toNative() *C.cef_drag_handler_t {
	return (*C.cef_drag_handler_t)(d)
}

func lookupDragHandlerProxy(obj *BaseRefCounted) DragHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(DragHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type DragHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *DragHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnDragEnter (on_drag_enter)
// Called when an external drag event enters the browser window. |dragData|
// contains the drag event data and |mask| represents the type of drag
// operation. Return false (0) for default drag handling behavior or true (1)
// to cancel the drag event.
func (d *DragHandler) OnDragEnter(browser *Browser, dragData *DragData, mask DragOperationsMask) int32 {
	return lookupDragHandlerProxy(d.Base()).OnDragEnter(d, browser, dragData, mask)
}

//nolint:gocritic
//export gocef_drag_handler_on_drag_enter
func gocef_drag_handler_on_drag_enter(self *C.cef_drag_handler_t, browser *C.cef_browser_t, dragData *C.cef_drag_data_t, mask C.cef_drag_operations_mask_t) C.int {
	me__ := (*DragHandler)(self)
	proxy__ := lookupDragHandlerProxy(me__.Base())
	return C.int(proxy__.OnDragEnter(me__, (*Browser)(browser), (*DragData)(dragData), DragOperationsMask(mask)))
}

// OnDraggableRegionsChanged (on_draggable_regions_changed)
// Called whenever draggable regions for the browser window change. These can
// be specified using the '-webkit-app-region: drag/no-drag' CSS-property. If
// draggable regions are never defined in a document this function will also
// never be called. If the last draggable region is removed from a document
// this function will be called with an NULL vector.
func (d *DragHandler) OnDraggableRegionsChanged(browser *Browser, frame *Frame, regionsCount uint64, regions *DraggableRegion) {
	lookupDragHandlerProxy(d.Base()).OnDraggableRegionsChanged(d, browser, frame, regionsCount, regions)
}

//nolint:gocritic
//export gocef_drag_handler_on_draggable_regions_changed
func gocef_drag_handler_on_draggable_regions_changed(self *C.cef_drag_handler_t, browser *C.cef_browser_t, frame *C.cef_frame_t, regionsCount C.size_t, regions *C.cef_draggable_region_t) {
	me__ := (*DragHandler)(self)
	proxy__ := lookupDragHandlerProxy(me__.Base())
	regions_ := regions.toGo()
	proxy__.OnDraggableRegionsChanged(me__, (*Browser)(browser), (*Frame)(frame), uint64(regionsCount), regions_)
}
