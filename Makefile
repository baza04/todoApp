.SILENT:

build:
	go build -o ./bin/app cmd/web/main.go

#run: build
#	./bin/app

run:
	docker-compose up --remove-orphans --build app

 run-postgres:
 	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d postgres

# restart-postgres:
# 	docker restart todo-db

# stop-postgres:
# 	docker stop todo-db

# postgres-cli:
# 	docker exec -it todo-db /bin/bash

create-new-migration:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

init-swagger:
	swag init -g cmd/web/main.go

test-cover:
	go test -coverprofile=coverage.out ./... -v -coverpkg=./...

html-cover: test-cover
	go tool cover -html=coverage.out

# build:
# 	go build -o ./bin/app ./cmd/web/main.go
# run: build
# 	./bin/app

# pkgs = $(shell go list ./... | fgrep -v /vendor)

# get-linters: # without golangci-linter
# 	go get golang.org/x/lint/golint
# 	go get honnef.co/go/tools/cmd/staticcheck
# 	go get github.com/kisielk/errcheck

# 	golint $(pkgs)
# 	go vet $(pkgs)
# 	staticcheck $(pkgs)
# 	errcheck $(pkgs)

lint:
	golangci-lint run -c .golangci.yml

gci-lint:
	golangci-lint run --fix --disable-all -E govet,gosimple,unused,stylecheck,unparam,staticcheck,errcheck,gofmt,deadcode

mock-gen:
	mockgen -source=./pkg/repository/repository.go -destination=./pkg/repository/repo_mock.go --package=repository
	mockgen -source=./pkg/service/service.go -destination=./pkg/service/service_mock.go --package=service
