provider "google" {
  project = var.project_id
  region  = var.project_region
}

resource "google_dns_managed_zone" "zone" {
  name        = var.zone_name
  dns_name    = var.dns_name
  description = var.zone_description
  project     = var.project_id
}

####################
# dev 環境
####################

# APIサーバー用のCNAMEレコード
resource "google_dns_record_set" "dev_api_cname" {
  name         = "dev-api.${var.dns_name}"
  managed_zone = google_dns_managed_zone.zone.name
  type         = "CNAME"
  ttl          = 300
  rrdatas      = ["ghs.googlehosted.com."]
}

data "terraform_remote_state" "dev_state" {
  backend = "gcs"
  config = {
    bucket = var.dev_terraform_state_bucket
    prefix = "dev" // dev環境のステートファイルのプレフィックスを追加
  }
}

resource "google_dns_record_set" "dev_web_a_record" {
  name         = "dev.${var.dns_name}"
  managed_zone = google_dns_managed_zone.zone.name
  type         = "A"
  ttl          = 300
  rrdatas      = [data.terraform_remote_state.dev_state.outputs.website_ip]
}
