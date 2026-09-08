FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO_ENABLED=0: el binario de producción solo usa el driver de Postgres
# (gorm.io/driver/postgres), que es pure Go. El driver de SQLite (que sí
# necesita cgo) solo se usa en los tests de integración, nunca en cmd/api.
RUN CGO_ENABLED=0 GOOS=linux go build -o /hospital-api ./cmd/api

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=builder /hospital-api /hospital-api
EXPOSE 8080
ENTRYPOINT ["/hospital-api"]
