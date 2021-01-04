package common

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aiaoyang/resourceManager/resource"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetAlarm 获取告警信息
func GetAlarm() ([]resource.Info, error) {
	var req = NewDescribeAlarmRequest()
	var resp = cms.CreateDescribeAlertHistoryListResponse()
	return Describe(GlobalClients, req, resp, resource.AlertType)
	// return describeAlarm()
}

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

// MyDescribeAlertHistoryListResponse 添加RDS查询响应结构体别名，方便为其添加Info方法
type MyDescribeAlertHistoryListResponse cms.DescribeAlertHistoryListResponse

// Info 将Alert response转换为Info信息
func (m MyDescribeAlertHistoryListResponse) Info(accountName string) (infos []resource.Info, err error) {
	infos = append(
		infos,
		resource.Info{
			Name:   accountName,
			Detail: fmt.Sprintf("%d", len(m.AlarmHistoryList.AlarmHistory)),
			Type:   resource.ResourceMap[int(resource.AlertType)],
		},
	)
	return
}

// AcsResponseToAlarmInfo 特例函数，针对告警的信息查询，将response转为Info
func AcsResponseToAlarmInfo(accountName string, response responses.AcsResponse) (result []resource.Info, err error) {
	res, ok := response.(*cms.DescribeAlertHistoryListResponse)
	if !ok {
		err = errRDSTransferError
		return
	}
	return MyDescribeAlertHistoryListResponse(*res).Info(accountName)
}
