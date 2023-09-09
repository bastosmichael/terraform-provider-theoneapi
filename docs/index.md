---
layout: "theoneapi"
page_title: "Provider: The One API"
description: |-
  The One API provider allows you to manage resources from The One API in your Terraform configuration.
---

# The One API Provider

The One API provider allows you to manage resources from The One API, such as books and movies related to The Lord of the Rings.

## Example Usage

```hcl
provider "theoneapi" {
  api_token = "YOUR_API_TOKEN_HERE"
}

resource "theoneapi_book" "example" {
  name = "The Fellowship of the Ring"
}
