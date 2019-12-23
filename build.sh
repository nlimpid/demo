#!/usr/bin/env bash

set -xe

go vet ./...

GOOS=linux CGO_ENABLED=0 go build -o app

docker build -t hub.bunny-tech.com/prod/test:git.$1 -f Dockerfile .
docker push hub.bunny-tech.com/prod/test:git.$1
