package logger

import "context"

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type Logger interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)

	With(fields ...Field) Logger
	WithContext(ctx context.Context) Logger
	Sync() error
}
