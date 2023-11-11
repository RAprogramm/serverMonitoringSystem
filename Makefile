run:
	./scripts/run.sh

build:
	go build -v ./...

.DEFAULT_GOAL := run
