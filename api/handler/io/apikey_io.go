package io

type CreateAPIKeyResponse struct {
	RawKey string `json:"raw_key"`
}

// APIKeyStatusResponse APIキーの存在状態を表すレスポンス
type APIKeyStatusResponse struct {
	Exists bool `json:"exists"` // APIキーが存在するかどうか
}
