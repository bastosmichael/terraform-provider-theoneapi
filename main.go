package main

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"internal/provider"
)

func main() {
	// Check if TF_PROVIDER_DEBUG environment variable is set to "true"
	if os.Getenv("TF_PROVIDER_DEBUG") == "true" {
		plugin.DebugMode = true
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
