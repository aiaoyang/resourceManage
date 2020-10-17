package aliyun

import (
	"fmt"
	"log"
	"sync"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type ECSInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Account   string `json:"account"`
	Index     string `json:"index"`
}

type ECSClient struct {
	*ecs.Client
	Name string
}
type ECSResponse struct {
	*ecs.DescribeInstancesResponse
	Name string
}

var ecsClients []ECSClient

func init() {

	for _, region := range regions {

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

func GetECS() ([]ECSInfo, error) {

	return testEcs()

}

func testEcs() ([]ECSInfo, error) {

	responsesChan := make(chan ECSResponse, len(ecsClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(ecsClients))

	for _, c := range ecsClients {

		go func(wg *sync.WaitGroup, ch chan ECSResponse, client ECSClient) {

			defer wg.Done()

			resp, err := client.DescribeInstances(NewDescribeECSRequest())

			if err != nil {

				log.Println(err)

				return

			}

			tmp := ECSResponse{resp, client.Name}

			ch <- tmp

		}(wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	ecses := make([]ECSInfo, 0)

	index := 0

	for responses := range responsesChan {

		for _, v := range responses.Instances.Instance {

			index++

			tmpEcs := ECSInfo{
				Name:      v.InstanceName,
				Index:     fmt.Sprintf("%d", index),
				Size:      fmt.Sprintf("%dC-%.1fG", v.Cpu, float32(v.Memory)/1024.0),
				EndOfTime: v.ExpiredTime,
				Account:   responses.Name,
				Type:      "ECS",
			}

			ecses = append(ecses, tmpEcs)

		}

	}

	return ecses, nil

}

func NewDescribeECSRequest() *ecs.DescribeInstancesRequest {

	request := ecs.CreateDescribeInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request

}
