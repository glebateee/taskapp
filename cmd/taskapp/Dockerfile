FROM golang:1.26.3-bookworm AS builder
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/taskapp/exe /app/cmd/taskapp

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/taskapp/exe /app
CMD ["/app/exe"]
