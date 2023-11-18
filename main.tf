terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

variable "region" {
  type    = string
  default = "asia-east2"
}

variable "zone" {
  type    = string
  default = "asia-east2-a"
}

variable "project_id" {
  type    = string
  default = "book-api-405216"
}

provider "google" {
  credentials = file("gcd_key.json")

  project = var.project_id
  region  = var.region
  zone    = var.zone
}

resource "google_sql_database_instance" "postgres_instance" {
  name             = "postgres-instance"
  database_version = "POSTGRES_12"
  project          = var.project_id
  region           = var.region
  deletion_protection = false

  # Allow external connections
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled = true
      authorized_networks {
        name  = "allow-compute-engine"
        value = "0.0.0.0/0"  # Replace with a specific IP range for security
      }
    }
  }
}

resource "google_sql_database" "book_api_db" {
  name     = "book-api-db"
  instance = google_sql_database_instance.postgres_instance.name
}

resource "google_sql_user" "admin_user" {
  name     = "admin"
  instance = google_sql_database_instance.postgres_instance.name
  password = "adminPassword"
}

resource "google_compute_instance" "app_instance" {
  name         = "book-api"
  machine_type = "n2-standard-2"
  zone         = "asia-east2-a"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  depends_on = [google_sql_database.book_api_db]

  network_interface {
    network = "default"
    access_config {
    }
  }

  metadata_startup_script = templatefile("./startup.sh", {
    db_username = google_sql_user.admin_user.name,
    db_password = google_sql_user.admin_user.password,
    db_host = google_sql_database_instance.postgres_instance.ip_address.0.ip_address,
    db_port = 5432,
    db_name = google_sql_database.book_api_db.name
  })
}

resource "google_compute_firewall" "allow-postgres" {
  name    = "allow-postgres"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["5432"]  # Adjust the port if necessary
  }

  source_ranges = ["0.0.0.0/0"]  # Replace with a specific IP range for security
}

resource "google_compute_firewall" "allow-api" {
  name    = "allow-8080"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["8080"]
  }

  source_ranges = ["0.0.0.0/0"]  # Replace with a specific IP range for security
}