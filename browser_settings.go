package cef

import (
	// #include <stdlib.h>
	// #include "include/capi/cef_app_capi.h"
	"C"
)

type BrowserSettings *C.cef_browser_settings_t

// NewBrowserSettings creates a new default BrowserSettings instance.
func NewBrowserSettings() BrowserSettings {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	settings.size = C.sizeof_struct__cef_browser_settings_t
	return BrowserSettings(settings)
}
