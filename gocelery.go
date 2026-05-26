// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"context"
	"time"
)

// CeleryClient provides API for sending celery tasks
type CeleryClient struct {
	broker  CeleryBroker
	backend CeleryBackend
	worker  *CeleryWorker
}

// CeleryBroker is interface for celery broker database
type CeleryBroker interface {
	SendCeleryMessage(*CeleryMessage) error
	GetTaskMessage() (*TaskMessage, error) // must be non-blocking
}

// CeleryBackend is interface for celery backend database
type CeleryBackend interface {
	GetResult(string) (*ResultMessage, error) // must be non-blocking
	SetResult(taskID string, result *ResultMessage) error
}

// NewCeleryClient creates new celery client
func NewCeleryClient(broker CeleryBroker, backend CeleryBackend, numWorkers int) (*CeleryClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register task
func (cc *CeleryClient) Register(name string, task interface{}) { _ = "STUB: not implemented"; return }

// StartWorkerWithContext starts celery workers with given parent context
func (cc *CeleryClient) StartWorkerWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// StartWorker starts celery workers
func (cc *CeleryClient) StartWorker() { _ = "STUB: not implemented"; return }

// StopWorker stops celery workers
func (cc *CeleryClient) StopWorker() { _ = "STUB: not implemented"; return }

// WaitForStopWorker waits for celery workers to terminate
func (cc *CeleryClient) WaitForStopWorker() { _ = "STUB: not implemented"; return }

// Delay gets asynchronous result
func (cc *CeleryClient) Delay(task string, args ...interface{}) (*AsyncResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelayKwargs gets asynchronous results with argument map
func (cc *CeleryClient) DelayKwargs(task string, args map[string]interface{}) (*AsyncResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *CeleryClient) delay(task *TaskMessage) (*AsyncResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CeleryTask is an interface that represents actual task
// Passing CeleryTask interface instead of function pointer
// avoids reflection and may have performance gain.
// ResultMessage must be obtained using GetResultMessage()
type CeleryTask interface {

	// ParseKwargs - define a method to parse kwargs
	ParseKwargs(map[string]interface{}) error

	// RunTask - define a method for execution
	RunTask() (interface{}, error)
}

// AsyncResult represents pending result
type AsyncResult struct {
	TaskID  string
	backend CeleryBackend
	result  *ResultMessage
}

// Get gets actual result from backend
// It blocks for period of time set by timeout and returns error if unavailable
func (ar *AsyncResult) Get(timeout time.Duration) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AsyncGet gets actual result from backend and returns nil if not available
func (ar *AsyncResult) AsyncGet() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// Ready checks if actual result is ready
func (ar *AsyncResult) Ready() (bool, error) { _ = "STUB: not implemented"; return false, nil }
