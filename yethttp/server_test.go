package yethttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pvormste/yetwebutils/yetlog"
)

func (serverWrapper *ServerWrapper) createTestInstance(t *testing.T) (testServerWrapper ServerWrapper, port int) {
	newWrapperInstance := *serverWrapper

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
	serverWrapper := NewServerWrapper(logger, 0, http.NewServeMux())

	serverWrapperTestInstance, testPort := serverWrapper.createTestInstance(t)
	ctx, cancelFunc := context.WithCancel(context.Background())

	t.Run("should successfully start a server", func(t *testing.T) {
		assert := assert.New(t)

		err := serverWrapperTestInstance.Serve(ctx)
		require.NoError(t, err)

		assert.Eventually(func() bool {
			return checkIfPortIsInUse(t, testPort)
		}, time.Second, 5*time.Millisecond)
	})

	t.Run("should stop server successfully", func(t *testing.T) {
		assert := assert.New(t)

		cancelFunc()
		assert.Eventually(func() bool {
			return !checkIfPortIsInUse(t, testPort)
		}, time.Second, 5*time.Millisecond)
	})
}

func checkIfPortIsInUse(t *testing.T, port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return true
	}

	if err := l.Close(); err != nil {
		t.Fatal(err)
	}

	return false
}
