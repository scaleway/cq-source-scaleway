A [Scaleway](https://scaleway.com/) source plugin for CloudQuery that loads data from Scaleway to any database, data warehouse or data lake supported by CloudQuery, such as PostgreSQL, BigQuery, Athena, and many more.

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
To enable this, `backend_options` must be set in the spec (as shown below). This is documented in the [Managing Incremental Tables](https://cloudquery.io/docs/advanced-topics/managing-incremental-tables) section.

## Configuration

The following source configuration file will sync all data from Scaleway to a PostgreSQL database. See [the CloudQuery Quickstart](https://cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: "scaleway"
  registry: "cloudquery"
  path: "scaleway/scaleway"
  version: "v2.0.0"
  # use this to enable incremental syncing
  # backend_options:
  #   table_name: "cq_state_scaleway"
  #   connection: "@@plugins.DESTINATION_NAME.connection"
  tables: 
    - "*"
  skip_tables:
    - "scaleway_ipfs_volumes"
    - "scaleway_marketplace_image_versions"
  destinations: 
    - "postgresql"
  spec:
    # plugin spec section
```

### Plugin Spec

- `regions` (`[]string`) (optional) (default: all regions)

  List of regions to query.

- `zones` (`[]string`) (optional) (default: to all zones)

  List of zones to query.

- `timeout_secs` (`integer` in seconds) (optional) (default: `10`)

  Timeout for requests against the Scaleway API endpoint.

- `concurrency` (`integer`) (optional) (default: `1000`)

  Best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

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

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).

### Publish a new version to the Cloudquery Hub

Tagging a new release should invoke the GitHub Actions workflow to publish to the [Cloudquery Hub](https://hub.cloudquery.io/).

`CQ_CI_CLOUDQUERY_API_KEY` should be set up in GitHub secrets (Settings `->` Security `->` Secrets and variables `->` Actions) beforehand, with a valid [API key](https://cloudquery.io/docs/deployment/generate-api-key).

After publishing the new version, it will [show up](https://hub.cloudquery.io/plugins/source/scaleway/scaleway) in the [hub](https://hub.cloudquery.io/).

For more information please refer to the official [Publishing a Plugin to the Hub](https://cloudquery.io/docs/developers/publishing-a-plugin-to-the-hub) guide.
