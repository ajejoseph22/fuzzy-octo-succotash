resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    Terraform = "true"
    Name      = "main VPC"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
  tags = {
    Terraform = "true",
    Name      = "main IGW"
  }
}

resource "aws_eip" "nat_eip" {
  vpc = true
  tags = {
    Terraform = "true",
    Name      = "main NAT EIP"
  }
}

resource "aws_nat_gateway" "main" {
  subnet_id = aws_subnet.public.id
  allocation_id = aws_eip.nat_eip.id
  tags = {
    Terraform = "true",
    Name      = "main NAT GW"
  }
}

data "aws_availability_zones" "azs" {}

resource "aws_subnet" "private" {
  vpc_id                  = aws_vpc.main.id
  cidr_block              = var.private_cidr
  availability_zone       = data.aws_availability_zones.azs.names[0]
  map_public_ip_on_launch = false
  tags = {
    Terraform = "true",
    Name      = "private subnet"
  }
}

resource "aws_subnet" "public" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.public_cidr
  availability_zone = data.aws_availability_zones.azs.names[0]
  tags = {
    Terraform = "true",
    Name      = "public subnet"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_nat_gateway.main.id
  }
}

resource "aws_route_table_association" "public" {
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.public.id
}

resource "aws_route_table_association" "private" {
  route_table_id = aws_route_table.private.id
  subnet_id      = aws_subnet.private.id
}