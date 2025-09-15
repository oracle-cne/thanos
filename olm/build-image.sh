#!/usr/bin/env bash -x
export GOPATH=$(go env GOPATH)

name="thanos"
version="{{{ .major }}}.{{{ .minor }}}.{{{ .patch }}}"
registry="container-registry.oracle.com/olcne"
docker_tag=${registry}/${name}:v${version}

pwd
ls -lrt
mkdir -p pkg/basicauth
cp olm/options.go pkg/basicauth/
patch -p0 < olm/basic_auth.patch
patch -p0 < olm/go_extra_flags.patch

docker build --pull \
    --build-arg https_proxy=${https_proxy} \
    -t ${docker_tag} -f ./olm/builds/Dockerfile .
