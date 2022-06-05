variable "region" {
  type = string
  default = "us-east-1"
}

variable "terraform_store_bucket" {
  type = string
  default = "terraform-store-22112755097"
}

variable "terraform_store_key" {
  type = string
  default = "network/terraform.tfstate"
}