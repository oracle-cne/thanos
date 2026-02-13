#!/usr/bin/env bash

mkdir -p bin
version="0.41.0

git config --global --add safe.directory $(pwd)

GIT_REVISION=$(git rev-parse HEAD)
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
ldflags="
        -X main.version=v${version}
        -X github.com/prometheus/common/version.Version=${version}
        -X github.com/prometheus/common/version.Revision=${GIT_REVISION}
        -X github.com/prometheus/common/version.Branch=HEAD
        -X github.com/prometheus/common/version.BuildUser=${USER}@${HOST}
        -X github.com/prometheus/common/version.BuildDate=${BUILD_DATE}"

go build -a -tags netgo -trimpath=false -v -o ./bin/thanos \
	-ldflags "${ldflags}" ./cmd/thanos
