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
