FROM golang:1.24 AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/app

FROM cgr.dev/chainguard/git:latest@sha256:f06036ea97419c784339638cf4a21b52d39df8dfd8aa0e0e73307fc1082c6043
WORKDIR /
COPY --from=builder /go/bin/app /app

ENTRYPOINT [ "/app" ]
