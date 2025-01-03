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

# Docker Caching

Caching is a critical feature in Docker that helps improve build performance and efficiency.
Docker provides multiple types of caching mechanisms, each suited for different use cases:

## Types of Docker Caching

1. Docker Layer Caching

   Docker uses a layer-based architecture where each instruction in the Dockerfile creates a layer.
   Layers are cached and reused during subsequent builds if they haven't changed.
   This type of caching helps speed up builds by skipping unchanged layers.

[//]: # (Todo Add Docerfile code example)

2. Docker Cache Mounts

   Docker cache mounts provide a flexible way to optimize builds by specifying a persistent cache location that can be used across builds.
   Unlike regular Docker layer caching, which depends on an exact match of the instruction and its dependencies, cache mounts allow for cumulative caching. This means even if a layer is rebuilt, **only new or changed data is processed**, while unchanged data is reused from the cache.

Comparison: Regular Layer Caching vs. Cache Mounts

| Aspect       | Layer Caching                                                  | Cache Mounts                                                     |
|--------------|----------------------------------------------------------------|------------------------------------------------------------------|
| Scope        | Caches the entire layer based on instruction and dependencies. | Caches specific directories or files persistently.               |
| Invalidation | Cache invalidates if files or instructions change.             | Cache remains cumulative; only updates new or changed data.      |
| Use Case     | General caching for Dockerfile instructions.                   | Targeted caching for tools like package managers or build tools. |
| Example      | `COPY` or `RUN` commands in Dockerfile                         | `--mount=type=cache` in `RUN` commands.                          |

# References

- https://docs.docker.com/build/cache/optimize/

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