FROM golang:1.26.0-alpine AS builder

WORKDIR /build

COPY go.mod go.sum

#RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /src/main ./cmd/service/

FROM alpine:latest

COPY --from=builder /src/main .

EXPOSE 8000

CMD ["./main"]
