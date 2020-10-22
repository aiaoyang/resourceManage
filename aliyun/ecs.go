package aliyun

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// ECSInfo ecs信息
type ECSInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

// ECSClient ecs请求客户端
type ECSClient struct {
	client interface{}

	Name string
}

// ECSResponse ecs请求返回值结构体
type ECSResponse struct {
	response interface{}
	Name     string
}

var ecsClients []ECSClient

func init() {

	for _, region := range config.GVC.Regions {

		for _, m := range config.GVC.Accounts {
			c, err := ecs.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}

			tmp := ECSClient{c, m.Name}

			ecsClients = append(ecsClients, tmp)

		}

	}

}

// GetECS 获取ecs信息
func GetECS() ([]ECSInfo, error) {

	return describeECS()

}

func describeECS() ([]ECSInfo, error) {

	responsesChan := make(chan ECSResponse, len(ecsClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(ecsClients))

	for _, c := range ecsClients {

		go func(
			wg *sync.WaitGroup,
			ch chan ECSResponse,
			client ECSClient,
		) {

			defer wg.Done()

			resp, err := client.client.(*ecs.Client).DescribeInstances(NewDescribeECSRequest())

			if err != nil {

				log.Println(err)

				return

			}

			defer resp.GetOriginHttpResponse().Body.Close()

			tmp := ECSResponse{resp, client.Name}

			ch <- tmp

		}(wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	ecses := make([]ECSInfo, 0)

	index := 0

	for responses := range responsesChan {

		for _, v := range responses.response.(*ecs.DescribeInstancesResponse).Instances.Instance {

			index++

			paseTime, err := time.Parse("2006-01-02T15:04Z", v.ExpiredTime)
			if err != nil {
				log.Fatal(err)
			}
			// log.Printf("time.end : %s\n", paseTime.String())

			s := green

			if time.Now().AddDate(0, 1, 0).After(paseTime) {
				s = yellow
			}
			if time.Now().AddDate(0, 0, 7).After(paseTime) {
				s = red
			}
			if time.Now().After(paseTime) {
				s = nearDead
			}

			tmpEcs := ECSInfo{
				Name:  v.InstanceName,
				Index: fmt.Sprintf("%d", index),
				Size:  fmt.Sprintf("%dC-%.1fG", v.Cpu, float32(v.Memory)/1024.0),
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpiredTime),
				Account: responses.Name,
				Type:    "ECS",
				Status:  s,
			}

			ecses = append(ecses, tmpEcs)

		}

	}

	return ecses, nil

}

// NewDescribeECSRequest 生成获取ecs信息的请求request
func NewDescribeECSRequest() *ecs.DescribeInstancesRequest {

	request := ecs.CreateDescribeInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request

}
