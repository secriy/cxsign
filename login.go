package cxsign

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"log"
)

type Account struct {
	Phone    string
	Password string
	Uid      string
	Name     string
	Fid      string
}

type nameJSON struct {
	Msg struct {
		Name string `json:"name"`
	} `json:"msg"`
}

// NewAccount is for create a new account model.
func NewAccount(phone, passwd string) *Account {
	return &Account{
		Phone:    phone,
		Password: passwd,
	}
}

// Login for login chaoxing account.
func (acc *Account) Login(c *http.Client) (ret bool) {
	loginURL := fmt.Sprintf("https://passport2-api.chaoxing.com/v11/loginregister?code=%s&cx_xxt_passport=json&uname=%s&loginType=1&roleSelect=true", acc.Password, acc.Phone)
	resp, err := c.Get(loginURL)
	if err != nil {
		log.Fatal(err)
		ret = false
		return
	}
	// read response body
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		ret = false
		return
	}
	defer resp.Body.Close()
	if strings.Contains(string(s), "验证通过") {
		ret = true
	}
	acc.GetInfo(c, resp.Cookies())
	return
}

// GetInfo is getting the infomation of the logined user.
func (acc *Account) GetInfo(c *http.Client, cookies []*http.Cookie) {
	for _, v := range cookies {
		switch v.Name {
		case "fid":
			acc.Fid = v.Value
		case "_uid":
			acc.Uid = v.Value
		}
	}
	acc.Name = getName(c)
}

// DoQrCodeSign is for running the function of QR Code sign.
func (acc *Account) DoQrCodeSign(c *http.Client, t interface{}) {
	param := &qrCodeParam{
		Fid:  acc.Fid,
		Uid:  acc.Uid,
		Name: acc.Name,
	}
	if v, ok := t.(string); ok {
		param.parseToken(v)
	} else if v, ok := t.(*os.File); ok {
		param.scanQrCode(v)
	} else {
		return
	}
	param.qrCodeSign(c)
}

// GetName is getting the name of the logined user.
func getName(c *http.Client) string {
	userURL := "https://sso.chaoxing.com/apis/login/userLogin4Uname.do"
	resp, err := c.Get(userURL)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	nameJSON := nameJSON{}
	_ = json.Unmarshal(s, &nameJSON)

	return nameJSON.Msg.Name
}
