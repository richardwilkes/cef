// Code generated - DO NOT EDIT.

#include "GetExtensionResourceCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_get_extension_resource_callback_proxy(cef_get_extension_resource_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cont = (void *)&gocef_get_extension_resource_callback_cont;
	self->cancel = (void *)&gocef_get_extension_resource_callback_cancel;
}
