FROM alpine:3.21

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD migration_prod.sh .
ADD prod.env .

RUN chmod +x migration_prod.sh

ENTRYPOINT ["bash", "migration_prod.sh"]