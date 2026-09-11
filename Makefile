fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

lint:
	golangci-lint run

clean:
	go clean

collector:
	go run go.opentelemetry.io/collector/cmd/builder@v0.156.0 --config builder-config.yaml
