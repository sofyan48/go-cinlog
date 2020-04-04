package entity

import (
	"time"
)

// SessionGlobal ...
type SessionGlobal struct {
	URL     string
	Service string
}

// SaveLogger ...
type SaveLogger struct {
	UUID   string                 `json:"uuid"`
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

// GetLoggerRequest ...
type GetLoggerRequest struct {
	UUID   string `json:"uuid" bson:"uuid"`
	Action string `json:"action" bson:"action"`
}

// GetAllLoggerRequest ...
type GetAllLoggerRequest struct {
	Action string `json:"action" bson:"action"`
}

// LoggerEventHistory ...
type LoggerEventHistory struct {
	UUID      string          `json:"__uuid"`
	Action    string          `json:"__action"`
	Offset    uint64          `json:"__offset"`
	History   []LoggerHistory `json:"history"`
	Status    string          `json:"status"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdateAt  *time.Time      `json:"update_at"`
}

// LoggerHistory ...
type LoggerHistory struct {
	Data map[string]interface{} `json:"data"`
}
