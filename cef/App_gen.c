// Code generated - DO NOT EDIT.

#include "App_gen.h"
#include "_cgo_export.h"

void gocef_set_app_proxy(cef_app_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_command_line_processing = (void *)&gocef_app_on_before_command_line_processing;
	self->on_register_custom_schemes = (void *)&gocef_app_on_register_custom_schemes;
	self->get_resource_bundle_handler = (void *)&gocef_app_get_resource_bundle_handler;
	self->get_browser_process_handler = (void *)&gocef_app_get_browser_process_handler;
	self->get_render_process_handler = (void *)&gocef_app_get_render_process_handler;
}
