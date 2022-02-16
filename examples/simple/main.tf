terraform {
  required_providers {
    tls-utils = {
      version = "0.1.0"
      source  = "hashicorp.com/ekristen/tls-utils"
    }
  }
}

provider "tls-utils" {}

data "host_thumbprint" "github" {
  address  = "github.com"
  provider = tls-utils
}

output "thumbprint_sha1" {
  value = data.host_thumbprint.github.id
}

output "thumbprint_lower" {
  value = data.host_thumbprint.github.lower
}
