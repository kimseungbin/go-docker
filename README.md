# Overview

This project demonstrates efficient Docker workflows, focusing on techniques to optimize container images and build
processes. Key highlights include:

1. Creating Small and Efficient Docker Images:
    - Use multi-stage builds to minimize image size.
    - Leverage image and dependency caching to speed up the build and reduce redundancy.
2. Demonstrating Caching with GitHub Workflows:
    - Utilize GitHub Actions to automate the build process
    - Compare the effectiveness of caching by running parallel builds:
        - One with caching enabled
        - Another without caching for comparison

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