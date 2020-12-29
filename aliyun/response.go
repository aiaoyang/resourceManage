package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

var (
	errECSTransferError = errors.New("ecs response 类型不为 DescribeInstancesResponse")

	errDomainTransferError = errors.New("domain response 类型不为 QueryDomainListResponse")

	errCertTransferError = errors.New("cert response 类型不为 CommonResponse")

	errRDSTransferError = errors.New("rds response 类型不为 MyDescribeDBInstancesResponse")
)

// ResponseToResult 通用响应转换函数  Response转为我们所需要Info
func ResponseToResult(accountName string, response responses.AcsResponse, resourceType ResourceType) (result []Info, err error) {

	switch resourceType {
	case EcsType:
		result, err = AcsResponseToEcsInfo(accountName, response)
		return
	case DomainType:
		result, err = AcsResponseToDoaminInfo(accountName, response)
		return
	case CertType:
		result, err = AcsResponseToCertInfo(accountName, response)
		return
	case RdsType:
		result, err = AcsResponseToRdsInfo(accountName, response)
		return
	default:
		return nil, fmt.Errorf("资源类型传参错误")
	}

}

// MyDescribeInstancesResponse 添加ecs查询响应结构体别名，方便为其添加Info方法
type MyDescribeInstancesResponse ecs.DescribeInstancesResponse

// Info 将Ecs response转换为Info信息
func (m MyDescribeInstancesResponse) Info(accountName string) (infos []Info, err error) {
	for _, v := range m.Instances.Instance {
		s := parseTime(v.ExpiredTime, ecsTimeFormat)
		infos = append(
			infos,
			Info{
				Name: v.InstanceName,
				// Index:  fmt.Sprintf("%d", index),
				Detail: fmt.Sprintf("%dC-%.1fG", v.Cpu, float32(v.Memory)/1024.0),
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpiredTime),
				Account: accountName,
				Type:    ResourceMap[int(EcsType)],
				Status:  s,
			},
		)
	}
	return
}

// AcsResponseToEcsInfo 特例函数，针对ecs的信息查询，将response转为Info
func AcsResponseToEcsInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*ecs.DescribeInstancesResponse)
	if !ok {
		err = errECSTransferError
		return
	}
	return MyDescribeInstancesResponse(*res).Info(accountName)
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

// AcsResponseToDoaminInfo 特例函数，针对Domain的信息查询，将response转为Info
func AcsResponseToDoaminInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*domain.QueryDomainListResponse)
	if !ok {
		err = errDomainTransferError
		return
	}
	return MyDescribeDomainResponse(*res).Info(accountName)
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

// AcsResponseToCertInfo 特例函数，针对Cert的信息查询，将response转为Info
func AcsResponseToCertInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*responses.CommonResponse)
	if !ok {
		err = errCertTransferError
		return
	}

	return MyCertResponse(*res).Info(accountName)
}

// MyDescribeDBInstancesResponse 添加RDS查询响应结构体别名，方便为其添加Info方法
type MyDescribeDBInstancesResponse rds.DescribeDBInstancesResponse

// Info 将RDS response转换为Info信息
func (m MyDescribeDBInstancesResponse) Info(accountName string) (infos []Info, err error) {
	for _, v := range m.Items.DBInstance {
		s := parseTime(v.ExpireTime, rdsTimeFormat)
		infos = append(
			infos,
			Info{
				Name:   fmt.Sprintf("%s/%s", v.DBInstanceDescription, v.DBInstanceId),
				Detail: fmt.Sprintf("%s", v.DBInstanceClass),
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpireTime),
				Account: accountName,
				Type:    ResourceMap[int(RdsType)],
				Status:  s,
			},
		)
	}
	return
}

// AcsResponseToRdsInfo 特例函数，针对rds的信息查询，将response转为Info
func AcsResponseToRdsInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*rds.DescribeDBInstancesResponse)
	if !ok {
		err = errRDSTransferError
		return
	}
	return MyDescribeDBInstancesResponse(*res).Info(accountName)
}
