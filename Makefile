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
	go test -cover ./...

.PHONY: test-e2e
test-e2e:
	go run main.go &
	go build -o e2e-tester ./tests/e2e && ./e2e-tester