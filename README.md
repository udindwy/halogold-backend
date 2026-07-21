# HaloGold Backend API

HaloGold adalah backend API untuk platform investasi emas digital yang memungkinkan pengguna melihat harga, membeli, dan menjual emas secara digital. Proyek ini dibangun mengutamakan *Clean Architecture* (Handler -> Service -> Repository -> Database) dengan *error handling* dan validasi request yang ketat.

---

## Tech Stack

- **Language:** Go (1.24+)
- **Framework:** Gin Web Framework
- **Database:** PostgreSQL
- **ORM:** GORM (dengan fitur AutoMigrate)
- **Validation:** go-playground/validator/v10
- **Configuration:** godotenv

---

## Installation

Pastikan [Golang](https://go.dev/doc/install) dan [PostgreSQL](https://www.postgresql.org/download/) telah terinstall di komputer Anda.

1. Clone repositori ini
   ```bash
   git clone https://github.com/udindwy/halogold-backend.git
   cd halogold-backend
   ```
2. Unduh semua dependensi Go
   ```bash
   go mod tidy
   ```

---

## Environment Configuration

Salin file `.env.example` (jika ada) menjadi `.env`, atau buat file `.env` di root direktori proyek Anda dengan konfigurasi berikut:

```env
APP_NAME=HaloGold API
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=halogold
DB_SSLMODE=disable
```

---

## Create Database

Sebelum menjalankan aplikasi, Anda harus membuat database kosong di dalam PostgreSQL Anda.

Anda bisa menggunakan psql CLI atau GUI (seperti pgAdmin / DBeaver):
```sql
CREATE DATABASE halogold;
```
*(Catatan: Aplikasi menggunakan GORM AutoMigrate, sehingga tabel `users` dan `transactions` akan otomatis di-generate ketika server dinyalakan pertama kali).*

---

## Run Application

Jalankan entry point utama aplikasi:

```bash
go run ./cmd/api
```

Server akan mulai berjalan di port HTTP `8080` (sesuai konfigurasi `.env`).

---

## Folder Structure

```text
halogold-backend/
├── cmd/
│   └── api/
│       └── main.go                  # Main entry point & Dependency Injection
├── config/
│   ├── config.go                    # Struct konfigurasi Env
│   └── database.go                  # Koneksi DB, Connection Pool & AutoMigrate
├── internal/
│   ├── dto/                         # Data Transfer Objects (Req/Res formats)
│   ├── handler/                     # HTTP Handlers (Gin)
│   ├── model/                       # GORM Entities
│   ├── repository/                  # Database Operations
│   ├── service/                     # Business Logic (Kalkulasi Gram/Harga)
│   └── validator/                   # Gin validation error formatter
├── pkg/
│   ├── logger/                      # Standard logger wrapper
│   └── response/                    # Standardized JSON response helper
├── routes/
│   └── routes.go                    # URL Route mapping
├── .env                             # Environment Variables
├── go.mod                           # Go dependencies
└── README.md
```

---

## API List

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/price` | Melihat harga emas terkini per gram |
| `POST` | `/buy` | Membeli emas berdasarkan nominal Rupiah |
| `POST` | `/sell` | Menjual emas berdasarkan Gram |
| `GET` | `/transactions` | Melihat riwayat seluruh transaksi |

---

## Example Requests & Responses

### 1. Get Gold Price
**Request:** `GET /price`
**Response:**
```json
{
  "price": 1945200
}
```

### 2. Buy Gold
**Request:** `POST /buy`
```json
{
  "amount": 500000
}
```
**Response:**
```json
{
  "gram": 0.2570429775858523,
  "price": 1945200
}
```
**Invalid Request Response:**
```json
{
  "error": "amount must be greater than zero"
}
```

### 3. Sell Gold
**Request:** `POST /sell`
```json
{
  "gram": 1
}
```
**Response:**
```json
{
  "amount": 1945200
}
```

### 4. Get Transactions History
**Request:** `GET /transactions`
**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "type": "BUY",
    "amount": 500000,
    "gram": 0.2570429775858523,
    "created_at": "2026-07-21T10:00:00Z"
  }
]
```

---

## Future Improvements

Berikut adalah fitur dan praktik pengembangan yang bisa ditingkatkan pada versi mendatang:
- **Authentication & JWT**: Melindungi endpoint pengguna menggunakan Token JWT dan Oauth2 Login.
- **Unit Testing Full Coverage**: Menggunakan *mock* (seperti `testify/mock`) untuk mencapai *code coverage* Service dan Handler hingga 100%.
- **Live Gold Price Integration**: Menarik data harga emas secara dinamis dan real-time dari API pihak eksternal (third party).
- **Dockerization**: Membungkus aplikasi Go dan database PostgreSQL menggunakan `docker-compose` agar mudah dijalankan (plug & play).
- **Pagination & Filtering**: Menambahkan query parameter (misal: `?limit=10&page=1`) pada Endpoint `GET /transactions` agar optimal jika data berjumlah besar.
- **Swagger Documentation**: Membuat spesifikasi OpenAPI untuk interaktivitas dokumentasi langsung pada URL aplikasi.
