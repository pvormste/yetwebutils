// Package yetstopwatch can be used to measure execution time of a function.
//
// Example:
//   func QueryDatabase(logger yetlog.Logger) error {
//     defer LogExecutionTimeFor("QueryDatabase()", yetstopwatch.Now(), logger)
//     // ... function logic starts here
//   }

package yetstopwatch
