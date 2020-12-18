module go.opentelemetry.io/otel/exporters/trace/jaeger

go 1.13

replace go.opentelemetry.io/otel => ../../..

require (
	github.com/apache/thrift v0.13.0
	github.com/google/go-cmp v0.5.4
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/otel v0.15.0 // indirect
	go.opentelemetry.io/otel/sdk v0.15.0
	google.golang.org/api v0.36.0
	google.golang.org/grpc v1.33.2
)
