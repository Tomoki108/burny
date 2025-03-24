variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "project_region" {
  description = "Google Cloud Project region"
  type        = string
}

###################
# backend関連の変数
###################

// NOTE: vscode autocomplete not work with object variable
// https://github.com/hashicorp/vscode-terraform/issues/1855
variable "secrets" {
  description = "Secrets for the project"
  type = object({
    db_name     = string
    db_user     = string
    db_password = string
  })
  sensitive = true
}

variable "api_domain" {
  description = "Cloud Run domain"
  type        = string
}

###################
# frontend関連の変数
###################

variable "frontend_bucket_name" {
  description = "GCS静的ウェブサイトのバケット名"
  type        = string
}

variable "frontend_bucket_location" {
  description = "GCSバケットのロケーション"
  type        = string
  default     = "asia-northeast1"
}

variable "web_domain" {
  description = "ウェブサイトのドメイン名"
  type        = string
}
