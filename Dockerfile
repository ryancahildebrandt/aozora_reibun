FROM golang:1.24-bookworm AS base

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o aozora_reibun

CMD ["/build/aozora_reibun"]
