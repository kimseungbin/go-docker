# Overview

This is an example project to set up a local development environment for api server of Go and Redis using Docker.
The language and the database are largely irrelevant to the purpose of this project.

## Specifications

- Language: Go 1.19.2
- Database: Redis 7.x
- Container: Docker and docker-compose
- Web Framework: [Echo](https://github.com/labstack/echo)

# Install

*This installation guide is exclusively for Mac.
To install in other OSs, install what is equivalent of requirements in the target OS.*

## 1. Install Docker

- Docker: `$ brew install docker`
- docker-compose: `$ brew install docker-compose`

**OR**

*Preferred*
- `$ brew cask install docker`

Above command installs not only Docker itself but also Docker Desktop on Mac, docker-compose and docker-machine too.

# Set language-specific settings

## Go

- with Go installed on local machine

```shell
go mode init github.com/[github_name]/[repo_name]
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
```

- fill out your own `github_name` and `repo_name`
- using repo URI for `init`'s argument is just a convention: it can be any plain text


## JavaScript (Node.js runtime)

- with Node.js and NPM installed on local machine

```shell
npm init
```


# Run

`docker-compose -f docker-compose.yml up`