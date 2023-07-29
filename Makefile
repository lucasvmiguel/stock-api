PORT:=8080
REGISTRY:=github.com/lucasvmiguel
API_IMAGE:=stock-api
VERSION:=latest
LOCAL_URL:=http://localhost:$(PORT)

test-unit: install generate-mocks
	ENV=TEST go test -cover $(shell go list ./... | grep -v test)

test-integration: install
	go clean -cache
	ENV=TEST go test -cover $(shell go list ./... | grep test)

test-stress:
	ab -n 1000 -c 10 -T 'application/json' -p ./test/stress/create_data.json $(LOCAL_URL)/api/v1/products > test/stress/create-report.txt &
	ab -n 1000 -c 10 $(LOCAL_URL)/api/v1/products > test/stress/get-paginated-report.txt &
	@echo "Waiting for stress tests to finish..."
	@sleep 10
	@echo "Stress tests finished! Check the reports in test/stress folder."

install:
	go mod tidy
	go get github.com/golang/mock/mockgen@v1.6.0

run-api: install
	go run cmd/api/main.go

build: install
	go build cmd/api/main.go

docker-build:
	docker build -t $(REGISTRY)/$(API_IMAGE):$(VERSION) -f cmd/api/Dockerfile .

persistence-up:
	docker-compose up

persistence-down:
	docker-compose down

generate-mocks:
	go run github.com/golang/mock/mockgen -source=./internal/product/handler/handler.go -package=handler -destination=./internal/product/handler/handler_mocks.go
	go run github.com/golang/mock/mockgen -source=./internal/product/service/service.go -package=service -destination=./internal/product/service/service_mocks.go

generate_tls_files:
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
