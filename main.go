package main

import (
	"github.com/bastosmichael/terraform-provider-theoneapi/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: theoneapi.Provider,
	})
}
