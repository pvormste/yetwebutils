package yetstopwatch

import (
	"testing"

	"github.com/pvormste/yetlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnable(t *testing.T) {
	t.Run("should enable with Enable()", func(t *testing.T) {
		require.False(t, enabled)

		Enable()
		assert.True(t, enabled)
	})

	t.Run("should disable with Disable()", func(t *testing.T) {
		require.True(t, enabled)

		Disable()
		assert.False(t, enabled)
	})

	t.Run("should set enabled with provided boolean", func(t *testing.T) {
		require.False(t, enabled)

		SetEnabled(true)
		assert.True(t, IsEnabled())
	})
}

func BenchmarkLogExecutionTimeFor(b *testing.B) {
	b.Run("disabled state", func(b *testing.B) {
		Disable()
		for n := 0; n < b.N; n++ {
			LogExecutionTimeFor("func", Now(), yetlog.NewNullLogger())
		}
	})

	b.Run("enabled state", func(b *testing.B) {
		Enable()
		for n := 0; n < b.N; n++ {
			LogExecutionTimeFor("func", Now(), yetlog.NewNullLogger())
		}
	})
}
