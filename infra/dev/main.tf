provider "google" {
  project = var.project_id
  region  = "asia-northeast1"
}

resource "google_artifact_registry_repository" "cloud_run_repo" {
  repository_id = "dev-cloud-run"
  location      = "asia-northeast1"
  description   = "Artifact Registry repository for Cloud Run application images"
  format        = "DOCKER"
}

resource "google_sql_database_instance" "postgres_instance" {
  name             = "dev-postgres-instance"
  database_version = "POSTGRES_13"
  region           = "asia-northeast1"

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "default" {
  name     = "burny_db"
  instance = google_sql_database_instance.postgres_instance.name
}

resource "google_sql_user" "default" {
  name     = "burny_user"
  instance = google_sql_database_instance.postgres_instance.name
  password = var.postgres_password
}

resource "google_service_account" "cloud_run_sa" {
  account_id   = "cloud-run-service"
  display_name = "Cloud Run Service Account"
}

resource "google_project_iam_member" "cloud_run_sa_cloudsql_client" {
  project = var.project_id
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.cloud_run_sa.email}"
}