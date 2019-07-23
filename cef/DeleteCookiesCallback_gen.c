// Code created from "callback.c.tmpl" - don't edit by hand

#include "DeleteCookiesCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_delete_cookies_callback_proxy(cef_delete_cookies_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_complete = (void *)&gocef_delete_cookies_callback_on_complete;
}
