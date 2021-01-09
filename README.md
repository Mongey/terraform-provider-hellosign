# `terraform-plugin-hellosign`

A [Terraform][1] plugin for managing [Hellosign][2] resources.

## Installation

Download and extract the [latest release](https://github.com/Mongey/terraform-provider-hellosign/releases/latest) to
your [terraform plugin directory][third-party-plugins] (typically `~/.terraform.d/plugins/`) or define the plugin in the required_providers block.

```hcl
terraform {
  required_providers {
    hellosign = {
      source = "Mongey/hellosign"
    }
  }
}
```

## Example

Configure the provider directly, or set the ENV variables `HELLOSIGN_API_KEY`

```hcl
terraform {
  required_providers {
    hellosign = {
      source = "Mongey/hellosign"
    }
  }
}

provider "hellosign" {
  api_key = "abc_123"
}

resource "hellosign_api_app" "example" {
  name         = "My Application"
  domain       = "example.com"
  callback_url = "https://example.com/hellosign/callback"
}
```

[1]: https://www.terraform.io
[2]: https://www.hellosign.com
[third-party-plugins]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
