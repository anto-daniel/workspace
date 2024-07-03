# Generate terraform script to create a new ec2 instance with ssh key pair name attach to it.
provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "example" {
  ami           = "ami-0221ff74a10bad05e"
  instance_type = "t2.micro"
  key_name      = "personal"
}

