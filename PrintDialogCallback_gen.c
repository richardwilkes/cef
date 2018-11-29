// Code generated - DO NOT EDIT.

#include "PrintDialogCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_print_dialog_callback_proxy(cef_print_dialog_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->cont = (void *)&gocef_print_dialog_callback_cont;
	self->cancel = (void *)&gocef_print_dialog_callback_cancel;
}
