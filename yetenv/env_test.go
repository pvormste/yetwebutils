package yetenv

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetEnvironment(t *testing.T) {
	t.Run("should return environment 'production'", func(t *testing.T) {
		values := []string{
			"production",
			"PRODUCTION",
		}

		for _, value := range values {
			value := value

			t.Run(fmt.Sprintf("ENVIRONMENT=%s", value), func(t *testing.T) {
				tt := NewGomegaWithT(t)

				err := os.Setenv("ENVIRONMENT", value)
				tt.Expect(err).ToNot(HaveOccurred())

				actualEnv := GetEnvironment()
				tt.Expect(actualEnv).To(Equal(Production))
			})
		}
	})

	t.Run("should return environemnt 'staging'", func(t *testing.T) {
		values := []string{
			"staging",
			"STAGING",
		}

		for _, value := range values {
			value := value

			t.Run(fmt.Sprintf("ENVIRONMENT=%s", value), func(t *testing.T) {
				tt := NewGomegaWithT(t)

				err := os.Setenv("ENVIRONMENT", value)
				tt.Expect(err).ToNot(HaveOccurred())

				actualEnv := GetEnvironment()
				tt.Expect(actualEnv).To(Equal(Staging))
			})
		}
	})

	t.Run("should return environemnt 'develop' for any other value of ENVIRONMENT", func(t *testing.T) {
		values := []string{
			"develop",
			"DEVELOP",
			"ANY",
		}

		for _, value := range values {
			value := value

			t.Run(fmt.Sprintf("ENVIRONMENT=%s", value), func(t *testing.T) {
				tt := NewGomegaWithT(t)

				err := os.Setenv("ENVIRONMENT", value)
				tt.Expect(err).ToNot(HaveOccurred())

				actualEnv := GetEnvironment()
				tt.Expect(actualEnv).To(Equal(Develop))
			})
		}
	})
}
