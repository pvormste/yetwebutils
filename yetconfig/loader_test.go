package yetconfig

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Name                 string
	Description          string
	TestEmbeddableConfig `config:"EMBEDDED"`
}

type TestEmbeddableConfig struct {
	Port int
}

func TestLoadEnvConfig(t *testing.T) {
	t.Run("should load all config values (from .env.dev, .env and variable)", func(t *testing.T) {
		assert := assert.New(t)

		err := os.Setenv("EMBEDDED__PORT", "8080")
		require.NoError(t, err)

		cfg := TestConfig{}
		LoadEnvConfig(&cfg, "./testdata")

		assert.Equal(cfg.Name, "STAGING APP")            // from .env
		assert.Equal(cfg.Description, "DEV DESCRIPTION") // from env.dev
		assert.Equal(cfg.Port, 8080)                     // from enviornment variable
	})

	t.Run("should not crash when a .env file is missing", func(t *testing.T) {
		assert := assert.New(t)

		err := os.Setenv("ENVIRONMENT", "production")
		require.NoError(t, err)

		cfg := TestConfig{}
		LoadEnvConfig(&cfg, "./testdata")

		assert.Equal(cfg.Name, "STAGING APP") // from .env
		assert.Equal(cfg.Description, "")     // because .env.prod is missing
		assert.Equal(cfg.Port, 8080)          // from enviornment variable
	})
}
