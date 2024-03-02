FROM busybox:stable AS puller

ARG PSQLDEF_DOWNLOAD_TARGET=https://github.com/sqldef/sqldef/releases/download/v0.16.15/psqldef_linux_amd64.tar.gz
RUN wget ${PSQLDEF_DOWNLOAD_TARGET} -O - | tar xvzf -

# FROM gcr.io/distroless/base-debian11:latest AS runner
FROM alpine:3.19.1 AS runner

COPY --from=puller /psqldef /usr/local/bin/psqldef
