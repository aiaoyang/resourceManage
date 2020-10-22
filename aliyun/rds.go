package aliyun

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// RDSInfo rds信息
type RDSInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

// RDSClient rds请求客户端
type RDSClient struct {
	*rds.Client
	Name string
}

// RDSResponse rds请求返回信息结构体
type RDSResponse struct {
	*rds.DescribeDBInstancesResponse
	Name string
}

var rdsClients []RDSClient

func init() {
	for _, region := range config.GVC.Regions {

		for _, m := range config.GVC.Accounts {
			c, err := rds.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}
			tmp := RDSClient{c, m.Name}

			rdsClients = append(rdsClients, tmp)

		}

	}
}

// GetRDS 获取rds信息
func GetRDS() ([]RDSInfo, error) {
	return descirbeRDS()
}
func descirbeRDS() ([]RDSInfo, error) {
	responsesChan := make(chan RDSResponse, len(rdsClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(rdsClients))

	for _, c := range rdsClients {

		go func(
			wg *sync.WaitGroup,
			ch chan RDSResponse,
			client RDSClient,
		) {

			defer wg.Done()

			resp, err := client.DescribeDBInstances(NewDescribeRDSRequest())

			if err != nil {

				log.Println(err)

				return

			}

			tmp := RDSResponse{resp, client.Name}

			ch <- tmp

		}(wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	rdses := make([]RDSInfo, 0)

	index := 0

	for responses := range responsesChan {

		for _, v := range responses.Items.DBInstance {

			index++

			s := green

			paseTime, err := time.Parse("2006-01-02T15:04:05Z", v.ExpireTime)

			if err != nil {
				log.Println(err)
			} else {

				if time.Now().AddDate(0, 1, 0).After(paseTime) {
					s = yellow
				}

				if time.Now().AddDate(0, 0, 7).After(paseTime) {
					s = red
				}

				if time.Now().After(paseTime) {
					s = nearDead
				}

			}
			// log.Printf("time.end : %s\n", paseTime.String())

			tmpRds := RDSInfo{
				Name:  fmt.Sprintf("%s/%s", v.DBInstanceDescription, v.DBInstanceId),
				Index: fmt.Sprintf("%d", index),
				Size:  fmt.Sprintf("%s", v.DBInstanceClass),
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpireTime),
				Account: responses.Name,
				Type:    "RDS",
				Status:  s,
			}

			rdses = append(rdses, tmpRds)

		}

	}

	return rdses, nil

}

// NewDescribeRDSRequest 生成rds信息查询请求request
func NewDescribeRDSRequest() *rds.DescribeDBInstancesRequest {

	request := rds.CreateDescribeDBInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request
}
