# yethttp

`yethttp` reduces the amount of boilerplate to write a simple http server with graceful shutdown.

## Usage

```go
type WebApp struct {
    yethttp.ServerWrapper
}

func NewWebApp(logger yetlog.Logger) WebApp {
    serverWrapper := yethttp.NewServerWrapper(logger, 8080, http.NewServeMux())

    return WebApp{
        ServerWrapper: serverWrapper,
    }
}

func main() {
    webApp := NewWebApp(yetlog.NewNullLogger())
    if err := webApp.Serve(context.Background()); err != nil {
        panic(err)
    }   
}
```