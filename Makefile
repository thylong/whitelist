.PHONY: build
build:
	go build -o app

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	go vet
	golint

.PHONY: test
test: test-func test-e2e test-bench

.PHONY: test-func
test-func:
	go test -cover ./...

.PHONY: test-bench
test-bench:
	go test -bench=. -tags=benchmark ./...

.PHONY: test-e2e
test-e2e:
	docker-compose run --rm tester
	docker-compose down --remove-orphans