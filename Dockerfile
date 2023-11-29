FROM golang:1.19-alpine as builder
RUN apk add build-base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -tags musl -o gfas cmd/gfas/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gfas ./
CMD ["/app/gfas"]