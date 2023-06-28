FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ./server/go.mod .
RUN go mod download

COPY . .

RUN go build -o node ./server/main.go

FROM alpine:latest

RUN apk add --update docker openrc

RUN rc-update add docker boot

RUN apk --no-cache add ca-certificates

ENV PORT=8080

WORKDIR /app

COPY --from=builder /app/node /app/node

ENTRYPOINT ["/app/node"]