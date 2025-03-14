variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "project_region" {
  description = "Google Cloud region"
  type        = string
  default     = "asia-northeast1"
}

variable "zone_name" {
  description = "Managed Zone ID (ex, example-com)"
  type        = string
  default     = "burny-page"
}

variable "zone_description" {
  description = "DNS Zone Description"
  type        = string
  default     = "DNS zone for burny.page."
}

variable "dns_name" {
  description = "DNS Zone Name（ex, example.com. suffixed with dot）"
  type        = string
  default     = "burny.page."
}
