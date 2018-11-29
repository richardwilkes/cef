#include "LoadHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_load_handler_proxy(cef_load_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_loading_state_change = (void *)&gocef_load_handler_on_loading_state_change;
	self->on_load_start = (void *)&gocef_load_handler_on_load_start;
	self->on_load_end = (void *)&gocef_load_handler_on_load_end;
	self->on_load_error = (void *)&gocef_load_handler_on_load_error;
}
