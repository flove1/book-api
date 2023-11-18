
#!/bin/bash
# Install Dockerapt-get update
apt-get install -y docker.io
apt-get install -y python3

newgrp docker

# Set enviroment variables:
export ADMIN_PASSWORD=curl -H "Metadata-flavor: Google" "http://metadata.google.internal/computeMetadata/v1/project/attributes/ADMIN_PASSWORD"
export AUTH_KEY=curl -H "Metadata-flavor: Google" "http://metadata.google.internal/computeMetadata/v1/project/attributes/AUTH_KEY"
export DB_USERNAME="${db_username}"
export DB_PASSWORD="${db_password}"
export DB_HOST="${db_host}"
export DB_PORT="${db_port}"
export DB_NAME="${db_name}"

# Pull and run Docker project
git clone https://github.com/flove1/book-api
cd book-api

docker build -t book-api .
docker run -d --name book-api-container -p 80:8080 book-api