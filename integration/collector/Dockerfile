FROM golang:1.26 AS build

RUN go install go.opentelemetry.io/collector/cmd/builder@v0.156.0

# Mirror the sibling checkout layout that builder-config.yaml expects.
COPY --from=go-stix . /src/go-stix
WORKDIR /src/otelcol-exporter-stix
COPY . .

RUN CGO_ENABLED=0 builder --config builder-config.yaml

FROM gcr.io/distroless/static-debian12

COPY --from=build /src/otelcol-exporter-stix/build/otelcol-stix /otelcol-stix

ENTRYPOINT ["/otelcol-stix"]
CMD ["--config", "/etc/otelcol/config.yaml"]
