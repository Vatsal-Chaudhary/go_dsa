.PHONY: build run test test-file test-func

build:
	go build -o ./bin/go_dsa

run: build
	./bin/go_dsa

test:
	@if [ -n "$(FILE)" ] || [ -n "$(FUNC)" ]; then ./scripts/test.sh "$(FILE)" "$(FUNC)"; else go test ./...; fi

test-file:
	@./scripts/test.sh "$(FILE)"

test-func:
	@./scripts/test.sh "$(FILE)" "$(FUNC)"
