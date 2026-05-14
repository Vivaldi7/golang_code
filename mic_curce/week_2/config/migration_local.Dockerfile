FROM alpine:3.13

RUN apk update &&\
    apk upgrade &&\
    apk add bash &&\
    rm -rf /var/cache/apk/*

#COPY C:/Users/836D~1/GOLang/mic_curce/week_3/postgres/goose_linux_x86_64 /bin/goose
COPY goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose


WORKDIR /root

RUN mkdir -p migrations
COPY migrations ./migrations
RUN ls -la migrations/
ADD migration_local.sh .
ADD local.env .

RUN chmod +x migration_local.sh


ENTRYPOINT [ "bash", "migration_local.sh" ]