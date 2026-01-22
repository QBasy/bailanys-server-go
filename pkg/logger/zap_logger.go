package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log *zap.Logger
}

func NewZapLogger(cfg Config, service string) *ZapLogger {
	cfg.ApplyDefaults()

	cores := []zapcore.Core{}

	if cfg.Outputs.Stdout {
		cores = append(cores, NewJSONCore(os.Stdout, ParseLevel(cfg.Level)))
	}

	core := zapcore.NewTee(cores...)

	l := zap.New(core).With(
		zap.String("service", service),
	)

	return &ZapLogger{log: l}
}
