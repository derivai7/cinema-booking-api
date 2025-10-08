## Bahtiar Rifa'i

Backend API untuk platform pembelian tiket bioskop online dengan fitur seat locking, schedule management, dan auto-refund.
**Backend Development Test - Mitra Kasih Perkasa 2025**

## Tech Stack

- **Language:** Go
- **Framework:** Gin + Go Validator
- **Database:** PostgreSQL + GORM
- **Auth:** JWT (golang-jwt/jwt)
- **Config:** Viper
---

## Quick Start

### 1. Setup Database

```bash
createdb cinema_booking

psql -U postgres -d cinema_booking -f schema.sql
```

### 2. Setup Environment

```bash
cp .env.example .env

# Edit sesuai konfigurasi Anda
nano .env
```

### 3. Install Dependencies

```bash
go mod download
go mod tidy
```

### 4. Run Application

```bash
# Development
go run cmd/api/main.go
./cinema-api
```

### Test Credentials

```
Admin:    admin@cinema.com    / password123
Staff:    staff@cinema.com    / password123
Customer: customer@cinema.com / password123
```

### Postman Collection

Import file: `cinema-booking-api.postman_collection.json`

### Test Flow

1. Login sebagai Admin/Staff/Customer
2. Get All Schedules (semua role bisa)
3. Create Schedule (Admin/Staff only)
4. Update Schedule (Admin/Staff only)
5. Delete Schedule (Admin/Staff only)

---

## API Endpoints

### Authentication

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/api/auth/login` | No   | User login |

### Schedules

| Method | Endpoint | Auth | Role | Description |
|--------|----------|------|------|-------------|
| GET | `/api/schedules` | Yes  | All | Get all schedules |
| GET | `/api/schedules/:id` | Yes  | All | Get schedule by ID |
| POST | `/api/schedules` | Yes  | Admin, Staff | Create schedule |
| PUT | `/api/schedules/:id` | Yes  | Admin, Staff | Update schedule |
| DELETE | `/api/schedules/:id` | Yes  | Admin, Staff | Delete schedule |

**Response Format:**
```json
{
  "success": true,
  "message": "Success message",
  "data": { ... }
}
```

---

## Documentation

### System Design

Lihat dokumentasi lengkap di [`docs/system-design/SYSTEM_DESIGN.md`](docs/system-design/SYSTEM_DESIGN.md)

### Database Schema

Schema lengkap ada di file [`docs/database/schema.sql`](docs/database/schema.sql)