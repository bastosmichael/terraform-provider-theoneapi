# The One API Terraform Provider (Terraform Plugin Framework)

This repository contains a Terraform provider that allows you to manage resources through the One API.

## Prerequisites

- [Go](https://golang.org/doc/install) >= 1.16
- [Terraform](https://www.terraform.io/downloads.html) >= 0.14.x
- Access to the One API.

## Installing The Provider

1. Clone the repository:
   ```bash
   git clone https://github.com/bastosmichael/terraform-provider-theoneapi.git
   ```

2. Move to the directory:
   ```bash
   cd terraform-provider-theoneapi
   ```

3. Build the provider:
   ```bash
   go build -o terraform-provider-theoneapi
   ```

4. Move the provider to your plugins directory:
   ```bash
   mkdir -p ~/.terraform.d/plugins/example.com/user/theoneapi/0.1.0/linux_amd64
   mv the-one-api-terraform-provider ~/.terraform.d/plugins/example.com/user/theoneapi/0.1.0/linux_amd64
   ```

## Setting Up The Provider

1. Configure the provider:

   In your Terraform configuration, reference the provider and supply the necessary credentials:

   ```hcl
   provider "theoneapi" {
     api_endpoint = "https://the-one-api.dev/"
     api_token    = "YOUR_API_TOKEN"
   }
   ```

## Running The Provider

To plan and apply your Terraform configuration:

1. Initialize your configuration:

   ```bash
   terraform init
   ```

2. Plan your changes:

   ```bash
   terraform plan
   ```

3. Apply your configuration:

   ```bash
   terraform apply
   ```

## Debugging

If you encounter any issues or unexpected behaviors, you can enable debug mode by setting the environment variable:

```bash
export TF_PROVIDER_DEBUG=true
```

Then run your Terraform commands.

## Running Tests

To run the provider's tests:

```bash
go test ./...
```

## Publishing the Provider

1. Tag your release:

   ```bash
   git tag v0.1.0
   git push --tags
   ```

2. Build a release binary for your platform:

   ```bash
   GOOS=linux GOARCH=amd64 go build -o the-one-api-terraform-provider_v0.1.0
   ```

3. Upload the binary to the GitHub release or any other distribution method you prefer.

Note: This is a basic guide. If your provider gains traction in the community, consider registering it with Terraform's public provider registry.

---

This README provides a basic overview for a Terraform provider. You might need to adjust paths, commands, or details to better fit your actual environment and provider's specifics.