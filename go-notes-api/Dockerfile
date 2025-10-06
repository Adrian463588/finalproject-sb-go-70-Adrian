# Gunakan base image Go yang sesuai dengan versi di go.mod
FROM golang:1.25-alpine

# Tentukan direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu
COPY go.mod go.sum ./

# Download semua dependencies
RUN go mod download

# Salin semua source code ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Tentukan port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
