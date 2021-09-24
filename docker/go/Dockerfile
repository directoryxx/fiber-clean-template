FROM golang:1.17-buster 

RUN apt update && apt install git \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    make

RUN mkdir /app

RUN go get github.com/gofiber/fiber/v2
# RUN go get github.com/oxequa/realize


WORKDIR /app
