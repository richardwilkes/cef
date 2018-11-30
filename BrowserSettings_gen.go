// Code generated - DO NOT EDIT.

package cef

import (
	// #include "capi_gen.h"
	"C"
)

// BrowserSettings (cef_browser_settings_t from include/internal/cef_types.h)
// Browser initialization settings. Specify NULL or 0 to get the recommended
// default values. The consequences of using custom values may not be well
// tested. Many of these and other settings can also configured using command-
// line switches.
type BrowserSettings struct {
	// Size (size)
	// Size of this structure.
	Size uint64
	// WindowlessFrameRate (windowless_frame_rate)
	// The maximum rate in frames per second (fps) that CefRenderHandler::OnPaint
	// will be called for a windowless browser. The actual fps may be lower if
	// the browser cannot generate frames at the requested rate. The minimum
	// value is 1 and the maximum value is 60 (default 30). This value can also be
	// changed dynamically via CefBrowserHost::SetWindowlessFrameRate.
	WindowlessFrameRate int32
	// StandardFontFamily (standard_font_family)
	// Font settings.
	StandardFontFamily string
	// FixedFontFamily (fixed_font_family)
	FixedFontFamily string
	// SerifFontFamily (serif_font_family)
	SerifFontFamily string
	// SansSerifFontFamily (sans_serif_font_family)
	SansSerifFontFamily string
	// CursiveFontFamily (cursive_font_family)
	CursiveFontFamily string
	// FantasyFontFamily (fantasy_font_family)
	FantasyFontFamily string
	// DefaultFontSize (default_font_size)
	DefaultFontSize int32
	// DefaultFixedFontSize (default_fixed_font_size)
	DefaultFixedFontSize int32
	// MinimumFontSize (minimum_font_size)
	MinimumFontSize int32
	// MinimumLogicalFontSize (minimum_logical_font_size)
	MinimumLogicalFontSize int32
	// DefaultEncoding (default_encoding)
	// Default encoding for Web content. If empty "ISO-8859-1" will be used. Also
	// configurable using the "default-encoding" command-line switch.
	DefaultEncoding string
	// RemoteFonts (remote_fonts)
	// Controls the loading of fonts from remote sources. Also configurable using
	// the "disable-remote-fonts" command-line switch.
	RemoteFonts State
	// Javascript (javascript)
	// Controls whether JavaScript can be executed. Also configurable using the
	// "disable-javascript" command-line switch.
	Javascript State
	// JavascriptCloseWindows (javascript_close_windows)
	// Controls whether JavaScript can be used to close windows that were not
	// opened via JavaScript. JavaScript can still be used to close windows that
	// were opened via JavaScript or that have no back/forward history. Also
	// configurable using the "disable-javascript-close-windows" command-line
	// switch.
	JavascriptCloseWindows State
	// JavascriptAccessClipboard (javascript_access_clipboard)
	// Controls whether JavaScript can access the clipboard. Also configurable
	// using the "disable-javascript-access-clipboard" command-line switch.
	JavascriptAccessClipboard State
	// JavascriptDomPaste (javascript_dom_paste)
	// Controls whether DOM pasting is supported in the editor via
	// execCommand("paste"). The |javascript_access_clipboard| setting must also
	// be enabled. Also configurable using the "disable-javascript-dom-paste"
	// command-line switch.
	JavascriptDomPaste State
	// Plugins (plugins)
	// Controls whether any plugins will be loaded. Also configurable using the
	// "disable-plugins" command-line switch.
	Plugins State
	// UniversalAccessFromFileUrls (universal_access_from_file_urls)
	// Controls whether file URLs will have access to all URLs. Also configurable
	// using the "allow-universal-access-from-files" command-line switch.
	UniversalAccessFromFileUrls State
	// FileAccessFromFileUrls (file_access_from_file_urls)
	// Controls whether file URLs will have access to other file URLs. Also
	// configurable using the "allow-access-from-files" command-line switch.
	FileAccessFromFileUrls State
	// WebSecurity (web_security)
	// Controls whether web security restrictions (same-origin policy) will be
	// enforced. Disabling this setting is not recommend as it will allow risky
	// security behavior such as cross-site scripting (XSS). Also configurable
	// using the "disable-web-security" command-line switch.
	WebSecurity State
	// ImageLoading (image_loading)
	// Controls whether image URLs will be loaded from the network. A cached image
	// will still be rendered if requested. Also configurable using the
	// "disable-image-loading" command-line switch.
	ImageLoading State
	// ImageShrinkStandaloneToFit (image_shrink_standalone_to_fit)
	// Controls whether standalone images will be shrunk to fit the page. Also
	// configurable using the "image-shrink-standalone-to-fit" command-line
	// switch.
	ImageShrinkStandaloneToFit State
	// TextAreaResize (text_area_resize)
	// Controls whether text areas can be resized. Also configurable using the
	// "disable-text-area-resize" command-line switch.
	TextAreaResize State
	// TabToLinks (tab_to_links)
	// Controls whether the tab key can advance focus to links. Also configurable
	// using the "disable-tab-to-links" command-line switch.
	TabToLinks State
	// LocalStorage (local_storage)
	// Controls whether local storage can be used. Also configurable using the
	// "disable-local-storage" command-line switch.
	LocalStorage State
	// Databases (databases)
	// Controls whether databases can be used. Also configurable using the
	// "disable-databases" command-line switch.
	Databases State
	// ApplicationCache (application_cache)
	// Controls whether the application cache can be used. Also configurable using
	// the "disable-application-cache" command-line switch.
	ApplicationCache State
	// Webgl (webgl)
	// Controls whether WebGL can be used. Note that WebGL requires hardware
	// support and may not work on all systems even when enabled. Also
	// configurable using the "disable-webgl" command-line switch.
	Webgl State
	// BackgroundColor (background_color)
	// Background color used for the browser before a document is loaded and when
	// no document color is specified. The alpha component must be either fully
	// opaque (0xFF) or fully transparent (0x00). If the alpha component is fully
	// opaque then the RGB components will be used as the background color. If the
	// alpha component is fully transparent for a windowed browser then the
	// CefSettings.background_color value will be used. If the alpha component is
	// fully transparent for a windowless (off-screen) browser then transparent
	// painting will be enabled.
	BackgroundColor Color
	// AcceptLanguageList (accept_language_list)
	// Comma delimited ordered list of language codes without any whitespace that
	// will be used in the "Accept-Language" HTTP header. May be set globally
	// using the CefBrowserSettings.accept_language_list value. If both values are
	// empty then "en-US,en" will be used.
	AcceptLanguageList string
}

// NewBrowserSettings creates a new BrowserSettings.
func NewBrowserSettings() *BrowserSettings {
	return &BrowserSettings{
		Size: C.sizeof_struct__cef_browser_settings_t,
	}
}

func (d *BrowserSettings) toNative(native *C.cef_browser_settings_t) *C.cef_browser_settings_t {
	native.size = C.size_t(d.Size)
	native.windowless_frame_rate = C.int(d.WindowlessFrameRate)
	setCEFStr(d.StandardFontFamily, &native.standard_font_family)
	setCEFStr(d.FixedFontFamily, &native.fixed_font_family)
	setCEFStr(d.SerifFontFamily, &native.serif_font_family)
	setCEFStr(d.SansSerifFontFamily, &native.sans_serif_font_family)
	setCEFStr(d.CursiveFontFamily, &native.cursive_font_family)
	setCEFStr(d.FantasyFontFamily, &native.fantasy_font_family)
	native.default_font_size = C.int(d.DefaultFontSize)
	native.default_fixed_font_size = C.int(d.DefaultFixedFontSize)
	native.minimum_font_size = C.int(d.MinimumFontSize)
	native.minimum_logical_font_size = C.int(d.MinimumLogicalFontSize)
	setCEFStr(d.DefaultEncoding, &native.default_encoding)
	native.remote_fonts = C.cef_state_t(d.RemoteFonts)
	native.javascript = C.cef_state_t(d.Javascript)
	native.javascript_close_windows = C.cef_state_t(d.JavascriptCloseWindows)
	native.javascript_access_clipboard = C.cef_state_t(d.JavascriptAccessClipboard)
	native.javascript_dom_paste = C.cef_state_t(d.JavascriptDomPaste)
	native.plugins = C.cef_state_t(d.Plugins)
	native.universal_access_from_file_urls = C.cef_state_t(d.UniversalAccessFromFileUrls)
	native.file_access_from_file_urls = C.cef_state_t(d.FileAccessFromFileUrls)
	native.web_security = C.cef_state_t(d.WebSecurity)
	native.image_loading = C.cef_state_t(d.ImageLoading)
	native.image_shrink_standalone_to_fit = C.cef_state_t(d.ImageShrinkStandaloneToFit)
	native.text_area_resize = C.cef_state_t(d.TextAreaResize)
	native.tab_to_links = C.cef_state_t(d.TabToLinks)
	native.local_storage = C.cef_state_t(d.LocalStorage)
	native.databases = C.cef_state_t(d.Databases)
	native.application_cache = C.cef_state_t(d.ApplicationCache)
	native.webgl = C.cef_state_t(d.Webgl)
	native.background_color = C.cef_color_t(d.BackgroundColor)
	setCEFStr(d.AcceptLanguageList, &native.accept_language_list)
	return native
}

func (d *BrowserSettings) fromNative(native *C.cef_browser_settings_t) *BrowserSettings {
	d.Size = uint64(native.size)
	d.WindowlessFrameRate = int32(native.windowless_frame_rate)
	d.StandardFontFamily = cefstrToString(&native.standard_font_family)
	d.FixedFontFamily = cefstrToString(&native.fixed_font_family)
	d.SerifFontFamily = cefstrToString(&native.serif_font_family)
	d.SansSerifFontFamily = cefstrToString(&native.sans_serif_font_family)
	d.CursiveFontFamily = cefstrToString(&native.cursive_font_family)
	d.FantasyFontFamily = cefstrToString(&native.fantasy_font_family)
	d.DefaultFontSize = int32(native.default_font_size)
	d.DefaultFixedFontSize = int32(native.default_fixed_font_size)
	d.MinimumFontSize = int32(native.minimum_font_size)
	d.MinimumLogicalFontSize = int32(native.minimum_logical_font_size)
	d.DefaultEncoding = cefstrToString(&native.default_encoding)
	d.RemoteFonts = State(native.remote_fonts)
	d.Javascript = State(native.javascript)
	d.JavascriptCloseWindows = State(native.javascript_close_windows)
	d.JavascriptAccessClipboard = State(native.javascript_access_clipboard)
	d.JavascriptDomPaste = State(native.javascript_dom_paste)
	d.Plugins = State(native.plugins)
	d.UniversalAccessFromFileUrls = State(native.universal_access_from_file_urls)
	d.FileAccessFromFileUrls = State(native.file_access_from_file_urls)
	d.WebSecurity = State(native.web_security)
	d.ImageLoading = State(native.image_loading)
	d.ImageShrinkStandaloneToFit = State(native.image_shrink_standalone_to_fit)
	d.TextAreaResize = State(native.text_area_resize)
	d.TabToLinks = State(native.tab_to_links)
	d.LocalStorage = State(native.local_storage)
	d.Databases = State(native.databases)
	d.ApplicationCache = State(native.application_cache)
	d.Webgl = State(native.webgl)
	d.BackgroundColor = Color(native.background_color)
	d.AcceptLanguageList = cefstrToString(&native.accept_language_list)
	return d
}
