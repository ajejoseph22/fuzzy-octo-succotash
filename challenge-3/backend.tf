terraform {
  backend "s3" {
    bucket = var.terraform_store_bucket
    key = var.terraform_store_key
#    todo: docs include this but buckets are global. do i really need?
#    https://www.terraform.io/language/settings/backends/s3
#    region = "us-east-1"
  }

  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      version = "~> 3.0.0"
    }
  }
}

data "terraform_remote_state" "networking" {
  backend = "s3"
  config {
    bucket = var.terraform_store_bucket
    key    = var.terraform_store_key
#    todo: docs include this but buckets are global. do i really need?
#    https://www.terraform.io/language/settings/backends/s3
#    region = "us-east-1"
  }
}