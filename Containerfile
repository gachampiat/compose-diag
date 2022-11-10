FROM golang:latest AS builder

WORKDIR /src

ENV CGO_ENABLED=0

COPY go.mod /src/

RUN go mod download

COPY . .

RUN  go build -a -o compose-diag


FROM alpine:edge

RUN apk add --no-cache -X http://dl-cdn.alpinelinux.org/alpine/edge/testing py3-nwdiag

RUN mkdir -p /app \
    && adduser -D compose \
    && chown -R compose:compose /app

USER compose

WORKDIR /app

COPY --from=builder /src/compose-diag .
COPY assets /app/assets

ENTRYPOINT [ "./compose-diag" ] 

