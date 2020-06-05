# latexd

Dockerized tooling for latex documents.

## Requirements

- [Docker](https://www.docker.com/products/docker-desktop)
- [Golang](https://golang.org/dl/)

## Docker

The `latexd` image is an ubuntu based docker container with all the tooling necesary to interact with latex.

To build the image:

```sh
make build
```

To open `latexd` shell:

```sh
make shell
```

## CLI

An user friendly cli to interact with latex documents using the `latexd` docker image.

To build the cli:

```sh
make cli
```

Run the cli:

```sh
./latedx -h
```