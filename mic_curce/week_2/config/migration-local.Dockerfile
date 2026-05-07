FROM alpine:3.13

RUN apk update &&\
    apk upgrade &&\
    apk add bash &&\
    rm -rf /var/cache/apk/*

COPY C:/Users/836D~1/GOLang/mic_curce/week_3/config/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose


WORKDIR /root

ADD migrations ./migrations
ADD migration-local.sh .
ADD local.env .

RUN chmod +x migration.sh


ENTRYPOINT [ "bash", "migration.sh" ]