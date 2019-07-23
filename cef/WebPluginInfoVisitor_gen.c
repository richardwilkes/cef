// Code created from "callback.c.tmpl" - don't edit by hand

#include "WebPluginInfoVisitor_gen.h"
#include "_cgo_export.h"

void gocef_set_web_plugin_info_visitor_proxy(cef_web_plugin_info_visitor_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->visit = (void *)&gocef_web_plugin_info_visitor_visit;
}
