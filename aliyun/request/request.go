package request

import (
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// NewDescribeAlarmRequest 生成告警请求request
func NewDescribeAlarmRequest() *cms.DescribeAlertHistoryListRequest {

	request := cms.CreateDescribeAlertHistoryListRequest()

	request.Scheme = "https"

	request.State = "ALARM"
	request.StartTime = strconv.FormatInt(time.Now().AddDate(0, 0, -1).Unix(), 10)
	request.EndTime = strconv.FormatInt(time.Now().Unix(), 10)

	request.PageSize = requests.NewInteger(1000)

	return request

}

// NewDescribeECSRequest 生成获取ecs信息的请求request
func NewDescribeECSRequest() *ecs.DescribeInstancesRequest {

	request := ecs.CreateDescribeInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request

}

// NewDescribeRDSRequest 生成rds信息查询请求request
func NewDescribeRDSRequest() *rds.DescribeDBInstancesRequest {

	request := rds.CreateDescribeDBInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request
}

// NewDescribeDomainRequest 生成获取domain信息的请求request
func NewDescribeDomainRequest() *domain.QueryDomainListRequest {

	request := domain.CreateQueryDomainListRequest()

	request.Scheme = "https"

	request.PageNum = requests.NewInteger(1)

	request.PageSize = requests.NewInteger(30)

	return request

}

// NewDescribeBalanceRequest 查询余额请求request
func NewDescribeBalanceRequest() *bssopenapi.QueryAccountBalanceRequest {

	request := bssopenapi.CreateQueryAccountBalanceRequest()

	request.Scheme = "https"

	return request

}
