resource "google_compute_instance" "jenkins" {
  name         = "jenkins-instance"
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
      // Ephemeral IP
    }
  }

  metadata_startup_script = <<-SCRIPT
    #!/bin/bash
    
    sudo wget -O /usr/share/keyrings/jenkins-keyring.asc \
        https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key
    echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
        https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
        /etc/apt/sources.list.d/jenkins.list > /dev/null
    sudo apt-get update

    sudo apt-get install -y default-jre wget
    sudo apt-get install -y jenkins
    sudo apt-get install -y golang
    sudo apt-get install -y docker
    sudo apt-get install -y docker-compose

    # Start Jenkins service
    systemctl start jenkins
    systemctl enable --now jenkins
  SCRIPT
}