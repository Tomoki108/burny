package e2e

import "encoding/json"

// レスポンスのJSONから動的なフィールドを削除する.
// デフォルトでid, created_at, updated_atを削除する.
func removeDynamicFields(res []byte, ignoreFields ...string) ([]byte, error) {
	var data interface{}
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}

	switch d := data.(type) {
	case []interface{}:
		removeDynamicFieldsFromArray(d, ignoreFields...)
	case map[string]interface{}:
		removeDynamicFieldsFromObject(d, ignoreFields...)
	}

	return json.MarshalIndent(data, "", "  ")
}

func removeDynamicFieldsFromArray(arr []interface{}, ignoreFields ...string) {
	ignoreFields = append(ignoreFields, "id", "created_at", "updated_at")
	for i, elem := range arr {
		if obj, ok := elem.(map[string]interface{}); ok {
			for _, field := range ignoreFields {
				delete(obj, field)
			}
			arr[i] = obj
		}
	}
}

func removeDynamicFieldsFromObject(obj map[string]interface{}, ignoreFields ...string) {
	ignoreFields = append(ignoreFields, "id", "created_at", "updated_at")
	for _, field := range ignoreFields {
		delete(obj, field)
	}
}
