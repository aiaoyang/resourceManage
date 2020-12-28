package main

import (
	"log"
	"time"

	"github.com/aiaoyang/resourceManager/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var err error

func main() {
	log.SetFlags(log.Llongfile)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"token", "Authorization", "Content-Type", "Origin", "Content-Length", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"token", "Access-Control-Allow-Origin", "Authorization"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ecs", handler.OnGetECS)
	r.GET("/rds", handler.OnGetRDS)
	r.GET("/alarm", handler.OnGetAlarm)
	r.GET("/life", handler.OnGetLifeTime)
	r.GET("/domain", handler.OnGetDomain)
	r.GET("/cert", handler.OnGetCert)
	r.Run(":9090")
}
