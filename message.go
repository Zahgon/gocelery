// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package gocelery

import (
	"reflect"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// CeleryMessage is actual message to be sent to Redis
type CeleryMessage struct {
	Body            string                 `json:"body"`
	Headers         map[string]interface{} `json:"headers,omitempty"`
	ContentType     string                 `json:"content-type"`
	Properties      CeleryProperties       `json:"properties"`
	ContentEncoding string                 `json:"content-encoding"`
}

func (cm *CeleryMessage) reset() { _ = "STUB: not implemented"; return }

var celeryMessagePool = sync.Pool{
	New: func() interface{} {
		return &CeleryMessage{
			Body:        "",
			Headers:     nil,
			ContentType: "application/json",
			Properties: CeleryProperties{
				BodyEncoding:  "base64",
				CorrelationID: uuid.Must(uuid.NewV4()).String(),
				ReplyTo:       uuid.Must(uuid.NewV4()).String(),
				DeliveryInfo: CeleryDeliveryInfo{
					Priority:   0,
					RoutingKey: "celery",
					Exchange:   "celery",
				},
				DeliveryMode: 2,
				DeliveryTag:  uuid.Must(uuid.NewV4()).String(),
			},
			ContentEncoding: "utf-8",
		}
	},
}

func getCeleryMessage(encodedTaskMessage string) *CeleryMessage {
	_ = "STUB: not implemented"
	return nil
}

func releaseCeleryMessage(v *CeleryMessage) { _ = "STUB: not implemented"; return }

// CeleryProperties represents properties json
type CeleryProperties struct {
	BodyEncoding  string             `json:"body_encoding"`
	CorrelationID string             `json:"correlation_id"`
	ReplyTo       string             `json:"reply_to"`
	DeliveryInfo  CeleryDeliveryInfo `json:"delivery_info"`
	DeliveryMode  int                `json:"delivery_mode"`
	DeliveryTag   string             `json:"delivery_tag"`
}

// CeleryDeliveryInfo represents deliveryinfo json
type CeleryDeliveryInfo struct {
	Priority   int    `json:"priority"`
	RoutingKey string `json:"routing_key"`
	Exchange   string `json:"exchange"`
}

// GetTaskMessage retrieve and decode task messages from broker
func (cm *CeleryMessage) GetTaskMessage() *TaskMessage {
	_ = "STUB: not implemented"
	// ensure content-type is 'application/json'
	return nil
}

// ensure body encoding is base64

// ensure content encoding is utf-8

// decode body

// TaskMessage is celery-compatible message
type TaskMessage struct {
	ID      string                 `json:"id"`
	Task    string                 `json:"task"`
	Args    []interface{}          `json:"args"`
	Kwargs  map[string]interface{} `json:"kwargs"`
	Retries int                    `json:"retries"`
	ETA     *string                `json:"eta"`
	Expires *time.Time             `json:"expires"`
}

func (tm *TaskMessage) reset() { _ = "STUB: not implemented"; return }

var taskMessagePool = sync.Pool{
	New: func() interface{} {
		eta := time.Now().Format(time.RFC3339)
		return &TaskMessage{
			ID:      uuid.Must(uuid.NewV4()).String(),
			Retries: 0,
			Kwargs:  nil,
			ETA:     &eta,
		}
	},
}

func getTaskMessage(task string) *TaskMessage { _ = "STUB: not implemented"; return nil }

func releaseTaskMessage(v *TaskMessage) { _ = "STUB: not implemented"; return }

// DecodeTaskMessage decodes base64 encrypted body and return TaskMessage object
func DecodeTaskMessage(encodedBody string) (*TaskMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Encode returns base64 json encoded string
func (tm *TaskMessage) Encode() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ResultMessage is return message received from broker
type ResultMessage struct {
	ID        string        `json:"task_id"`
	Status    string        `json:"status"`
	Traceback interface{}   `json:"traceback"`
	Result    interface{}   `json:"result"`
	Children  []interface{} `json:"children"`
}

func (rm *ResultMessage) reset() { _ = "STUB: not implemented"; return }

var resultMessagePool = sync.Pool{
	New: func() interface{} {
		return &ResultMessage{
			Status:    "SUCCESS",
			Traceback: nil,
			Children:  nil,
		}
	},
}

func getResultMessage(val interface{}) *ResultMessage { _ = "STUB: not implemented"; return nil }

func getReflectionResultMessage(val *reflect.Value) *ResultMessage {
	_ = "STUB: not implemented"
	return nil
}

func releaseResultMessage(v *ResultMessage) { _ = "STUB: not implemented"; return }
