// Code generated - DO NOT EDIT.

#include "RunFileDialogCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_run_file_dialog_callback_proxy(cef_run_file_dialog_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_file_dialog_dismissed = (void *)&gocef_run_file_dialog_callback_on_file_dialog_dismissed;
}
