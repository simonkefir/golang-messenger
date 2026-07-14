FROM golang:1.26-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/exe ./cmd/messenger

FROM alpine:3.24.1
RUN apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/exe .
CMD ["./exe"]