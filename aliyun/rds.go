package aliyun

import (
	"fmt"
	"log"
	"sync"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

type RDSInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Account   string `json:"account"`
	Index     string `json:"index"`
}

type RDSClient struct {
	*rds.Client
	Name string
}

type RDSResponse struct {
	*rds.DescribeDBInstancesResponse
	Name string
}

var rdsClients []RDSClient

func init() {
	for _, region := range regions {

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

func GetRDS() ([]RDSInfo, error) {
	return testRDS()
}

func testRDS() ([]RDSInfo, error) {
	responsesChan := make(chan RDSResponse, len(rdsClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(rdsClients))

	for _, c := range rdsClients {

		go func(wg *sync.WaitGroup, ch chan RDSResponse, client RDSClient) {

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

			tmpRds := RDSInfo{
				Name:      fmt.Sprintf("%s/%s", v.DBInstanceDescription, v.DBInstanceId),
				Index:     fmt.Sprintf("%d", index),
				Size:      fmt.Sprintf("%s", v.DBInstanceClass),
				EndOfTime: v.ExpireTime,
				Account:   responses.Name,
				Type:      "RDS",
			}

			rdses = append(rdses, tmpRds)

		}

	}

	return rdses, nil

}

func NewDescribeRDSRequest() *rds.DescribeDBInstancesRequest {

	request := rds.CreateDescribeDBInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request
}
