package tcp

import "fmt"

// MagicHeader
const (
	MagicValue uint16 = 0xB1F1
)

// Version
const (
	VersionV1 uint8 = 0x01
)

// Opcode
const (
	OpcodeIPv4 uint16 = 0xA001 // 返回 IPv4
	OpcodeIPv6 uint16 = 0xA002 // 返回 IPv6
	OpcodeBoth uint16 = 0xA003 // 返回 IPv4 + IPv6
	OpcodeJSON uint16 = 0xA004 // 返回 JSON 格式
)

// TokenSize
const (
	TokenSize = 16
)

// Status
const (
	StatusOK             uint8 = 0 // 正常，返回有效 IPv4
	StatusUnauthorized   uint8 = 1 // token 无效/未授权
	StatusRateLimited    uint8 = 2 // 触发限流策略
	StatusInvalidRequest uint8 = 3 // 请求格式错误、Magic错、Opcode不支持等
	StatusServerError    uint8 = 4 // 服务器内部错误（例如无法解析 IP）
)

// ResponseSize
const ResponseSize = 8 // 响应包大小，固定为 8 字节

var (
	ErrMagic             = fmt.Errorf("invalid magic")
	ErrVersion           = fmt.Errorf("invalid version")
	ErrInvalidIP         = fmt.Errorf("invalid ip")
	ErrRateLimitExceeded = fmt.Errorf("rate limit exceeded")
)
