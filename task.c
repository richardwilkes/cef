#include "ref.h"
#include "task.h"
#include "_cgo_export.h"

void gocef_execute_task(cef_task_t *self) {
	taskCallback(((gocef_task_t *)self)->id);
}

cef_task_t *gocef_new_task(int id) {
	gocef_task_t *task = (gocef_task_t *)gocef_refcnt_alloc(sizeof(gocef_task_t));
	task->id = id;
	task->task.execute = gocef_execute_task;
	return &task->task;
}