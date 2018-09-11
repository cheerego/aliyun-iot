package iot

import (
	"io/ioutil"
	"net/url"
	"strings"
	"time"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"math/rand"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Version         string
	RegionId        string
}

func (c *Client) Send(request map[string]string) string {
	request["Format"] = "JSON"
	request["Version"] = c.Version
	request["SignatureMethod"] = "HMAC-SHA1"
	request["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	request["SignatureVersion"] = "1.0"
	request["SignatureNonce"] = RandStr(5)
	request["RegionId"] = c.RegionId
	request["AccessKeyId"] = c.AccessKeyId

	keys := make([]string, 0)
	for key := range request {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	kv := make([]string, 0)

	for _, value := range keys {
		kv = append(kv, escaper(value)+"="+escaper(request[value]))
	}
	needleString := strings.Join(kv, "&")

	var stringToSign string = "GET&" + escaper("/") + "&" + escaper(needleString)

	h := hmac.New(sha1.New, []byte(c.AccessKeySecret+"&"))
	h.Write([]byte(stringToSign))
	Signature := escaper(base64.StdEncoding.EncodeToString(h.Sum(nil)))

	resp, erro := http.Get("https://iot." + c.RegionId + ".aliyuncs.com?" + needleString + "&Signature=" + Signature)
	if erro != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}

func escaper(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func RandStr(strlen int) string {
	rand.Seed(time.Now().Unix())
	data := make([]byte, strlen)
	var num int
	for i := 0; i < strlen; i++ {
		num = rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		data[i] = byte(num)
	}
	return string(data)
}
