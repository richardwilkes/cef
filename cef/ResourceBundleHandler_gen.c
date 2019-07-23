// Code created from "callback.c.tmpl" - don't edit by hand

#include "ResourceBundleHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_resource_bundle_handler_proxy(cef_resource_bundle_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->get_localized_string = (void *)&gocef_resource_bundle_handler_get_localized_string;
	self->get_data_resource = (void *)&gocef_resource_bundle_handler_get_data_resource;
	self->get_data_resource_for_scale = (void *)&gocef_resource_bundle_handler_get_data_resource_for_scale;
}
