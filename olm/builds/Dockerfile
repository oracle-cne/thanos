FROM container-registry.oracle.com/os/oraclelinux:8 as builder

WORKDIR $GOPATH/src/github.com/thanos-io/thanos

RUN dnf update -y && \
    dnf install hostname make which git -y && \
    dnf install golang -y && \
    dnf clean all && \
    rm -rf /var/cache/yum/*

COPY . $GOPATH/src/github.com/thanos-io/thanos

RUN sh olm/go-build.sh

# -----------------------------------------------------------------------------

FROM container-registry.oracle.com/os/oraclelinux:8-slim

RUN microdnf update -y && \
    microdnf clean all && \
    rm -rf /var/cache/yum/*

COPY LICENSE SECURITY.md THIRD_PARTY_LICENSES.txt /licenses/
COPY olm/SECURITY.md /licenses/ORACLE_SECURITY.md

COPY --from=builder $GOPATH/src/github.com/thanos-io/thanos/bin/thanos /bin/thanos

USER nobody

ENTRYPOINT [ "/bin/thanos" ]
