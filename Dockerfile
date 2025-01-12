FROM golang:1.23.3-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

FROM builder AS test
RUN go test -v ./...

FROM builder AS build
RUN go build -o /go/bin/app ./cmd/main.go

FROM alpine:edge
WORKDIR /usr/src/app
COPY --from=build /go/bin/app /usr/local/bin/app

CMD ["app"]
