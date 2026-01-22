package logger

import (
	"io"

	"go.uber.org/zap/zapcore"
)

func NewJSONCore(w io.Writer, level zapcore.Level) zapcore.Core {
	enc := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:    "ts",
		LevelKey:   "level",
		MessageKey: "message",

		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	})

	return zapcore.NewCore(enc, zapcore.AddSync(w), level)
}
