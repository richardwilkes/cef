// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "Client_gen.h"
	"C"
)

// ClientProxy defines methods required for using Client.
type ClientProxy interface {
	GetAudioHandler(self *Client) *AudioHandler
	GetContextMenuHandler(self *Client) *ContextMenuHandler
	GetDialogHandler(self *Client) *DialogHandler
	GetDisplayHandler(self *Client) *DisplayHandler
	GetDownloadHandler(self *Client) *DownloadHandler
	GetDragHandler(self *Client) *DragHandler
	GetFindHandler(self *Client) *FindHandler
	GetFocusHandler(self *Client) *FocusHandler
	GetJsdialogHandler(self *Client) *JsdialogHandler
	GetKeyboardHandler(self *Client) *KeyboardHandler
	GetLifeSpanHandler(self *Client) *LifeSpanHandler
	GetLoadHandler(self *Client) *LoadHandler
	GetRenderHandler(self *Client) *RenderHandler
	GetRequestHandler(self *Client) *RequestHandler
	OnProcessMessageReceived(self *Client, browser *Browser, source_process ProcessID, message *ProcessMessage) int32
}

// Client (cef_client_t from include/capi/cef_client_capi.h)
// Implement this structure to provide handler implementations.
type Client C.cef_client_t

// NewClient creates a new Client with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewClient(proxy ClientProxy) *Client {
	result := (*Client)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_client_t, proxy)))
	if proxy != nil {
		C.gocef_set_client_proxy(result.toNative())
	}
	return result
}

func (d *Client) toNative() *C.cef_client_t {
	return (*C.cef_client_t)(d)
}

func lookupClientProxy(obj *BaseRefCounted) ClientProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(ClientProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type ClientProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *Client) Base() *BaseRefCounted {
	return (*BaseRefCounted)(&d.base)
}

// GetAudioHandler (get_audio_handler)
// Return the handler for audio rendering events.
func (d *Client) GetAudioHandler() *AudioHandler {
	return lookupClientProxy(d.Base()).GetAudioHandler(d)
}

//nolint:gocritic
//export gocef_client_get_audio_handler
func gocef_client_get_audio_handler(self *C.cef_client_t) *C.cef_audio_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetAudioHandler(me__)).toNative()
}

// GetContextMenuHandler (get_context_menu_handler)
// Return the handler for context menus. If no handler is provided the default
// implementation will be used.
func (d *Client) GetContextMenuHandler() *ContextMenuHandler {
	return lookupClientProxy(d.Base()).GetContextMenuHandler(d)
}

//nolint:gocritic
//export gocef_client_get_context_menu_handler
func gocef_client_get_context_menu_handler(self *C.cef_client_t) *C.cef_context_menu_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetContextMenuHandler(me__)).toNative()
}

// GetDialogHandler (get_dialog_handler)
// Return the handler for dialogs. If no handler is provided the default
// implementation will be used.
func (d *Client) GetDialogHandler() *DialogHandler {
	return lookupClientProxy(d.Base()).GetDialogHandler(d)
}

//nolint:gocritic
//export gocef_client_get_dialog_handler
func gocef_client_get_dialog_handler(self *C.cef_client_t) *C.cef_dialog_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetDialogHandler(me__)).toNative()
}

// GetDisplayHandler (get_display_handler)
// Return the handler for browser display state events.
func (d *Client) GetDisplayHandler() *DisplayHandler {
	return lookupClientProxy(d.Base()).GetDisplayHandler(d)
}

//nolint:gocritic
//export gocef_client_get_display_handler
func gocef_client_get_display_handler(self *C.cef_client_t) *C.cef_display_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetDisplayHandler(me__)).toNative()
}

// GetDownloadHandler (get_download_handler)
// Return the handler for download events. If no handler is returned downloads
// will not be allowed.
func (d *Client) GetDownloadHandler() *DownloadHandler {
	return lookupClientProxy(d.Base()).GetDownloadHandler(d)
}

//nolint:gocritic
//export gocef_client_get_download_handler
func gocef_client_get_download_handler(self *C.cef_client_t) *C.cef_download_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetDownloadHandler(me__)).toNative()
}

// GetDragHandler (get_drag_handler)
// Return the handler for drag events.
func (d *Client) GetDragHandler() *DragHandler {
	return lookupClientProxy(d.Base()).GetDragHandler(d)
}

//nolint:gocritic
//export gocef_client_get_drag_handler
func gocef_client_get_drag_handler(self *C.cef_client_t) *C.cef_drag_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetDragHandler(me__)).toNative()
}

// GetFindHandler (get_find_handler)
// Return the handler for find result events.
func (d *Client) GetFindHandler() *FindHandler {
	return lookupClientProxy(d.Base()).GetFindHandler(d)
}

//nolint:gocritic
//export gocef_client_get_find_handler
func gocef_client_get_find_handler(self *C.cef_client_t) *C.cef_find_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetFindHandler(me__)).toNative()
}

// GetFocusHandler (get_focus_handler)
// Return the handler for focus events.
func (d *Client) GetFocusHandler() *FocusHandler {
	return lookupClientProxy(d.Base()).GetFocusHandler(d)
}

//nolint:gocritic
//export gocef_client_get_focus_handler
func gocef_client_get_focus_handler(self *C.cef_client_t) *C.cef_focus_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetFocusHandler(me__)).toNative()
}

// GetJsdialogHandler (get_jsdialog_handler)
// Return the handler for JavaScript dialogs. If no handler is provided the
// default implementation will be used.
func (d *Client) GetJsdialogHandler() *JsdialogHandler {
	return lookupClientProxy(d.Base()).GetJsdialogHandler(d)
}

//nolint:gocritic
//export gocef_client_get_jsdialog_handler
func gocef_client_get_jsdialog_handler(self *C.cef_client_t) *C.cef_jsdialog_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetJsdialogHandler(me__)).toNative()
}

// GetKeyboardHandler (get_keyboard_handler)
// Return the handler for keyboard events.
func (d *Client) GetKeyboardHandler() *KeyboardHandler {
	return lookupClientProxy(d.Base()).GetKeyboardHandler(d)
}

//nolint:gocritic
//export gocef_client_get_keyboard_handler
func gocef_client_get_keyboard_handler(self *C.cef_client_t) *C.cef_keyboard_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetKeyboardHandler(me__)).toNative()
}

// GetLifeSpanHandler (get_life_span_handler)
// Return the handler for browser life span events.
func (d *Client) GetLifeSpanHandler() *LifeSpanHandler {
	return lookupClientProxy(d.Base()).GetLifeSpanHandler(d)
}

//nolint:gocritic
//export gocef_client_get_life_span_handler
func gocef_client_get_life_span_handler(self *C.cef_client_t) *C.cef_life_span_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetLifeSpanHandler(me__)).toNative()
}

// GetLoadHandler (get_load_handler)
// Return the handler for browser load status events.
func (d *Client) GetLoadHandler() *LoadHandler {
	return lookupClientProxy(d.Base()).GetLoadHandler(d)
}

//nolint:gocritic
//export gocef_client_get_load_handler
func gocef_client_get_load_handler(self *C.cef_client_t) *C.cef_load_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetLoadHandler(me__)).toNative()
}

// GetRenderHandler (get_render_handler)
// Return the handler for off-screen rendering events.
func (d *Client) GetRenderHandler() *RenderHandler {
	return lookupClientProxy(d.Base()).GetRenderHandler(d)
}

//nolint:gocritic
//export gocef_client_get_render_handler
func gocef_client_get_render_handler(self *C.cef_client_t) *C.cef_render_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetRenderHandler(me__)).toNative()
}

// GetRequestHandler (get_request_handler)
// Return the handler for browser request events.
func (d *Client) GetRequestHandler() *RequestHandler {
	return lookupClientProxy(d.Base()).GetRequestHandler(d)
}

//nolint:gocritic
//export gocef_client_get_request_handler
func gocef_client_get_request_handler(self *C.cef_client_t) *C.cef_request_handler_t {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return (proxy__.GetRequestHandler(me__)).toNative()
}

// OnProcessMessageReceived (on_process_message_received)
// Called when a new message is received from a different process. Return true
// (1) if the message was handled or false (0) otherwise. Do not keep a
// reference to or attempt to access the message outside of this callback.
func (d *Client) OnProcessMessageReceived(browser *Browser, source_process ProcessID, message *ProcessMessage) int32 {
	return lookupClientProxy(d.Base()).OnProcessMessageReceived(d, browser, source_process, message)
}

//nolint:gocritic
//export gocef_client_on_process_message_received
func gocef_client_on_process_message_received(self *C.cef_client_t, browser *C.cef_browser_t, source_process C.cef_process_id_t, message *C.cef_process_message_t) C.int {
	me__ := (*Client)(self)
	proxy__ := lookupClientProxy(me__.Base())
	return C.int(proxy__.OnProcessMessageReceived(me__, (*Browser)(browser), ProcessID(source_process), (*ProcessMessage)(message)))
}
