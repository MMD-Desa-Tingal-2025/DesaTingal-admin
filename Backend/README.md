
---

# Health Check dan Logging Endpoint

Dokumen ini menjelaskan cara melakukan pengecekan status layanan aplikasi serta melihat log aktivitas yang berkaitan dengan database.

## API Health Check

Gunakan perintah berikut untuk memeriksa apakah layanan API berjalan dengan normal:

```bash
curl http://localhost:8080/api/v1/health
```

## Log Aktivitas Database

Untuk menampilkan informasi log aplikasi yang berhubungan dengan database, gunakan perintah berikut:

```bash
curl http://localhost:8080/api/v1/logs
```

## Web Health Check

Untuk melakukan pengecekan status aplikasi melalui halaman web, gunakan perintah berikut:

```bash
curl http://localhost:8080/web/health
```

## Menjalankan Aplikasi

Gunakan perintah berikut untuk menjalankan aplikasi:

```bash
go run cdm/server/main.go
```

---

## Contoh Log Aplikasi

Berikut adalah contoh log yang muncul ketika aplikasi berhasil dijalankan:

```
2024/01/20 10:30:00 Successfully connected to PostgreSQL database
2024/01/20 10:30:00 Table application_logs created/verified successfully
2024/01/20 10:30:00 Database migration completed successfully
```

## Contoh Log Error

Contoh log berikut menunjukkan kegagalan saat mencoba terhubung ke database:

```
2024/01/20 10:30:00 Failed to connect to database: dial tcp 127.0.0.1:5432: connect: connection refused
```

---