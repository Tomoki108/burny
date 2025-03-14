module "backend" {
  source                 = "../modules/backend"
  project_id             = var.project_id
  project_region         = var.project_region
  cloud_run_domain       = var.cloud_run_domain
  cloud_run_service_name = var.cloud_run_service_name
  secrets                = var.secrets
  github_repository      = var.github_repository
}

module "frontend" {
  source          = "../modules/frontend"
  project_id      = var.project_id
  bucket_name     = var.frontend_bucket_name
  bucket_location = var.frontend_bucket_location
  enable_cdn      = var.enable_cdn
}

# フロントエンドのデプロイURLを出力
output "frontend_url" {
  value = module.frontend.bucket_url
}
