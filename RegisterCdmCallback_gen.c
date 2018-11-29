// Code generated - DO NOT EDIT.

#include "RegisterCdmCallback_gen.h"
#include "_cgo_export.h"

void gocef_set_register_cdm_callback_proxy(cef_register_cdm_callback_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_cdm_registration_complete = (void *)&gocef_register_cdm_callback_on_cdm_registration_complete;
}
