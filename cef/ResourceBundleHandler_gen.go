// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "ResourceBundleHandler_gen.h"
	"C"
)

// ResourceBundleHandlerProxy defines methods required for using ResourceBundleHandler.
type ResourceBundleHandlerProxy interface {
	GetLocalizedString(self *ResourceBundleHandler, stringID int32, string_r *string) int32
	GetDataResource(self *ResourceBundleHandler, resourceID int32, data *unsafe.Pointer, dataSize *uint64) int32
	GetDataResourceForScale(self *ResourceBundleHandler, resourceID int32, scaleFactor ScaleFactor, data *unsafe.Pointer, dataSize *uint64) int32
}

// ResourceBundleHandler (cef_resource_bundle_handler_t from include/capi/cef_resource_bundle_handler_capi.h)
// Structure used to implement a custom resource bundle structure. See
// CefSettings for additional options related to resource bundle loading. The
// functions of this structure may be called on multiple threads.
type ResourceBundleHandler C.cef_resource_bundle_handler_t

// NewResourceBundleHandler creates a new ResourceBundleHandler with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewResourceBundleHandler(proxy ResourceBundleHandlerProxy) *ResourceBundleHandler {
	result := (*ResourceBundleHandler)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_resource_bundle_handler_t, proxy)))
	if proxy != nil {
		C.gocef_set_resource_bundle_handler_proxy(result.toNative())
	}
	return result
}

func (d *ResourceBundleHandler) toNative() *C.cef_resource_bundle_handler_t {
	return (*C.cef_resource_bundle_handler_t)(d)
}

func lookupResourceBundleHandlerProxy(obj *BaseRefCounted) ResourceBundleHandlerProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ResourceBundleHandlerProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ResourceBundleHandlerProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *ResourceBundleHandler) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetLocalizedString (get_localized_string)
// Called to retrieve a localized translation for the specified |string_id|.
// To provide the translation set |string| to the translation string and
// return true (1). To use the default translation return false (0). Include
// cef_pack_strings.h for a listing of valid string ID values.
func (d *ResourceBundleHandler) GetLocalizedString(stringID int32, string_r *string) int32 {
	return lookupResourceBundleHandlerProxy(d.Base()).GetLocalizedString(d, stringID, string_r)
}

//nolint:gocritic
//export gocef_resource_bundle_handler_get_localized_string
func gocef_resource_bundle_handler_get_localized_string(self *C.cef_resource_bundle_handler_t, stringID C.int, string_r *C.cef_string_t) C.int {
	me__ := (*ResourceBundleHandler)(self)
	proxy__ := lookupResourceBundleHandlerProxy(me__.Base())
	string_r_ := cefstrToString(string_r)
	return C.int(proxy__.GetLocalizedString(me__, int32(stringID), &string_r_))
}

// GetDataResource (get_data_resource)
// Called to retrieve data for the specified scale independent |resource_id|.
// To provide the resource data set |data| and |data_size| to the data pointer
// and size respectively and return true (1). To use the default resource data
// return false (0). The resource data will not be copied and must remain
// resident in memory. Include cef_pack_resources.h for a listing of valid
// resource ID values.
func (d *ResourceBundleHandler) GetDataResource(resourceID int32, data *unsafe.Pointer, dataSize *uint64) int32 {
	return lookupResourceBundleHandlerProxy(d.Base()).GetDataResource(d, resourceID, data, dataSize)
}

//nolint:gocritic
//export gocef_resource_bundle_handler_get_data_resource
func gocef_resource_bundle_handler_get_data_resource(self *C.cef_resource_bundle_handler_t, resourceID C.int, data *unsafe.Pointer, dataSize *C.size_t) C.int {
	me__ := (*ResourceBundleHandler)(self)
	proxy__ := lookupResourceBundleHandlerProxy(me__.Base())
	return C.int(proxy__.GetDataResource(me__, int32(resourceID), data, (*uint64)(dataSize)))
}

// GetDataResourceForScale (get_data_resource_for_scale)
// Called to retrieve data for the specified |resource_id| nearest the scale
// factor |scale_factor|. To provide the resource data set |data| and
// |data_size| to the data pointer and size respectively and return true (1).
// To use the default resource data return false (0). The resource data will
// not be copied and must remain resident in memory. Include
// cef_pack_resources.h for a listing of valid resource ID values.
func (d *ResourceBundleHandler) GetDataResourceForScale(resourceID int32, scaleFactor ScaleFactor, data *unsafe.Pointer, dataSize *uint64) int32 {
	return lookupResourceBundleHandlerProxy(d.Base()).GetDataResourceForScale(d, resourceID, scaleFactor, data, dataSize)
}

//nolint:gocritic
//export gocef_resource_bundle_handler_get_data_resource_for_scale
func gocef_resource_bundle_handler_get_data_resource_for_scale(self *C.cef_resource_bundle_handler_t, resourceID C.int, scaleFactor C.cef_scale_factor_t, data *unsafe.Pointer, dataSize *C.size_t) C.int {
	me__ := (*ResourceBundleHandler)(self)
	proxy__ := lookupResourceBundleHandlerProxy(me__.Base())
	return C.int(proxy__.GetDataResourceForScale(me__, int32(resourceID), ScaleFactor(scaleFactor), data, (*uint64)(dataSize)))
}
