output "bucket_url" {
  description = "The URL of the created GCS bucket"
  value       = "https://storage.googleapis.com/${google_storage_bucket.static_website.name}"
}

output "backend_bucket_name" {
  description = "The name of the backend bucket for CDN (if enabled)"
  value       = google_compute_backend_bucket.static_website_backend.name
}

output "custom_domain_ip" {
  description = "The global IP address for the custom domain (if enabled)"
  value       = google_compute_global_address.website_ip.address
}

output "custom_domain" {
  description = "The custom domain for the website (if enabled)"
  value       = var.web_domain
}
