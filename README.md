# Go base

A basic project structure with configuration using environment variables and a convenient http client package to reduce boilerplate.

### Features
- Define configuration variables in `config/config.go`
- Use the HTTP client with simple calls to `http.GET` and `http.POST`

## Installation
- Clone this repository in a location from your GOPATH
- Replace all occurrences of `Vinelab/go-base` with the name your `vendor/project`

## Vendoring Dependencies
Using Docker image `vinelab/go-vndr`:
1. Run `docker run -it -v /path/to/project:/go/src/github.com/[vendor]/[project] vinelab/go-vndr bash` to run a container with the code mounted and a new SSH session established
    > 1. Changes in code on the local file system will be directly available in the container (mounted)
    > 2. Shutting down the container will not cause any harm to the code, except the vendor packages installed using `vndr` will need to be re-installed after running another container (see step n.2)
2. Run (within the container) `vndr`
