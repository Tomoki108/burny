module "backend" {
  source                 = "../modules/backend"
  project_id             = var.project_id
  project_region         = var.project_region
  cloud_run_domain       = var.cloud_run_domain
  cloud_run_service_name = var.cloud_run_service_name
  secrets                = var.secrets
  github_repository      = var.github_repository
}