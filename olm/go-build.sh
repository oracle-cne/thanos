#!/usr/bin/env bash

set -euo pipefail

log() {
	echo "go-build.sh: $*"
}

mkdir -p bin
version="${THANOS_VERSION:-${VERSION:-${1:-}}}"
build_host="${HOST:-$(hostname)}"
build_user="${USER:-$(id -un)}@${build_host}"
go_source="${GOPATH_SRC:-$(pwd)}"

if [[ -z "${version}" ]]; then
	log "unable to determine Thanos version; set THANOS_VERSION or pass it as the first argument" >&2
	exit 1
fi

log "starting thanos build"

GIT_REVISION=$(git rev-parse HEAD)
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
ldflags="
        -X github.com/prometheus/common/version.Version=${version}
        -X github.com/prometheus/common/version.Revision=${GIT_REVISION}
        -X github.com/prometheus/common/version.Branch=HEAD
        -X github.com/prometheus/common/version.BuildUser=${build_user}
        -X github.com/prometheus/common/version.BuildDate=${BUILD_DATE}"

log "git_revision=${GIT_REVISION}"
log "build_date=${BUILD_DATE}"
log "build_user=${build_user}"
log "go_version=$(go version)"
log "thanos_version=${version}"
log "compiling Thanos binary"
go build -a -tags netgo,slicelabels -trimpath=false -v -o ./bin/thanos \
	-ldflags "${ldflags}" \
	"${go_source}"/cmd/thanos

log "verifying thanos binary"
./bin/thanos --version
log "completed thanos build"
