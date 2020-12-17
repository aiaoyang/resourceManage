package aliyun

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/aliyun/request"
	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// AlarmInfo 告警信息
type AlarmInfo struct {
	Account string `json:"account"`
	Count   int    `json:"count"`
}

// AlarmClient 告警请求客户端
type AlarmClient struct {
	client *cms.Client

	Name string
}

// AlarmResponse 告警返回信息结构体
type AlarmResponse struct {
	response *cms.DescribeAlertHistoryListResponse
	Name     string
}

var alarmClients []AlarmClient

func init() {

	for _, region := range config.GVC.Regions {
		for _, m := range config.GVC.Accounts {

			c, err := cms.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}

			tmp := AlarmClient{c, m.Name}

			alarmClients = append(alarmClients, tmp)

		}
	}

}

// GetAlarm 获取告警信息
func GetAlarm() ([]AlarmInfo, error) {
	return describeAlarm()
}

func describeAlarm() ([]AlarmInfo, error) {

	responsesChan := make(chan AlarmResponse, len(alarmClients))

	wg := &sync.WaitGroup{}

	wg.Add(len(alarmClients))

	/*
		func (AlarmClients,AlarmResponse,DescribeAlarmRequest)
	*/
	for _, c := range alarmClients {

		go func(wg *sync.WaitGroup, ch chan AlarmResponse, client AlarmClient) {

			defer wg.Done()

			req := request.NewDescribeAlarmRequest()

			req.SetScheme("https")

			resp, err := client.client.DescribeAlertHistoryList(req)

			if err != nil {

				log.Println(err)

				return

			}

			defer resp.GetOriginHttpResponse().Body.Close()

			tmp := AlarmResponse{resp, client.Name}

			fmt.Println(tmp.response)

			ch <- tmp

		}(wg, responsesChan, c)

	}

	wg.Wait()

	close(responsesChan)

	alarmsMap := make(map[string]int, 0)

	for responses := range responsesChan {
		alarmHistory := responses.response.AlarmHistoryList.AlarmHistory
		alarmsMap[responses.Name] += len(alarmHistory)
	}
	alarms := []AlarmInfo{}
	for accountName, count := range alarmsMap {
		alarms = append(alarms, AlarmInfo{accountName, count})
	}
	fmt.Println(alarms)
	return alarms, nil

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
