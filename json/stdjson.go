//go:build !jsoniter && !go_json
// +build !jsoniter,!go_json

package json

import (
	"encoding/json"
)

var (
	Name          = "std_json"
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
)
