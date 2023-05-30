# Source code + build dependencies base image
FROM golang:1.20-alpine AS source

WORKDIR /app

RUN go env -w CGO_ENABLED="0"

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

# Dev image with live reloading + debugger
FROM source AS dev

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

# Test image
FROM source AS test

RUN go test ./...

## Build image
FROM source AS run

RUN go build -ldflags="-s -w" -o /bin/app ./cmd/app/
CMD ["/bin/app"]
