package aliyun

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type DomainInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

type DomainClient struct {
	client *domain.Client

	Name string
}

type DomainResponse struct {
	response *domain.QueryDomainListResponse
	Name     string
}

var domainClients []DomainClient

func init() {

	for _, region := range config.GVC.Regions {

		for _, m := range config.GVC.Accounts {
			c, err := domain.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}

			tmp := DomainClient{c, m.Name}

			domainClients = append(domainClients, tmp)

		}

	}

}

func GetDomain() ([]DomainInfo, error) {
	return descirbeDomain()
}

func descirbeDomain() ([]DomainInfo, error) {

	responsesChan := make(chan DomainResponse, len(domainClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(domainClients))

	for _, c := range domainClients {

		go func(
			wg *sync.WaitGroup,
			ch chan DomainResponse,
			client DomainClient,
		) {

			defer wg.Done()

			resp, err := client.client.QueryDomainList(NewDescribeDomainRequest())

			if err != nil {

				log.Println(err)

				return

			}

			defer resp.GetOriginHttpResponse().Body.Close()

			tmp := DomainResponse{resp, client.Name}

			ch <- tmp

		}(wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	domains := make([]DomainInfo, 0)

	index := 0

	domainToAccountMap := make(map[string]string)
	for responses := range responsesChan {
		for _, domain := range responses.response.Data.Domain {

			paseTime, err := time.Parse("2006-01-02 15:04:05", domain.ExpirationDate)
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
			if _, ok := domainToAccountMap[domain.DomainName]; ok {
				continue
			} else {
				domainToAccountMap[domain.DomainName] = responses.Name
				index++

				tmpDomain := DomainInfo{
					Name:  domain.DomainName,
					Index: fmt.Sprintf("%d", index),
					EndOfTime: func(endOfTime string) string {
						if endOfTime == "" {
							return "error"
						}
						return endOfTime
					}(domain.ExpirationDate),
					Account: domainToAccountMap[domain.DomainName],
					Type:    "Domain",
					Status:  s,
				}

				domains = append(domains, tmpDomain)
			}

		}

	}

	return domains, nil

}

// NewDescribeDomainRequest 生成获取domain信息的请求request
func NewDescribeDomainRequest() *domain.QueryDomainListRequest {

	request := domain.CreateQueryDomainListRequest()

	request.Scheme = "https"

	request.PageNum = requests.NewInteger(1)

	request.PageSize = requests.NewInteger(30)

	return request

}
