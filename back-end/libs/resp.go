package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type R struct {
	C    *gin.Context
	Code int
	Msg  string
	Data gin.H
}

func Resp(r R) {
	if r.Data == nil {
		r.Data = gin.H{}
	}
	if r.Msg == "" {
		r.Msg = codes[r.Code]
	}
	r.C.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
		"data": r.Data,
	})
}
