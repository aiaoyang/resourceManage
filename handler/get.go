package handler

import (
	"log"
	"net/http"

	"github.com/aiaoyang/resourceManager/apis"
	"github.com/gin-gonic/gin"
)

// OnGetECS 查询ecs列表
func OnGetECS(c *gin.Context) {
	ecses, err := apis.GetECS()
	errHandler(c, ecses, err)

}

// OnGetRDS 查询rds列表
func OnGetRDS(c *gin.Context) {
	rdses, err := apis.GetRDS()
	errHandler(c, rdses, err)

}

// OnGetAlarm 查询告警列表
func OnGetAlarm(c *gin.Context) {
	alarm, err := apis.GetAlarm()
	errHandler(c, alarm, err)

}

// OnGetDomain 查询域名列表
func OnGetDomain(c *gin.Context) {
	domain, err := apis.GetDomain()
	errHandler(c, domain, err)

}

// OnGetLifeTime 查询剩余时间
func OnGetLifeTime(c *gin.Context) {

}

// OnGetCert 查询证书列表
func OnGetCert(c *gin.Context) {
	cert, err := apis.GetCert()
	errHandler(c, cert, err)
}

// errHandler 错误处理
func errHandler(c *gin.Context, payload string, err error) {
	if err != nil {
		log.Printf("err : %s\n", err.Error())
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errmsg": err,
				"code":   -1,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"msg":  payload,
			"code": 0,
		},
	)

}
