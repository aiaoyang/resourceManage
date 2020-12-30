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
