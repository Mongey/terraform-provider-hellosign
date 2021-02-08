package provider

import (
	"log"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func accProvider() map[string]*schema.Provider {
	log.Println("[INFO] Setting up override for a provider")
	provider := Provider()

	return map[string]*schema.Provider{
		"hellosign": provider,
	}
}

func TestAcc_APIAPP(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers:  accProvider(),
		IsUnitTest: false,
		Steps: []r.TestStep{
			{
				Config: testResourceAPIApp_initialConfig,
			},
			{
				ResourceName:      "hellosign_api_app.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testResourceAPIApp_initialConfig = `
resource "hellosign_api_app" "test" {
  name         = "Terraform Webhook test"
  domain       = "example.com"
  callback_url = "https://example.com/hellosign/callback"
}
`
