#!/bin/bash

# Docker installation
apt-get update
apt-get install -y docker.io

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/download/v2.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Pull and run Docker Compose project
git clone https://github.com/flove1/book-api
cd book-api

newgrp docker

# Set enviroment variables:
export ADMIN_PASSWORD=`curl \
    -H "Metadata-flavor: Google" \
    "http://metadata.google.internal/computeMetadata/v1/project/attributes/ADMIN_PASSWORD"`

export AUTH_KEY=`curl \
    -H "Metadata-flavor: Google" \
    "http://metadata.google.internal/computeMetadata/v1/project/attributes/AUTH_KEY"`

export DB_USERNAME="${db_username}"
export DB_PASSWORD="${db_password}"
export DB_HOST="${db_host}"
export DB_PORT="${db_port}"
export DB_NAME="${db_name}"

docker-compose up -d