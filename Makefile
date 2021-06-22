all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o build/openweather_exporter ./cmd/exporter
