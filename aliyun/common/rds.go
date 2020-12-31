package common

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// GetRDS 查询RDS列表
func GetRDS() (infos []Info, err error) {
	var req = rds.CreateDescribeDBInstancesRequest()
	var resp = rds.CreateDescribeDBInstancesResponse()
	req.PageSize = requests.NewInteger(100)

	return Describe(GlobalClients, req, resp, RdsType)
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
