FROM container-registry.oracle.com/os/oraclelinux:9 AS builder

ARG GOLANG_VERSION
ARG THANOS_VERSION

RUN dnf -y install --enablerepo=ol9_codeready_builder git golang-${GOLANG_VERSION} hostname make && \
    dnf clean all && \
    go version

# Copy the contents of the repository
COPY . /workspace/go/src/github.com/thanos-io/thanos

ENV GOPATH=/workspace/go CGO_ENABLED=0 GOTOOLCHAIN=local
ENV GOPATH_SRC=$GOPATH/src/github.com/thanos-io/thanos
WORKDIR /workspace/go/src/github.com/thanos-io/thanos

RUN git config --global --add safe.directory $(pwd) && \
    bash olm/go-build.sh "${THANOS_VERSION}"

# -----------------------------------------------------------------------------

FROM container-registry.oracle.com/os/oraclelinux:9-slim

COPY LICENSE SECURITY.md THIRD_PARTY_LICENSES.txt /usr/share/licenses/thanos/
COPY olm/SECURITY.md /usr/share/licenses/thanos/ORACLE_SECURITY.md
COPY --from=builder /workspace/go/src/github.com/thanos-io/thanos/bin/thanos /usr/bin/thanos

RUN microdnf update -y && \
    microdnf clean all && \
    rm -rf /var/cache/yum/* && \
    test -d /LICENSES || ln -s /usr/share/licenses /LICENSES

ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

USER nobody

ENTRYPOINT [ "/usr/bin/thanos" ]
