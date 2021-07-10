FROM golang:1.13-alpine
ARG HAIFAHRUL_UID=1000
ARG HAIFAHRUL_GID=1000
ENV APP_NAME go-playground

COPY build/assets/ghz /usr/local/bin
COPY build/assets/repositories /etc/apk/repositories

RUN chmod 644 /etc/apk/repositories && \
    apk update && apk add --no-cache sudo git
RUN apk add --no-cache make tzdata curl && \
    chmod a+x /usr/local/bin/ghz

COPY build/assets/sudoers /etc/sudoers

RUN chmod 0400 /etc/sudoers && \
    addgroup haifahrul -g ${HAIFAHRUL_GID} && \
    addgroup sudo && \
    adduser -h /home/haifahrul -s /bin/sh -u ${HAIFAHRUL_UID} -G users -D haifahrul && \
    adduser haifahrul haifahrul && \
    adduser haifahrul sudo

USER haifahrul
WORKDIR "/go/src/github.com/haifahrul/${APP_NAME}"

COPY --chown=haifahrul:users .netrc /tmp/.netrc

COPY --chown=haifahrul:users build/dev/evans_linux_amd64.tar.gz /tmp

RUN sudo chown -R haifahrul:haifahrul /go && \
    cp /tmp/.netrc ~/.netrc && \
    chmod 600 ~/.netrc && \
    export GO111MODULE=off && \
    echo "Installing additional tools, please wait..." && \
    go get -u golang.org/x/lint/golint && echo "- Success installing golint" && \
    go get github.com/githubnemo/CompileDaemon && echo "- Success installing CompileDaemon" && \
    unset GO111MODULE


ENTRYPOINT ["sh", "./build/dev/entrypoint.sh"]
