package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// GetECS 查询ecs列表
func GetECS() ([]Info, error) {
	var resp = ecs.CreateDescribeInstancesResponse()
	var req = ecs.CreateDescribeInstancesRequest()
	// (每页100)
	req.PageSize = requests.NewInteger(100)

	return Describe(GlobalClients, req, resp, EcsType)
}
