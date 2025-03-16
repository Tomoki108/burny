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

variable "web_domain" {
  description = "The domain name for the website"
  type        = string
}
