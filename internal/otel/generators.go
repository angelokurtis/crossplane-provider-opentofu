//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../controller/namespaced/workspace/workspace.go" -s "connector" -i "ExternalConnector" -y "ExternalConnector defines how to establish a connection between a Crossplane managed resource and its corresponding external system." -p "otel" -d "false" -o "external_connector.go"
//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../controller/namespaced/workspace/workspace.go" -s "external" -i "ExternalClient" -y "ExternalClient defines the interface for interacting with external resources." -p "otel" -d "false" -o "external_client.go"
//go:generate go run github.com/vburenin/ifacemaker@v1.3.0 -f "../opentofu/opentofu.go" -s "Harness" -i "OpenTofuClient" -y "OpenTofuClient defines the interface for interacting with an OpenTofu backend." -p "otel" -m "github.com/upbound/provider-opentofu/internal/opentofu" -d "false" -o "opentofu_client.go"

package otel
