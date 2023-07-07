package main

import (
	"fmt"
	"log"

	"testing"
	"time"
)

type logEntry struct {
	Severity     string            `json:"severity"`
	Message      interface{}       `json:"message"`
	Timestamp    time.Time         `json:"timestamp"`
	Labels       map[string]string `json:"logging.googleapis.com/labels,omitempty"`
	SpanID       string            `json:"logging.googleapis.com/spanId,omitempty"`
	TraceID      string            `json:"logging.googleapis.com/trace,omitempty"`
	TraceSampled bool              `json:"logging.googleapis.com/trace_sampled,omitempty"`
}

func TestServer(t *testing.T) {
	t.Cleanup(cleanup)
	Log()
}

func cleanup() {
	time.Sleep(time.Millisecond * 1500)
}

func Log() bool {
	entry := logEntry{
		Severity:     "",
		Message:      nil,
		Timestamp:    time.Time{},
		Labels:       nil,
		SpanID:       "",
		TraceID:      "",
		TraceSampled: false,
	}
	log.Println(entry)
	fmt.Println(entry)
	return true
}
