#include "JsdialogHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_jsdialog_handler_proxy(cef_jsdialog_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_jsdialog = (void *)&gocef_jsdialog_handler_on_jsdialog;
	self->on_before_unload_dialog = (void *)&gocef_jsdialog_handler_on_before_unload_dialog;
	self->on_reset_dialog_state = (void *)&gocef_jsdialog_handler_on_reset_dialog_state;
	self->on_dialog_closed = (void *)&gocef_jsdialog_handler_on_dialog_closed;
}
