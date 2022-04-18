package generator

import qrcode "github.com/skip2/go-qrcode"

type QRCodeGenerator interface {
	GenerateQRCode(content string, level qrcode.RecoveryLevel, size int) ([]byte, error)
}

type qrCodeGeneratorImpl struct{}

func NewQRCodeGeneratorImpl() *qrCodeGeneratorImpl {
	return &qrCodeGeneratorImpl{}
}

func (q *qrCodeGeneratorImpl) GenerateQRCode(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	return qrcode.Encode(content, level, size)
}
