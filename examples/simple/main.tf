terraform {
  required_providers {
    tls-utils = {
      source = "hashicorp.com/ekristen/tls-utils"
    }
  }
}

provider "tls-utils" {}

data "host_thumbprint" "github" {
  address  = "github.com"
  provider = tls-utils
}

output "thumbprint_id" {
  value = data.host_thumbprint.github.id
}

output "thumbprint_sha1" {
  value = data.host_thumbprint.github.sha1
}

output "thumbprint_md5" {
  value = data.host_thumbprint.github.md5
}
