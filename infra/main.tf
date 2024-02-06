terraform {
  required_version = ">= 0.13"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

data "aws_ami" "ubuntu_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}
resource "aws_internet_gateway" "my_internet_gateway" {
  vpc_id = aws_vpc.my_vpc.id

  tags = {
    Name = "my-internet-gateway"
  }
}

resource "aws_route_table" "my_route_table" {
  vpc_id = aws_vpc.my_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.my_internet_gateway.id
  }

  tags = {
    Name = "my-route-table"
  }
}

resource "aws_route_table_association" "my_route_table_association" {
  subnet_id      = aws_subnet.my_subnet.id
  route_table_id = aws_route_table.my_route_table.id
}

resource "aws_vpc" "my_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = {
    project = "Kataterm"
  }
}

resource "aws_subnet" "my_subnet" {
  vpc_id                  = aws_vpc.my_vpc.id
  cidr_block              = "10.0.0.0/24"
  map_public_ip_on_launch = true
  tags = {
    project = "Kataterm"
  }
}

resource "aws_security_group" "my_security_group" {
  name        = "my-security-group"
  description = "Allow inbound SSH and HTTP traffic"
  vpc_id      = aws_vpc.my_vpc.id
  tags = {
    project = "Kataterm"
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["152.203.164.230/32"]
  }

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"]
  }
  egress {
    from_port   = 53
    to_port     = 53
    protocol    = "udp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 53
    to_port     = 53
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 80
    to_port     = 65000
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

resource "aws_eip" "my_eip" {
  instance = aws_instance.my_ec2_instance[0].id
}

resource "aws_eip_association" "my_eip_association" {
  instance_id   = aws_instance.my_ec2_instance[0].id
  allocation_id = aws_eip.my_eip.id
}
resource "aws_instance" "my_ec2_instance" {
  count                  = 3
  ami                    = "ami-029bdee89471523f0"
  instance_type          = "t2.medium"
  key_name               = "my-key-pair"
  subnet_id              = aws_subnet.my_subnet.id
  vpc_security_group_ids = [aws_security_group.my_security_group.id]
  user_data              = <<-EOF
#!/bin/bash

yum install amazon-linux-extras -y docker
yum install -y git
echo "User-data commands executed successfully."
  EOF
  root_block_device {
    volume_size = 40
  }

  tags = {
    project = "Kataterm"
    Name    = "my-ec2-instance-${count.index + 1}"
  }

  disable_api_termination = true
}

output "ssh_connect_command" {
  value      = "ssh -i ~/.ssh/suretro ec2-user@${aws_instance.my_ec2_instance[0].public_ip}"
  depends_on = [aws_instance.my_ec2_instance]
}



resource "aws_key_pair" "my_key_pair" {
  key_name   = "my-key-pair"
  public_key = file("~/.ssh/suretro.pub")
  tags = {
    project = "Kataterm"
  }
}
