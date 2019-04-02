# txt2web

## About

A small and simple http file server that serves textfiles.

## USAGE

1. Build a Docker image based on `andimenge/txt2web` with your text file like this:

``` Dockerfile
FROM andimenge/txt2web:latest
COPY my-text-file.txt /textfile
```

2. Run your container. Default container port is `8080`.

## Links

[Docker Hub](https://cloud.docker.com/repository/docker/andimenge/txt2web)