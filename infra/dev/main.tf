module "github" {
  source         = "../modules/github"
  project_id     = var.project_id
  project_region = var.project_region
}

module "backend" {
  source                  = "../modules/backend"
  project_id              = var.project_id
  project_region          = var.project_region
  cloud_run_domain        = var.cloud_run_domain
  cloud_run_service_name  = var.cloud_run_service_name
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

output "frontend_url" {
  value = module.frontend.bucket_url
}

output "github_workload_identity_provider_id" {
  value = module.github.workload_identity_provider_id
}

