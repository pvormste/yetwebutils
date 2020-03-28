[![GitHub license](https://img.shields.io/github/license/pvormste/yetwebutils)](https://github.com/pvormste/yetwebutils/blob/master/LICENSE) ![](https://github.com/pvormste/yetwebutils/workflows/lint/badge.svg?branch=master) ![](https://github.com/pvormste/yetwebutils/workflows/tests/badge.svg?branch=master)

# Yet 'Another' Web Utils (yetwebutils)

`yetwebutils` is a highly opinionated collection of go packages, which reduces the amount of boilerplate code for simple web projects.

But essentially it acts like another standard lib for creating web applications.

_This project is more like an experiment on how much web functionality can be abstracted in own reusable packages and I'm very curious how this will evolve. The main goal is to reduce boilerplate code for greenfield web projects. Nervertheless feel free to use it for your own needs!_

## Install

```bash
go get -u github.com/pvormste/yetwebutils
```

## Packages

All packages are prefixed with `yet`.

| package | Description | GoDoc |
| ------- | ----------- | ----- |
| [yethttp](https://github.com/pvormste/yet-web-utils/tree/master/yethttp) | Provides helpers for working with http | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yethttp) |
| [yetnet](https://github.com/pvormste/yet-web-utils/tree/master/yetnet) | Provides helper code for working with low level network like TCP, UDP and Ports | [docs](https://godoc.org/github.com/pvormste/yetwebutils/yetnet) |
 
