package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENTRY_TOKEN", nil),
				Description: "The authentication token used to connect to Sentry",
			},
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENTRY_BASE_URL", "https://app.getsentry.com/api/"),
				Description: "The Sentry Base API URL",
			},
		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}