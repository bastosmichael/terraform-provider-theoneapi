package theoneapi

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "api_token": {
                Type:     schema.TypeString,
                Required: true,
                DefaultFunc: schema.EnvDefaultFunc("THE_ONE_API_TOKEN", nil),
                Description: "API Token for theOneAPI.",
            },
        },

        ResourcesMap: map[string]*schema.Resource{
            "theoneapi_book":  resourceBook(),
            "theoneapi_movie": resourceMovie(),
        },

        ConfigureFunc: providerConfigure,
    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    apiToken := d.Get("api_token").(string)

    client := NewClient(apiToken)

    return client, nil
}

type Client struct {
    APIToken string
}

func NewClient(apiToken string) *Client {
    return &Client{APIToken: apiToken}
}

func resourceBook() *schema.Resource {
    return &schema.Resource{
        Create: resourceBookCreate,
        Read:   resourceBookRead,
        Update: resourceBookUpdate,
        Delete: resourceBookDelete,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceBookCreate(d *schema.ResourceData, m interface{}) error {
    // Logic to create a book
    // Use the API client and call the appropriate endpoint
    return resourceBookRead(d, m)
}

func resourceBookRead(d *schema.ResourceData, m interface{}) error {
    // Logic to read a book
    return nil
}

func resourceBookUpdate(d *schema.ResourceData, m interface{}) error {
    // Logic to update a book
    return resourceBookRead(d, m)
}

func resourceBookDelete(d *schema.ResourceData, m interface{}) error {
    // Logic to delete a book
    return nil
}

func resourceMovie() *schema.Resource {
    return &schema.Resource{
        Create: resourceMovieCreate,
        Read:   resourceMovieRead,
        Update: resourceMovieUpdate,
        Delete: resourceMovieDelete,

        Schema: map[string]*schema.Schema{
            "name": {
                Type:     schema.TypeString,
                Required: true,
            },
            "runtime_in_minutes": {
                Type:     schema.TypeInt,
                Optional: true,
            },
            // Add more fields as necessary
        },
    }
}

func resourceMovieCreate(d *schema.ResourceData, m interface{}) error {
    // Logic to create a movie
    return resourceMovieRead(d, m)
}

func resourceMovieRead(d *schema.ResourceData, m interface{}) error {
    // Logic to read a movie
    return nil
}

func resourceMovieUpdate(d *schema.ResourceData, m interface{}) error {
    // Logic to update a movie
    return resourceMovieRead(d, m)
}

func resourceMovieDelete(d *schema.ResourceData, m interface{}) error {
    // Logic to delete a movie
    return nil
}
