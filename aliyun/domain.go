package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// GetDomain 查询域名
func GetDomain() (infos []Info, err error) {
	var resp = domain.CreateQueryDomainListResponse()

	var req = domain.CreateQueryDomainListRequest()
	req.PageNum = requests.NewInteger(1)
	req.PageSize = requests.NewInteger(30)

	return Describe(GlobalClients, req, resp, DomainType)
}

// AcsResponseToDoaminInfo 特例函数，针对Domain的信息查询，将response转为Info
func AcsResponseToDoaminInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
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
func (m MyDescribeDomainResponse) Info(accountName string) (infos []Info, err error) {
	for _, v := range m.Data.Domain {
		s := parseTime(v.ExpirationDate, domainTimeFormat)
		infos = append(
			infos,
			Info{
				Name: v.DomainName,
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpirationDate),
				Account: accountName,
				Type:    ResourceMap[int(DomainType)],
				Status:  s,
			},
		)
	}
	return
}
