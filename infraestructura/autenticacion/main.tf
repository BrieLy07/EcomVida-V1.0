provider "aws" {
  region = "us-east-1"
}

resource "aws_vpc" "vida_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "vida-vpc"
  }
}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.vida_vpc.id
  tags = {
    Name = "vida-gateway"
  }
}

resource "aws_subnet" "vida_subnet" {
  vpc_id     = aws_vpc.vida_vpc.id
  cidr_block = "10.0.1.0/24"
  availability_zone = "us-east-1a"

  tags = {
    Name = "vida-subnet"
  }
}

resource "aws_security_group" "vida_sg" {
  name        = "vida-sg"
  description = "Allow HTTP and SSH"
  vpc_id      = aws_vpc.vida_vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
