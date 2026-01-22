package logger

type Field struct {
	Key   string
	Value any
}

func String(k, v string) Field    { return Field{k, v} }
func Int(k string, v int) Field   { return Field{k, v} }
func Bool(k string, v bool) Field { return Field{k, v} }
func Any(k string, v any) Field   { return Field{k, v} }
func Err(err error) Field         { return Field{"error", err} }
