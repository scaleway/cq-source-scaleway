# CloudQuery Scaleway Source Plugin

[![test](https://github.com/scaleway/cq-source-scaleway/actions/workflows/test.yaml/badge.svg)](https://github.com/scaleway/cq-source-scaleway/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scaleway/cq-source-scaleway)](https://goreportcard.com/report/github.com/scaleway/cq-source-scaleway)

A [Scaleway](https://scaleway.com/) source plugin for CloudQuery that loads data from Scaleway to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync all data points for `mywebsite.com` to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "scaleway"
  path: "scaleway/scaleway"
  version: "${VERSION}"
  # backend: "local" # use this to enable incremental syncing
  tables: 
    ["*"]
  destinations: 
    - "postgresql"
  spec:
    # plugin spec section
```

### Plugin Spec

- `timeout_secs` (integer in seconds, optional. Default: 10):
  Timeout for requests against the Scaleway API endpoint.
