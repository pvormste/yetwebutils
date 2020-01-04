package yetconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectEnvFileByEnvironment(t *testing.T) {
	t.Run("should return .env.prod for production", func(t *testing.T) {
		assert := assert.New(t)

		err := os.Setenv("ENVIRONMENT", "production")
		require.NoError(t, err)

		actualEnvFile := selectEnvFileByEnvironment()
		assert.Equal(envFileProduction, actualEnvFile)
	})

	t.Run("should return .env.staging for staging", func(t *testing.T) {
		assert := assert.New(t)

		err := os.Setenv("ENVIRONMENT", "staging")
		require.NoError(t, err)

		actualEnvFile := selectEnvFileByEnvironment()
		assert.Equal(envFileStaging, actualEnvFile)
	})

	t.Run("should return .env.dev for develop", func(t *testing.T) {
		assert := assert.New(t)

		err := os.Setenv("ENVIRONMENT", "develop")
		require.NoError(t, err)

		actualEnvFile := selectEnvFileByEnvironment()
		assert.Equal(envFileDevelop, actualEnvFile)
	})
}

func TestMergeEnvFilePath(t *testing.T) {
	t.Run("should create a file path to the current directory when path is only dot", func(t *testing.T) {
		assert := assert.New(t)

		actualFilePath := mergeEnvFilePath(".env", ".")
		assert.Equal(".env", actualFilePath)
	})

	t.Run("should successfully create a file path when path is provided", func(t *testing.T) {
		assert := assert.New(t)

		actualFilePath := mergeEnvFilePath(".env", "./testdata")
		assert.Equal("testdata/.env", actualFilePath)
	})
}
