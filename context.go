package slogzap

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type contextKey string

var (
	fields contextKey = "slog_fields"
)

func WithValue(parent context.Context, key string, val any) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	if v, ok := parent.Value(fields).(*sync.Map); ok {
		mapCopy := copySyncMap(v)
		mapCopy.Store(key, val)
		return context.WithValue(parent, fields, mapCopy)
	}
	v := &sync.Map{}
	v.Store(key, val)
	return context.WithValue(parent, fields, v)
}
func StartTrace(parent context.Context) context.Context {
	traceId := uuid.New().String()
	return WithValue(parent, "trace_id", traceId)
}
func copySyncMap(m *sync.Map) *sync.Map {
	var cp sync.Map
	m.Range(func(k, v interface{}) bool {
		cp.Store(k, v)
		return true
	})
	return &cp
}
