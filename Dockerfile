FROM golang:1.12-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/src/7factor.io/
ENV GO111MODULE=on

# install deps
COPY ./src/7factor.io/go.* ./
RUN go mod download

# Copy src
COPY ./src/7factor.io ./

# build binary and install
RUN go install ./...

FROM bash:5.0.2

COPY --from=builder /go/bin/cmd /go/bin/cmd

# make binary default executable
CMD ["/go/bin/cmd"]
