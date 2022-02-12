PHONY:

run-srv: build-srv
	./bin/srv

run-cli: build-cli
	./bin/cli

build-srv:
	go build -o ./bin/srv ./server/main/main.go

build-cli:
	go build -o ./bin/cli ./client/main/main.go