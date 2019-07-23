// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_string_userfree_t gocef_web_plugin_info_get_name(cef_web_plugin_info_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_web_plugin_info_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_web_plugin_info_get_path(cef_web_plugin_info_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_web_plugin_info_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_web_plugin_info_get_version(cef_web_plugin_info_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_web_plugin_info_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_web_plugin_info_get_description(cef_web_plugin_info_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_web_plugin_info_t *)) { return callback__(self); }
	"C"
)

// WebPluginInfo (cef_web_plugin_info_t from include/capi/cef_web_plugin_capi.h)
// Information about a specific web plugin.
type WebPluginInfo C.cef_web_plugin_info_t

func (d *WebPluginInfo) toNative() *C.cef_web_plugin_info_t {
	return (*C.cef_web_plugin_info_t)(d)
}

// Base (base)
// Base structure.
func (d *WebPluginInfo) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetName (get_name)
// Returns the plugin name (i.e. Flash).
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *WebPluginInfo) GetName() string {
	return cefuserfreestrToString(C.gocef_web_plugin_info_get_name(d.toNative(), d.get_name))
}

// GetPath (get_path)
// Returns the plugin file path (DLL/bundle/library).
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *WebPluginInfo) GetPath() string {
	return cefuserfreestrToString(C.gocef_web_plugin_info_get_path(d.toNative(), d.get_path))
}

// GetVersion (get_version)
// Returns the version of the plugin (may be OS-specific).
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *WebPluginInfo) GetVersion() string {
	return cefuserfreestrToString(C.gocef_web_plugin_info_get_version(d.toNative(), d.get_version))
}

// GetDescription (get_description)
// Returns a description of the plugin from the version information.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *WebPluginInfo) GetDescription() string {
	return cefuserfreestrToString(C.gocef_web_plugin_info_get_description(d.toNative(), d.get_description))
}
