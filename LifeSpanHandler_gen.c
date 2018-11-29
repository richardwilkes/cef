#include "LifeSpanHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_life_span_handler_proxy(cef_life_span_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_before_popup = (void *)&gocef_life_span_handler_on_before_popup;
	self->on_after_created = (void *)&gocef_life_span_handler_on_after_created;
	self->do_close = (void *)&gocef_life_span_handler_do_close;
	self->on_before_close = (void *)&gocef_life_span_handler_on_before_close;
}
