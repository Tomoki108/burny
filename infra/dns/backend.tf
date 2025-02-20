terraform {
  backend "gcs" {
    bucket = "burny-dns-tfstate"
  }
}