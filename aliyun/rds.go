package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// GetRDS 查询RDS列表
func GetRDS() (infos []Info, err error) {
	var req = rds.CreateDescribeDBInstancesRequest()
	var resp = rds.CreateDescribeDBInstancesResponse()
	req.PageSize = requests.NewInteger(100)

	return Describe(GlobalClients, req, resp, RdsType)
}
