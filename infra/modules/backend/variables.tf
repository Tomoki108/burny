variable "project_id" {
  type = string
}

variable "project_region" {
  type = string
}

variable "api_domain" {
  type = string
}

variable "secrets" {
  type = object({
    db_name : string,
    db_user : string,
    db_password : string
  })
}

variable "github_actions_sa_email" {
  type = string
}
