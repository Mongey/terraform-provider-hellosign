---
layout: ""
page_title: "Provider: HelloSign"
description: |-
  The HelloSign provider provides resources to interact with a HelloSign API.
---

# HelloSign Provider

The HelloSign provider provides resources to interact with a HelloSign API.

## Example Usage

```terraform
provider "hellosign" {
  api_key = "abc_123" # optionally use HELLOSIGN_API_KEY env var
}
```

## Schema

### Optional

- **api_key** (String, Sensitive)
