FROM alpine:latest as cert
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

FROM scratch
COPY --from=cert /etc/ssl /etc/ssl
COPY yamllint-checktyle /
ENTRYPOINT ["/yamllint-checkstyle"]