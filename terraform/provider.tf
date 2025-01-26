provider "aws" {
    access_key = var.db_aws_access_key
    secret_key = var.db_aws_secret_key
    region = var.db_aws_region
}

data "aws_ami" "amazon" {
    most_recent = true
    filter {
        name = "name"
        values = ["amzn2-ami-hvm*"]
    }

    filter {
        name = "virtualization-type"
        values = ["hvm"]
    }

    owners = ["amazon"]
}