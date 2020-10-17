package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiaoyang/resourceManager/aliyun"
	"github.com/aiaoyang/resourceManager/apis"
	"github.com/gin-gonic/gin"
)

func OnGetECS(c *gin.Context) {
	ecses, err := apis.GetECS()
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
			"msg":  ecses,
			"code": 0,
		},
	)

}

func OnGetRDS(c *gin.Context) {
	rdses, err := apis.GetRDS()
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
			"msg":  rdses,
			"code": 0,
		},
	)

}

func OnGetAlarm(c *gin.Context) {
	alarm, err := aliyun.GetAlarm()
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

	res, err := json.Marshal(alarm)
	c.JSON(
		http.StatusOK,
		gin.H{
			"msg":  fmt.Sprintf("%s", res),
			"code": 0,
		},
	)

}

func OnGetLifeTime(c *gin.Context) {
	lifeTimeInstances, err := apis.GetLifeTime()
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
			"msg":  lifeTimeInstances,
			"code": 0,
		},
	)
}
