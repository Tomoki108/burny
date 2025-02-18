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