FROM golang:1.17-buster

WORKDIR /go/src/github.com/directoryxx/go-clean-architecture-web-application-boilerplate/app

RUN apt update && apt install git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    make
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/urfave/cli/v2
RUN GO111MODULE=on go get github.com/oxequa/realize
RUN go get entgo.io/ent/cmd/ent


CMD realize start