PORT:=8080
REGISTRY:=github.com/lucasvmiguel
API_IMAGE:=stock-api
VERSION:=latest
LOCAL_URL:=http://localhost:8080

test-unit: generate-mocks
	ENV=TEST go test -cover $(shell go list ./... | grep -v test)

test-integration:
	go clean -cache
	ENV=TEST go test -cover $(shell go list ./... | grep test)

test-stress:
	ab -n 1000 -c 10 -T 'application/json' -p ./test/stress/create_data.json $(LOCAL_URL)/api/v1/products > test/stress/create-report.txt &
	ab -n 1000 -c 10 $(LOCAL_URL)/api/v1/products > test/stress/get-paginated-report.txt &
	@echo "Waiting for stress tests to finish..."
	@sleep 10
	@echo "Stress tests finished! Check the reports in test/stress folder."

run:
	go run cmd/api/main.go

build:
	go build cmd/api/main.go

docker-build:
	docker build -t $(REGISTRY)/$(API_IMAGE):$(VERSION) -f cmd/api/Dockerfile .

docker-run:
	docker run --rm -p $(PORT):$(PORT) $(REGISTRY)/$(API_IMAGE):$(VERSION)

persistence-up:
	docker-compose up

persistence-down:
	docker-compose down

generate-mocks:
	go run github.com/golang/mock/mockgen -source=./internal/product/handler/handler.go -package=handler -destination=./internal/product/handler/handler_mocks.go
	go run github.com/golang/mock/mockgen -source=./internal/product/service/service.go -package=service -destination=./internal/product/service/service_mocks.go
