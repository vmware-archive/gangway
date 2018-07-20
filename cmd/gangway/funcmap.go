package main

import (
	"encoding/base64"
	"text/template"
)

var genericMap = map[string]interface{}{
	"base64enc": base64encode,
}

func FuncMap() template.FuncMap {
	return template.FuncMap(genericMap)
}

func base64encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}
