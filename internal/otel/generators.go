// Generate the ExternalConnector interface from the 'connector' struct in the workspace controller.
//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../controller/namespaced/workspace/workspace.go" -s "connector" -i "ExternalConnector" -y "ExternalConnector defines how to establish a connection between a Crossplane managed resource and its corresponding external system." -p "otel" -d "false" -o "external_connector.go"

// Generate the ExternalClient interface from the 'external' struct in the workspace controller.
//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../controller/namespaced/workspace/workspace.go" -s "external" -i "ExternalClient" -y "ExternalClient defines the interface for interacting with external resources." -p "otel" -d "false" -o "external_client.go"

// Generate the OpenTofuClient interface from the 'Harness' struct in opentofu.go.
//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../opentofu/opentofu.go" -s "Harness" -i "OpenTofuClient" -y "OpenTofuClient defines the interface for interacting with an OpenTofu backend." -p "otel" -m "github.com/upbound/provider-opentofu/internal/opentofu" -d "false" -o "opentofu_client.go"

// Generate a wrapper for ExternalClient that adds OpenTelemetry instrumentation.
//go:generate go run github.com/hexdigest/gowrap/cmd/gowrap@v1.4.3 gen -g -p "github.com/upbound/provider-opentofu/internal/otel" -i "ExternalClient" -t "opentelemetry.gotmpl" -o "instrumented_external_client.go"

// Generate a wrapper for ExternalConnector with OpenTelemetry instrumentation.
//go:generate go run github.com/hexdigest/gowrap/cmd/gowrap@v1.4.3 gen -g -p "github.com/upbound/provider-opentofu/internal/otel" -i "ExternalConnector" -t "opentelemetry.gotmpl" -o "instrumented_external_connector.go"

// Generate a wrapper for OpenTofuClient with OpenTelemetry instrumentation.
//go:generate go run github.com/hexdigest/gowrap/cmd/gowrap@v1.4.3 gen -g -p "github.com/upbound/provider-opentofu/internal/otel" -i "OpenTofuClient" -t "opentelemetry.gotmpl" -o "instrumented_opentofu_client.go"

package otel
