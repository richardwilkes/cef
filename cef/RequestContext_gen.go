// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	// int gocef_request_context_is_same(cef_request_context_t * self, cef_request_context_t * other, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_request_context_t *)) { return callback__(self, other); }
	// int gocef_request_context_is_sharing_with(cef_request_context_t * self, cef_request_context_t * other, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_request_context_t *)) { return callback__(self, other); }
	// int gocef_request_context_is_global(cef_request_context_t * self, int (CEF_CALLBACK *callback__)(cef_request_context_t *)) { return callback__(self); }
	// cef_request_context_handler_t * gocef_request_context_get_handler(cef_request_context_t * self, cef_request_context_handler_t * (CEF_CALLBACK *callback__)(cef_request_context_t *)) { return callback__(self); }
	// cef_string_userfree_t gocef_request_context_get_cache_path(cef_request_context_t * self, cef_string_userfree_t (CEF_CALLBACK *callback__)(cef_request_context_t *)) { return callback__(self); }
	// cef_cookie_manager_t * gocef_request_context_get_default_cookie_manager(cef_request_context_t * self, cef_completion_callback_t * callback, cef_cookie_manager_t * (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_completion_callback_t *)) { return callback__(self, callback); }
	// int gocef_request_context_register_scheme_handler_factory(cef_request_context_t * self, cef_string_t * scheme_name, cef_string_t * domain_name, cef_scheme_handler_factory_t * factory, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *, cef_string_t *, cef_scheme_handler_factory_t *)) { return callback__(self, scheme_name, domain_name, factory); }
	// int gocef_request_context_clear_scheme_handler_factories(cef_request_context_t * self, int (CEF_CALLBACK *callback__)(cef_request_context_t *)) { return callback__(self); }
	// void gocef_request_context_purge_plugin_list_cache(cef_request_context_t * self, int reload_pages, void (CEF_CALLBACK *callback__)(cef_request_context_t *, int)) { return callback__(self, reload_pages); }
	// int gocef_request_context_has_preference(cef_request_context_t * self, cef_string_t * name, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, name); }
	// cef_value_t * gocef_request_context_get_preference(cef_request_context_t * self, cef_string_t * name, cef_value_t * (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, name); }
	// cef_dictionary_value_t * gocef_request_context_get_all_preferences(cef_request_context_t * self, int include_defaults, cef_dictionary_value_t * (CEF_CALLBACK *callback__)(cef_request_context_t *, int)) { return callback__(self, include_defaults); }
	// int gocef_request_context_can_set_preference(cef_request_context_t * self, cef_string_t * name, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, name); }
	// int gocef_request_context_set_preference(cef_request_context_t * self, cef_string_t * name, cef_value_t * value, cef_string_t * error_r, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *, cef_value_t *, cef_string_t *)) { return callback__(self, name, value, error_r); }
	// void gocef_request_context_clear_certificate_exceptions(cef_request_context_t * self, cef_completion_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_completion_callback_t *)) { return callback__(self, callback); }
	// void gocef_request_context_close_all_connections(cef_request_context_t * self, cef_completion_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_completion_callback_t *)) { return callback__(self, callback); }
	// void gocef_request_context_resolve_host(cef_request_context_t * self, cef_string_t * origin, cef_resolve_callback_t * callback, void (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *, cef_resolve_callback_t *)) { return callback__(self, origin, callback); }
	// void gocef_request_context_load_extension(cef_request_context_t * self, cef_string_t * root_directory, cef_dictionary_value_t * manifest, cef_extension_handler_t * handler, void (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *, cef_dictionary_value_t *, cef_extension_handler_t *)) { return callback__(self, root_directory, manifest, handler); }
	// int gocef_request_context_did_load_extension(cef_request_context_t * self, cef_string_t * extension_id, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, extension_id); }
	// int gocef_request_context_has_extension(cef_request_context_t * self, cef_string_t * extension_id, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, extension_id); }
	// int gocef_request_context_get_extensions(cef_request_context_t * self, cef_string_list_t extension_ids, int (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_list_t)) { return callback__(self, extension_ids); }
	// cef_extension_t * gocef_request_context_get_extension(cef_request_context_t * self, cef_string_t * extension_id, cef_extension_t * (CEF_CALLBACK *callback__)(cef_request_context_t *, cef_string_t *)) { return callback__(self, extension_id); }
	"C"
)

// RequestContext (cef_request_context_t from include/capi/cef_request_context_capi.h)
// A request context provides request handling for a set of related browser or
// URL request objects. A request context can be specified when creating a new
// browser via the cef_browser_host_t static factory functions or when creating
// a new URL request via the cef_urlrequest_t static factory functions. Browser
// objects with different request contexts will never be hosted in the same
// render process. Browser objects with the same request context may or may not
// be hosted in the same render process depending on the process model. Browser
// objects created indirectly via the JavaScript window.open function or
// targeted links will share the same render process and the same request
// context as the source browser. When running in single-process mode there is
// only a single render process (the main process) and so all browsers created
// in single-process mode will share the same request context. This will be the
// first request context passed into a cef_browser_host_t static factory
// function and all other request context objects will be ignored.
type RequestContext C.cef_request_context_t

func (d *RequestContext) toNative() *C.cef_request_context_t {
	return (*C.cef_request_context_t)(d)
}

// Base (base)
// Base structure.
func (d *RequestContext) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// IsSame (is_same)
// Returns true (1) if this object is pointing to the same context as |that|
// object.
func (d *RequestContext) IsSame(other *RequestContext) int32 {
	return int32(C.gocef_request_context_is_same(d.toNative(), other.toNative(), d.is_same))
}

// IsSharingWith (is_sharing_with)
// Returns true (1) if this object is sharing the same storage as |that|
// object.
func (d *RequestContext) IsSharingWith(other *RequestContext) int32 {
	return int32(C.gocef_request_context_is_sharing_with(d.toNative(), other.toNative(), d.is_sharing_with))
}

// IsGlobal (is_global)
// Returns true (1) if this object is the global context. The global context
// is used by default when creating a browser or URL request with a NULL
// context argument.
func (d *RequestContext) IsGlobal() int32 {
	return int32(C.gocef_request_context_is_global(d.toNative(), d.is_global))
}

// GetHandler (get_handler)
// Returns the handler for this context if any.
func (d *RequestContext) GetHandler() *RequestContextHandler {
	return (*RequestContextHandler)(C.gocef_request_context_get_handler(d.toNative(), d.get_handler))
}

// GetCachePath (get_cache_path)
// Returns the cache path for this object. If NULL an "incognito mode" in-
// memory cache is being used.
// The resulting string must be freed by calling cef_string_userfree_free().
func (d *RequestContext) GetCachePath() string {
	return cefuserfreestrToString(C.gocef_request_context_get_cache_path(d.toNative(), d.get_cache_path))
}

// GetDefaultCookieManager (get_default_cookie_manager)
// Returns the default cookie manager for this object. This will be the global
// cookie manager if this object is the global request context. Otherwise,
// this will be the default cookie manager used when this request context does
// not receive a value via cef_request_tContextHandler::get_cookie_manager().
// If |callback| is non-NULL it will be executed asnychronously on the IO
// thread after the manager's storage has been initialized.
func (d *RequestContext) GetDefaultCookieManager(callback *CompletionCallback) *CookieManager {
	return (*CookieManager)(C.gocef_request_context_get_default_cookie_manager(d.toNative(), callback.toNative(), d.get_default_cookie_manager))
}

// RegisterSchemeHandlerFactory (register_scheme_handler_factory)
// Register a scheme handler factory for the specified |scheme_name| and
// optional |domain_name|. An NULL |domain_name| value for a standard scheme
// will cause the factory to match all domain names. The |domain_name| value
// will be ignored for non-standard schemes. If |scheme_name| is a built-in
// scheme and no handler is returned by |factory| then the built-in scheme
// handler factory will be called. If |scheme_name| is a custom scheme then
// you must also implement the cef_app_t::on_register_custom_schemes()
// function in all processes. This function may be called multiple times to
// change or remove the factory that matches the specified |scheme_name| and
// optional |domain_name|. Returns false (0) if an error occurs. This function
// may be called on any thread in the browser process.
func (d *RequestContext) RegisterSchemeHandlerFactory(scheme_name, domain_name string, factory *SchemeHandlerFactory) int32 {
	scheme_name_ := C.cef_string_userfree_alloc()
	setCEFStr(scheme_name, scheme_name_)
	defer func() {
		C.cef_string_userfree_free(scheme_name_)
	}()
	domain_name_ := C.cef_string_userfree_alloc()
	setCEFStr(domain_name, domain_name_)
	defer func() {
		C.cef_string_userfree_free(domain_name_)
	}()
	return int32(C.gocef_request_context_register_scheme_handler_factory(d.toNative(), (*C.cef_string_t)(scheme_name_), (*C.cef_string_t)(domain_name_), factory.toNative(), d.register_scheme_handler_factory))
}

// ClearSchemeHandlerFactories (clear_scheme_handler_factories)
// Clear all registered scheme handler factories. Returns false (0) on error.
// This function may be called on any thread in the browser process.
func (d *RequestContext) ClearSchemeHandlerFactories() int32 {
	return int32(C.gocef_request_context_clear_scheme_handler_factories(d.toNative(), d.clear_scheme_handler_factories))
}

// PurgePluginListCache (purge_plugin_list_cache)
// Tells all renderer processes associated with this context to throw away
// their plugin list cache. If |reload_pages| is true (1) they will also
// reload all pages with plugins.
// cef_request_tContextHandler::OnBeforePluginLoad may be called to rebuild
// the plugin list cache.
func (d *RequestContext) PurgePluginListCache(reload_pages int32) {
	C.gocef_request_context_purge_plugin_list_cache(d.toNative(), C.int(reload_pages), d.purge_plugin_list_cache)
}

// HasPreference (has_preference)
// Returns true (1) if a preference with the specified |name| exists. This
// function must be called on the browser process UI thread.
func (d *RequestContext) HasPreference(name string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return int32(C.gocef_request_context_has_preference(d.toNative(), (*C.cef_string_t)(name_), d.has_preference))
}

// GetPreference (get_preference)
// Returns the value for the preference with the specified |name|. Returns
// NULL if the preference does not exist. The returned object contains a copy
// of the underlying preference value and modifications to the returned object
// will not modify the underlying preference value. This function must be
// called on the browser process UI thread.
func (d *RequestContext) GetPreference(name string) *Value {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return (*Value)(C.gocef_request_context_get_preference(d.toNative(), (*C.cef_string_t)(name_), d.get_preference))
}

// GetAllPreferences (get_all_preferences)
// Returns all preferences as a dictionary. If |include_defaults| is true (1)
// then preferences currently at their default value will be included. The
// returned object contains a copy of the underlying preference values and
// modifications to the returned object will not modify the underlying
// preference values. This function must be called on the browser process UI
// thread.
func (d *RequestContext) GetAllPreferences(include_defaults int32) *DictionaryValue {
	return (*DictionaryValue)(C.gocef_request_context_get_all_preferences(d.toNative(), C.int(include_defaults), d.get_all_preferences))
}

// CanSetPreference (can_set_preference)
// Returns true (1) if the preference with the specified |name| can be
// modified using SetPreference. As one example preferences set via the
// command-line usually cannot be modified. This function must be called on
// the browser process UI thread.
func (d *RequestContext) CanSetPreference(name string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	return int32(C.gocef_request_context_can_set_preference(d.toNative(), (*C.cef_string_t)(name_), d.can_set_preference))
}

// SetPreference (set_preference)
// Set the |value| associated with preference |name|. Returns true (1) if the
// value is set successfully and false (0) otherwise. If |value| is NULL the
// preference will be restored to its default value. If setting the preference
// fails then |error| will be populated with a detailed description of the
// problem. This function must be called on the browser process UI thread.
func (d *RequestContext) SetPreference(name string, value *Value, error_r *string) int32 {
	name_ := C.cef_string_userfree_alloc()
	setCEFStr(name, name_)
	defer func() {
		C.cef_string_userfree_free(name_)
	}()
	error_r_ := C.cef_string_userfree_alloc()
	setCEFStr(*error_r, error_r_)
	defer func() {
		*error_r = cefstrToString(error_r_)
		C.cef_string_userfree_free(error_r_)
	}()
	return int32(C.gocef_request_context_set_preference(d.toNative(), (*C.cef_string_t)(name_), value.toNative(), (*C.cef_string_t)(error_r_), d.set_preference))
}

// ClearCertificateExceptions (clear_certificate_exceptions)
// Clears all certificate exceptions that were added as part of handling
// cef_request_tHandler::on_certificate_error(). If you call this it is
// recommended that you also call close_all_connections() or you risk not
// being prompted again for server certificates if you reconnect quickly. If
// |callback| is non-NULL it will be executed on the UI thread after
// completion.
func (d *RequestContext) ClearCertificateExceptions(callback *CompletionCallback) {
	C.gocef_request_context_clear_certificate_exceptions(d.toNative(), callback.toNative(), d.clear_certificate_exceptions)
}

// CloseAllConnections (close_all_connections)
// Clears all active and idle connections that Chromium currently has. This is
// only recommended if you have released all other CEF objects but don't yet
// want to call Cefshutdown(). If |callback| is non-NULL it will be executed
// on the UI thread after completion.
func (d *RequestContext) CloseAllConnections(callback *CompletionCallback) {
	C.gocef_request_context_close_all_connections(d.toNative(), callback.toNative(), d.close_all_connections)
}

// ResolveHost (resolve_host)
// Attempts to resolve |origin| to a list of associated IP addresses.
// |callback| will be executed on the UI thread after completion.
func (d *RequestContext) ResolveHost(origin string, callback *ResolveCallback) {
	origin_ := C.cef_string_userfree_alloc()
	setCEFStr(origin, origin_)
	defer func() {
		C.cef_string_userfree_free(origin_)
	}()
	C.gocef_request_context_resolve_host(d.toNative(), (*C.cef_string_t)(origin_), callback.toNative(), d.resolve_host)
}

// LoadExtension (load_extension)
// Load an extension.
//
// If extension resources will be read from disk using the default load
// implementation then |root_directory| should be the absolute path to the
// extension resources directory and |manifest| should be NULL. If extension
// resources will be provided by the client (e.g. via cef_request_tHandler
// and/or cef_extension_tHandler) then |root_directory| should be a path
// component unique to the extension (if not absolute this will be internally
// prefixed with the PK_DIR_RESOURCES path) and |manifest| should contain the
// contents that would otherwise be read from the "manifest.json" file on
// disk.
//
// The loaded extension will be accessible in all contexts sharing the same
// storage (HasExtension returns true (1)). However, only the context on which
// this function was called is considered the loader (DidLoadExtension returns
// true (1)) and only the loader will receive cef_request_tContextHandler
// callbacks for the extension.
//
// cef_extension_tHandler::OnExtensionLoaded will be called on load success or
// cef_extension_tHandler::OnExtensionLoadFailed will be called on load
// failure.
//
// If the extension specifies a background script via the "background"
// manifest key then cef_extension_tHandler::OnBeforeBackgroundBrowser will be
// called to create the background browser. See that function for additional
// information about background scripts.
//
// For visible extension views the client application should evaluate the
// manifest to determine the correct extension URL to load and then pass that
// URL to the cef_browser_host_t::CreateBrowser* function after the extension
// has loaded. For example, the client can look for the "browser_action"
// manifest key as documented at
// https://developer.chrome.com/extensions/browserAction. Extension URLs take
// the form "chrome-extension://<extension_id>/<path>".
//
// Browsers that host extensions differ from normal browsers as follows:
//  - Can access chrome.* JavaScript APIs if allowed by the manifest. Visit
//    chrome://extensions-support for the list of extension APIs currently
//    supported by CEF.
//  - Main frame navigation to non-extension content is blocked.
//  - Pinch-zooming is disabled.
//  - CefBrowserHost::GetExtension returns the hosted extension.
//  - CefBrowserHost::IsBackgroundHost returns true for background hosts.
//
// See https://developer.chrome.com/extensions for extension implementation
// and usage documentation.
func (d *RequestContext) LoadExtension(root_directory string, manifest *DictionaryValue, handler *ExtensionHandler) {
	root_directory_ := C.cef_string_userfree_alloc()
	setCEFStr(root_directory, root_directory_)
	defer func() {
		C.cef_string_userfree_free(root_directory_)
	}()
	C.gocef_request_context_load_extension(d.toNative(), (*C.cef_string_t)(root_directory_), manifest.toNative(), handler.toNative(), d.load_extension)
}

// DidLoadExtension (did_load_extension)
// Returns true (1) if this context was used to load the extension identified
// by |extension_id|. Other contexts sharing the same storage will also have
// access to the extension (see HasExtension). This function must be called on
// the browser process UI thread.
func (d *RequestContext) DidLoadExtension(extension_id string) int32 {
	extension_id_ := C.cef_string_userfree_alloc()
	setCEFStr(extension_id, extension_id_)
	defer func() {
		C.cef_string_userfree_free(extension_id_)
	}()
	return int32(C.gocef_request_context_did_load_extension(d.toNative(), (*C.cef_string_t)(extension_id_), d.did_load_extension))
}

// HasExtension (has_extension)
// Returns true (1) if this context has access to the extension identified by
// |extension_id|. This may not be the context that was used to load the
// extension (see DidLoadExtension). This function must be called on the
// browser process UI thread.
func (d *RequestContext) HasExtension(extension_id string) int32 {
	extension_id_ := C.cef_string_userfree_alloc()
	setCEFStr(extension_id, extension_id_)
	defer func() {
		C.cef_string_userfree_free(extension_id_)
	}()
	return int32(C.gocef_request_context_has_extension(d.toNative(), (*C.cef_string_t)(extension_id_), d.has_extension))
}

// GetExtensions (get_extensions)
// Retrieve the list of all extensions that this context has access to (see
// HasExtension). |extension_ids| will be populated with the list of extension
// ID values. Returns true (1) on success. This function must be called on the
// browser process UI thread.
func (d *RequestContext) GetExtensions(extension_ids StringList) int32 {
	return int32(C.gocef_request_context_get_extensions(d.toNative(), C.cef_string_list_t(extension_ids), d.get_extensions))
}

// GetExtension (get_extension)
// Returns the extension matching |extension_id| or NULL if no matching
// extension is accessible in this context (see HasExtension). This function
// must be called on the browser process UI thread.
func (d *RequestContext) GetExtension(extension_id string) *Extension {
	extension_id_ := C.cef_string_userfree_alloc()
	setCEFStr(extension_id, extension_id_)
	defer func() {
		C.cef_string_userfree_free(extension_id_)
	}()
	return (*Extension)(C.gocef_request_context_get_extension(d.toNative(), (*C.cef_string_t)(extension_id_), d.get_extension))
}
