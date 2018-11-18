#ifndef GOCEF_TASK_H_
#define GOCEF_TASK_H_
#pragma once

#include "include/capi/cef_task_capi.h"

typedef struct _gocef_task_t {
	cef_task_t task;
	int        id;
} gocef_task_t;

cef_task_t *gocef_new_task(int id);

#endif // GOCEF_TASK_H_
