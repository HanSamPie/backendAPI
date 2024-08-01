terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.39.1"
    }
  }
}
provider "google" {
  #project = "acme-app"
  region = "europe-west10-b"
}