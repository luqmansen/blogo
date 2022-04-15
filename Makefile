GOOS=darwin
GOPATH:=$(shell go env GOPATH)
GIT_COMMIT_SHA := $(shell git rev-list -1 HEAD)

ROOT := github.com/luqmansen/blogo

.PHONY: dev
RUN:
	nodemon --exec go run cmd/app/main.go --signal SIGTERM

test:
	go clean -testcache
	go test ./... -v

build:
	rm -rf build/app
	GOOS=$(GOOS) CGO_ENABLED=0 go build -ldflags "-s -w -X $(ROOT)/pkg/version.GitCommitSha=$(GIT_COMMIT_SHA)" \
		-o build/app/app cmd/app/main.go

create-migration:
	docker run --rm -v `pwd`/db/migration:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(name)

migrate-up:
	docker run --rm -v `pwd`/db/migration:/migrations --network host migrate/migrate  -path=/migrations/ -database "mysql://root:root@(0.0.0.0:3306)/blogo" up

migrate-down:
	docker run --rm -v `pwd`/db/migration:/migrations --network host migrate/migrate  -path=/migrations/ -database "mysql://root:root@(0.0.0.0:3306)/blogo" down -all


gen-mock:
	docker run --rm -v $(PWD):/app -w /app luqmansen/docker-mockgen mockgen -source=pkg/services/services.go -destination=mock/service_mock.go -package=mock

gen-server:
	oapi-codegen --package server --generate gin api/openapi/swagger.json > internal/server/server.gen.go
	oapi-codegen --package server --generate spec api/openapi/swagger.json > internal/server/spec.gen.go
	oapi-codegen --package server --generate types api/openapi/swagger.json > internal/server/types.gen.go

