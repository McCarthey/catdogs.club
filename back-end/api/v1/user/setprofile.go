package user

import (
	"catdogs-service/libs"

	client "catdogs.club/back-end/client/user"
	"catdogs.club/back-end/logging"
	pb "catdogs.club/back-end/pb"
	"github.com/gin-gonic/gin"
)

func SetProfile(c *gin.Context) {
	var profile ProfileParam
	if err := c.ShouldBind(&profile); err != nil {
		logging.Error("Bind profile err", err.Error())
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}

	rsp, err := client.SetProfile(&pb.SetProfileReq{
		Name:     profile.Name,
		Gender:   profile.Gender,
		Age:      profile.Age,
		PhoneNum: profile.PhoneNum,
		Email:    profile.Email,
		Birthday: profile.Birthday,
		City:     profile.City,
		Address:  profile.Address,
	})
	if err != nil {
		logging.Error("Call SetProfile err: ", err)
	}
	libs.Resp(libs.R{Code: 0})
}
