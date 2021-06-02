.PHONY: clean

run: build
	./orca
clean:
	rm orca

build:
	go build

windows: orca.go
	env GOOS=windows GOARCH=amd64 go build
macos: orca.go
	env GOOS=darwin GOARCH=amd64 go build
