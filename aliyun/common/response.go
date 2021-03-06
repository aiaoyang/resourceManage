package common

import (
	"errors"
	"fmt"

	"github.com/aiaoyang/resourceManager/resource"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

var (
	errECSTransferError = errors.New("ecs response 类型不为 DescribeInstancesResponse")

	errDomainTransferError = errors.New("domain response 类型不为 QueryDomainListResponse")

	errCertTransferError = errors.New("cert response 类型不为 MyDescribeDBInstancesResponse")

	errRDSTransferError = errors.New("rds response 类型不为 CommonResponse")
)

// ResponseToResult 通用响应转换函数  Response转为我们所需要Info
func ResponseToResult(accountName string, response responses.AcsResponse, resourceType resource.Type) (result []resource.Info, err error) {

	switch resourceType {
	case resource.EcsType:
		return AcsResponseToEcsInfo(accountName, response)
	case resource.DomainType:
		return AcsResponseToDoaminInfo(accountName, response)
	case resource.CertType:
		return AcsResponseToCertInfo(accountName, response)
	case resource.RdsType:
		return AcsResponseToRdsInfo(accountName, response)
	case resource.AlertType:
		return AcsResponseToAlarmInfo(accountName, response)
	default:
		return nil, fmt.Errorf("资源类型传参错误")
	}
}
