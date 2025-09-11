FROM container-registry.oracle.com/os/oraclelinux:8

WORKDIR $GOPATH/src/github.com/thanos-io/thanos

RUN dnf update -y && \
    dnf install make which git -y && \
    dnf install golang -y && \
    dnf clean all && \
    rm -rf /var/cache/yum/*

COPY . $GOPATH/src/github.com/thanos-io/thanos

RUN make build

# -----------------------------------------------------------------------------

FROM container-registry.oracle.com/os/oraclelinux:8-slim

RUN microdnf update -y && \
    microdnf clean all && \
    rm -rf /var/cache/yum/*

COPY LICENSE SECURITY.md THIRD_PARTY_LICENSES.txt /licenses/
COPY olm/SECURITY.md /licenses/ORACLE_SECURITY.md

COPY --from=builder /root/go/bin/thanos /bin/thanos

USER nobody

ENTRYPOINT [ "/bin/thanos" ]
