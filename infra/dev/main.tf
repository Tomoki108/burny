provider "google" {
  project = var.project_id
  region  = var.project_region
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

resource "google_service_account" "cloud_run_sa" {
  account_id   = "cloud-run-service"
  display_name = "Cloud Run Service Account"
}

resource "google_project_iam_binding" "cloud_run_sa_iam" {
  for_each = toset([
    "roles/cloudsql.client",
    "roles/secretmanager.secretAccessor",
  ])

  project = var.project_id
  role    = each.value
  members = [
    "serviceAccount:${google_service_account.cloud_run_sa.email}"
  ]
}

resource "google_service_account" "github_actions_sa" {
  account_id   = "github-actions-service"
  display_name = "GitHub Actions Service Account"
}

resource "google_project_iam_binding" "github_actions_sa_iam" {
  for_each = toset([
    "roles/secretmanager.secretAccessor",
    "roles/iam.serviceAccountTokenCreator",
    "roles/artifactregistry.createOnPushWriter",
  ])

  project = var.project_id
  role    = each.value
  members = [
    "serviceAccount:${google_service_account.github_actions_sa.email}"
  ]
}

resource "google_iam_workload_identity_pool" "github_pool" {
  project                   = var.project_id
  workload_identity_pool_id = "github-pool"
  display_name              = "My Workload Identity Pool"
  description               = "A pool to federate identities from GitHub Actions"
}

# For GitHub Actions, set the issuer URI to GitHub's OIDC endpoint.
resource "google_iam_workload_identity_pool_provider" "github_provider" {
  project                            = var.project_id
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  provider                           = google
  display_name                       = "GitHub OIDC Provider"
  description                        = "OIDC provider for GitHub Actions federation"

  attribute_mapping = {
    "google.subject"             = "assertion.sub"
    "attribute.repository"       = "assertion.repository"
    "attribute.repository_owner" = "assertion.repository_owner"

  }
  attribute_condition = "attribute.repository == assertion.repository && attribute.repository_owner == assertion.repository_owner"

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}


locals {
  backend_secret_ids = ["db_name", "db_user", "db_password", "db_instance_connection_name"]
}

resource "google_secret_manager_secret" "backend-secrets" {
  for_each  = { for idx, id in local.backend_secret_ids : idx => id }
  secret_id = each.value
  replication {
    auto {}
  }
}
