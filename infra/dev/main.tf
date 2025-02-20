provider "google" {
  project = var.project_id
  region  = var.project_region
}

locals {
  backend_secret_ids = ["db_name", "db_user", "db_password", "db_instance_connection_name"]
}

resource "google_artifact_registry_repository" "cloud_run_repo" {
  repository_id = "dev-cloud-run"
  location      = var.project_region
  description   = "Artifact Registry repository for Cloud Run application images"
  format        = "DOCKER"
}

resource "google_sql_database_instance" "postgres_instance" {
  name             = "dev-postgres-instance"
  database_version = "POSTGRES_13"
  region           = var.project_region

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "default" {
  name     = var.secrets.db_name
  instance = google_sql_database_instance.postgres_instance.name
}

resource "google_sql_user" "default" {
  name     = var.secrets.db_user
  instance = google_sql_database_instance.postgres_instance.name
  password = var.secrets.db_password
}

locals {
  iam_assignments = {
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
    # Cloud Run Service Account
    "${google_service_account.cloud_run_sa.email}_roles_run_admin" = {
      role   = "roles/run.admin"
      member = "serviceAccount:${google_service_account.cloud_run_sa.email}"
    },
    "${google_service_account.cloud_run_sa.email}_roles_cloudsql" = {
      role   = "roles/cloudsql.client"
      member = "serviceAccount:${google_service_account.cloud_run_sa.email}"
    },
    "${google_service_account.cloud_run_sa.email}_roles_secretmanager" = {
      role   = "roles/secretmanager.secretAccessor"
      member = "serviceAccount:${google_service_account.cloud_run_sa.email}"
    },
  }
}

resource "google_service_account" "cloud_run_sa" {
  account_id   = "cloud-run-service"
  display_name = "Cloud Run Service Account"
}

resource "google_service_account" "github_actions_sa" {
  account_id   = "github-actions-service"
  display_name = "GitHub Actions Service Account"
}

resource "google_project_iam_member" "iam_member" {
  for_each = local.iam_assignments
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

resource "google_service_account_iam_member" "github_actions_act_as_cloud_run_sa" {
  # github_actions_saに、cloud_run_saの権限を代理実行できる権限を付与
  member             = "serviceAccount:${google_service_account.github_actions_sa.email}"
  service_account_id = google_service_account.cloud_run_sa.name
  role               = "roles/iam.serviceAccountUser"
}


resource "google_iam_workload_identity_pool" "github_pool" {
  project                   = var.project_id
  workload_identity_pool_id = "github-pool"
  display_name              = "My Workload Identity Pool"
  description               = "A pool to federate identities from GitHub Actions"
}

resource "google_iam_workload_identity_pool_provider" "github_provider" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  display_name                       = "GitHub OIDC Provider"
  description                        = "OIDC provider for GitHub Actions federation"

  attribute_mapping = {
    "google.subject" = "assertion.sub"
  }
  attribute_condition = "assertion.repository == '${var.github_repository}'"

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

resource "google_secret_manager_secret" "backend-secrets" {
  for_each  = { for idx, id in local.backend_secret_ids : idx => id }
  secret_id = each.value
  replication {
    auto {}
  }
}