//go:build sonic
// +build sonic

package json

import (
	"github.com/bytedance/sonic"
)

var (
	Name          = "sonic"
	Marshal       = sonic.ConfigStd.Marshal
	Unmarshal     = sonic.ConfigStd.Unmarshal
	MarshalIndent = sonic.ConfigStd.MarshalIndent
	NewDecoder    = sonic.ConfigStd.NewDecoder
	NewEncoder    = sonic.ConfigStd.NewEncoder
)
