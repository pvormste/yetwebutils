# yetlog

`yetlog` provides a logger interface.

## Usage

```go
func DoSomething(logger yetlog.Logger) {
    logger.Warn("this function is useless", "name", "DoSomething")
}
```