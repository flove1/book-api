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