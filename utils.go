package goarxml

import (
	"encoding/json"
	"strings"
)

func ToJson(data interface{}) string {
	bstr, _ := json.MarshalIndent(data, "", "")
	return string(bstr)
}

func GetLastName(path string) string {
	if len(path) == 0 {
		return ""
	}
	parts := strings.Split(path, "/")
	partsLength := len(parts)
	return parts[partsLength-1]
}
