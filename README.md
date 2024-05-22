### PostgreSQL
Run docker compose up from projects root dir and wait for PostgreSQL to be ready.

Make sure the port 5432 is free.

### App
From inside apps folder, run:
- go mod tidy
- go run main.go -c ../common/env_vars
