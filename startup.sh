#!/bin/bash
apt-get update
apt-get install -y docker.io

newgrp docker

# Pull and run Docker project
git clone https://github.com/flove1/book-api
cd book-api

docker build -t book-api .
docker run \
    -e ADMIN_PASSWORD=$ADMIN_PASSWORD' \
    -e AUTH_KEY=$AUTH_KEY' \
    -e DB_USERNAME='${db_username}' \
    -e DB_PASSWORD='${db_password}' \
    -e DB_HOST='${db_host}' \
    -e DB_PORT='${db_port}' \
    -e DB_NAME='${db_name}' \
    --name book-api-container \
    -p 80:8080 \
    -v ./configs:/app/configs book-api