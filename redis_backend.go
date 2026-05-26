// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"github.com/gomodule/redigo/redis"
)

// RedisCeleryBackend is celery backend for redis
type RedisCeleryBackend struct {
	*redis.Pool
}

// NewRedisBackend creates new RedisCeleryBackend with given redis pool.
// RedisCeleryBackend can be initialized manually as well.
func NewRedisBackend(conn *redis.Pool) *RedisCeleryBackend { _ = "STUB: not implemented"; return nil }

// NewRedisCeleryBackend creates new RedisCeleryBackend
//
// Deprecated: NewRedisCeleryBackend exists for historical compatibility
// and should not be used. Pool should be initialized outside of gocelery package.
func NewRedisCeleryBackend(uri string) *RedisCeleryBackend { _ = "STUB: not implemented"; return nil }

// GetResult queries redis backend to get asynchronous result
func (cb *RedisCeleryBackend) GetResult(taskID string) (*ResultMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetResult pushes result back into redis backend
func (cb *RedisCeleryBackend) SetResult(taskID string, result *ResultMessage) error {
	_ = "STUB: not implemented"
	return nil
}
