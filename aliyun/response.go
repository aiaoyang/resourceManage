package aliyun

import (
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
		return AcsResponseToEcsInfo(accountName, response)
	case DomainType:
		return AcsResponseToDoaminInfo(accountName, response)
	case CertType:
		return AcsResponseToCertInfo(accountName, response)
	case RdsType:
		return AcsResponseToRdsInfo(accountName, response)
	case AlertType:
		return AcsResponseToAlarmInfo(accountName, response)
	default:
		return nil, fmt.Errorf("资源类型传参错误")
	}

}
