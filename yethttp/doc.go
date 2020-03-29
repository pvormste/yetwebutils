// Package yethttp provides helpers for working with http.
//
// Example:
//  type WebApp struct {
//    yethttp.EmbeddableServerWrapper
//  }
//
//  func NewWebApp(logger yetlog.Logger) WebApp {
//    serverWrapper := yethttp.NewEmbeddableServerWrapper(logger, 8080)
//
//    return WebApp{
//      EmbeddableServerWrapper: serverWrapper,
//    }
//  }
//
//  func main() {
//    webApp := NewWebApp(yetlog.NewNullLogger())
//    if err := webApp.Serve(context.Background()); err != nil {
//      panic(err)
//    }
//
//    if err := webApp.WaitForShutdown(context.Background()); err != nil {
//      panic(err)
//    }
//  }
package yethttp
