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

// EmbeddableServerWrapper wraps a http.Server and can handle server startup, routing and graceful shutdown.
type EmbeddableServerWrapper struct {
	HttpServer *http.Server
	HttpPort   int
	logger     yetlog.Logger
}

// NewEmbeddableServerWrapper returns a new EmbeddableServerWrapper.
func NewEmbeddableServerWrapper(logger yetlog.Logger, port int) EmbeddableServerWrapper {
	wrapper := EmbeddableServerWrapper{
		HttpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", port),
		},
		HttpPort: port,
		logger:   logger,
	}

	return wrapper
}

// Serve starts the server and listens for new connections.
func (e *EmbeddableServerWrapper) Serve(ctx context.Context) error {
	e.HttpServer.Handler = e.Routes()

	c := make(chan error)
	go func() {
		e.logger.Info("starting server", "port", e.HttpPort)
		if err := e.StartServer(); err != nil {
			c <- err
		}
	}()

	go func() {
		<-ctx.Done()
		if err := e.GracefulShutdown(ctx); err != nil {
			e.logger.Error("could not shutdown http server gracefully", "port", e.HttpPort)
		}
	}()

	select {
	case err := <-c:
		return err
	default:
		return nil
	}
}

// StartServer starts the underlying http.Server by using httpServer.ListenAndServe(). This method can be overwritten
// to be able to use framework specific function calls.
func (e *EmbeddableServerWrapper) StartServer() error {
	return e.HttpServer.ListenAndServe()
}

// GracefulShutdown will shutdown the underlying http.Server gracefully. This method can be overwritten to be able
// to use framework specific function calls.
func (e *EmbeddableServerWrapper) GracefulShutdown(ctx context.Context) error {
	e.logger.Info("shutting down http server gracefully", "port", e.HttpPort)
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

// Routes is used to define the server routes and handlers. It uses the http.DefaultServeMux by default which shouldn't
// be used in any productive application. This method is meant to be overwritten by the user of this package.
func (e *EmbeddableServerWrapper) Routes() http.Handler {
	return http.DefaultServeMux
}
