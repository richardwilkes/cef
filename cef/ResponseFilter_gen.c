// Code generated - DO NOT EDIT.

#include "ResponseFilter_gen.h"
#include "_cgo_export.h"

void gocef_set_response_filter_proxy(cef_response_filter_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->init_filter = (void *)&gocef_response_filter_init_filter;
	self->filter = (void *)&gocef_response_filter_filter;
}
