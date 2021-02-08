package provider

import (
	"context"
	"log"

	"github.com/Mongey/terraform-provider-hellosign/internal/hellosign"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	log.Printf("[INFO] Creating HelloSign Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HELLOSIGN_API_KEY", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"hellosign_api_app": apiAppResource(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log.Printf("[INFO] Configuring HelloSign client")
	apiKey := d.Get("api_key").(string)
	cfg := &hellosign.Config{APIKey: apiKey}
	c := hellosign.NewClient(cfg)

	return c, nil
}
