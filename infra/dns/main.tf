provider "google" {
  project     = var.project_id
  region      = var.region
}

resource "google_dns_managed_zone" "zone" {
  name        = var.zone_name 
  dns_name    = var.dns_name
  description = var.description
  project     = var.project_id
}