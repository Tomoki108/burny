output "github_actions_sa" {
  description = "The GitHub Actions service account"
  value       = google_service_account.github_actions_sa
}

output "workload_identity_pool" {
  description = "The Workload Identity Pool for GitHub Actions"
  value       = google_iam_workload_identity_pool.github_pool
}

output "workload_identity_provider_id" {
  description = "The Workload Identity Provider ID for GitHub Actions"
  value       = google_iam_workload_identity_pool_provider.github_provider.id
}
