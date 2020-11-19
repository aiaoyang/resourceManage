package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aiaoyang/resourceManager/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var err error
var (
	host   = "127.0.0.1"
	port   = 5432
	user   = "postgres"
	dbname = "testdb"
)

func main() {
	log.SetFlags(log.Llongfile)
	r := gin.Default()
	// pgsqlStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	// db, err := gorm.Open("postgres", pgsqlStr)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db.Model(apis.MyResource{}).AutoMigrate()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"token", "Authorization", "Content-Type", "Origin", "Content-Length", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"token", "Access-Control-Allow-Origin", "Authorization"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// r.GET("/test",
	// 	func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"msg": "hello world",
	// 		})
	// 	},
	// )
	r.GET("/resources",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"type": "domain",
				"data": func() string {
					type Resource struct {
						Type string `json:"type"`
						Name string `json:"name"`
					}
					r := Resource{
						Type: "ecs",
						Name: "碧蓝游戏网关",
					}
					b, err := json.Marshal(r)
					if err != nil {
						log.Fatal(err)
					}
					return fmt.Sprintf("%s", b)
				}(),
			})
		},
	)

	r.GET("/ecs", handler.OnGetECS)
	r.GET("/rds", handler.OnGetRDS)
	r.GET("/alarm", handler.OnGetAlarm)
	r.GET("/life", handler.OnGetLifeTime)
	r.GET("/domain", handler.OnGetDomain)
	r.Run(":9090")
}
