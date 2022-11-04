FROM golang:1.17.9-alpine3.14

RUN apk update && apk upgrade
# RUN apk add --no-cache --virtual .build-deps --no-progress -q \
#     bash \
#     curl \
#     busybox-extras \
#     make \
#     git \
#     tzdata && \
    # cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# RUN apk update && apk add --no-cache coreutils

WORKDIR /src

RUN ls -ls

RUN mkdir -p /src/golangApp
COPY . /src/golangApp
WORKDIR /src/golangApp

RUN go mod tidy -compat=1.17

RUN go build

EXPOSE 9000

# CMD "./golangApp"