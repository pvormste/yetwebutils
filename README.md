[![GitHub license](https://img.shields.io/github/license/pvormste/yetwebutils)](https://github.com/pvormste/yetwebutils/blob/master/LICENSE)

# Yet 'Another' Web Utils (yetwebutils)

`yetwebutils` is a highly opinionated collection of go packages, which reduces the amount of boilerplate code for simple web projects.
Remember: **It is highly opinionated. HIGHLY.** :-)

But essentially it acts like a second standard lib for creating web applications.

## Packages

All packages are prefixed with `yet`.

| package | Description | GoDoc |
| ------- | ----------- | ----- |
| [yetconfig](https://github.com/pvormste/yet-web-utils/tree/master/yetconfig) | Helper for loading config from environment | |
| [yetenv](https://github.com/pvormste/yet-web-utils/tree/master/yetenv) | Provides some logic to determine the environment (develop, staging, production) | |
| [yethttp](https://github.com/pvormste/yet-web-utils/tree/master/yethttp) | Provides helpers for working with http | |
| [yetlog](https://github.com/pvormste/yet-web-utils/tree/master/yetlog) | Provides a logger agnostic logging interface which is used by most `yet` packages | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetlog) |
| [yetnet](https://github.com/pvormste/yet-web-utils/tree/master/yetnet) | Provides helper code for working with low level network like TCP, UDP and Ports | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetnet) |
 