// Code generated - DO NOT EDIT.

#include "ReadHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_read_handler_proxy(cef_read_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->read = (void *)&gocef_read_handler_read;
	self->seek = (void *)&gocef_read_handler_seek;
	self->tell = (void *)&gocef_read_handler_tell;
	self->eof = (void *)&gocef_read_handler_eof;
	self->may_block = (void *)&gocef_read_handler_may_block;
}
