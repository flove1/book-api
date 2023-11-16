terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

provider "google" {
  credentials = file("gcd_key.json")

  project = "book-api-405216"
  region  = "asia-east2"
  zone    = "asia-east2-a"
}

resource "google_compute_instance" "app_instance" {
  name         = "app-instance"
  machine_type = "n1-standard-1"
  zone         = "asia-east2-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    network = "default"
    access_config {
      // Ephemeral IP
    }
  }

  metadata_startup_script = <<-SCRIPT
    #!/bin/bash
    # Install Docker
    apt-get update
    apt-get install -y docker.io

    # Install Docker Compose
    apt-get install -y docker-compose

    # Pull and run your Docker Compose project
    git clone https://github.com/flove1/book-api
    cd book-api
    docker-compose up -d
  SCRIPT
}