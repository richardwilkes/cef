#include "WriteHandler_gen.h"
#include "_cgo_export.h"

void gocef_set_write_handler_proxy(cef_write_handler_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->write = (void *)&gocef_write_handler_write;
	self->seek = (void *)&gocef_write_handler_seek;
	self->tell = (void *)&gocef_write_handler_tell;
	self->flush = (void *)&gocef_write_handler_flush;
	self->may_block = (void *)&gocef_write_handler_may_block;
}
