#include "PDFPrintCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_pdf_print_callback_proxy(cef_pdf_print_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_pdf_print_finished = (void *)&gocef_pdf_print_callback_on_pdf_print_finished;
}
