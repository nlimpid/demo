#!/usr/bin/env bash

set -xe

go vet ./...

GOOS=linux CGO_ENABLED=0 go build -o app

docker build -t docker.pkg.github.com/nlimpid/demo/demo:git.$1 -f Dockerfile .
docker push docker.pkg.github.com/nlimpid/demo/demo:git.$1
