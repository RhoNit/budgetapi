#!bin/bash

export AWS_ACCESS_KEY_ID="${DB_EC2_AWS_ACCESS_KEY}"
export AWS_SECRET_ACCESS_KEY="${DB_EC2_AWS_SECRET_KEY}"
export AWS_REGION="${DB_EC2_AWS_REGION}"

# Add PostgreSQL repository for the latest version
sudo amazon-linux-extras enable postgresql14

sudo yum update -y
sudo yum install -y docker curl cat postgresql #postgresql-server postgresql-contrib

sudo systemctl enable docker.service
sudo systemctl start docker.service

sudo usermod -a -G docker ec2-user

docker run --name budget_db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=budget \
  -p 5432:5432 \
  --restart unless-stopped \
  -d postgres:alpine

echo "waiting for PostgreSQL to start..."
until docker exec budget_db pg_isready -U postgres; do
  sleep 2
done
echo "PostgreSQL is ready..."

psql --version
