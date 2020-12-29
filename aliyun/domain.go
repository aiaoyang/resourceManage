package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
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
