package iot

import (
	"net/url"
	"strings"
	"time"
	"math/rand"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"io/ioutil"
)

type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	Version         string
	RegionId        string
}

func (c *Client) Send(request map[string]string) (string, error) {
	request["Format"] = "JSON"
	request["Version"] = c.Version
	request["SignatureMethod"] = "HMAC-SHA1"
	request["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	request["SignatureVersion"] = "1.0"
	request["SignatureNonce"] = GetRandomString(14)
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
		return "", erro
	}
	defer resp.Body.Close()

	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		return "", erro
	}

	return string(body), nil
}

func escaper(str string) string {
	str = url.QueryEscape(str)
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
