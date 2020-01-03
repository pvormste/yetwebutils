# yetenv

`yetenv` is small util package which helps to determine on which environment the application is running. It reads the `ENVIRONMENT` variable.

The `ENVIRONMENT` values are not case-sensitives.

| `ENVIRONMENT` value | Constant |
| ------------------- | -------- |
| `production` | yetenv.Production |
| `staging` | yetenv.Staging |
| any other value | yetenv.Develop |
