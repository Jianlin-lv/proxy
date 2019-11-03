#
# Cilium incremental build. Should be fast given builder-deps is up-to-date!
#
# cilium-builder tag is the Git SHA of the compatible build image commit.
# Keeping the old images available will allow older versions to be built
# while allowing the new versions to make changes that are not backwards compatible.
#
FROM quay.io/cilium/cilium-envoy-builder:96eb614d0fa4012abc9850ca58f33c28cbbd8f61 as builder
LABEL maintainer="maintainer@cilium.io"
WORKDIR /go/src/github.com/cilium/cilium/envoy
COPY . ./
ARG V
#
# Please do not add any dependency updates before the 'make install' here,
# as that will mess with caching for incremental builds!
#
RUN make PKG_BUILD=1 V=$V DESTDIR=/tmp/install cilium-envoy install

#
# Extract installed cilium-envoy binaries to an otherwise empty image
#
FROM scratch
LABEL maintainer="maintainer@cilium.io"
COPY --from=builder /tmp/install /
