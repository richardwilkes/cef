// Code created from "callback.go.tmpl" - don't edit by hand

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

import (
	// #include "BrowserViewDelegate_gen.h"
	"C"
)

// BrowserViewDelegateProxy defines methods required for using BrowserViewDelegate.
type BrowserViewDelegateProxy interface {
	OnBrowserCreated(self *BrowserViewDelegate, browser_view *BrowserView, browser *Browser)
	OnBrowserDestroyed(self *BrowserViewDelegate, browser_view *BrowserView, browser *Browser)
	GetDelegateForPopupBrowserView(self *BrowserViewDelegate, browser_view *BrowserView, settings *BrowserSettings, client *Client, is_devtools int32) *BrowserViewDelegate
	OnPopupBrowserViewCreated(self *BrowserViewDelegate, browser_view, popup_browser_view *BrowserView, is_devtools int32) int32
}

// BrowserViewDelegate (cef_browser_view_delegate_t from include/capi/views/cef_browser_view_delegate_capi.h)
// Implement this structure to handle BrowserView events. The functions of this
// structure will be called on the browser process UI thread unless otherwise
// indicated.
type BrowserViewDelegate C.cef_browser_view_delegate_t

// NewBrowserViewDelegate creates a new BrowserViewDelegate with the specified proxy. Passing
// in nil will result in default handling, if applicable.
func NewBrowserViewDelegate(proxy BrowserViewDelegateProxy) *BrowserViewDelegate {
	result := (*BrowserViewDelegate)(unsafe.Pointer(newRefCntObj(C.sizeof_struct__cef_browser_view_delegate_t, proxy)))
	if proxy != nil {
		C.gocef_set_browser_view_delegate_proxy(result.toNative())
	}
	return result
}

func (d *BrowserViewDelegate) toNative() *C.cef_browser_view_delegate_t {
	return (*C.cef_browser_view_delegate_t)(d)
}

func lookupBrowserViewDelegateProxy(obj *BaseRefCounted) BrowserViewDelegateProxy {
	proxy, exists := lookupProxy(obj)
	if !exists {
		jot.Fatal(1, errs.New("Proxy not found for ID"))
	}
	actual, ok := proxy.(BrowserViewDelegateProxy)
	if !ok {
		jot.Fatal(1, errs.New("Proxy was not of type BrowserViewDelegateProxy"))
	}
	return actual
}

// Base (base)
// Base structure.
func (d *BrowserViewDelegate) Base() *ViewDelegate {
	return (*ViewDelegate)(&d.base)
}

// OnBrowserCreated (on_browser_created)
// Called when |browser| associated with |browser_view| is created. This
// function will be called after cef_life_span_handler_t::on_after_created()
// is called for |browser| and before on_popup_browser_view_created() is
// called for |browser|'s parent delegate if |browser| is a popup.
func (d *BrowserViewDelegate) OnBrowserCreated(browser_view *BrowserView, browser *Browser) {
	lookupBrowserViewDelegateProxy(d.Base().Base()).OnBrowserCreated(d, browser_view, browser)
}

//nolint:gocritic
//export gocef_browser_view_delegate_on_browser_created
func gocef_browser_view_delegate_on_browser_created(self *C.cef_browser_view_delegate_t, browser_view *C.cef_browser_view_t, browser *C.cef_browser_t) {
	me__ := (*BrowserViewDelegate)(self)
	proxy__ := lookupBrowserViewDelegateProxy(me__.Base().Base())
	proxy__.OnBrowserCreated(me__, (*BrowserView)(browser_view), (*Browser)(browser))
}

// OnBrowserDestroyed (on_browser_destroyed)
// Called when |browser| associated with |browser_view| is destroyed. Release
// all references to |browser| and do not attempt to execute any functions on
// |browser| after this callback returns. This function will be called before
// cef_life_span_handler_t::on_before_close() is called for |browser|.
func (d *BrowserViewDelegate) OnBrowserDestroyed(browser_view *BrowserView, browser *Browser) {
	lookupBrowserViewDelegateProxy(d.Base().Base()).OnBrowserDestroyed(d, browser_view, browser)
}

//nolint:gocritic
//export gocef_browser_view_delegate_on_browser_destroyed
func gocef_browser_view_delegate_on_browser_destroyed(self *C.cef_browser_view_delegate_t, browser_view *C.cef_browser_view_t, browser *C.cef_browser_t) {
	me__ := (*BrowserViewDelegate)(self)
	proxy__ := lookupBrowserViewDelegateProxy(me__.Base().Base())
	proxy__.OnBrowserDestroyed(me__, (*BrowserView)(browser_view), (*Browser)(browser))
}

// GetDelegateForPopupBrowserView (get_delegate_for_popup_browser_view)
// Called before a new popup BrowserView is created. The popup originated from
// |browser_view|. |settings| and |client| are the values returned from
// cef_life_span_handler_t::on_before_popup(). |is_devtools| will be true (1)
// if the popup will be a DevTools browser. Return the delegate that will be
// used for the new popup BrowserView.
func (d *BrowserViewDelegate) GetDelegateForPopupBrowserView(browser_view *BrowserView, settings *BrowserSettings, client *Client, is_devtools int32) *BrowserViewDelegate {
	return lookupBrowserViewDelegateProxy(d.Base().Base()).GetDelegateForPopupBrowserView(d, browser_view, settings, client, is_devtools)
}

//nolint:gocritic
//export gocef_browser_view_delegate_get_delegate_for_popup_browser_view
func gocef_browser_view_delegate_get_delegate_for_popup_browser_view(self *C.cef_browser_view_delegate_t, browser_view *C.cef_browser_view_t, settings *C.cef_browser_settings_t, client *C.cef_client_t, is_devtools C.int) *C.cef_browser_view_delegate_t {
	me__ := (*BrowserViewDelegate)(self)
	proxy__ := lookupBrowserViewDelegateProxy(me__.Base().Base())
	settings_ := settings.toGo()
	return (proxy__.GetDelegateForPopupBrowserView(me__, (*BrowserView)(browser_view), settings_, (*Client)(client), int32(is_devtools))).toNative()
}

// OnPopupBrowserViewCreated (on_popup_browser_view_created)
// Called after |popup_browser_view| is created. This function will be called
// after cef_life_span_handler_t::on_after_created() and on_browser_created()
// are called for the new popup browser. The popup originated from
// |browser_view|. |is_devtools| will be true (1) if the popup is a DevTools
// browser. Optionally add |popup_browser_view| to the views hierarchy
// yourself and return true (1). Otherwise return false (0) and a default
// cef_window_t will be created for the popup.
func (d *BrowserViewDelegate) OnPopupBrowserViewCreated(browser_view, popup_browser_view *BrowserView, is_devtools int32) int32 {
	return lookupBrowserViewDelegateProxy(d.Base().Base()).OnPopupBrowserViewCreated(d, browser_view, popup_browser_view, is_devtools)
}

//nolint:gocritic
//export gocef_browser_view_delegate_on_popup_browser_view_created
func gocef_browser_view_delegate_on_popup_browser_view_created(self *C.cef_browser_view_delegate_t, browser_view *C.cef_browser_view_t, popup_browser_view *C.cef_browser_view_t, is_devtools C.int) C.int {
	me__ := (*BrowserViewDelegate)(self)
	proxy__ := lookupBrowserViewDelegateProxy(me__.Base().Base())
	return C.int(proxy__.OnPopupBrowserViewCreated(me__, (*BrowserView)(browser_view), (*BrowserView)(popup_browser_view), int32(is_devtools)))
}
