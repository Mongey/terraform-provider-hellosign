package ccloud

import (
	"context"
	"log"

	"github.com/Mongey/terraform-provider-hellosign/internal/hellosign"
	hs "github.com/StefanNyman/hellosign"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func apiAppResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: apiAppCreate,
		ReadContext:   apiAppRead,
		UpdateContext: apiAppUpdate,
		DeleteContext: apiAppDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name you want to assign to the ApiApp.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "The domain name the ApiApp will be associated with.",
			},
			"callback_url": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "The URL at which the ApiApp should receive event callbacks.",
			},
			"custom_logo_file": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    false,
				Description: "An image file to use as a custom logo in embedded contexts.",
			},
			"owner_account": {
				Type:     schema.TypeList,
				MaxItems: 1,
				ForceNew: false,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Description: "",
							Computed:    true,
							Type:        schema.TypeString,
							Optional:    true,
						},
						"email_address": {
							Description: "",
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
			"oauth": {
				Type:     schema.TypeList,
				MaxItems: 1,
				ForceNew: false,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"callback_url": {
							Description: "The callback URL to be used for OAuth flows. (Required if oauth[scopes] is provided)",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"scopes": {
							Description: "A list of OAuth scopes to be granted to the app. (Required if oauth[callback_url] is provided)",
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        schema.TypeString,
						},
						"secret": {
							Description: "",
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
					},
				},
			},
			"white_labeling_options": {
				Type:     schema.TypeMap,
				ForceNew: false,
				Optional: true,
			},
			"is_approved": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func apiAppCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//c := meta.(*hellosign.Client)
	//var diag diag.Diagnostics

	return nil //diag.FromErr(err)
}

func apiAppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*hellosign.Client)
	id := d.Id()

	app, err := c.App(id)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("client_id", id)
	d.Set("name", app.Name)
	d.Set("callback_url", app.CallbackURL)
	d.Set("domain", app.Domain)
	d.Set("is_approved", app.IsApproved)
	log.Printf("[INFO] oauth: %v", app.Oauth)
	log.Printf("[INFO] app: %v", app)

	if app.Oauth != nil {
		d.Set("oauth", []interface{}{flattenOuathBlock(app)})
	} else {
		d.Set("oauth", []interface{}{})
	}
	d.Set("owner_account", []interface{}{flatternOwnerAccount(app)})

	return nil
}

func apiAppUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//c := meta.(*hellosign.Client)

	//return diag.FromErr(err)
	return nil
}

func apiAppDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*hellosign.Client)
	err := c.DeleteApp(d.Id())

	return diag.FromErr(err)
}

func flatternOwnerAccount(app *hs.APIApp) interface{} {
	return map[string]interface{}{
		"account_id":    app.OwnerAccount.AccountID,
		"email_address": app.OwnerAccount.EmailAddress,
	}
}

func flattenOuathBlock(app *hs.APIApp) interface{} {
	return map[string]interface{}{
		"callback_url": app.Oauth.CallbackURL,
		"scopes":       app.Oauth.Scopes,
		"secret":       app.Oauth.Secret,
	}
}
