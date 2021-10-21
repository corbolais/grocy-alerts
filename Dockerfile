##
# Build golang application
##

ARG VERSION="development"
FROM golang:alpine AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download && \
  apk add --no-cache gcc musl-dev
COPY . .
RUN go build -ldflags="-X 'main.version=${VERSION}'" -o grocy-alerts .
WORKDIR /dist
RUN cp /build/grocy-alerts .

##
# Build a small image
##

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY --from=builder /dist/grocy-alerts /

ENTRYPOINT ["/grocy-alerts"]