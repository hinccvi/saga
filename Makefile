MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v -e server -e test -e middleware -e entity -e constants -e mocks -e errors -e proto -e jwt | sort -r )
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
HOST_IP ?= $(shell ifconfig | grep -E "([0-9]{1,3}\.){3}[0-9]{1,3}" | grep -v 127.0.0.1 | awk '{print $$2}' | cut -f2 -d: |head -n1)

CONFIG_FILE ?= ./config/local.yml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := migrate -path migrations -database "$(APP_DSN)"
DOCKER_REPOSITORY := hinccvi/saga
MOCKERY := mockery --name=Repository -r --output=./internal/mocks

PID_FILE := './.pid'
FSWATCH_FILE := './fswatch.cfg'

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run-customer
run-customer: ## run the customer service
	go run ${LDFLAGS} cmd/customer/main.go

.PHONY: run-order
run-order: ## run the order service
	go run ${LDFLAGS} cmd/order/main.go

.PHONY: run-order-history
run-order-history: ## run the order history service
	go run ${LDFLAGS} cmd/orderHistory/main.go

.PHONY: build
build:  ## build the arm API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o server $(MODULE)/cmd/server

build-amd64:  ## build the amd64 API server binary
	CGO_ENABLED=0 GOARCH=amd64 go build ${LDFLAGS} -a -o server $(MODULE)/cmd/server

.PHONY: build-docker
build-docker: ## build the program as a arm docker image
	docker build -f cmd/server/Dockerfile -t $(DOCKER_REPOSITORY):$(VERSION) .

build-docker-amd64: ## build the program as a amd64 docker image
	docker build --platform=linux/amd64 -f cmd/server/Dockerfile -t $(DOCKER_REPOSITORY):$(VERSION) .

.PHONY: push-docker
push-docker: ## push docker image to dockerhub
	docker push $(DOCKER_REPOSITORY):$(VERSION)

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: postgres-start
postgres-start: ## start the postgres image
	@mkdir -p testdata/postgres
	docker run --rm --name postgres -v $(shell pwd)/testdata:/testdata \
		--network app-tier \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -d -p 5432:5432 postgres \
		-c wal_level=logical \
		-c max_wal_senders=8 \
		-c max_replication_slots=4 \
		-c listen_addresses=*

.PHONY: postgres-stop
postgres-stop: ## stop the postgres image
	docker stop postgres

.PHONY: mongo-start
mongo-start: ## start the mongo image
	docker run --rm --name mongo -d -p 27017:27017 --network app-tier \
		-v $(shell pwd)/testdata/mongo:/data/db -d \
		-e MONGO_INITDB_ROOT_USERNAME=mongo \
		-e MONGO_INITDB_ROOT_PASSWORD=mongo \
		mongo

.PHONY: mongo-stop
mongo-stop: ## stop the mongo image
	docker stop mongo

.PHONY: zookeeper-start
zookeeper-start: ## start the zookeeper
	docker run --rm -dp 2181:2181 --name zookeeper \
		--network app-tier \
		debezium/zookeeper:latest

.PHONY: kafka-start
kafka-start: ## start the kafka server
	@mkdir -p testdata/kafka
	docker run --rm -dp 9092:9092 --name kafka \
    --network app-tier \
		-e KAFKA_ADVERTISED_HOST_NAME=$(HOST_IP) \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
		-e KAFKA_AUTO_CREATE_TOPICS_ENABLE='true' \
		-e KAFKA_DELETE_TOPIC_ENABLE='true' \
		-e ZOOKEEPER_CONNECT=zookeeper:2181 \
		-v $(shell pwd)/testdata/kafka:/kafka/data \
    debezium/kafka:latest

.PHONY: connector-start
connector-start: ## start the kafka connector
	docker run --rm -dp 8083:8083 --name connector \
		--network app-tier \
		-e BOOTSTRAP_SERVERS=kafka:9092 \
		-e GROUP_ID=pg \
		-e CONFIG_STORAGE_TOPIC=pg_connect_configs \
		-e OFFSET_STORAGE_TOPIC=pg_connect_offsets \
		-e STATUS_STORATE-TOPIC=pg_connect_statuses \
		debezium/connect:latest

.PHONY: zookeeper-stop
zookeeper-stop: ## stop the zookeeper server
	docker stop kafka 

.PHONY: kafka-stop
kafka-stop: ## stop the kafka server
	docker stop zookeeper

.PHONY: connector-stop
connector-stop: ## stop the connector server
	docker stop connector

.PHONY: testdata
testdata: ## populate the database with test data
	make migrate-reset
	@echo "Populating test data..."
	@docker exec -it postgres psql "$(APP_DSN)" -f /testdata/testdata.sql

.PHONY: lint
lint: ## run golangci-lint on all Go package (requires golangci-lint)
	@golangci-lint run

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir migrations/ -seq $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: mockery
mockery: ## mock code autogenerator 
	@read -p "Enter the repository name: " repo; \
	struct_name="$$(tr '[:lower:]' '[:upper:]' <<< $${repo:0:1})$${repo:1}"; \
	$(MOCKERY) --dir=./internal/$${repo// /_}/repository/ --filename=$${repo// /_}Repository.go --structname=$${struct_name}Repository

.PHONY: protoc
protoc: ## compile protobuf file
	@protoc \
  --go_out=./proto/pb --go_opt=paths=source_relative \
  --go-grpc_out=./proto/pb --go-grpc_opt=paths=source_relative \
  --validate_out=lang=go,paths=source_relative:./proto/pb \
  -Iproto/ $$(find proto -iname "*.proto") \
  -I $$GOPATH/src/protoc-gen-validate/
