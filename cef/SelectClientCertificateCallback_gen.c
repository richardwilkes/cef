// Code generated - DO NOT EDIT.

#include "SelectClientCertificateCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_select_client_certificate_callback_proxy(cef_select_client_certificate_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->select = (void *)&gocef_select_client_certificate_callback__select;
}
