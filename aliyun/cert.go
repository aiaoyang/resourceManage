package aliyun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetCert 查询证书
func GetCert() (infos []Info, err error) {
	var req = NewGetCertListRequest("cn-hangzhou")
	var resp = responses.NewCommonResponse()
	req.TransToAcsRequest()

	// log.Println(resp.String())
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

// AcsResponseToCertInfo 特例函数，针对Cert的信息查询，将response转为Info
func AcsResponseToCertInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*responses.CommonResponse)
	if !ok {
		err = errCertTransferError
		return
	}

	return MyCertResponse(*res).Info(accountName)
}

// MyCertResponse 证书响应结构体
type MyCertResponse responses.CommonResponse

// Info 将证书 响应转换为Info信息
func (m MyCertResponse) Info(accountName string) (infos []Info, err error) {
	// Certificate 证书信息
	type certificate struct {
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
	// CertResp common response
	type CertResp struct {
		TotalCount      int           `json:"TotalCount"`
		RequestID       string        `json:"RequestId"`
		CurrentPage     int           `json:"CurrentPage"`
		CertificateList []certificate `json:"CertificateList"`
	}
	resp := CertResp{}
	// log.Printf("%s\n\n", responses.response.GetHttpContentString())
	err = json.Unmarshal(m.GetHttpContentBytes(), &resp)

	if err != nil {
		return
	}

	certToAccountMap := make(map[string]bool)
	for _, cert := range resp.CertificateList {

		s := parseTime(cert.EndDate, certTimeFormat)

		if _, ok := certToAccountMap[cert.Sans]; ok {
			continue
		} else {
			certToAccountMap[cert.Sans] = true

			tmpCert := Info{
				Name: cert.Sans,
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "error"
					}
					return endOfTime
				}(cert.EndDate),
				Account: accountName,
				Type:    ResourceMap[int(CertType)],
				Status:  s,
			}

			infos = append(infos, tmpCert)
		}

	}
	return
}
