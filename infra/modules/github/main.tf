provider "google" {
  project = var.project_id
  region  = var.project_region
}

####################
# IAM
####################
resource "google_service_account" "github_actions_sa" {
  account_id   = "github-actions-service"
  display_name = "GitHub Actions Service Account"
}

locals {
  github_iam_assignments = {
    # GitHub Actions Service Account
    "${google_service_account.github_actions_sa.email}_roles_secretmanager" = {
      role   = "roles/secretmanager.secretAccessor"
      member = "serviceAccount:${google_service_account.github_actions_sa.email}"
    },
    "${google_service_account.github_actions_sa.email}_roles_artifactregistry" = {
      role   = "roles/artifactregistry.writer"
      member = "serviceAccount:${google_service_account.github_actions_sa.email}"
    },
    "${google_service_account.github_actions_sa.email}_roles_run_admin" = {
      role   = "roles/run.admin"
      member = "serviceAccount:${google_service_account.github_actions_sa.email}"
    },
    "${google_service_account.github_actions_sa.email}_roles_storage_object" = {
      role   = "roles/storage.objectUser"
      member = "serviceAccount:${google_service_account.github_actions_sa.email}"
    },
  }
}

resource "google_project_iam_member" "github_iam_member" {
  for_each = local.github_iam_assignments
  project  = var.project_id
  member   = each.value.member
  role     = each.value.role
}

resource "google_service_account_iam_member" "github_actions_workload_identity" {
  # workload_identity_poolに、github_actions_saの権限を借用できる権限を付与
  member             = "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_pool.name}/*"
  service_account_id = google_service_account.github_actions_sa.name
  role               = "roles/iam.workloadIdentityUser"
}

resource "google_iam_workload_identity_pool" "github_pool" {
  project                   = var.project_id
  workload_identity_pool_id = "github-actions-pool-01"
  display_name              = "GitHub Actions Pool"
  description               = "A pool to federate identities from GitHub Actions"
}

resource "google_iam_workload_identity_pool_provider" "github_provider" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-actions-provider-01"
  display_name                       = "GitHub OIDC Provider"
  description                        = "OIDC provider for GitHub Actions federation"
  attribute_mapping = {
    "google.subject" = "assertion.sub"
  }
  attribute_condition = "assertion.repository == 'Tomoki108/burny'"
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}
