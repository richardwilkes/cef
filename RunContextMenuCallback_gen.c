#include "RunContextMenuCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_run_context_menu_callback_proxy(cef_run_context_menu_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cont = (void *)&gocef_run_context_menu_callback_cont;
	self->cancel = (void *)&gocef_run_context_menu_callback_cancel;
}
