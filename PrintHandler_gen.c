#include "PrintHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_print_handler_proxy(cef_print_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_print_start = (void *)&gocef_print_handler_on_print_start;
	self->on_print_settings = (void *)&gocef_print_handler_on_print_settings;
	self->on_print_dialog = (void *)&gocef_print_handler_on_print_dialog;
	self->on_print_job = (void *)&gocef_print_handler_on_print_job;
	self->on_print_reset = (void *)&gocef_print_handler_on_print_reset;
	self->get_pdf_paper_size = (void *)&gocef_print_handler_get_pdf_paper_size;
}
