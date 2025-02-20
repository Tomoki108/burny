variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "region" {
  description = "Google Cloud region"
  type        = string
  default     = "asia-northeast1"
}

variable "zone_name" {
  description = "Managed Zone ID (ex, example-com)"
  type        = string
}

variable "dns_name" {
  description = "DNS Zone Name（ex, example.com. suffixed with dot）"
  type        = string
}

variable "description" {
  description = "DNS Zone Description"
  type        = string
  default     = "DNS zone for example.com"
}