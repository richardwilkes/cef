package cef

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// State represents the state of a setting.
type State int

// Possible states.
const (
	StateDefault State = iota
	StateEnabled
	StateDisabled
)

// Color holds a 32-bit ARGB color value, not premultiplied.
type Color uint32

// BrowserSettings holds browser initialization settings.
//
// Defined in include/internal/cef_types.h
type BrowserSettings struct {
	native *C.cef_browser_settings_t
}

// NewBrowserSettings creates a new default BrowserSettings instance.
func NewBrowserSettings() *BrowserSettings {
	settings := &BrowserSettings{native: (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_struct__cef_browser_settings_t))}
	settings.native.size = C.sizeof_struct__cef_browser_settings_t
	return settings
}

// SetWindowlessFrameRate sets the maximum rate in frames per second (fps)
// rendering will be called for a windowless browser. The actual fps may be
// lower if the browser cannot generate frames at the requested rate. The
// minimum value is 1 and the maximum value is 60 (default 30).
func (b *BrowserSettings) SetWindowlessFrameRate(rate int) {
	if rate < 1 {
		rate = 1
	} else if rate > 60 {
		rate = 60
	}
	b.native.windowless_frame_rate = C.int(rate)
}

// SetStandardFontFamily sets the standard font family.
func (b *BrowserSettings) SetStandardFontFamily(family string) {
	setCEFStr(family, &b.native.standard_font_family)
}

// SetFixedFontFamily sets the fixed font family.
func (b *BrowserSettings) SetFixedFontFamily(family string) {
	setCEFStr(family, &b.native.fixed_font_family)
}

// SetSerifFontFamily sets the serif font family.
func (b *BrowserSettings) SetSerifFontFamily(family string) {
	setCEFStr(family, &b.native.serif_font_family)
}

// SetSansSerifFontFamily sets the sans serif font family.
func (b *BrowserSettings) SetSansSerifFontFamily(family string) {
	setCEFStr(family, &b.native.sans_serif_font_family)
}

// SetCursiveFontFamily sets the cursive font family.
func (b *BrowserSettings) SetCursiveFontFamily(family string) {
	setCEFStr(family, &b.native.cursive_font_family)
}

// SetFantasyFontFamily sets the fantasy font family.
func (b *BrowserSettings) SetFantasyFontFamily(family string) {
	setCEFStr(family, &b.native.fantasy_font_family)
}

// SetDefaultFontSize sets the default font size.
func (b *BrowserSettings) SetDefaultFontSize(size int) {
	if size < 1 {
		size = 1
	}
	b.native.default_font_size = C.int(size)
}

// SetDefaultFixedFontSize sets the default fixed font size.
func (b *BrowserSettings) SetDefaultFixedFontSize(size int) {
	if size < 1 {
		size = 1
	}
	b.native.default_fixed_font_size = C.int(size)
}

// SetMinimumFontSize sets the minimum font size.
func (b *BrowserSettings) SetMinimumFontSize(size int) {
	if size < 1 {
		size = 1
	}
	b.native.minimum_font_size = C.int(size)
}

// SetMinimumLogicalFontSize sets the minimum logical font size.
func (b *BrowserSettings) SetMinimumLogicalFontSize(size int) {
	if size < 1 {
		size = 1
	}
	b.native.minimum_logical_font_size = C.int(size)
}

// SetDefaultEncoding controls the default encoding for Web content. Defaults
// to "ISO-8859-1".
func (b *BrowserSettings) SetDefaultEncoding(encoding string) {
	setCEFStr(encoding, &b.native.default_encoding)
}

// SetRemoteFonts controls the loading of fonts from remote sources. Defaults
// to yes.
func (b *BrowserSettings) SetRemoteFonts(state State) {
	b.native.remote_fonts = C.cef_state_t(state)
}

// SetJavaScript controls whether JavaScript can be executed. Defaults to yes.
func (b *BrowserSettings) SetJavaScript(state State) {
	b.native.javascript = C.cef_state_t(state)
}

// SetJavaScriptCanCloseWindows controls whether JavaScript can be used to
// close windows that were not opened via JavaScript. JavaScript can still be
// used to close windows that were opened via JavaScript or that have no
// back/forward history. Defaults to yes.
func (b *BrowserSettings) SetJavaScriptCanCloseWindows(state State) {
	b.native.javascript_close_windows = C.cef_state_t(state)
}

// SetJavaScriptCanAccessClipboard controls whether JavaScript can access the
// clipboard. Defaults to yes.
func (b *BrowserSettings) SetJavaScriptCanAccessClipboard(state State) {
	b.native.javascript_access_clipboard = C.cef_state_t(state)
}

// SetJavaScriptCanDOMPaste controls whether DOM pasting is supported in the
// editor via execCommand("paste"). SetJavaScriptCanAccessClipboard must also
// be enabled. Defaults to yes.
func (b *BrowserSettings) SetJavaScriptCanDOMPaste(state State) {
	b.native.javascript_dom_paste = C.cef_state_t(state)
}

// SetPlugins controls whether any plugins will be loaded. Defaults to yes.
func (b *BrowserSettings) SetPlugins(state State) {
	b.native.plugins = C.cef_state_t(state)
}

// SetUniversalAccessFromFileURLs controls whether file URLs will have access
// to all URLs. Defaults to no.
func (b *BrowserSettings) SetUniversalAccessFromFileURLs(state State) {
	b.native.universal_access_from_file_urls = C.cef_state_t(state)
}

// SetFileAccessFromFileURLs controls whether file URLs will have access to
// other file URLs. Defaults to no.
func (b *BrowserSettings) SetFileAccessFromFileURLs(state State) {
	b.native.file_access_from_file_urls = C.cef_state_t(state)
}

// SetWebSecurity controls whether web security restrictions (same-origin
// policy) will be enforced. Disabling this setting is not recommend as it
// will allow risky security behavior such as cross-site scripting (XSS).
// Defaults to yes.
func (b *BrowserSettings) SetWebSecurity(state State) {
	b.native.web_security = C.cef_state_t(state)
}

// SetImageLoading controls whether image URLs will be loaded from the
// network. A cached image will still be rendered if requested. Defaults to
// yes.
func (b *BrowserSettings) SetImageLoading(state State) {
	b.native.image_loading = C.cef_state_t(state)
}

// SetImageShrinkStandaloneToFit controls whether standalone images will be
// shrunk to fit the page. Defaults to yes.
func (b *BrowserSettings) SetImageShrinkStandaloneToFit(state State) {
	b.native.image_shrink_standalone_to_fit = C.cef_state_t(state)
}

// SetTextAreaResize controls whether text areas can be resized. Defaults to
// yes.
func (b *BrowserSettings) SetTextAreaResize(state State) {
	b.native.text_area_resize = C.cef_state_t(state)
}

// SetTabToLinks controls whether the tab key can advance focus to links.
// Defaults to yes.
func (b *BrowserSettings) SetTabToLinks(state State) {
	b.native.tab_to_links = C.cef_state_t(state)
}

// SetLocalStorage controls whether local storage can be used. Default is yes.
func (b *BrowserSettings) SetLocalStorage(state State) {
	b.native.local_storage = C.cef_state_t(state)
}

// SetDatabases controls whether databases can be used. Default is yes.
func (b *BrowserSettings) SetDatabases(state State) {
	b.native.databases = C.cef_state_t(state)
}

// SetApplicationCache controls whether the application cache can be used.
// Default is yes.
func (b *BrowserSettings) SetApplicationCache(state State) {
	b.native.application_cache = C.cef_state_t(state)
}

// SetWebGL controls whether WebGL can be used. Note that WebGL requires
// hardware support and may not work on all systems even when enabled. Default
// is yes.
func (b *BrowserSettings) SetWebGL(state State) {
	b.native.webgl = C.cef_state_t(state)
}

// SetBackgroundColor sets the background color used for the browser before a
// document is loaded and when no document color is specified. The alpha
// component must be either fully opaque (0xFF) or fully transparent (0x00).
// If the alpha component is fully opaque then the RGB components will be used
// as the background color. If the alpha component is fully transparent for a
// windowed browser then the background color value will be used. If the alpha
// component is fully transparent for a windowless (off-screen) browser then
// transparent painting will be enabled. Defaults to the value in Settings.
func (b *BrowserSettings) SetBackgroundColor(color Color) {
	b.native.background_color = C.cef_color_t(color)
}

// SetAcceptLanguageList sets a comma delimited ordered list of language codes
// without any whitespace that will be used in the "Accept-Language" HTTP
// header. Defaults to the value in Settings.
func (b *BrowserSettings) SetAcceptLanguageList(list string) {
	setCEFStr(list, &b.native.accept_language_list)
}
