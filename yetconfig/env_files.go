package yetconfig

import (
	"path"

	"github.com/pvormste/yetenv"
)

const (
	envFileProduction = ".env.prod"
	envFileStaging    = ".env.staging"
	envFileDevelop    = ".env.dev"
	envFileCustom     = ".env"
)

func selectEnvFileByEnvironment() string {
	env := yetenv.GetEnvironment()

	switch env {
	case yetenv.Production:
		return envFileProduction
	case yetenv.Staging:
		return envFileStaging
	}

	return envFileDevelop
}

func mergeEnvFilePath(envFileName string, pathToEnvFiles string) string {
	if pathToEnvFiles == "." {
		pathToEnvFiles = "./"
	}

	return path.Join(pathToEnvFiles, envFileName)
}
