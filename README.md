
# Intro

This package contains a WIP runner for [Devfiles](https://docs.devfile.io/).

It takes as input a Devfile, converts it to a Compose project and executes it using a Docker daemon.
The binary can be used as a Docker CLI plugin.

## Building

Requires at least go `1.16` because of transitive dependencies (e.g. Docker libraries)

```sh
go build -ldflags="-s -w" -o docker-devenv main.go
```

## Running

```sh
./devrunner --help
```

## Testing

```sh
go test ./...
```

### Installing as a Docker CLI plugin

```sh
mkdir -p ~/.docker/cli-plugins
cp docker-devenv ~/.docker/cli-plugins
```

## Roadmap

* implement all other devfile spec
  * ✅ Docker image
  * ❌ Dockerfile
  * ❌ memory/cpu limits
  * ✅ volumes
  * ✅ environment variable
  * ✅ mounting source code
  * ❌ adding IDE runtime components
  * ❌ command execution
  * ❌ devfile events
* support other definition formats (e.g. convert to Devfile):
  * ❌ `gitpod.yml`
  * ❌ `devcontainer.json`
* infer suitable definition based on the type of source code (e.g. node if it has `package.json`)
  * ✅ Java Maven
  * ✅ Go
  * ✅ Python
* implement lifecycle operations:
  * ✅ start
  * ✅ stop
  * ✅ update
  * ✅ command execution
  * ✅ describe
  * ❌ kubernetes operations
