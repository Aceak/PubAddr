package tcp

const (
	MagicValue   uint16 = 0xB1F1
	VersionV1    uint8  = 0x01
	OpcodeIPv4   uint16 = 0xA001
	TokenSize           = 16
	ResponseSize        = 8
)

const (
	StatusOK             uint8 = 0
	StatusUnauthorized   uint8 = 1
	StatusRateLimited    uint8 = 2
	StatusInvalidRequest uint8 = 3
	StatusServerError    uint8 = 4
)
