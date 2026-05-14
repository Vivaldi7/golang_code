FROM alpine:3.13

RUN apk update &&\
    apk upgrade &&\
    apk add bash &&\
    rm -rf /var/cache/apk/*

#COPY goose_linux_x86_64 /bin/goose

COPY goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose


WORKDIR /root

RUN mkdir -p migrations
COPY migrations ./migrations
RUN ls -la migrations/
ADD migration_prod.sh .
ADD prod.env .

RUN chmod +x migration_prod.sh


ENTRYPOINT [ "bash", "migration_prod.sh" ]