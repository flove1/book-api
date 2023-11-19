resource "google_compute_instance" "app_instance" {
  name         = "book-api"
  machine_type = "n2-standard-2"
  zone         = var.zone

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
    ports    = ["5432"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_firewall" "allow-api" {
  name    = "allow-api"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["8080"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_firewall" "allow-node-exporter" {
  name    = "allow-node-exporter"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["9100"]
  }

  source_ranges = ["0.0.0.0/0"]
}