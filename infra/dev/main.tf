module "github" {
  source            = "../modules/github"
  project_id        = var.project_id
  project_region    = var.project_region
  github_repository = var.github_repository
}

module "backend" {
  source                  = "../modules/backend"
  project_id              = var.project_id
  project_region          = var.project_region
  cloud_run_domain        = var.cloud_run_domain
  cloud_run_service_name  = var.cloud_run_service_name
  secrets                 = var.secrets
  github_repository       = var.github_repository
  github_actions_sa_email = module.github.github_actions_sa.email
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

# # GitHub Workload Identity Pool ID を出力
# output "github_workload_identity_pool_id" {
#   value = module.github.workload_identity_pool.id
# }

# # GitHub Actions サービスアカウントのメールアドレスを出力
# output "github_actions_sa_email" {
#   value = module.github.github_actions_sa.email
# }
