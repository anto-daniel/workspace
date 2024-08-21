# Generate terraform script to create a new ec2 instance with ssh key pair name attach to it.
provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "example" {
  ami           = "ami-0b0ea68c435eb488d"
  instance_type = "t2.micro"
  key_name      = "personal"
}

