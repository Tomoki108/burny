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

// This is a DNS record for Google Search Console verification (needed to use domain with resources)
resource "google_dns_record_set" "ownership_proof_txt" {
  name         = "burny.page."
  managed_zone = google_dns_managed_zone.zone.name
  type         = "TXT"
  ttl          = 300
  rrdatas      = [var.ownership_proof_txt_rrdata]
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

resource "google_dns_record_set" "prod_api_cname" {
  name         = var.prod_api_cname_name
  managed_zone = google_dns_managed_zone.zone.name
  type         = "CNAME"
  ttl          = 300
  rrdatas      = ["ghs.googlehosted.com."]
}

data "terraform_remote_state" "prod_state" {
  backend = "gcs"
  config = {
    bucket = var.prod_terraform_state_bucket
  }
}

resource "google_dns_record_set" "prod_web_a_record" {
  name         = var.prod_web_a_name
  managed_zone = google_dns_managed_zone.zone.name
  type         = "A"
  ttl          = 300
  rrdatas      = [data.terraform_remote_state.prod_state.outputs.website_ip]
}


####################
# mailer (AWS SES)
####################

resource "google_dns_record_set" "mailer_records" {
  for_each     = { for idx, record in var.mailer_records : idx => record }
  name         = each.value.name
  managed_zone = google_dns_managed_zone.zone.name
  type         = each.value.type
  ttl          = 300
  rrdatas      = [each.value.rrdata]
}

