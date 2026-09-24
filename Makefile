APP_NAME=pr-sentinel

.PHONY: build test run clean

build:
	go build -o bin/$(APP_NAME) ./cmd/sentinel

test:
	go test ./...

run:
	go run ./cmd/sentinel

clean:
	rm -rf bin/
	rm -f pr-sentinel-report.md