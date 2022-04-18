package generator

import qrcode "github.com/skip2/go-qrcode"

func Generate(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	return qrcode.Encode(content, level, size)
}
