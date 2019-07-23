package post

import (
	"fmt"

	client "catdogs.club/back-end/client/post"
	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/logging"
	pb "catdogs.club/back-end/pb"
	"github.com/gin-gonic/gin"
)

func SetPostHandler(c *gin.Context) {
	var setpost SetPost
	if err := c.ShouldBind(&setpost); err != nil {
		fmt.Println(err)
		logging.Error("bind setpost err: ", err)
		libs.Resp(libs.R{
			C:    c,
			Code: -3000,
		})
		return
	}
	rsp, err := client.SetPost(&pb.SetPostReq{
		Title:   setpost.Title,
		Content: setpost.Content,
		Author:  setpost.Author,
	})
	if err != nil {
		logging.Error("Call SetPost err: ", err)
	}
	libs.Resp(libs.R{
		C:    c,
		Code: int(rsp.Code),
		Msg:  rsp.Msg,
	})
}
