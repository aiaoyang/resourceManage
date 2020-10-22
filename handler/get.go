package handler

import (
	"net/http"

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
	alarm, err := apis.GetAlarm()
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
			"msg":  alarm,
			"code": 0,
		},
	)

}

func OnGetDomain(c *gin.Context) {
	domain, err := apis.GetDomain()
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
			"msg":  domain,
			"code": 0,
		},
	)

}
func OnGetLifeTime(c *gin.Context) {
	// lifeTimeInstances, err := apis.GetLifeTime()
	// if err != nil {
	// 	c.AbortWithStatusJSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"errmsg": err,
	// 			"code":   -1,
	// 		},
	// 	)
	// 	return
	// }

	// c.JSON(
	// 	http.StatusOK,
	// 	gin.H{
	// 		"msg":  lifeTimeInstances,
	// 		"code": 0,
	// 	},
	// )
}
