// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"github.com/streadway/amqp"
)

// AMQPCeleryBackend CeleryBackend for AMQP
type AMQPCeleryBackend struct {
	*amqp.Channel
	Connection *amqp.Connection
	Host       string
}

// NewAMQPCeleryBackend creates new AMQPCeleryBackend
func NewAMQPCeleryBackend(host string) *AMQPCeleryBackend { _ = "STUB: not implemented"; return nil }

// NewAMQPCeleryBackendByConnAndChannel creates new AMQPCeleryBackend by AMQP connection and channel
func NewAMQPCeleryBackendByConnAndChannel(conn *amqp.Connection, channel *amqp.Channel) *AMQPCeleryBackend {
	_ = "STUB: not implemented"
	return nil
}

// Reconnect reconnects to AMQP server
func (b *AMQPCeleryBackend) Reconnect() { _ = "STUB: not implemented"; return }

// GetResult retrieves result from AMQP queue
func (b *AMQPCeleryBackend) GetResult(taskID string) (*ResultMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// name
// durable
// autoDelete
// exclusive
// noWait
// args

// open channel temporarily

// SetResult sets result back to AMQP queue
func (b *AMQPCeleryBackend) SetResult(taskID string, result *ResultMessage) error {
	_ = "STUB: not implemented"
	return nil

	//queueName := taskID
}

// autodelete is automatically set to true by python
// (406) PRECONDITION_FAILED - inequivalent arg 'durable' for queue 'bc58c0d895c7421eb7cb2b9bbbd8b36f' in vhost '/': received 'true' but current is 'false'

// name
// durable
// autoDelete
// exclusive
// noWait
// args
