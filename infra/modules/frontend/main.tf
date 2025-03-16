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

# ロードバランサー設定
resource "google_compute_backend_bucket" "static_website_backend" {
  project     = var.project_id
  name        = "${var.bucket_name}-backend"
  bucket_name = google_storage_bucket.static_website.name
  enable_cdn  = true

  # CDNキャッシュの設定
  cdn_policy {
    cache_mode  = "CACHE_ALL_STATIC"
    client_ttl  = 3600
    default_ttl = 3600
    max_ttl     = 86400
  }
}

# SSL証明書の作成
resource "google_compute_managed_ssl_certificate" "website_cert" {
  name    = "${var.bucket_name}-cert"
  project = var.project_id

  managed {
    domains = [var.web_domain]
  }
}

# カスタムドメイン用のグローバルIPアドレスの確保
resource "google_compute_global_address" "website_ip" {
  name    = "${var.bucket_name}-ip"
  project = var.project_id
}

# カスタムドメイン用のHTTPSプロキシ
resource "google_compute_target_https_proxy" "website_https_proxy" {
  name             = "${var.bucket_name}-https-proxy"
  url_map          = google_compute_url_map.website_url_map.id
  ssl_certificates = [google_compute_managed_ssl_certificate.website_cert.id]
  project          = var.project_id
}

# URLマップの作成
resource "google_compute_url_map" "website_url_map" {
  name            = "${var.bucket_name}-url-map"
  default_service = google_compute_backend_bucket.static_website_backend.id
  project         = var.project_id

  # SPA向けのカスタムルート設定
  host_rule {
    hosts        = [var.web_domain]
    path_matcher = "spa-routes"
  }

  path_matcher {
    name            = "spa-routes"
    default_service = google_compute_backend_bucket.static_website_backend.id

    # 静的アセットのパスルール
    path_rule {
      paths   = ["/assets/*"]
      service = google_compute_backend_bucket.static_website_backend.id
    }

    # 基本的にはdefault_serviceにすべてのリクエストを渡し、
    # サーバーサイドの静的ウェブサイト設定でindex.htmlへのフォールバックを処理
  }
}

# HTTPからHTTPSへのリダイレクト設定
resource "google_compute_url_map" "http_redirect" {
  name    = "${var.bucket_name}-http-redirect"
  project = var.project_id

  default_url_redirect {
    https_redirect         = true
    redirect_response_code = "MOVED_PERMANENTLY_DEFAULT" # 301リダイレクト
    strip_query            = false
  }
}

# HTTP用プロキシ（HTTPSへリダイレクトするため）
resource "google_compute_target_http_proxy" "website_http_proxy" {
  name    = "${var.bucket_name}-http-proxy"
  url_map = google_compute_url_map.http_redirect.id
  project = var.project_id
}

# HTTPSのグローバルフォワーディングルール
resource "google_compute_global_forwarding_rule" "website_https_rule" {
  name       = "${var.bucket_name}-https-rule"
  target     = google_compute_target_https_proxy.website_https_proxy.id
  ip_address = google_compute_global_address.website_ip.address
  port_range = "443"
  project    = var.project_id
}

# HTTPのグローバルフォワーディングルール（HTTPSへリダイレクト）
resource "google_compute_global_forwarding_rule" "website_http_rule" {
  name       = "${var.bucket_name}-http-rule"
  target     = google_compute_target_http_proxy.website_http_proxy.id
  ip_address = google_compute_global_address.website_ip.address
  port_range = "80"
  project    = var.project_id
}
