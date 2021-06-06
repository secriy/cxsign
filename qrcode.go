package cxsign

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

type qrCodeParam struct {
	Enc      string
	ActiveID string
	AppType  string
	Name     string
	Uid      string
	Fid      string
}

// QrCodeSign is for signing with the QR Code.
func (q *qrCodeParam) qrCodeSign(c *http.Client) {
	signURL := fmt.Sprintf("https://mobilelearn.chaoxing.com/pptSign/stuSignajax?enc=%s&name=%s&activeId=%s&uid=%s&clientip=&useragent=&latitude=-1&longitude=-1&fid=%s&appType=%s", q.Enc, url.QueryEscape(q.Name), q.ActiveID, q.Uid, q.Fid, q.AppType)
	resp, err := c.Get(signURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(r))
}

// Scan the QR Code
func (q *qrCodeParam) scanQrCode(fi *os.File) {
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Fatal(err)
		return
	}
	q.parseToken(qrmatrix.Content)
}

// Parse token and bind to param
func (q *qrCodeParam) parseToken(content string) {
	reg, _ := regexp.Compile(`SIGNIN:aid=(\w*)&source=(\w*)&Code=(\w*)&enc=(\w*)`)
	sub := reg.FindStringSubmatch(content)
	q.ActiveID = sub[1]
	q.AppType = sub[2]
	q.Enc = sub[4]
}
