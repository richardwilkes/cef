package cef

import (
	// #include <stdlib.h>
	// #include "include/capi/cef_app_capi.h"
	"C"
)

type Settings *C.cef_settings_t

// NewSettings creates a new default Settings instance.
func NewSettings() Settings {
	settings := (*C.cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))
	settings.size = C.sizeof_struct__cef_settings_t
	settings.no_sandbox = 1
	settings.command_line_args_disabled = 1
	return Settings(settings)
}
