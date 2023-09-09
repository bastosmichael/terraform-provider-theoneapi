package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/bastosmichael/terraform-provider-theoneapi/theoneapi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: theoneapi.Provider,
	})
}
