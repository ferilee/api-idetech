# IdeTech

Initial monorepo scaffold for the IdeTech platform based on `blue_print_idetech.md`.

## Structure

```text
backend/   Go API (Chi, tenant-aware bootstrap)
frontend/  Next.js app scaffold
deploy/    Docker and environment assets
docs/      Product and engineering notes
```

## Current scope

- tenant-aware API bootstrap foundation;
- PostgreSQL migration baseline;
- frontend landing page with tenant bootstrap flow;
- local Docker Compose for app dependencies.

## Next milestones

1. auth and JWT flow;
2. tenant and user persistence with PostgreSQL;
3. material upload to MinIO;
4. quiz generation pipeline.
