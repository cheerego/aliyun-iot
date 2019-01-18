package iot

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
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
	request["SignatureNonce"] = RandStringBytesMaskImprSrc(14)
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

//https://codeday.me/bug/20170607/22281.html
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
