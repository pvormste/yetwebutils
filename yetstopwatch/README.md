# yetstopwatch

`yetstopwatch` can be used to measure execution time of a function. [Documentation on pkg.go.dev](https://pkg.go.dev/github.com/pvormste/yetwebutils/yetstopwatch)

## Usage

```go
func QueryDatabase(logger yetlog.Logger) error {
    defer LogExecutionTimeFor("QueryDatabase()", yetstopwatch.Now(), logger)
    // ... function logic starts here
}
```

## Benchmarks

Be careful when using this in your hot path.
Here are some Benchmark results:

```
BenchmarkLogExecutionTimeFor/disabled_state-8           139697289                8.53 ns/op            0 B/op          0 allocs/op
BenchmarkLogExecutionTimeFor/enabled_state-8             8390904               140 ns/op              96 B/op          3 allocs/op
```