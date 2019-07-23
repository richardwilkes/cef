// Code created from "class.go.tmpl" - don't edit by hand

package cef

import (
	// #include "capi_gen.h"
	// cef_string_userfree_t gocef_extension_get_identifier(cef_extension_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_extension_get_path(cef_extension_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// cef_dictionary_value_t * gocef_extension_get_manifest(cef_extension_t * self, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// int gocef_extension_is_same(cef_extension_t * self, cef_extension_t * that, int (CEF_CALLBACK *callback__)(cef_extension_t *, cef_extension_t *)) { return callback__(self, that); }
	// cef_extension_handler_t * gocef_extension_get_handler(cef_extension_t * self, cef_extension_handler_t * (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// cef_request_context_t * gocef_extension_get_loader_context(cef_extension_t * self, cef_request_context_t * (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// int gocef_extension_is_loaded(cef_extension_t * self, int (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	// void gocef_extension_unload(cef_extension_t * self, void (CEF_CALLBACK *callback__)(cef_extension_t *)) { return callback__(self); }
	"C"
)

// Extension (cef_extension_t from include/capi/cef_extension_capi.h)
// Object representing an extension. Methods may be called on any thread unless
// otherwise indicated.
type Extension C.cef_extension_t

func (d *Extension) toNative() *C.cef_extension_t {
	return (*C.cef_extension_t)(d)
}

// Base (base)
// Base structure.
func (d *Extension) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetIdentifier (get_identifier)
// Returns the unique extension identifier. This is calculated based on the
// extension public key, if available, or on the extension path. See
// https://developer.chrome.com/extensions/manifest/key for details.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Extension) GetIdentifier() string {
	return cefuserfreestrToString(C.gocef_extension_get_identifier(d.toNative(), d.get_identifier))
}

// GetPath (get_path)
// Returns the absolute path to the extension directory on disk. This value
// will be prefixed with PK_DIR_RESOURCES if a relative path was passed to
// cef_request_tContext::LoadExtension.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *Extension) GetPath() string {
	return cefuserfreestrToString(C.gocef_extension_get_path(d.toNative(), d.get_path))
}

// GetManifest (get_manifest)
// Returns the extension manifest contents as a cef_dictionary_value_t object.
// See https://developer.chrome.com/extensions/manifest for details.
func (d *Extension) GetManifest() *DictionaryValue {
	return (*DictionaryValue)(C.gocef_extension_get_manifest(d.toNative(), d.get_manifest))
}

// IsSame (is_same)
// Returns true (1) if this object is the same extension as |that| object.
// Extensions are considered the same if identifier, path and loader context
// match.
func (d *Extension) IsSame(that *Extension) int32 {
	return int32(C.gocef_extension_is_same(d.toNative(), that.toNative(), d.is_same))
}

// GetHandler (get_handler)
// Returns the handler for this extension. Will return NULL for internal
// extensions or if no handler was passed to
// cef_request_tContext::LoadExtension.
func (d *Extension) GetHandler() *ExtensionHandler {
	return (*ExtensionHandler)(C.gocef_extension_get_handler(d.toNative(), d.get_handler))
}

// GetLoaderContext (get_loader_context)
// Returns the request context that loaded this extension. Will return NULL
// for internal extensions or if the extension has been unloaded. See the
// cef_request_tContext::LoadExtension documentation for more information
// about loader contexts. Must be called on the browser process UI thread.
func (d *Extension) GetLoaderContext() *RequestContext {
	return (*RequestContext)(C.gocef_extension_get_loader_context(d.toNative(), d.get_loader_context))
}

// IsLoaded (is_loaded)
// Returns true (1) if this extension is currently loaded. Must be called on
// the browser process UI thread.
func (d *Extension) IsLoaded() int32 {
	return int32(C.gocef_extension_is_loaded(d.toNative(), d.is_loaded))
}

// Unload (unload)
// Unload this extension if it is not an internal extension and is currently
// loaded. Will result in a call to
// cef_extension_tHandler::OnExtensionUnloaded on success.
func (d *Extension) Unload() {
	C.gocef_extension_unload(d.toNative(), d.unload)
}
