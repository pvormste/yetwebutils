package yetlog

// Logger is universal interface for implementing a logger.
type Logger interface {
	Debug(message string, fields ...interface{})
	Info(message string, fields ...interface{})
	Warn(message string, fields ...interface{})
	Error(message string, fields ...interface{})
	Fatal(message string, fields ...interface{})
}
