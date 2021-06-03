## 示例

```go
c := NewClient()
// User account
li := LoginInfo{
    phone:    "xxxxx",
    password: "xxxxx",
}
// Login
user := &User{}
if !li.Login(c, user) {
    log.Println("登录失败")
    return
}
// Do qrcode sign
csp := &sign.CodeSignParam{
    Name: user.name,
    Uid:  user.uid,
    Fid:  user.fid,
}
file, _ := os.Open("qrcode.png")
csp.DoSign(c, file)
```
