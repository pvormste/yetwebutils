package yethttp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pvormste/yetwebutils/yetlog"
)

// ServerWrapper wraps a http.Server and can handle server startup, routing and graceful shutdown.
type ServerWrapper struct {
	HttpServer *http.Server
	HttpPort   int
	logger     yetlog.Logger
}

// NewServerWrapper returns a new ServerWrapper. It uses the yetlog.Logger interface for logging and needs a port
// and a mux as http.Handler.
func NewServerWrapper(logger yetlog.Logger, port int, mux http.Handler) ServerWrapper {
	wrapper := ServerWrapper{
		HttpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
		HttpPort: port,
		logger:   logger,
	}

	return wrapper
}

// Serve starts the server and listens for new connections.
func (serverWrapper *ServerWrapper) Serve(ctx context.Context) error {
	c := make(chan error)
	go func() {
		serverWrapper.logger.Info("starting server", "port", serverWrapper.HttpPort)
		if err := serverWrapper.HttpServer.ListenAndServe(); err != nil {
			c <- err
		}
	}()

	go func() {
		<-ctx.Done()
		if err := serverWrapper.GracefulShutdown(ctx); err != nil {
			serverWrapper.logger.Error("could not shutdown http server gracefully", "port", serverWrapper.HttpPort)
		}
	}()

	select {
	case err := <-c:
		return err
	default:
		return nil
	}
}

// GracefulShutdown will shutdown the underlying http server gracefully.
func (serverWrapper *ServerWrapper) GracefulShutdown(ctx context.Context) error {
	serverWrapper.logger.Info("shutting down http server gracefully", "port", serverWrapper.HttpPort)
	return serverWrapper.HttpServer.Shutdown(ctx)
}

// WaitForShutdown blocks the go routine and will only continue when it gets a kill signal (SIGINt, SIGTERM, ...).
func (serverWrapper *ServerWrapper) WaitForShutdown(ctx context.Context) error {
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-kill
	if err := serverWrapper.GracefulShutdown(ctx); err != nil {
		return fmt.Errorf("wait-for-shutdown: %w", err)
	}

	return nil
}
