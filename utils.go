package goarxml

import "encoding/json"

func ToJson(data interface{}) string {
    bstr, _ := json.MarshalIndent(data,"", "")
    return string(bstr)
}
