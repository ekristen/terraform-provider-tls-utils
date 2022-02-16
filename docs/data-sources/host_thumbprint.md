---
page_title: "host_thumbprint Data Source - terraform-provider-tls-utils"
subcategory: ""
description: |-
  The host_thumbprint data source allows you to obtain the certificate thumbprint
---

# Data Source `host_thumbprint`

The host_thumbprint data source allows you to obtain the certificate thumbprint

## Example Usage

```terraform
data "tls-utils-host_thumbprint" "github" {
  address  = "github.com"
}
```

## Argument Reference

- `address` - (Required) The address of the host to extract the thumbprint from.
- `port` - [default: 443] (Optioanl) The port to connect to on the host.
- `insecure` - [default: true] (Optional) Boolean that can be set to true to disable SSL certificate verification.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `id` - The fingerprint of the host certificate.
- `sha1` - The SHA1 hash of the certificate.
- `md5` - The MD5 hash of the certificate.
