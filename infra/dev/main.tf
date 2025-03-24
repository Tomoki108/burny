module "github" {
  source         = "../modules/github"
  project_id     = var.project_id
  project_region = var.project_region
}

module "backend" {
  source                  = "../modules/backend"
  project_id              = var.project_id
  project_region          = var.project_region
  api_domain              = var.api_domain
  web_base_url            = "https://${var.web_domain}"
  enable_db_backup        = false
  secrets                 = var.secrets
  github_actions_sa_email = module.github.github_actions_sa.email
}

module "frontend" {
  source          = "../modules/frontend"
  project_id      = var.project_id
  bucket_name     = var.frontend_bucket_name
  bucket_location = var.frontend_bucket_location
  web_domain      = var.web_domain
}

