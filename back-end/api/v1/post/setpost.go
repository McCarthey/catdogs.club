package post

import (
	client "catdogs.club/back-end/client/post"
	"catdogs.club/back-end/libs"
	"catdogs.club/back-end/logging"
	pb "catdogs.club/back-end/pb"
	"github.com/gin-gonic/gin"
)

// @Tags 文章
// @Summary 发布文章
// @Description 发布文章
// @Accept  json
// @Produce  json
// @Param title formData string true "标题"
// @Param content formData string true "内容"
// @Param author formData string true "作者"
// @Success 0 {string} string "{"code": 0, "data": {}, "msg": "success"}"
// @Failure -999 {string} string "服务器出问题"
// @Failure -3000 {string} string "参数错误"
// @Router /setpost [post]
func SetPostHandler(c *gin.Context) {
	var setpost SetPost
	if err := c.ShouldBind(&setpost); err != nil {
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
		Code: int(rsp.Rsp.Code),
		Msg:  rsp.Rsp.Msg,
	})
}
