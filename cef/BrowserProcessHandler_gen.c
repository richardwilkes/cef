// Code created from "callback.c.tmpl" - don't edit by hand

#include "BrowserProcessHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_browser_process_handler_proxy(cef_browser_process_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_context_initialized = (void *)&gocef_browser_process_handler_on_context_initialized;
	self->on_before_child_process_launch = (void *)&gocef_browser_process_handler_on_before_child_process_launch;
	self->on_render_process_thread_created = (void *)&gocef_browser_process_handler_on_render_process_thread_created;
	self->get_print_handler = (void *)&gocef_browser_process_handler_get_print_handler;
	self->on_schedule_message_pump_work = (void *)&gocef_browser_process_handler_on_schedule_message_pump_work;
}
