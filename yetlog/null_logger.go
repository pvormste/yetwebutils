package yetlog

// NullLoger is a logger implementation which does nothing. Can be used for testing.
type NullLogger struct{}

// NewNullLogger creates a new NullLogger.
func NewNullLogger() Logger {
	return NullLogger{}
}

func (l NullLogger) Debug(message string, fields ...interface{}) {}

func (l NullLogger) Info(message string, fields ...interface{}) {}

func (l NullLogger) Warn(message string, fields ...interface{}) {}

func (l NullLogger) Error(message string, fields ...interface{}) {}

func (l NullLogger) Fatal(message string, fields ...interface{}) {}
