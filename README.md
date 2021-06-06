## Usage

```go
import (
    "github.com/Secriy/cxsign"
)

// create a http client
c := cxsign.NewClient()
// add a account
acc := cxsign.NewAccount("phone", "password")
// account login
if acc.Login(c) {
    fmt.Println("登录成功")
    // sign by text content of the QR Code
    acc.DoQrCodeSign(c, "SIGNIN:aid=x00000xxxxxxx&source=xx&Code=xxx000xxxxxxx&enc=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
    // sign by image of the QR Code
    file, _ := os.Open("qrcode.png")
    acc.DoQrCodeSign(c, file)
}
```
