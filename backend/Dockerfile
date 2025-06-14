# --- BUILD STAGE
ARG GO_VERSION=1.24

FROM golang:${GO_VERSION}-bookworm AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o /run-app .

# --- DEVELOPMENT STAGE
FROM golang:${GO_VERSION}-alpine AS dev

RUN apk add --no-cache git curl && \
    go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]

# --- DEPLOY STAGE
FROM debian:bookworm AS final

RUN useradd --no-create-home --shell /bin/false appuser

COPY --from=builder /run-app /usr/local/bin/run-app

USER appuser

CMD ["run-app"]
