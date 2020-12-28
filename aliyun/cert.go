package aliyun

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// CertResp common response
type CertResp struct {
	TotalCount      int           `json:"TotalCount"`
	RequestID       string        `json:"RequestId"`
	CurrentPage     int           `json:"CurrentPage"`
	CertificateList []Certificate `json:"CertificateList"`
}

// Certificate 证书信息
type Certificate struct {
	// 是否在阿里云购买
	BuyInAliyun bool `json:"byInAliyun"`

	// 城市
	City string `json:"city"`

	//
	Common string `json:"common"`

	//
	Country string `json:"conutry"`

	// 到期日
	EndDate string `json:"endDate"`

	// 是否过期
	Expired bool `json:"expired"`

	//
	Fingerprint string `json:"fingerprint"`

	//
	ID int64 `json:"id"`

	//
	Issuer string `json:"issuer"`

	//
	Name string `json:"name"`

	//
	OrgName string `json:"orgName"`

	//
	Province string `json:"province"`

	// 域名
	Sans string `json:"sans"`

	// 证书申请日期
	StartDate string `json:"startDate"`
}

// CertInfo 用于返回给前端的证书信息
type CertInfo struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

// CertResponse 自定义 证书请求响应response
type CertResponse struct {
	response *responses.CommonResponse
	Name     string
}

// CertClient 自定义证书请求客户端
type CertClient struct {
	client *domain.Client

	Name string
}

var certClients []CertClient

func init() {
	for _, region := range config.GVC.Regions {

		for _, m := range config.GVC.Accounts {
			c, err := domain.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}

			tmp := CertClient{c, m.Name}

			certClients = append(certClients, tmp)

		}

	}
}

// GetCertList 证书列表查询
func GetCertList() ([]CertInfo, error) {
	return getCertList()
}
func getCertList() ([]CertInfo, error) {

	ctx := context.Background()

	responsesChan := make(chan CertResponse, len(certClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(certClients))

	for _, c := range certClients {

		go func(
			ctx context.Context,
			wg *sync.WaitGroup,
			ch chan CertResponse,
			client CertClient,
		) {

			defer wg.Done()

			resp, err := client.client.ProcessCommonRequest(NewGetCertListRequest("cn-hangzhou"))

			if err != nil {

				log.Println(err)

				return

			}

			defer resp.GetOriginHttpResponse().Body.Close()

			tmp := CertResponse{resp, client.Name}

			ch <- tmp

		}(ctx, wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	certs := make([]CertInfo, 0)

	index := 0

	certToAccountMap := make(map[string]string)
	for responses := range responsesChan {
		resp := CertResp{}
		// log.Printf("%s\n\n", responses.response.GetHttpContentString())
		err := json.Unmarshal(responses.response.GetHttpContentBytes(), &resp)

		if err != nil {
			return nil, err
		}

		for _, cert := range resp.CertificateList {
			log.Printf("certName: %s\n\n", cert.Sans)
		}
		log.Printf("cert list is : %v\n\n", resp.CertificateList)
		for _, cert := range resp.CertificateList {

			paseTime, err := time.Parse("2006-01-02", cert.EndDate)
			if err != nil {
				return nil, err
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
			if _, ok := certToAccountMap[cert.Sans]; ok {
				continue
			} else {
				certToAccountMap[cert.Sans] = responses.Name
				index++

				tmpCert := CertInfo{
					Name:  cert.Sans,
					Index: fmt.Sprintf("%d", index),
					EndOfTime: func(endOfTime string) string {
						if endOfTime == "" {
							return "error"
						}
						return endOfTime
					}(cert.EndDate),
					Account: certToAccountMap[cert.Sans],
					Type:    "Cert",
					Status:  s,
				}

				certs = append(certs, tmpCert)
			}

		}

	}

	return certs, nil

}

// NewGetCertListRequest 生成获取证书列表信息的请求request
func NewGetCertListRequest(region string) *requests.CommonRequest {

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "cas.aliyuncs.com"
	request.Version = "2018-07-13"
	request.ApiName = "DescribeUserCertificateList"

	request.QueryParams["RegionId"] = region
	request.QueryParams["ShowSize"] = "30"
	request.QueryParams["CurrentPage"] = "1"
	return request

}
