package cxsign

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"log"
)

type LoginInfo struct {
	phone    string
	password string
}

type User struct {
	uid  string
	name string
	fid  string
}

type nameJSON struct {
	Msg struct {
		Name string `json:"name"`
	} `json:"msg"`
}

// Login for login chaoxing
func (li *LoginInfo) Login(c *http.Client, user *User) (ret bool) {
	loginURL := fmt.Sprintf("https://passport2-api.chaoxing.com/v11/loginregister?code=%s&cx_xxt_passport=json&uname=%s&loginType=1&roleSelect=true", li.password, li.phone)
	resp, err := c.Get(loginURL)
	if err != nil {
		log.Fatal(err)
		ret = false
		return
	}
	// Read response body
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
	user.GetInfo(c, resp.Cookies())
	return
}

// GetName
func GetName(c *http.Client) string {
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

func (u *User) GetInfo(c *http.Client, cookies []*http.Cookie) {
	for _, v := range cookies {
		switch v.Name {
		case "fid":
			u.fid = v.Value
		case "_uid":
			u.uid = v.Value
		}
	}
	u.name = GetName(c)
}
