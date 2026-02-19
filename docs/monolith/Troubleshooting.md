# Troubleshooting

## Remote Module Repeatedly Downloaded

### Problem Description

Remote modules are downloaded on every reconciliation, causing:

- High network egress costs (up to 7.2GB/day per workspace)
- Slower reconciliation times due to download latency
- Unnecessary bandwidth usage

**Example scenario:**
- 50MB remote module
- Reconciles every 10 minutes (6 times/hour)
- Result: 300MB/hour = 7.2GB/day of repeated downloads

### Solution

Set `remotePullPolicy: IfNotPresent` in the Workspace spec to download the module once and reuse it:

```yaml
apiVersion: opentofu.upbound.io/v1beta1
kind: Workspace
metadata:
  name: example
spec:
  forProvider:
    source: Remote
    module: git::https://github.com/terraform-aws-modules/terraform-aws-vpc?ref=v5.1.2
    remotePullPolicy: IfNotPresent  # Add this line
```

**Benefits:**
- Significant network reduction (example: 50MB module at 10-min poll = 7.2GB/day -> 50MB/day, 98.6% savings)
- Faster reconciliation after initial download
- Module automatically re-downloaded if URL changes

**Best practices:**
- Always use with pinned module versions (e.g., `?ref=v5.1.2`, not `?ref=main`)
- Verify the module URL includes a specific version tag or commit SHA
- Monitor logs to confirm download skipping behavior

**Verification:**
```bash
# Check workspace logs
kubectl logs -n upbound-system deploy/provider-opentofu-* | grep "Remote module"

# First reconciliation: "Remote module downloaded"
# Subsequent: "Remote module already present, skipping download"

# Check status field
kubectl get workspace <name> -o jsonpath='{.status.atProvider.remoteSource}'
```

**When to use Always policy (default):**
- Development workspaces with frequently changing modules
- Modules without pinned versions (floating refs like `main` or `develop`)
- When module freshness is more important than cost

See the [Remote Module Pull Policy](RemoteModulePullPolicy.md) documentation for detailed information.
