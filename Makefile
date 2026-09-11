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

COMPOSE = docker compose -f integration/docker-compose.yaml

e2e:
	$(COMPOSE) up -d --build --wait --force-recreate
	go test -tags integration -count=1 ./integration/...; status=$$?; \
	[ $$status -eq 0 ] || $(COMPOSE) logs collector; \
	$(COMPOSE) down; exit $$status
