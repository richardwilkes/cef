// Code generated - DO NOT EDIT.

#include "WebPluginUnstableCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_web_plugin_unstable_callback_proxy(cef_web_plugin_unstable_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->is_unstable = (void *)&gocef_web_plugin_unstable_callback_is_unstable;
}
