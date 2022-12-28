PORT:=8080
REGISTRY:=github.com/lucasvmiguel
API_IMAGE:=stock-api
VERSION:=latest
ENV:=DEVELOPMENT

test-unit:
	scripts/mockgen.sh
	ENV=TEST go test -cover $(shell go list ./... | grep -v integration)

test-integration:
	go clean -cache
	ENV=TEST go test -cover $(shell go list ./... | grep integration)

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
