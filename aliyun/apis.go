package aliyun

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

var regions map[int]string

func init() {
	regions = make(map[int]string)
	regions[0] = "cn-shanghai"
	regions[1] = "cn-hangzhou"
	regions[2] = "ap-southeast-1"
}

type Alarm struct {
	Account    string `json:"account"`
	AlarmCount int    `json:"alarmCount"`
	Region     string `json:"region"`
	// State      string `json:"state"`
}

func GetAlarm() ([]*Alarm, error) {

	return nil, nil
	// alarmsChan := make(chan *Alarm, 4)
	// wg := &sync.WaitGroup{}
	// wg.Add(4)
	// for _, account := range []Account{account1, account2, account3, account4} {
	// 	go func(wg *sync.WaitGroup, ch chan *Alarm, acc Account) {
	// 		alarmTmp, err := getSingleAccountAlarm(acc)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		alarmsChan <- alarmTmp
	// 		wg.Done()
	// 	}(wg, alarmsChan, account)
	// }
	// wg.Wait()
	// close(alarmsChan)

	// // time.Sleep(time.Second * 10)
	// // close(alarmsChan)
	// alarms := []*Alarm{}
	// for alarm := range alarmsChan {
	// 	alarms = append(alarms, alarm)
	// }
	// log.Printf("alarms: %v\n", alarms)
	// return alarms, nil
}

func testAlarm() ([]Alarm, error) {
	log.SetFlags(log.Llongfile | log.Ldate)
	client, err := cms.NewClientWithAccessKey("cn-shanghai", "0", "0")

	request := cms.CreateDescribeAlertHistoryListRequest()
	request.Scheme = "https"
	request.State = "ALARM"

	response, err := client.DescribeAlertHistoryList(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	alarms := []Alarm{}
	for index, v := range response.AlarmHistoryList.AlarmHistory {
		index++
		fmt.Printf("%v\n", v)

		tmp := Alarm{
			Account:    "yongshiwl",
			AlarmCount: 1,
			// State: v.State,
		}
		alarms = append(alarms, tmp)
	}
	defer response.GetOriginHttpResponse().Body.Close()
	log.Println(alarms)
	return alarms, nil
}

type Account struct {
	Name      string
	AccessID  string
	AccessKEY string
}

func getSingleAccountAlarm(account Account) (alarm *Alarm, err error) {
	log.SetFlags(log.Llongfile | log.Ldate)

	clients := make([]*cms.Client, 0)

	client, err := cms.NewClientWithAccessKey("cn-shanghai", account.AccessID, account.AccessKEY)

	if err != nil {
		return nil, err
	}

	clients = append(clients, client)

	request := cms.CreateDescribeAlertHistoryListRequest()

	request.Scheme = "https"

	request.State = "ALARM"
	request.StartTime = strconv.FormatInt(time.Now().AddDate(0, 0, -1).Unix(), 10)
	request.EndTime = strconv.FormatInt(time.Now().Unix(), 10)

	tmp := Alarm{}

	response, err := client.DescribeAlertHistoryList(request)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer response.GetOriginHttpResponse().Body.Close()

	tmp.Account = account.Name

	tmp.AlarmCount = len(response.AlarmHistoryList.AlarmHistory)

	return &tmp, nil

}

func GetLifeTime() (string, error) {
	return getLifeTime()
}

func getLifeTime() (string, error) {

	client, err := bssopenapi.NewClientWithAccessKey("cn-hangzhou", "0", "0")

	request := bssopenapi.CreateQueryAvailableInstancesRequest()

	request.Scheme = "https"

	response, err := client.QueryAvailableInstances(request)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer response.GetOriginHttpResponse().Body.Close()

	fmt.Printf("response is %#v\n", response.Data.InstanceList)

	res, err := json.Marshal(response.Data.InstanceList)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", res), nil
}
