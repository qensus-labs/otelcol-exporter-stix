# OpenTelemetry Collector STIX Exporter

The STIX Exporter converts OpenTelemetry Logs into STIX 2.1 cyber threat intelligence bundles.

## Features

- OpenTelemetry Collector exporter
- STIX 2.1 compliant output
- Built on github.com/qensus-labs/go-stix
- Deterministic STIX IDs
- Native OpenTelemetry semantic convention mapping

## Status

🚧 Early development

## Development

The exporter depends on unreleased `go-stix` changes, so check out
[go-stix](https://github.com/qensus-labs/go-stix) next to this repository:

```text
qensus-labs/
├── go-stix/
└── otelcol-exporter-stix/
```

Build a Collector (`build/otelcol-stix`) containing the OTLP receiver and the
STIX exporter:

```bash
make collector
```

Run the end-to-end test (requires Docker). It starts a
[Medallion](https://github.com/oasis-open/cti-taxii-server) TAXII 2.1 server
and the Collector, sends OTLP logs, and checks the resulting STIX objects
arrive in the TAXII collection:

```bash
make e2e
```

To explore manually, start the stack with
`docker compose -f integration/docker-compose.yaml up --build` and send OTLP/JSON
logs to `http://localhost:4318/v1/logs`. Medallion listens on
`http://localhost:9088` (`admin` / `Password01`).

## License

Apache License 2.0