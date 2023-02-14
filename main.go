package main

import (
	"github.com/scaleway/cq-source-scaleway/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
