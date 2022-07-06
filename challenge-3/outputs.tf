output "vpc_id" {
  value = resource.aws_vpc.main.id
}

output "public_subnet_id" {
  value = resource.aws_subnet.public.id
}

output "private_subnet_id" {
  value = resource.aws_subnet.private.id
}