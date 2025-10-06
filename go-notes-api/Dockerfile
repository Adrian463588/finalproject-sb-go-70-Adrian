# Gunakan base image Go
# Memulai dengan "dapur" bersih yang sudah terinstal Go versi 1.21
FROM golang:1.21-alpine

# Tentukan direktori kerja
# Di dalam dapur itu, kita akan bekerja di dalam sebuah nampan bernama /app
WORKDIR /app

# Salin file go.mod dan go.sum
# Bawa "daftar belanjaan" (dependencies) kita ke nampan terlebih dahulu
COPY go.mod go.sum ./

# Download semua dependencies
# Suruh Go untuk "membeli" semua bahan sesuai daftar belanjaan
RUN go mod download

# Salin semua source code
# Sekarang, bawa semua "buku resep" (kode .go kita) ke dalam nampan
COPY . .

# Build aplikasi
# "Masak" semua bahan dan resep menjadi satu kue jadi bernama "main"
RUN go build -o main .

# Expose port
# Beri label pada kotak kue: "Sajikan melalui lubang nomor 8080"
EXPOSE 8080

# Perintah untuk menjalankan aplikasi
# Tulis instruksi terakhir di kotak: "Untuk memulai, jalankan 'main'"
CMD ["./main"]