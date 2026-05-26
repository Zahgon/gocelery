// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"context"
	"reflect"
	"sync"
	"time"
)

// CeleryWorker represents distributed task worker
type CeleryWorker struct {
	broker          CeleryBroker
	backend         CeleryBackend
	numWorkers      int
	registeredTasks map[string]interface{}
	taskLock        sync.RWMutex
	cancel          context.CancelFunc
	workWG          sync.WaitGroup
	rateLimitPeriod time.Duration
}

// NewCeleryWorker returns new celery worker
func NewCeleryWorker(broker CeleryBroker, backend CeleryBackend, numWorkers int) *CeleryWorker {
	_ = "STUB: not implemented"
	return nil
}

// StartWorkerWithContext starts celery worker(s) with given parent context
func (w *CeleryWorker) StartWorkerWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// process task request

// run task

// push result to backend

// StartWorker starts celery workers
func (w *CeleryWorker) StartWorker() { _ = "STUB: not implemented"; return }

// StopWorker stops celery workers
func (w *CeleryWorker) StopWorker() { _ = "STUB: not implemented"; return }

// StopWait waits for celery workers to terminate
func (w *CeleryWorker) StopWait() {
	_ = "STUB: not implemented"

	// GetNumWorkers returns number of currently running workers
	return
}

func (w *CeleryWorker) GetNumWorkers() int { _ = "STUB: not implemented"; return 0 }

// Register registers tasks (functions)
func (w *CeleryWorker) Register(name string, task interface{}) { _ = "STUB: not implemented"; return }

// GetTask retrieves registered task
func (w *CeleryWorker) GetTask(name string) interface{} { _ = "STUB: not implemented"; return nil }

// RunTask runs celery task
func (w *CeleryWorker) RunTask(message *TaskMessage) (*ResultMessage, error) {
	_ = "STUB: not implemented"

	// ignore if the message is expired
	return nil, nil
}

// check for malformed task message - args cannot be nil

// get task

// convert to task interface

// use reflection to execute function ptr

func runTaskFunc(taskFunc *reflect.Value, message *TaskMessage) (*ResultMessage, error) {
	_ = "STUB: not implemented"

	// check number of arguments
	return nil, nil
}

// construct arguments

// special case - convert float64 to int if applicable
// this is due to json limitation where all numbers are converted to float64

// call method
