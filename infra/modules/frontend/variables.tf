variable "project_id" {
  description = "The Google Cloud Project ID"
  type        = string
}

variable "bucket_name" {
  description = "The name of the GCS bucket to create for static website hosting"
  type        = string
}

variable "bucket_location" {
  description = "The location of the GCS bucket"
  type        = string
  default     = "asia-northeast1"
}

variable "website_domain" {
  description = "The domain name for the website (optional, if using custom domain)"
  type        = string
  default     = null
}

variable "enable_cdn" {
  description = "Enable Cloud CDN for the static website"
  type        = bool
  default     = false
}
