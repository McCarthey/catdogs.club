package user

import (
	"fmt"

	client "catdogs.club/back-end/client/user"

	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/logging"
	pb "catdogs.club/back-end/pb"
	"github.com/gin-gonic/gin"
)

// @Tags 用户
// @Summary 注册接口
// @Description 注册接口
// @Accept  json
// @Produce  json
// @Param email formData string true "邮箱账号"
// @Param password formData string true "密码"
// @Success 0 {string} string "{"code": 0, "data": {}, "msg": "success"}"
// @Failure -999 {string} string "服务器出问题"
// @Failure -1000 {string} string "用户已存在"
// @Router /register [post]
func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err)
		logging.Info("bind user err", err.Error())
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}

	rsp, err := client.Regist(&pb.RegisterReq{
		Email:    user.Email,
		Password: user.Password,
	})
	logging.Info("Call Regist Err: ", err)
	libs.Resp(libs.R{
		C:    c,
		Code: int(rsp.Rsp.Code),
		Msg:  rsp.Rsp.Msg,
		Data: gin.H{
			"token": rsp.Token,
		},
	})
}
