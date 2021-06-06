package cxsign

import (
	"fmt"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	c := NewClient()
	acc := NewAccount("phone", "password")
	if acc.Login(c) {
		fmt.Println("登录成功")
		// sign by text content of the QR Code
		acc.DoQrCodeSign(c, "SIGNIN:aid=x00000xxxxxxx&source=xx&Code=xxx000xxxxxxx&enc=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		// sign by image of the QR Code
		file, _ := os.Open("code.jfif")
		acc.DoQrCodeSign(c, file)
	}
}
