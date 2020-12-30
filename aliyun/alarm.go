package aliyun

import (
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// GetAlarm 获取告警信息
func GetAlarm() ([]Info, error) {
	var req = NewDescribeAlarmRequest()
	var resp = cms.CreateDescribeAlertHistoryListResponse()
	return Describe(GlobalClients, req, resp, AlertType)
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
