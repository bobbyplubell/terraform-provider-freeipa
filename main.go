package main

import (
	"github.com/camptocamp/terraform-provider-freeipa/freeipa"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: freeipa.Provider,
	})
}
