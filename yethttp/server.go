package yethttp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pvormste/yetlog"
)

// StartServerFunc is a function which can be used to customize the server start behavior.
type StartServerFunc func(server *http.Server) error

// DefaultStartServerFunc provides a default implementation for starting a http.Server which basically calls
// its ListenAndServe() method.
var DefaultStartServerFunc = func(server *http.Server) error {
	return server.ListenAndServe()
}

// RoutesFunc is used to define the server routes and handlers.
type RoutesFunc func() http.Handler

// DefaultRoutesFunc uses the http.DefaultServeMux which shouldn't be used in
// any productive application. This method is meant to be used for prototyping or testing.
var DefaultRoutesFunc = func() http.Handler {
	return http.DefaultServeMux
}

// EmbeddableServerWrapper wraps a http.Server and can handle server startup, routing and graceful shutdown.
type EmbeddableServerWrapper struct {
	HttpServer      *http.Server
	Port            int
	startServerFunc StartServerFunc
	logger          yetlog.Logger
}

// NewEmbeddableServerWrapper returns a new EmbeddableServerWrapper.
func NewEmbeddableServerWrapper(logger yetlog.Logger, port int) EmbeddableServerWrapper {
	wrapper := EmbeddableServerWrapper{
		HttpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
		Port:            port,
		startServerFunc: DefaultStartServerFunc,
		logger:          logger,
	}

	return wrapper
}

// Serve starts the server and listens for new connections.
func (e *EmbeddableServerWrapper) Serve(ctx context.Context, routesFunc RoutesFunc) error {
	e.HttpServer.Handler = routesFunc()

	c := make(chan error)
	go func() {
		e.logger.Info("starting server", "port", e.Port)
		if err := e.startServerFunc(e.HttpServer); err != nil {
			c <- err
		}
	}()

	go func() {
		<-ctx.Done()
		if err := e.GracefulShutdown(ctx); err != nil {
			e.logger.Error("could not shutdown http server gracefully", "port", e.Port)
		}
	}()

	select {
	case err := <-c:
		return err
	default:
		return nil
	}
}

// SetStartServer is used to overwrite the internal StartServerFunc. It should be called before using
// the Serve() method.
func (e *EmbeddableServerWrapper) SetStartServerFunc(startServerFunc StartServerFunc) {
	e.startServerFunc = startServerFunc
}

// GracefulShutdown will shutdown the underlying http.Server gracefully. This method can be overwritten to be able
// to use framework specific function calls.
func (e *EmbeddableServerWrapper) GracefulShutdown(ctx context.Context) error {
	e.logger.Info("shutting down http server gracefully", "port", e.Port)
	return e.HttpServer.Shutdown(ctx)
}

// WaitForShutdown blocks the go routine and will only continue when it receives a kill signal (SIGINT, SIGTERM, ...).
func (e *EmbeddableServerWrapper) WaitForShutdown(ctx context.Context) error {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-kill
	if err := e.GracefulShutdown(ctx); err != nil {
		return fmt.Errorf("wait-for-shutdown: %w", err)
	}

	return nil
}
