data "template_file" "db_ec2_ud" {
    template = file("user-data-db.sh")
    vars = {
        DB_EC2_AWS_ACCESS_KEY = var.db_aws_access_key
        DB_EC2_AWS_SECRET_KEY = var.db_aws_secret_key
        DB_EC2_AWS_REGION = var.db_aws_region
    }
}

resource "aws_vpc" "main_vpc" {
  cidr_block = "10.0.0.0/16"
  enable_dns_support = true
  enable_dns_hostnames = true

  tags = {
    Name = "main_vpc"
  }
}

# Create Public Subnet
resource "aws_subnet" "public_subnet" {
  vpc_id            = aws_vpc.main_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-2a"
  map_public_ip_on_launch = true
  tags = {
    Name = "public_subnet"
  }
}

# Internet Gateway (IGW) for Public Subnet Access
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.main_vpc.id

  tags = {
    Name = "main_igw"
  }
}

resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.main_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "public_route_table"
  }
}

resource "aws_route_table_association" "public_subnet_association" {
  subnet_id      = aws_subnet.public_subnet.id
  route_table_id = aws_route_table.public_route_table.id
}

resource "aws_security_group" "allow_postgres_port" {
  vpc_id = aws_vpc.main_vpc.id

  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
  }

  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
  }

  tags = {
    Name = "allow_postgres_port"
  }
}

resource "aws_instance" "db_ec2_server" {
    ami = data.aws_ami.amazon.id
    instance_type = "t2.micro"
    subnet_id = aws_subnet.public_subnet.id
    security_groups = [aws_security_group.allow_postgres_port.id]
    user_data = data.template_file.db_ec2_ud.rendered
    key_name = "terraform-key"

    tags = {
        Name = "db_ec2_server"
    }
}


output "db_ec2_public_ip" {
  description = "Public IP of the Database EC2 instance"
  value       = aws_instance.db_ec2_server.public_ip
}