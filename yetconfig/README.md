# yetconfig

`yetconfig` is a config loader for environment config. It follows the principle from [12 Factor App](https://12factor.net/config) which requires the config to be stored in the environment.

So this package utilizes the package [github.com/JeremyLoy/config](https://github.com/JeremyLoy/config) which can be used to load the config from the environment.

## How it works

The env config will be loaded in 3 stages being the latest one which overrides the stages before.

#### Stage 1: Predefined .env files per environment

If you want to have some default configurations which is does not contain any sensitive data, you can store them by environment and will be loaded first. See following table:

| env file | loaded on environment |
| -------- | --------------------- |
| `.env.dev` | `yetenv.Develop` |
| `.env.staging` | `yetenv.Staging` |
| `.env.prod` | `yetenv.Production` |

**NOTE:** If env file is missing on load you will get a warning logged. This is because how the library handles file I/O errors.

#### Stage 2: Custom .env files (overrides Stage 1)

A simple `.env` file will override all defined values from stage 1.

**NOTE:** As in stage 1 you will get a warning when the file is missing.

Example:

```
# .env.dev

PORT=8080
```

```
# .env

PORT=80
```

The actual value for `PORT` will be `80`.

#### Stage 3: Variables in the environment (overrides Stage 2 and Stage 1)

All variables which are present in the environment will override the values from .env files.

Example:
```
$ export PORT=8001
```

The value for `PORT` will be `8001` regardless of how the value was set before.

## Usage

NOTE: You can use `"."` to load the .env files from current working directory.

```go
type Config struct {
    Port int
}

cfg := Config{}
LoadEnvConfig(&cfg, ".")

// Use the config
fmt.Println(cfg.Port)
```