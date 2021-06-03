package sign

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/tuotoo/qrcode"
)

type CodeSignParam struct {
	Enc      string
	Name     string
	ActiveID string
	Uid      string
	Fid      string
	AppType  string
}

// Sign is a function of qrcode sign task
func (param *CodeSignParam) Sign(c *http.Client) {
	signURL := fmt.Sprintf("https://mobilelearn.chaoxing.com/pptSign/stuSignajax?enc=%s&name=%s&activeId=%s&uid=%s&clientip=&useragent=&latitude=-1&longitude=-1&fid=%s&appType=%s", param.Enc, url.QueryEscape(param.Name), param.ActiveID, param.Uid, param.Fid, param.AppType)
	resp, err := c.Get(signURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))
}

// Scan qrcode
func (param *CodeSignParam) Scan(fi *os.File) {
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Fatal(err)
		return
	}
	reg, _ := regexp.Compile(`SIGNIN:aid=(\w*)&source=(\w*)&Code=(\w*)&enc=(\w*)`)
	sub := reg.FindStringSubmatch(qrmatrix.Content)
	param.ActiveID = sub[1]
	param.AppType = sub[2]
	param.Enc = sub[4]
}

// DoSign is a function to do qrcode sign
func (param *CodeSignParam) DoSign(c *http.Client, fi *os.File) {
	param.Scan(fi)
	param.Sign(c)
}
