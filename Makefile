.PHONY: build
build:
	go build -o app

.PHONY: run
run:
	go run main.go

.PHONY: test
test: test-func test-e2e

.PHONY: test-func
test-func:
	go test -cover -bench=. ./...

.PHONY: test-e2e
test-e2e:
	docker-compose run --rm tester
	docker-compose down --remove-orphans