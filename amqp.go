// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"github.com/streadway/amqp"
)

// deliveryAck acknowledges delivery message with retries on error
func deliveryAck(delivery amqp.Delivery) { _ = "STUB: not implemented"; return }
