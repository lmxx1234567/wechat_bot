package client

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//GetUUID get uuid for QRLogin
func GetUUID() (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://login.wx.qq.com/jslogin", nil)
	if err != nil {
		log.Print(err)
		return "", err
	}

	q := url.Values{
		"appid":        {"wx782c26e4c19acffb"},
		"redirect_uri": {"https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage"},
		"fun":          {"fun"},
		"lang":         {"zh_CN"},
		"_":            {strconv.FormatInt(time.Now().Unix(), 10)},
	}
	req.URL.RawQuery = q.Encode()
	log.Print(req.URL.RawQuery)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString = string(bodyBytes)
	}
	resp.Body.Close()
	log.Println(bodyString[len(bodyString)-14 : len(bodyString)-2])
	return bodyString[len(bodyString)-14 : len(bodyString)-2], nil
}

//GetQRcode get QRcode by uuid
func GetQRcode(uuid string) (string, error) {
	resp, err := http.Get("https://login.weixin.qq.com/qrcode/" + uuid)
	if err != nil {
		log.Print(err)
		return "", err
	}
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		img, _, err := image.Decode(bytes.NewReader(bodyBytes))
		if err != nil {
			return "", err
		}
		str, err := getBase64(&img)
		if err != nil {
			return "", err
		}
		return str, nil
	}
	return "", errors.New("StatusCode : " + strconv.Itoa(resp.StatusCode))

}

// writeImageWithTemplate encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func getBase64(img *image.Image) (string, error) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return str, nil
}
