provider "google" {
  project = var.project_id
  region  = var.project_region
}

resource "google_dns_managed_zone" "zone" {
  name        = "burny-page"
  dns_name    = "burny.page."
  description = "DNS zone for burny.page."
  project     = var.project_id
}

####################
# dev environment
####################

resource "google_dns_record_set" "dev_api_cname" {
  name         = var.dev_api_cname_name
  managed_zone = google_dns_managed_zone.zone.name
  type         = "CNAME"
  ttl          = 300
  rrdatas      = ["ghs.googlehosted.com."]
}

data "terraform_remote_state" "dev_state" {
  backend = "gcs"
  config = {
    bucket = var.dev_terraform_state_bucket
    prefix = "dev"
  }
}

resource "google_dns_record_set" "dev_web_a_record" {
  name         = var.dev_web_a_name
  managed_zone = google_dns_managed_zone.zone.name
  type         = "A"
  ttl          = 300
  rrdatas      = [data.terraform_remote_state.dev_state.outputs.website_ip]
}

####################
# prod environment
####################


####################
# mailer (AWS SES)
####################

resource "google_dns_record_set" "mailer_records" {
  for_each     = { for idx, record in var.mailer_records : idx => record }
  name         = each.value.name
  managed_zone = google_dns_managed_zone.zone.name
  type         = each.value.type
  ttl          = 300
  rrdatas      = each.value.rrdatas
}

