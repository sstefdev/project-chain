build:
	go build -o ./bin/project-chain

run: build
	./bin/project-chain
	
test:
	go test -v ./...