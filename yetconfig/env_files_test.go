package yetconfig

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestSelectEnvFileByEnvironment(t *testing.T) {
	t.Run("should return .env.prod for production", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		err := os.Setenv("ENVIRONMENT", "production")
		tt.Expect(err).To(Not(HaveOccurred()))

		actualEnvFile := selectEnvFileByEnvironment()
		tt.Expect(actualEnvFile).To(Equal(envFileProduction))
	})

	t.Run("should return .env.staging for staging", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		err := os.Setenv("ENVIRONMENT", "staging")
		tt.Expect(err).To(Not(HaveOccurred()))

		actualEnvFile := selectEnvFileByEnvironment()
		tt.Expect(actualEnvFile).To(Equal(envFileStaging))
	})

	t.Run("should return .env.dev for develop", func(t *testing.T) {
		tt := NewGomegaWithT(t)

		err := os.Setenv("ENVIRONMENT", "develop")
		tt.Expect(err).To(Not(HaveOccurred()))

		actualEnvFile := selectEnvFileByEnvironment()
		tt.Expect(actualEnvFile).To(Equal(envFileDevelop))
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
