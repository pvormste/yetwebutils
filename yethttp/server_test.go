package yethttp

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/pvormste/yetlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pvormste/yetwebutils/yetnet"
)

func (e *EmbeddableServerWrapper) createTestInstance(t *testing.T) (testServerWrapper EmbeddableServerWrapper, port int) {
	newWrapperInstance := *e

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	defer func() {
		if err := listener.Close(); err != nil {
			require.NoError(t, err)
		}
	}()

	addr := listener.Addr()
	port = addr.(*net.TCPAddr).Port
	newWrapperInstance.HttpServer.Addr = fmt.Sprintf(":%d", port)

	return newWrapperInstance, port
}

func TestApplication_Serve(t *testing.T) {
	logger := yetlog.NewNullLogger()
	serverWrapper := NewEmbeddableServerWrapper(logger, 0)

	serverWrapperTestInstance, testPort := serverWrapper.createTestInstance(t)
	ctx, cancelFunc := context.WithCancel(context.Background())

	t.Run("should successfully start a server", func(t *testing.T) {
		assert := assert.New(t)

		err := serverWrapperTestInstance.Serve(ctx, DefaultRoutesFunc)
		require.NoError(t, err)

		assert.Eventually(func() bool {
			isInUse, err := yetnet.IsPortInUse(testPort)
			require.NoError(t, err)
			return isInUse

		}, time.Second, 5*time.Millisecond)
	})

	t.Run("should stop server successfully and execute afterShutdownFunc", func(t *testing.T) {
		assert := assert.New(t)

		shutdownChan := make(chan bool)
		afterShutdownFunc := func() {
			close(shutdownChan)
		}
		serverWrapperTestInstance.AddAfterShutdownFunc(afterShutdownFunc)

		cancelFunc()
		assert.Eventually(func() bool {
			<-shutdownChan
			isOpen, err := yetnet.IsPortOpen(testPort)
			require.NoError(t, err)
			return isOpen
		}, time.Second, 5*time.Millisecond)
	})
}
