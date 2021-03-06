package handler

import (
	"net/http"

	"github.com/aiaoyang/resourceManager/apis"
	"github.com/gin-gonic/gin"
)

type getResource func() (string, error)

// OnGetECS 查询ecs列表
func OnGetECS(c *gin.Context) {
	errHandler(c, apis.GetECS)

}

// OnGetRDS 查询rds列表
func OnGetRDS(c *gin.Context) {
	errHandler(c, apis.GetRDS)

}

// OnGetAlarm 查询告警列表
func OnGetAlarm(c *gin.Context) {
	errHandler(c, apis.GetAlarm)

}

// OnGetDomain 查询域名列表
func OnGetDomain(c *gin.Context) {
	errHandler(c, apis.GetDomain)

}

// OnGetLifeTime 查询剩余时间
func OnGetLifeTime(c *gin.Context) {

}

// OnGetCert 查询证书列表
func OnGetCert(c *gin.Context) {
	errHandler(c, apis.GetCert)
}

// errHandler 错误处理
func errHandler(c *gin.Context, fn getResource) {
	payload, err := fn()
	if err != nil {
		if payload == "" {

			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"errmsg": err.Error(),
					"code":   -1,
				},
			)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"msg":    payload,
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
