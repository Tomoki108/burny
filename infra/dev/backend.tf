terraform {
  backend "gcs" {
    bucket  = "burny-tfstate"
    prefix  = "dev"
  }
}