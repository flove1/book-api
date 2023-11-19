resource "google_compute_instance" "monitoring" {
  name         = "monitoring"
  machine_type = "n2-standard-2"
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
    }
  }

  network_interface {
    network = "default"
    access_config {
    }
  }

  metadata_startup_script = templatefile("./monitoring.sh", {
    app_ip = google_compute_instance.app_instance.network_interface.0.access_config.0.nat_ip
  })
}

resource "google_compute_firewall" "allow-grafana" {
  name    = "allow-grafana"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["3000"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_compute_firewall" "allow-prometheus" {
  name    = "allow-prometheus"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["9090"]
  }

  source_ranges = ["0.0.0.0/0"]
}