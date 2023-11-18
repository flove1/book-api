#!/bin/bash
apt-get update
apt-get install -y docker.io

newgrp docker

# Set enviroment variables:
export ADMIN_PASSWORD=`curl \
    -H "Metadata-flavor: Google" \
    "http://metadata.google.internal/computeMetadata/v1/project/attributes/ADMIN_PASSWORD"`
export AUTH_KEY=`curl \ 
    -H "Metadata-flavor: Google" \
    "http://metadata.google.internal/computeMetadata/v1/project/attributes/AUTH_KEY"`

# Run migrations
export DB_URL='postgresql://${db_username}:${db_password}@${db_host}:${db_port}/${db_name}'
docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database $DB_URL  up

# Pull and run Docker project
git clone https://github.com/flove1/book-api
cd book-api

docker pull migrate/migrate

docker build -t book-api .
docker run \
    -e ADMIN_PASSWORD='$ADMIN_PASSWORD' \
    -e AUTH_KEY='$AUTH_KEY' \
    -e DB_USERNAME='${db_username}' \
    -e DB_PASSWORD='${db_password}' \
    -e DB_HOST='${db_host}' \
    -e DB_PORT='${db_port}' \
    -e DB_NAME='${db_name}' \
    --name book-api-container \
    -p 8080:8080 \
    -v ./configs:/app/configs book-api