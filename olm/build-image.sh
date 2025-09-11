#!/usr/bin/env bash
export GOPATH=$(go env GOPATH)

name="thanos"
version="0.39.2"
registry="container-registry.oracle.com/olcne"
docker_tag=${registry}/${name}:v${version}

docker build --pull \
    --build-arg https_proxy=${https_proxy} \
    -t ${docker_tag} -f ./olm/builds/Dockerfile .
