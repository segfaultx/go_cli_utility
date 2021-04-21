.PHONY: clean

run: build
	./Orchestrator_CLI
clean:
	rm Orchestrator_CLI

build:
	go build

windows: Orchestrator_CLI.go
	env GOOS=windows GOARCH=amd64 go build

