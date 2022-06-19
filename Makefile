.PHONY:
.SILENT:
.DEFAULT_GOAL := run

test:
	go test --short -coverprofile=cover.out -v ./...
	make test.coverage
test.coverage:
	go tool cover -func=cover.out | grep "total"
build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go
run:
	go run ./cmd/app/main.go
swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g internal/app/app.go
gen:
	mockgen -source=internal/services/services.go -destination=internal/services/mocks/mock.go