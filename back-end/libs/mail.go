package libs

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"catdogs.club/back-end/logging"

	configs "catdogs.club/back-end/configs/common"
)

var (
	Aliyun map[string]interface{}
)

func PercentEncode(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2a", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func CreateSignature(secret, stringToSign string) string {
	key := []byte(secret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(stringToSign))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s
}

func GetUtcTime() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)
	return s
}

func SendMail(to, sub, content string) error {
	utc := GetUtcTime()
	param := url.Values{}
	param.Add("Action", "SingleSendMail")
	param.Add("AccountName", configs.FromEmail)
	param.Add("ReplyToAddress", "true")
	param.Add("AddressType", "1")
	param.Add("ToAddress", to)
	param.Add("FromAlias", configs.SignName)
	param.Add("Subject", sub)
	param.Add("HtmlBody", content)
	param.Add("Format", "JSON")
	param.Add("Version", "2015-11-23")
	param.Add("SignatureMethod", "HMAC-SHA1")
	param.Add("SignatureNonce", utc)
	param.Add("SignatureVersion", "1.0")
	param.Add("AccessKeyId", configs.AccessKeyId)
	param.Add("Timestamp", utc)
	fmt.Println(param.Encode())
	percent := PercentEncode(param.Encode())
	stringToSign := "GET" + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(percent)
	signature := CreateSignature(configs.AccessSecret, stringToSign)
	param.Add("Signature", signature)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://dm.aliyuncs.com/?" + PercentEncode(param.Encode()))
	if err != nil {
		logging.Error("SendMail Error: ", err)
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Error("Read SendMail Resp Error: ", err)
		return err
	}
	json.Unmarshal(data, &Aliyun)
	logging.Info("SendMail Result: ", &Aliyun)
	return nil
}
