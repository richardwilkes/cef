#include "CookieVisitor_gen.h"
#include "_cgo_export.h"

void gocef_set_cookie_visitor_proxy(cef_cookie_visitor_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->visit = (void *)&gocef_cookie_visitor_visit;
}
