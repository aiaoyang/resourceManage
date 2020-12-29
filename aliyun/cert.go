package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func GetCert() (infos []Info, err error) {
	var (
		req  = NewGetCertListRequest("cn-hangzhou")
		resp = responses.NewCommonResponse()
	)
	return Describe(GlobalClients, req, resp, CertType)
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
