terraform {
  backend "gcs" {
    bucket = "burny-tfstate-prod"
  }
}
