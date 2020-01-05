[![GoDoc](https://godoc.org/github.com/pvormste/yetwebutils/yetnet?status.svg)](https://godoc.org/github.com/pvormste/yetwebutils/yetnet)

# yetnet

`yetnet` contains helper code for working with low level network like TCP, UDP and Ports.

# Usage

```go
isOpen, err := yetnet.IsPortOpen(8080)
```