// Code created from "callback.c.tmpl" - don't edit by hand

#include "CookieAccessFilter_gen.h"
#include "_cgo_export.h"

void gocef_set_cookie_access_filter_proxy(cef_cookie_access_filter_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->can_send_cookie = (void *)&gocef_cookie_access_filter_can_send_cookie;
	self->can_save_cookie = (void *)&gocef_cookie_access_filter_can_save_cookie;
}
