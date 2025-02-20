output "managed_zone_name" {
  description = "created Managed Zone name"
  value       = google_dns_managed_zone.zone.name
}

output "name_servers" {
  description = "name server list provided by Cloud DNS"
  value       = google_dns_managed_zone.zone.name_servers
}