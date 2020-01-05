// Package yethttp provides helpers for working with http.
//
// Example:
//  type WebApp struct {
//    yethttp.ServerWrapper
//  }
//
//  func NewWebApp(logger yetlog.Logger) WebApp {
//    serverWrapper := yethttp.NewServerWrapper(logger, 8080, http.NewServeMux())
//
//    return WebApp{
//      ServerWrapper: serverWrapper,
//    }
//  }
//
//  func main() {
//    webApp := NewWebApp(yetlog.NewNullLogger())
//    if err := webApp.Serve(context.Background()); err != nil {
//      panic(err)
//    }
//  }
package yethttp
