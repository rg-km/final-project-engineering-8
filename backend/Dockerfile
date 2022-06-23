FROM golang:1.17-alpine


RUN set -ex && \
    apk add --no-cache gcc musl-dev

RUN set -ex && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/cc1obj && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto1 && \
    rm -f /usr/libexec/gcc/x86_64-alpine-linux-musl/6.4.0/lto-wrapper && \
    rm -f /usr/bin/x86_64-alpine-linux-musl-gcj


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . ./

RUN ls

EXPOSE 5000

ENV PORT=5000

RUN echo "===> Port:" $PORT

RUN go build -o halloguru-api

CMD ./halloguru-api
