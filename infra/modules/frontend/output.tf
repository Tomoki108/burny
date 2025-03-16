output "website_ip" {
  description = "The global IP address for the custom domain"
  value       = google_compute_global_address.website_ip.address
}
