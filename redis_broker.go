// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"github.com/gomodule/redigo/redis"
)

// RedisCeleryBroker is celery broker for redis
type RedisCeleryBroker struct {
	*redis.Pool
	QueueName string
}

// NewRedisBroker creates new RedisCeleryBroker with given redis connection pool
func NewRedisBroker(conn *redis.Pool) *RedisCeleryBroker { _ = "STUB: not implemented"; return nil }

// NewRedisCeleryBroker creates new RedisCeleryBroker based on given uri
//
// Deprecated: NewRedisCeleryBroker exists for historical compatibility
// and should not be used. Use NewRedisBroker instead to create new RedisCeleryBroker.
func NewRedisCeleryBroker(uri string) *RedisCeleryBroker { _ = "STUB: not implemented"; return nil }

// SendCeleryMessage sends CeleryMessage to redis queue
func (cb *RedisCeleryBroker) SendCeleryMessage(message *CeleryMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCeleryMessage retrieves celery message from redis queue
func (cb *RedisCeleryBroker) GetCeleryMessage() (*CeleryMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTaskMessage retrieves task message from redis queue
func (cb *RedisCeleryBroker) GetTaskMessage() (*TaskMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRedisPool creates pool of redis connections from given connection string
//
// Deprecated: newRedisPool exists for historical compatibility
// and should not be used. Pool should be initialized outside of gocelery package.
func NewRedisPool(uri string) *redis.Pool { _ = "STUB: not implemented"; return nil }
