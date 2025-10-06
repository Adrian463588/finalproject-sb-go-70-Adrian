# Stage 1: build the Go binary
FROM golang:1.25-alpine AS builder

# Install git (jika modul menggunakan git) dan ca-certificates
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Salin go.mod / go.sum dahulu agar caching layer dependency optimal
COPY go.mod go.sum ./
RUN go mod download

# Salin source code
COPY . .

# Build binary statically (optional flags bisa ditambahkan)
RUN CGO_ENABLED=0 GOOS=linux go build -o /notes-app

# Stage 2: final minimal image
FROM alpine:latest

# Tambahkan sertifikat SSL / CA supaya HTTPs request bisa bekerja
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Salin binary dari stage builder
COPY --from=builder /notes-app .

# Jika kamu punya file konfigurasi atau migration scripts, bisa di-copy juga
# COPY --from=builder /app/migrations ./migrations

# Expose port aplikasi kamu
EXPOSE 8080

# Jalankan binary
CMD [ "./notes-app" ]
