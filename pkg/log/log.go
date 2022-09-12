package log

type Fields = map[string]interface{}

type Logger interface {
	Trace(message string, fields Fields)
	Debug(message string, fields Fields)
	Info(message string, fields Fields)
	Warn(message string, fields Fields)
	Error(message string, err error, fields Fields)
	Fatal(message string, err error, fields Fields)
	Panic(message string, err error, fields Fields)
}
