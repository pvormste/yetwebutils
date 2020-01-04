package yetconfig

import (
	"log"

	"github.com/JeremyLoy/config"
)

func LoadEnvConfig(targetConfig interface{}, pathToEnvFiles string) {
	envFile := selectEnvFileByEnvironment()

	envFilePath := mergeEnvFilePath(envFile, pathToEnvFiles)
	customEnvFilePath := mergeEnvFilePath(envFileCustom, pathToEnvFiles)

	if err := config.From(envFilePath).From(customEnvFilePath).FromEnv().To(targetConfig); err != nil {
		log.Printf("WARNING: some errors occured on config load: %s", err.Error())
	}
}
