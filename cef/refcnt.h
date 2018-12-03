#ifndef GOCEF_REFCNT_H_
#define GOCEF_REFCNT_H_
#pragma once

#include "include/capi/cef_base_capi.h"

cef_base_ref_counted_t *gocef_refcnt_alloc(size_t size);
uint32_t gocef_refcnt_id(cef_base_ref_counted_t *p);

#endif // GOCEF_REFCNT_H_
