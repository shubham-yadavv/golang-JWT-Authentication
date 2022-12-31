build:
	go build -o bin/main main.go

run: build
	./bin/main

dev:
	nodemon --exec go run main.go --signal SIGTERM