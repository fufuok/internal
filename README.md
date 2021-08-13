# Go-Internal

常用的内部包代码集, 按需摘取使用

## JSON

`json` 使用 `gin` 类似的可选组织方式:

- `go build .` 默认使用 `json-iterator/go`
- `go build -tags=stdjson .` 使用标准 JSON 库 `encoding/json`
- `go build -tags=gojson .` 使用 `goccy/go-json`







*ff*