FROM golang

ENV GO111MODULE on
ENV CGO_ENABLED 1
ENV GOOS linux
ENV GOARCH amd64

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy

WORKDIR src/main

RUN go build -o ../../bin/

WORKDIR /usr/src/app/bin

CMD ["sudo", "chmod", "+x", "main"]

CMD ["./main"]