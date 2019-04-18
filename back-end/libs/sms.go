package libs

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

const (
	accessKeyId  = "LTAI7GgefI5FIuhS"
	accessSecret = "WTrOHELCEaYDzapYeOFamaDBHIWsfr"
	method       = "POST"
	scheme       = "https"
	domain       = "dysmsapi.aliyuncs.com"
	version      = "2017-05-25"
	apiName      = "SendSms"
	regionId     = "cn-hangzhou"
	signName     = "yoko博客"
	templateCode = "SMS_163845182"
)

func SendSms(phone, code string) error {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		fmt.Println(err)
		return err
	}

	request := requests.NewCommonRequest()
	request.Method = method
	request.Scheme = scheme
	request.Domain = domain
	request.Version = version
	request.ApiName = apiName
	request.QueryParams["RegionId"] = regionId
	request.QueryParams["SignName"] = signName
	request.QueryParams["TemplateCode"] = templateCode
	request.QueryParams["PhoneNumbers"] = phone
	request.QueryParams["TemplateParam"] = "{\"code\":\"" + code + "\"}"

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Print(response.GetHttpContentString())
	return nil
}
