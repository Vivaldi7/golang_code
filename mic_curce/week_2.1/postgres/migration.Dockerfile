FROM alpine:3.13

RUN apk update &&\
    apk upgrade &&\
    apk add bash &&\
    rm -rf /var/cache/apk/*

COPY goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

#ADD http://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
#RUN chmod +x /bin/goose

#ADD http://github.com/pressly/goose/releases/download/v3.14.0/goose_windows_x86_64.exe /bin/goose.exe
#RUN chmod +x /bin/goose.exe

WORKDIR /root

RUN mkdir -p migrations
COPY migrations ./migrations
RUN ls -la migrations/
ADD migration.sh .
ADD .env .

RUN chmod +x migration.sh

WORKDIR /root

ENTRYPOINT [ "bash", "migration.sh" ]