package cef

import (
	// #include "task.h"
	"C"
	"sync"
)

var (
	nextTaskID     C.int = 1
	nextTaskIDLock sync.Mutex
	taskMap        = make(map[C.int]func())
)

// PostUITask schedules a task for execution on the UI thread.
func PostUITask(task func()) {
	nextTaskIDLock.Lock()
	id := nextTaskID
	nextTaskID++
	taskMap[id] = task
	nextTaskIDLock.Unlock()
	C.cef_post_task(C.TID_UI, C.gocef_new_task(id))
}

//export taskCallback
func taskCallback(id C.int) {
	nextTaskIDLock.Lock()
	task := taskMap[id]
	delete(taskMap, id)
	nextTaskIDLock.Unlock()
	if task != nil {
		task()
	}
}
