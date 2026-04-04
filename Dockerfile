FROM golang:1.23-alpine AS builder

WORKDIR /app

ENV GOTOOLCHAIN=auto

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]