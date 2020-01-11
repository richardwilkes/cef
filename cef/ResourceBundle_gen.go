// Code created from "class.go.tmpl" - don't edit by hand

package cef

import "unsafe"

import (
	// #include "capi_gen.h"
	// cef_string_userfree_t gocef_resource_bundle_get_localized_string(cef_resource_bundle_t * self, int stringID, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_resource_bundle_t *, int)) { return callback__(self, stringID); }
	// int gocef_resource_bundle_get_data_resource(cef_resource_bundle_t * self, int resourceID, void ** data, size_t * dataSize, int (CEF_CALLBACK *callback__)(cef_resource_bundle_t *, int, void **, size_t *)) { return callback__(self, resourceID, data, dataSize); }
	// int gocef_resource_bundle_get_data_resource_for_scale(cef_resource_bundle_t * self, int resourceID, cef_scale_factor_t scaleFactor, void ** data, size_t * dataSize, int (CEF_CALLBACK *callback__)(cef_resource_bundle_t *, int, cef_scale_factor_t, void **, size_t *)) { return callback__(self, resourceID, scaleFactor, data, dataSize); }
	"C"
)

// ResourceBundle (cef_resource_bundle_t from include/capi/cef_resource_bundle_capi.h)
// Structure used for retrieving resources from the resource bundle (*.pak)
// files loaded by CEF during startup or via the cef_resource_bundle_tHandler
// returned from cef_app_t::GetResourceBundleHandler. See CefSettings for
// additional options related to resource bundle loading. The functions of this
// structure may be called on any thread unless otherwise indicated.
type ResourceBundle C.cef_resource_bundle_t

func (d *ResourceBundle) toNative() *C.cef_resource_bundle_t {
	return (*C.cef_resource_bundle_t)(d)
}

// Base (base)
// Base structure.
func (d *ResourceBundle) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetLocalizedString (get_localized_string)
// Returns the localized string for the specified |string_id| or an NULL
// string if the value is not found. Include cef_pack_strings.h for a listing
// of valid string ID values.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *ResourceBundle) GetLocalizedString(stringID int32) string {
	return cefuserfreestrToString(C.gocef_resource_bundle_get_localized_string(d.toNative(), C.int(stringID), d.get_localized_string))
}

// GetDataResource (get_data_resource)
// Retrieves the contents of the specified scale independent |resource_id|. If
// the value is found then |data| and |data_size| will be populated and this
// function will return true (1). If the value is not found then this function
// will return false (0). The returned |data| pointer will remain resident in
// memory and should not be freed. Include cef_pack_resources.h for a listing
// of valid resource ID values.
func (d *ResourceBundle) GetDataResource(resourceID int32, data *unsafe.Pointer, dataSize *uint64) int32 {
	return int32(C.gocef_resource_bundle_get_data_resource(d.toNative(), C.int(resourceID), data, (*C.size_t)(dataSize), d.get_data_resource))
}

// GetDataResourceForScale (get_data_resource_for_scale)
// Retrieves the contents of the specified |resource_id| nearest the scale
// factor |scale_factor|. Use a |scale_factor| value of SCALE_FACTOR_NONE for
// scale independent resources or call GetDataResource instead. If the value
// is found then |data| and |data_size| will be populated and this function
// will return true (1). If the value is not found then this function will
// return false (0). The returned |data| pointer will remain resident in
// memory and should not be freed. Include cef_pack_resources.h for a listing
// of valid resource ID values.
func (d *ResourceBundle) GetDataResourceForScale(resourceID int32, scaleFactor ScaleFactor, data *unsafe.Pointer, dataSize *uint64) int32 {
	return int32(C.gocef_resource_bundle_get_data_resource_for_scale(d.toNative(), C.int(resourceID), C.cef_scale_factor_t(scaleFactor), data, (*C.size_t)(dataSize), d.get_data_resource_for_scale))
}
