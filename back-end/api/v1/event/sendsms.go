package event

import (
	"fmt"

	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/models"
	"github.com/gin-gonic/gin"
)

type SmsParam struct {
	PhoneNum string `form:"phonenum"`
}

func SendSms(c *gin.Context) {
	var smsParam SmsParam
	err := c.ShouldBind(&smsParam)
	if err != nil {
		fmt.Println(err)
	}
	code := libs.RandInt(4)
	go libs.SendSms(smsParam.PhoneNum, code)
	go saveCode(&smsParam, code)
	libs.Resp(libs.R{
		C:    c,
		Code: 0,
		Msg:  "success",
	})
}

func saveCode(param *SmsParam, code string) {
	v := models.VerifyCode{
		PhoneNum: param.PhoneNum,
		Code:     code,
	}
	err := v.Set()
	if err != nil {
		fmt.Println(err)
	}
}
