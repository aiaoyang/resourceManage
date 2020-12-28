package handler

import (
	"net/http"

	"github.com/aiaoyang/resourceManager/apis"
	"github.com/gin-gonic/gin"
)

func OnGetECS(c *gin.Context) {
	ecses, err := apis.GetECS()
	errHandler(c, ecses, err)

}

func OnGetRDS(c *gin.Context) {
	rdses, err := apis.GetRDS()
	errHandler(c, rdses, err)

}

func OnGetAlarm(c *gin.Context) {
	alarm, err := apis.GetAlarm()
	errHandler(c, alarm, err)

}

func OnGetDomain(c *gin.Context) {
	domain, err := apis.GetDomain()
	errHandler(c, domain, err)

}
func OnGetLifeTime(c *gin.Context) {

}

func errHandler(c *gin.Context, payload string, err error) {
	if err != nil {
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
