# Local Setup

Dokumen ini menjelaskan cara menjalankan backend `api-idetech` secara lokal.

## Prasyarat

- Go `1.22.2` atau kompatibel
- Docker dan Docker Compose

## Menjalankan Dependency

```bash
docker compose -f /home/pgun/dev/idetech/api-split/deploy/docker-compose.yml up -d postgres redis minio
```

Jika port default bentrok, sesuaikan mapping di:

`/home/pgun/dev/idetech/api-split/deploy/docker-compose.yml`

## Menjalankan Migrasi

```bash
docker exec -i deploy-postgres-1 psql -U idetech -d idetech < /home/pgun/dev/idetech/api-split/backend/migrations/0001_init.sql
```

## Menjalankan API

```bash
cd /home/pgun/dev/idetech/api-split/backend

APP_ENV=development \
APP_PORT=8080 \
APP_BASE_URL=http://localhost:8080 \
APP_ALLOWED_ORIGINS=http://localhost:3000,http://demo.localhost:3000 \
POSTGRES_HOST=127.0.0.1 \
POSTGRES_PORT=5432 \
POSTGRES_DB=idetech \
POSTGRES_USER=idetech \
POSTGRES_PASSWORD=idetech \
POSTGRES_SSLMODE=disable \
JWT_ISSUER=idetech-api \
JWT_AUDIENCE=idetech-web \
JWT_SECRET=change-me \
go run ./cmd/api
```

Catatan:

- jika `POSTGRES_HOST` kosong, backend fallback ke repository memory;
- jika `POSTGRES_HOST` diisi, backend memakai PostgreSQL dan seed demo tenant/user akan dipastikan tersedia.

## Endpoint Saat Ini

- `GET /healthz`
- `GET /api/v1/tenant/bootstrap`
- `POST /api/v1/auth/login`
- `GET /api/v1/auth/me`
- `GET /api/v1/users`

## Kredensial Demo

- tenant: `demo`
- guru: `guru.demo` / `demo123`
- admin: `admin.demo` / `admin123`

## Build Check

```bash
cd /home/pgun/dev/idetech/api-split/backend
GOTOOLCHAIN=local go build ./...
```
