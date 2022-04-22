# todoApp on Golang

## Run app in docker compose

```bash
docker compose up -d
```

## Run app
```bash
make run
```

## Run Postgres Docker container
```bash
make run-postgres
```

## Run Postgres Docker CLI
```bash
docker ps
docker exec -it <container_id> /bin/bash
psql -U postgres
```

## Postgres migrations
```bash
make migrate-up
```

```bash
make migrate-down
```

## Generate mocks
```bash
make mock-gen
```