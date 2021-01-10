---
page_title: "hellosign_api_app Resource - terraform-provider-hellosign"
subcategory: ""
description: |-
  
---

# Resource `hellosign_api_app`





## Schema

### Required

- **domain** (String) The domain name the ApiApp will be associated with.
- **name** (String) The name you want to assign to the ApiApp.

### Optional

- **callback_url** (String) The URL at which the ApiApp should receive event callbacks.
- **custom_logo_file** (String) An image file to use as a custom logo in embedded contexts.
- **id** (String) The ID of this resource.
- **oauth** (Block List, Max: 1) (see [below for nested schema](#nestedblock--oauth))
- **owner_account** (Block List, Max: 1) (see [below for nested schema](#nestedblock--owner_account))
- **white_labeling_options** (Map of String)

### Read-only

- **client_id** (String)
- **is_approved** (Boolean)

<a id="nestedblock--oauth"></a>
### Nested Schema for `oauth`

Optional:

- **callback_url** (String) The callback URL to be used for OAuth flows. (Required if oauth[scopes] is provided)
- **scopes** (List of String) A list of OAuth scopes to be granted to the app. (Required if oauth[callback_url] is provided)

Read-only:

- **secret** (String, Sensitive)


<a id="nestedblock--owner_account"></a>
### Nested Schema for `owner_account`

Optional:

- **account_id** (String)
- **email_address** (String)


