package e2e

import "encoding/json"

// レスポンスのJSONから動的なフィールドを削除する
// デフォルトでid, created_at, updated_atを削除する
func removeDynamicFields(res []byte, ignoreFields ...string) ([]byte, error) {
	ignoreFields = append(ignoreFields, "id", "created_at", "updated_at")

	var m map[string]interface{}
	if err := json.Unmarshal(res, &m); err != nil {
		return nil, err
	}
	// 除外するフィールドを削除
	for _, field := range ignoreFields {
		delete(m, field)
	}
	return json.MarshalIndent(m, "", "  ")
}
