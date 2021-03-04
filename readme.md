# kubodin

[![Release](https://img.shields.io/github/release/khorevaa/kubodin.svg?style=for-the-badge)](https://github.com/khorevaa/kubodin/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build status](https://img.shields.io/github/workflow/status/khorevaa/kubodin/goreleaser?style=for-the-badge)](https://github.com/khorevaa/kubodin/actions?workflow=goreleaser)
[![Codecov branch](https://img.shields.io/codecov/c/github/khorevaa/kubodin/master.svg?style=for-the-badge)](https://codecov.io/gh/khorevaa/kubodin)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/khorevaa/kubodin)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=for-the-badge)](https://saythanks.io/to/khorevaa)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## Features

* **Not need installed** `rac.exe` 
* Very **small & fast** app design for kubernetes and docker usage
* Full API documentation

## Limitings 

* Only for Kubernetes usage
* Only one 1C.Enterprise server (add on startup), with name `default`
* Only some allowed API
  * Infobases (`api/v1/app/default/infobases`)
    * Create infobase
    * Delete infobase
    * List short info of infobases
  * Cluster (`api/v1/app/default/clusters`)
    * List cluster
    * Cluster info
  * App (`api/v1/app/default`)
    * Get info
    * List
* Only in memory cache
* Only in memory database

## Installation


## How to use

### Docker
```shell

docker run -p 3001:3001 ghcr.io/khorevaa/kubodin:latest \
      --port :3001 --server 1cserver:1545

```
### Github Release

## License