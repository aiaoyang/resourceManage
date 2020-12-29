package aliyun

import (
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var (
	errECSTransferError = errors.New("response 类型不为 DescribeInstancesResponse")

	errDomainTransferError = errors.New("response 类型不为 QueryDomainListResponse")
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
