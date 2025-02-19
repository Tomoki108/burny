variable "project_id" {
  description = "Google Cloud Project ID"
  type        = string
}

variable "project_region" {
  description = "Google Cloud Project region"
  type        = string
}

// NOTE: vscode autocomplete not work with object variable
// https://github.com/hashicorp/vscode-terraform/issues/1855
variable "secrets" {
  description = "Secrets for the project"
  type = object({
    db_name                     = string
    db_user                     = string
    db_password                 = string
    db_instance_connection_name = string
  })
  sensitive = true
}

variable "github_repository" {
  description = "GitHub repository name"
  type        = string
}

variable "github_repository_owner" {
  description = "GitHub repository owner"
  type        = string
}
