terraform {
  # NOTE: prefix（stateファイルを入れるデフォルトディレクトリ）は無しで、バケット名をburny-tfstate-devにした方が良かった
  backend "gcs" {
    bucket = "burny-tfstate"
    prefix = "dev"
  }
}
