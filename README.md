[![GitHub license](https://img.shields.io/github/license/pvormste/yetwebutils)](https://github.com/pvormste/yetwebutils/blob/master/LICENSE) ![](https://github.com/pvormste/yetwebutils/workflows/lint/badge.svg?branch=master) ![](https://github.com/pvormste/yetwebutils/workflows/tests/badge.svg?branch=master)

# Yet 'Another' Web Utils (yetwebutils)

`yetwebutils` is a highly opinionated collection of go packages, which reduces the amount of boilerplate code for simple web projects.

But essentially it acts like another standard lib for creating web applications.

_This project is more like an experiment on how much web functionality can be abstracted in own reusable packages and I'm very curious how this will evolve. The main goal is to reduce boilerplate code for greenfield web projects. Nervertheless feel free to use it for your own needs!_

## Packages

All packages are prefixed with `yet`.

| package | Description | GoDoc |
| ------- | ----------- | ----- |
| [yetconfig](https://github.com/pvormste/yet-web-utils/tree/master/yetconfig) | Helpers for loading config from environment | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetconfig) |
| [yetenv](https://github.com/pvormste/yet-web-utils/tree/master/yetenv) | Provides some logic to determine the environment (develop, staging, production) | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetenv) |
| [yethttp](https://github.com/pvormste/yet-web-utils/tree/master/yethttp) | Provides helpers for working with http | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yethttp) |
| [yetlog](https://github.com/pvormste/yet-web-utils/tree/master/yetlog) | Provides a logger agnostic logging interface which is used by most `yet` packages | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetlog) |
| [yetnet](https://github.com/pvormste/yet-web-utils/tree/master/yetnet) | Provides helper code for working with low level network like TCP, UDP and Ports | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetnet) |
 
