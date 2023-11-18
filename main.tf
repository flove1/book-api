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

  project = var.project_id
  region  = var.region
  zone    = var.zone
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