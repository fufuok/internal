# Go-Internal

常用的内部包代码集, 按需摘取使用.

## JSON

`json` 可选的组织方式:

- `go build .` 默认使用 `goccy/go-json`
- `go build -tags=std_json .` 使用标准 JSON 库 `encoding/json`
- `go build -tags=sonic .` 使用 `bytedance/sonic`
- 若 gin 使用 `goccy/go-json`, 程序自身使用 `bytedance/sonic`, 则: `go build -tags='sonic go_json'`

## Zstandard

`zstd` 压缩/解压速度与压缩率俱佳的实时压缩算法, 纯 Go 和 CGO 实现, 不兼容官方 gzip.

```go
package zstd // import "github.com/fufuok/xy-data-router/internal/zstd"

Package zstd Ref: VictoriaMetrics/lib/encoding/zstd/

func Compress(dst, src []byte) []byte
func CompressLevel(dst, src []byte, compressionLevel int) []byte
func Decompress(dst, src []byte) ([]byte, error)
```







*ff*