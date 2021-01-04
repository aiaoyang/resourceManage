package common

import (
	"github.com/aiaoyang/resourceManager/resource"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// GetDomain 查询域名
func GetDomain() (infos []resource.Info, err error) {

	var createRequestFunc = func() requests.AcsRequest {
		var req = domain.CreateQueryDomainListRequest()
		req.PageNum = requests.NewInteger(1)
		req.PageSize = requests.NewInteger(30)
		return req
	}

	var createResponseFunc = func() responses.AcsResponse {
		return domain.CreateQueryDomainListResponse()
	}

	return Describe(GlobalClients, createRequestFunc, createResponseFunc, resource.DomainType)
}

// AcsResponseToDoaminInfo 特例函数，针对Domain的信息查询，将response转为Info
func AcsResponseToDoaminInfo(accountName string, response responses.AcsResponse) (result []resource.Info, err error) {
	res, ok := response.(*domain.QueryDomainListResponse)
	if !ok {
		err = errDomainTransferError
		return
	}
	return MyDescribeDomainResponse(*res).Info(accountName)
}

// MyDescribeDomainResponse 查询域名列表
type MyDescribeDomainResponse domain.QueryDomainListResponse

// Info 将Domain response转换为Info信息
func (m MyDescribeDomainResponse) Info(accountName string) (infos []resource.Info, err error) {
	for _, v := range m.Data.Domain {
		s := parseTime(v.ExpirationDate, domainTimeFormat)
		infos = append(
			infos,
			resource.Info{
				Name: v.DomainName,
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpirationDate),
				Account: accountName,
				Type:    resource.ResourceMap[int(resource.DomainType)],
				Status:  s,
			},
		)
	}
	return
}
