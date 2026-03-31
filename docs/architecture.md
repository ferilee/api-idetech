# Architecture Notes

This repository follows the constraints defined in `blue_print_idetech.md`:

- Go + Chi for the backend API;
- Next.js App Router for the frontend;
- PostgreSQL for relational data;
- Redis for fast state and cache;
- MinIO for tenant-scoped object storage;
- shared-schema multi-tenancy with tenant-aware rows.

The current bootstrap uses an in-memory tenant repository so the API can run before
database integration is completed. The migration in `backend/migrations` is the source
of truth for the first persistence milestone.
