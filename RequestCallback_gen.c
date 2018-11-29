#include "RequestCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_request_callback_proxy(cef_request_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cont = (void *)&gocef_request_callback_cont;
	self->cancel = (void *)&gocef_request_callback_cancel;
}
