build:
	go build -o ./bin/go_dsa

run: build
	./bin/go_dsa	

test:
	go test ./...