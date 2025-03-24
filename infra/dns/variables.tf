variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "project_region" {
  description = "Google Cloud region"
  type        = string
}

variable "ownership_proof_txt_rrdata" {
  description = "Google Search Console ownership proof TXT record data"
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

####################
# prod environment
####################

variable "prod_api_cname_name" {
  description = "CNAME record name for prod api"
  type        = string
}

variable "prod_web_a_name" {
  description = "A record name for prod web"
  type        = string
}

variable "prod_terraform_state_bucket" {
  description = "Terraform state bucket for prod environment"
  type        = string
}

####################
# mailer (AWS SES)
####################

variable "mailer_records" {
  description = "dns records for mailer"
  type = list(object({
    name   = string
    type   = string
    rrdata = string
  }))
}
