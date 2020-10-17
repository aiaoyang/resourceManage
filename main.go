package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aiaoyang/resourceManager/apis"
	"github.com/aiaoyang/resourceManager/handler"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
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
	pgsqlStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := gorm.Open("postgres", pgsqlStr)

	if err != nil {
		log.Fatal(err)
	}
	db.Model(apis.MyResource{}).AutoMigrate()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"token", "Authorization", "Content-Type", "Origin", "Content-Length", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"token", "Access-Control-Allow-Origin", "Authorization"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/test",
		func(c *gin.Context) {
			res := &apis.MyResource{}
			err = c.Bind(&res)
			if err != nil {
				log.Fatal(err)
			}
			r, err := json.Marshal(res.Values)

			if err != nil {
				log.Fatal(err)
			}

			res.Value = postgres.Jsonb{RawMessage: r}
			db.Debug().NewRecord(res)
			db.Debug().Create(res)
		},
	)
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
	r.Run(":9090")
}

// QueryECS 查询ecs信息
func QueryECS(regin, accessID, accessKey string) ([]byte, error) {

	client, err := bssopenapi.NewClientWithAccessKey(regin, accessID, accessKey)
	if err != nil {
		return nil, err
	}

	request := bssopenapi.CreateQueryAvailableInstancesRequest()
	request.Scheme = "https"

	request.PageNum = requests.NewInteger(1)
	request.ProductCode = "ecs"

	response, err := client.QueryAvailableInstances(request)
	if err != nil {
		return nil, err
	}
	return json.Marshal(response)

}
