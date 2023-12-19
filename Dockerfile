
FROM golang:1.20

WORKDIR /proxypod

COPY . ./

# Build Go model
RUN CGO_ENABLED=0 go build

FROM alpine:3.18
RUN mkdir -p /usr/local/bin
COPY --from=0  /proxypod/proxypod /usr/local/bin/proxypod

# Install bash
RUN apk add bash
