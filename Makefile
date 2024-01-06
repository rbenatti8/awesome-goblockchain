build:
	go build -o bin/goblockchain -v

run: build
	./bin/goblockchain

test:
	go test -v ./...