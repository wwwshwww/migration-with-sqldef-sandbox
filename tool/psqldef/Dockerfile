FROM alpine:3.19.1 AS puller

ENV PSQLDEF_DOWNLOAD_TARGET=https://github.com/sqldef/sqldef/releases/download/v0.16.15/psqldef_linux_amd64.tar.gz

RUN apk add --no-cache wget tar
RUN wget ${PSQLDEF_DOWNLOAD_TARGET} -O - | tar xvzf -

FROM alpine:3.19.1 AS runner

COPY --from=puller /psqldef /usr/local/bin/psqldef
