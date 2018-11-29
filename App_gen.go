package cef

import (
	// #include "App_gen.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

// AppProxy defines methods required for using App.
type AppProxy interface {
	OnBeforeCommandLineProcessing(self *App, process_type string, command_line *CommandLine)
	OnRegisterCustomSchemes(self *App, registrar *SchemeRegistrar)
	GetResourceBundleHandler(self *App) *ResourceBundleHandler
	GetBrowserProcessHandler(self *App) *BrowserProcessHandler
	GetRenderProcessHandler(self *App) *RenderProcessHandler
}

// App (cef_app_t from include/capi/cef_app_capi.h)
// Implement this structure to provide handler implementations. Methods will be
// called by the process and/or thread indicated.
type App C.cef_app_t

// NewApp creates a new App with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewApp(proxy AppProxy) *App {
	result := (*App)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_app_t, proxy)))
	if proxy != nil {
		C.gocef_set_app_proxy(result.toNative())
	}
	return result
}

func (d *App) toNative() *C.cef_app_t {
	return (*C.cef_app_t)(d)
}

func lookupAppProxy(obj *BaseRefCounted) AppProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(AppProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type AppProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *App) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// OnBeforeCommandLineProcessing (on_before_command_line_processing)
// Provides an opportunity to view and/or modify command-line arguments before
// processing by CEF and Chromium. The |process_type| value will be NULL for
// the browser process. Do not keep a reference to the cef_command_line_t
// object passed to this function. The CefSettings.command_line_args_disabled
// value can be used to start with an NULL command-line object. Any values
// specified in CefSettings that equate to command-line arguments will be set
// before this function is called. Be cautious when using this function to
// modify command-line arguments for non-browser processes as this may result
// in undefined behavior including crashes.
func (d *App) OnBeforeCommandLineProcessing(process_type string, command_line *CommandLine) {
	lookupAppProxy(d.Base()).OnBeforeCommandLineProcessing(d, process_type, command_line)
}

//export gocef_app_on_before_command_line_processing
func gocef_app_on_before_command_line_processing(self *C.cef_app_t, process_type *C.cef_string_t, command_line *C.cef_command_line_t) {
	me__ := (*App)(self)
	proxy__ := lookupAppProxy(me__.Base())
	proxy__.OnBeforeCommandLineProcessing(me__, cefstrToString(process_type), (*CommandLine)(command_line))
}

// OnRegisterCustomSchemes (on_register_custom_schemes)
// Provides an opportunity to register custom schemes. Do not keep a reference
// to the |registrar| object. This function is called on the main thread for
// each process and the registered schemes should be the same across all
// processes.
func (d *App) OnRegisterCustomSchemes(registrar *SchemeRegistrar) {
	lookupAppProxy(d.Base()).OnRegisterCustomSchemes(d, registrar)
}

//export gocef_app_on_register_custom_schemes
func gocef_app_on_register_custom_schemes(self *C.cef_app_t, registrar *C.cef_scheme_registrar_t) {
	me__ := (*App)(self)
	proxy__ := lookupAppProxy(me__.Base())
	proxy__.OnRegisterCustomSchemes(me__, (*SchemeRegistrar)(registrar))
}

// GetResourceBundleHandler (get_resource_bundle_handler)
// Return the handler for resource bundle events. If
// CefSettings.pack_loading_disabled is true (1) a handler must be returned.
// If no handler is returned resources will be loaded from pack files. This
// function is called by the browser and render processes on multiple threads.
func (d *App) GetResourceBundleHandler() *ResourceBundleHandler {
	return lookupAppProxy(d.Base()).GetResourceBundleHandler(d)
}

//export gocef_app_get_resource_bundle_handler
func gocef_app_get_resource_bundle_handler(self *C.cef_app_t) *C.cef_resource_bundle_handler_t {
	me__ := (*App)(self)
	proxy__ := lookupAppProxy(me__.Base())
	return (proxy__.GetResourceBundleHandler(me__)).toNative()
}

// GetBrowserProcessHandler (get_browser_process_handler)
// Return the handler for functionality specific to the browser process. This
// function is called on multiple threads in the browser process.
func (d *App) GetBrowserProcessHandler() *BrowserProcessHandler {
	return lookupAppProxy(d.Base()).GetBrowserProcessHandler(d)
}

//export gocef_app_get_browser_process_handler
func gocef_app_get_browser_process_handler(self *C.cef_app_t) *C.cef_browser_process_handler_t {
	me__ := (*App)(self)
	proxy__ := lookupAppProxy(me__.Base())
	return (proxy__.GetBrowserProcessHandler(me__)).toNative()
}

// GetRenderProcessHandler (get_render_process_handler)
// Return the handler for functionality specific to the render process. This
// function is called on the render process main thread.
func (d *App) GetRenderProcessHandler() *RenderProcessHandler {
	return lookupAppProxy(d.Base()).GetRenderProcessHandler(d)
}

//export gocef_app_get_render_process_handler
func gocef_app_get_render_process_handler(self *C.cef_app_t) *C.cef_render_process_handler_t {
	me__ := (*App)(self)
	proxy__ := lookupAppProxy(me__.Base())
	return (proxy__.GetRenderProcessHandler(me__)).toNative()
}
