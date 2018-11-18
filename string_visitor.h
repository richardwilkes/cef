#ifndef GOCEF_STRING_VISITOR_H_
#define GOCEF_STRING_VISITOR_H_
#pragma once

#include "include/capi/cef_string_visitor_capi.h"

typedef struct _gocef_async_string_visitor_t {
	cef_string_visitor_t visitor;
	int                  (CEF_CALLBACK *release)(cef_base_ref_counted_t *self);
	int32_t              id;
} gocef_async_string_visitor_t;

cef_string_visitor_t *gocef_new_async_string_visitor(int32_t id);

#endif // GOCEF_STRING_VISITOR_H_
