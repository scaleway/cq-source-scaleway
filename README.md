# CloudQuery Scaleway Source Plugin

[![test](https://github.com/scaleway/cq-source-scaleway/actions/workflows/test.yaml/badge.svg)](https://github.com/scaleway/cq-source-scaleway/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scaleway/cq-source-scaleway)](https://goreportcard.com/report/github.com/scaleway/cq-source-scaleway)

A [Scaleway](https://scaleway.com/) source plugin for CloudQuery that loads data from Scaleway to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

 - [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
 - [Supported Tables](docs/tables/README.md)

## Authentication

Credentials are used from the default config file and the environment. Get your credentials from the [IAM dashboard](https://console.scaleway.com/iam/api-keys)

Set credentials in environment variables or use the default config file at `~/.config/scw/config.yaml`. Config format and locations for other platforms are documented in the [SDK Docs](https://github.com/scaleway/scaleway-sdk-go/tree/master/scw#scaleway-config)

The environment variables to set are:
  - `SCW_ACCESS_KEY`
  - `SCW_SECRET_KEY`
  - `SCW_DEFAULT_ORGANIZATION_ID`

Env vars override config values if both are set. By default all regions and zones are queried.

## Incremental Syncing

The Scaleway plugin supports incremental syncing. This means that only new data will be fetched from Scaleway and loaded into your destination for supported tables (support depending on API endpoint). This is done by keeping track of the last item fetched and only fetching data that has been created since then.
To enable this, `backend` option must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](https://www.cloudquery.io/docs/advanced-topics/managing-incremental-tables) section.

## Configuration

The following source configuration file will sync all data from Scaleway to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "scaleway"
  path: "scaleway/scaleway"
  version: "${VERSION}"
  # backend: "local" # use this to enable incremental syncing
  tables: 
    - "*"
  skip_tables:
    - "scaleway_ipfs_volumes"
    - "scaleway_marketplace_image_versions"
  destinations: 
    - "postgresql"
  backend: local
  spec:
    # plugin spec section
```

### Plugin Spec

- `regions` (list of string, optional. Defaults to all regions):
  List of regions to query.

- `zones` (list of string, optional. Defaults to all zones):
  List of zones to query.

- `timeout_secs` (integer in seconds, optional. Default: 10):
  Timeout for requests against the Scaleway API endpoint.

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Follow [this link](https://github.com/scaleway/cq-source-scaleway/releases/new) to draft a new release.
2. Click `Choose a tag` and enter the new version number:
![image](https://user-images.githubusercontent.com/26760571/219360662-0ad1f83d-84c9-47c8-afb9-fe774ce03dcc.png)
3. Click `Create new tag: <version> on publish` assuming it's a new tag.
4. Click `Generate release notes` to automatically generate release notes.
5. Click `Publish release` to publish the release.

> Once the tag is pushed, a new GitHub Actions workflow will be triggered to build and upload the release binaries to the release
