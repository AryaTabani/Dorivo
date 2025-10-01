FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /dorivo-app ./main.go

FROM alpine:latest

COPY --from=builder /dorivo-app /dorivo-app

EXPOSE 8080

ENTRYPOINT ["/dorivo-app"]