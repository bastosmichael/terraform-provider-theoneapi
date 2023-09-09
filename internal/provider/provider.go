package theoneapi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("THEONEAPI_ENDPOINT", nil),
				Description: "The endpoint for TheOneAPI.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"theoneapi_book": resourceBook(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIEndpoint: d.Get("api_endpoint").(string),
	}

	return config.Client()
}
