provider "google" {
  project     = var.project_id
  region      = var.project_region
}

resource "google_dns_managed_zone" "zone" {
  name        = var.zone_name 
  dns_name    = var.dns_name
  description = var.zone_description
  project     = var.project_id
}

resource "google_dns_record_set" "dev_subdomain_cname" {
  name         = var.dns_subdomain_name_dev
  managed_zone = google_dns_managed_zone.zone.name
  type         = "CNAME"
  ttl          = 300
  rrdatas      = ["ghs.googlehosted.com."]
}