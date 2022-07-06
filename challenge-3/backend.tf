terraform {
  backend "s3" {
    bucket = "terraform-store-22112755097"
    key = "network/terraform.tfstate"
    region = "us-east-1"
  }

  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      version = "~> 3.0.0"
    }
  }
}