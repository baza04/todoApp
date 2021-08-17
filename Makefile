.SILENT:

run:
	go run ./cmd/web/main.go

run-postgres:
	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

restart-postgres:
	docker restart todo-db

create-new-migration:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down
	