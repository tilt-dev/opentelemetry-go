module go.opentelemetry.io/otel/exporters/trace/zipkin

go 1.13

replace go.opentelemetry.io/otel => ../../..

require (
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/otel v0.15.0 // indirect
	go.opentelemetry.io/otel/sdk v0.15.0
	google.golang.org/grpc v1.30.0
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
