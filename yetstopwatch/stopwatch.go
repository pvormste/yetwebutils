package yetstopwatch

import (
	"time"

	"github.com/pvormste/yetlog"
)

// Time provides an own alias type for wrapping a Time implementation.
type Time time.Time

// String can be used to get a string representation of the time.
func (t Time) String() string {
	return time.Time(t).String()
}

var enabled bool

// Enable enables the stopwatch package.
func Enable() {
	enabled = true
}

// Disable disables the stopwatch package.
func Disable() {
	enabled = false
}

// SetEnabled can be used to explictly set the enabled value.
func SetEnabled(newEnabledValue bool) {
	enabled = newEnabledValue
}

// IsEnabled returns the actual value of enabled.
func IsEnabled() bool {
	return enabled
}

// Now can be used to return the current time when calling this function.
func Now() Time {
	return Time(time.Now())
}

// LogExecutionTimeFor logs the execution time by using the provided logger. It only works when the package is enabled.
func LogExecutionTimeFor(name string, startTme Time, logger yetlog.Logger) {
	if !enabled {
		return
	}

	elapsed := time.Since(time.Time(startTme))
	logger.Debug("execution time", "name", name, "time", elapsed.String())
}
