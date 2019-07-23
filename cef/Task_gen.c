// Code created from "callback.c.tmpl" - don't edit by hand

#include "Task_gen.h"
#include "_cgo_export.h"

void gocef_set_task_proxy(cef_task_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->execute = (void *)&gocef_task_execute;
}
