variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "project_region" {
  description = "Google Cloud region"
  type        = string
}


####################
# dev environment
####################

variable "dev_api_cname_name" {
  description = "CNAME record name for dev api"
  type        = string
}

variable "dev_web_a_name" {
  description = "A record name for dev web"
  type        = string
}

variable "dev_terraform_state_bucket" {
  description = "Terraform state bucket for dev environment"
  type        = string
}
