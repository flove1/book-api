terraform {
  backend "gcs" {
    bucket  = "book-api-terraform-state-bucket"
    prefix  = "terraform/state-file.tfstate"
  }
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

variable "gcp_key_path" {
  description = "Path to the key file of GCP"
}

provider "google" {
  credentials = file(var.gcp_key_path)

  project = var.project_id
  region  = var.region
  zone    = var.zone
}