package logger

import (
	"context"

	"go.uber.org/zap"
)

func (l *ZapLogger) With(fields ...Field) Logger {
	return &ZapLogger{log: l.log.With(toZap(fields)...)}
}

func (l *ZapLogger) WithContext(ctx context.Context) Logger {
	if v := ctx.Value("trace_id"); v != nil {
		if s, ok := v.(string); ok && s != "" {
			return l.With(String("trace_id", s))
		}
	}
	return l
}

func (l *ZapLogger) Sync() error {
	return l.log.Sync()
}

func toZap(fields []Field) []zap.Field {
	out := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		out = append(out, zap.Any(f.Key, f.Value))
	}
	return out
}

func splitArgs(args []any) (msgArgs []any, fields []Field) {
	for _, a := range args {
		if f, ok := a.(Field); ok {
			fields = append(fields, f)
		} else {
			msgArgs = append(msgArgs, a)
		}
	}
	return
}
