variable "project_id" {
  type = string
}
variable "project_region" {
  type = string
}
variable "cloud_run_domain" {
  type = string
}
variable "cloud_run_service_name" {
  type = string
}
variable "secrets" {
  description = "DB secrets in an object with db_name, db_user and db_password."
  type = object({
    db_name : string,
    db_user : string,
    db_password : string
  })
}
variable "github_repository" {
  type = string
}
