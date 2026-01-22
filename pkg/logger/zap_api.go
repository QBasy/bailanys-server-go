package logger

import "fmt"

func (l *ZapLogger) Debug(format string, args ...any) {
	msgArgs, fields := splitArgs(args)
	l.log.Debug(fmt.Sprintf(format, msgArgs...), toZap(fields)...)
}

func (l *ZapLogger) Info(format string, args ...any) {
	msgArgs, fields := splitArgs(args)
	l.log.Info(fmt.Sprintf(format, msgArgs...), toZap(fields)...)
}

func (l *ZapLogger) Warn(format string, args ...any) {
	msgArgs, fields := splitArgs(args)
	l.log.Warn(fmt.Sprintf(format, msgArgs...), toZap(fields)...)
}

func (l *ZapLogger) Error(format string, args ...any) {
	msgArgs, fields := splitArgs(args)
	l.log.Error(fmt.Sprintf(format, msgArgs...), toZap(fields)...)
}
