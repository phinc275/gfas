FROM golang:1.19-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o gfas cmd/gfas/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/gfas ./
CMD ["/app/gfas"]