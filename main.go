package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/scaleway/cq-source-scaleway/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
