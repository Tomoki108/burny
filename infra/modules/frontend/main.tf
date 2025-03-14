resource "google_storage_bucket" "static_website" {
  name          = var.bucket_name
  project       = var.project_id
  location      = var.bucket_location
  force_destroy = true # 開発環境ではバケット削除を容易にするため

  # 静的ウェブサイトの設定
  website {
    main_page_suffix = "index.html"
    not_found_page   = "index.html" # SPAアプリケーションの場合、404はindex.htmlにリダイレクト
  }

  # バケットのACL設定
  uniform_bucket_level_access = true

  # CORS設定（必要に応じて）
  cors {
    origin          = ["*"] # 本番環境では特定のドメインに限定することをお勧めします
    method          = ["GET", "HEAD", "OPTIONS"]
    response_header = ["Content-Type", "Access-Control-Allow-Origin"]
    max_age_seconds = 3600
  }
}

# バケットを公開するためのIAMポリシー設定
resource "google_storage_bucket_iam_binding" "public_access" {
  bucket = google_storage_bucket.static_website.name
  role   = "roles/storage.objectViewer"
  members = [
    "allUsers", # パブリックアクセスを許可
  ]
}

# CDNが有効な場合のロードバランサー設定（オプション）
resource "google_compute_backend_bucket" "static_website_backend" {
  count       = var.enable_cdn ? 1 : 0
  name        = "${var.bucket_name}-backend"
  bucket_name = google_storage_bucket.static_website.name
  enable_cdn  = var.enable_cdn

  # CDNキャッシュの設定
  cdn_policy {
    cache_mode  = "CACHE_ALL_STATIC"
    client_ttl  = 3600
    default_ttl = 3600
    max_ttl     = 86400
  }
}

# 出力変数の定義
output "bucket_url" {
  description = "The URL of the created GCS bucket"
  value       = "https://storage.googleapis.com/${google_storage_bucket.static_website.name}"
}

output "backend_bucket_name" {
  description = "The name of the backend bucket for CDN (if enabled)"
  value       = var.enable_cdn ? google_compute_backend_bucket.static_website_backend[0].name : null
}
