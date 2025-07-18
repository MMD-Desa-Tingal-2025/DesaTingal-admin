
---

## Health Check dan Logging Endpoint

Gunakan perintah berikut untuk melakukan pengecekan status aplikasi dan melihat log aktivitas terkait database.

### Basic API Health Check

Memeriksa apakah layanan API berjalan dengan normal:

```bash
curl http://localhost:8080/api/v1/health
```

### Detail Log Database

Menampilkan informasi log aplikasi yang berhubungan dengan database:

```bash
curl http://localhost:8080/api/v1/logs
```

### Web Health Check

Melakukan pengecekan status aplikasi melalui halaman web:

```bash
curl http://localhost:8080/web/health
```

---

## Contoh Log Aplikasi

Contoh log yang muncul saat aplikasi berhasil dijalankan:

```
2024/01/20 10:30:00 Successfully connected to PostgreSQL database
2024/01/20 10:30:00 Table application_logs created/verified successfully
2024/01/20 10:30:00 Database migration completed successfully
```

### Contoh Log Error

Contoh log jika terjadi kegagalan koneksi ke database:

```
2024/01/20 10:30:00 Failed to connect to database: dial tcp 127.0.0.1:5432: connect: connection refused
```

---