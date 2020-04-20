[![GoDoc](https://godoc.org/github.com/pvormste/yetwebutils/yethttp?status.svg)](https://godoc.org/github.com/pvormste/yetwebutils/yethttp)

# yethttp

`yethttp` reduces the amount of boilerplate to write a simple http server with graceful shutdown.

## Usage

```go
type WebApp struct {
    yethttp.EmbeddableServerWrapper
}

func NewWebApp(logger yetlog.Logger) WebApp {
    serverWrapper := yethttp.NewEmbeddableServerWrapper(logger, 8080)

    return WebApp{
        EmbeddableServerWrapper: serverWrapper,
    }
}

func main() {
    webApp := NewWebApp(yetlog.NewNullLogger())
    if err := webApp.Serve(context.Background(), yethttp.DefaultRoutesFunc); err != nil {
        panic(err)
    }

    if err := webApp.WaitForShutdown(context.Background()); err != nil {
        panic(err)
    }      
}
```