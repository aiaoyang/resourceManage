<<<<<<< HEAD
=======
package common

import (
	"fmt"

	"github.com/aiaoyang/resourceManager/resource"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// GetECS 查询ecs列表
func GetECS() ([]resource.Info, error) {
	var resp = ecs.CreateDescribeInstancesResponse()
	var req = ecs.CreateDescribeInstancesRequest()
	// (每页100)
	req.PageSize = requests.NewInteger(100)
	return Describe(GlobalClients, req, resp, resource.EcsType)
}

// AcsResponseToEcsInfo 特例函数，针对ecs的信息查询，将response转为Info
func AcsResponseToEcsInfo(accountName string, response responses.AcsResponse) (result []resource.Info, err error) {
	res, ok := response.(*ecs.DescribeInstancesResponse)
	if !ok {
		err = errECSTransferError
		return
	}
	return MyDescribeInstancesResponse(*res).Info(accountName)
}

// MyDescribeInstancesResponse 添加ecs查询响应结构体别名，方便为其添加Info方法
type MyDescribeInstancesResponse ecs.DescribeInstancesResponse

// Info 将Ecs response转换为Info信息
func (m MyDescribeInstancesResponse) Info(accountName string) (infos []resource.Info, err error) {
	for _, v := range m.Instances.Instance {
		s := parseTime(v.ExpiredTime, ecsTimeFormat)
		infos = append(
			infos,
			resource.Info{
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
				Type:    resource.ResourceMap[int(resource.EcsType)],
				Status:  s,
			},
		)
	}
	return
}
>>>>>>> dev
