[![GoDoc](https://godoc.org/github.com/pvormste/yetwebutils/yetlog?status.svg)](https://godoc.org/github.com/pvormste/yetwebutils/yetlog)

# yetlog

`yetlog` provides a logger interface.

## Usage

```go
func DoSomething(logger yetlog.Logger) {
    logger.Warn("this function is useless", "name", "DoSomething")
}
```