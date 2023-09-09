// +build tools

package tools

import (
	// This is an example; please replace with actual required tools for your project.
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// This file imports packages that are used when running go generate, or used
// during the development process but not otherwise depended on by built code.
