# Gunakan golang:1.16 sebagai base image
FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Menjalankan migrasi database saat container dimulai
CMD ["./db.migrate", "reset"]
CMD ["./db.migrate", "up"]

# Expose port yang digunakan oleh aplikasi
EXPOSE 8080

# Menjalankan aplikasi Go API setelah migrasi selesai
CMD ["./main", "api"]
