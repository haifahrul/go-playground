#!/bin/bash

APP_NAME='go-playground'

go mod vendor

CompileDaemon \
    -exclude-dir="vendor" \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -mod=vendor -o $APP_NAME ./cmd/$APP_NAME/..." \
    -command="./$APP_NAME"
