// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"github.com/streadway/amqp"
)

// AMQPExchange stores AMQP Exchange configuration
type AMQPExchange struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
}

// NewAMQPExchange creates new AMQPExchange
func NewAMQPExchange(name string) *AMQPExchange { _ = "STUB: not implemented"; return nil }

// AMQPQueue stores AMQP Queue configuration
type AMQPQueue struct {
	Name       string
	Durable    bool
	AutoDelete bool
}

// NewAMQPQueue creates new AMQPQueue
func NewAMQPQueue(name string) *AMQPQueue { _ = "STUB: not implemented"; return nil }

// AMQPCeleryBroker is RedisBroker for AMQP
type AMQPCeleryBroker struct {
	*amqp.Channel
	Connection       *amqp.Connection
	Exchange         *AMQPExchange
	Queue            *AMQPQueue
	consumingChannel <-chan amqp.Delivery
	Rate             int
}

// NewAMQPConnection creates new AMQP channel
func NewAMQPConnection(host string) (*amqp.Connection, *amqp.Channel) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewAMQPCeleryBroker creates new AMQPCeleryBroker
func NewAMQPCeleryBroker(host string) *AMQPCeleryBroker { _ = "STUB: not implemented"; return nil }

// NewAMQPCeleryBrokerByConnAndChannel creates new AMQPCeleryBroker using AMQP conn and channel
func NewAMQPCeleryBrokerByConnAndChannel(conn *amqp.Connection, channel *amqp.Channel) *AMQPCeleryBroker {
	_ = "STUB: not implemented"
	return nil
}

// StartConsumingChannel spawns receiving channel on AMQP queue
func (b *AMQPCeleryBroker) StartConsumingChannel() error { _ = "STUB: not implemented"; return nil }

// SendCeleryMessage sends CeleryMessage to broker
func (b *AMQPCeleryBroker) SendCeleryMessage(message *CeleryMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// name
// durable
// autoDelete
// exclusive
// noWait
// args

// GetTaskMessage retrieves task message from AMQP queue
func (b *AMQPCeleryBroker) GetTaskMessage() (*TaskMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateExchange declares AMQP exchange with stored configuration
func (b *AMQPCeleryBroker) CreateExchange() error { _ = "STUB: not implemented"; return nil }

// CreateQueue declares AMQP Queue with stored configuration
func (b *AMQPCeleryBroker) CreateQueue() error { _ = "STUB: not implemented"; return nil }
