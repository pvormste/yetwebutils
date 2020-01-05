package yetnet

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsPortInUse(t *testing.T) {
	l, err := net.Listen("tcp", ":0")
	assert.NoError(t, err)

	port := l.Addr().(*net.TCPAddr).Port

	t.Run("should return true for IsPortInUse()", func(t *testing.T) {
		actualIsInUse, err := IsPortInUse(port)

		assert.NoError(t, err)
		assert.True(t, actualIsInUse)
	})

	t.Run("should return true for IsPortOpen()", func(t *testing.T) {
		err := l.Close()
		require.NoError(t, err)

		actualIsOpen, err := IsPortOpen(port)

		assert.NoError(t, err)
		assert.True(t, actualIsOpen)
	})
}
