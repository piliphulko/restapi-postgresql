# syntax=docker/dockerfile:1

FROM vault:1.13.3

RUN apk update && apk add --no-cache ca-certificates

COPY config/certificate.crt /usr/local/share/ca-certificates/certificate.crt

RUN update-ca-certificates

COPY vault/entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]