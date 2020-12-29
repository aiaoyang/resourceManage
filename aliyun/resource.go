package aliyun

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// Info 信息
type Info struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Detail    string `json:"detail"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

// Client ecs请求客户端
type Client struct {
	*ecs.Client
	AccountName string
}

// Clients 客户端列表
var Clients []Client

func init() {

	for _, region := range config.GVC.Regions {

		for _, m := range config.GVC.Accounts {
			c, err := ecs.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)

			if err != nil {
				log.Fatal(err)
			}

			tmp := Client{c, m.Name}

			Clients = append(Clients, tmp)

		}

	}

}

// DescribeECSRequest 生成获取ecs信息的请求request
func DescribeECSRequest() *ecs.DescribeInstancesRequest {

	request := ecs.CreateDescribeInstancesRequest()

	request.Scheme = "https"

	request.PageSize = requests.NewInteger(100)

	return request

}

// MydescribeInstancesResponse 添加ecs查询响应结构体别名，方便为其添加Info方法
type MydescribeInstancesResponse ecs.DescribeInstancesResponse

// Info 将response转换为Info信息
func (m MydescribeInstancesResponse) Info(name string) (infos []Info) {
	for _, v := range m.Instances.Instance {

		s := parseTime(v.ExpiredTime, ecsTimeFormat)

		infos = append(
			infos,
			Info{
				Name: v.InstanceName,
				// Index:  fmt.Sprintf("%d", index),
				Detail: fmt.Sprintf("%dC-%.1fG", v.Cpu, float32(v.Memory)/1024.0),
				EndOfTime: func(endOfTime string) string {
					if endOfTime == "" {
						return "后付费"
					}
					return endOfTime
				}(v.ExpiredTime),
				Account: name,
				Type:    "ECS",
				Status:  s,
			},
		)
	}
	return
}

// Name 返回客户端的账号名
func (c Client) Name() string {
	return c.AccountName
}

// AliyunClient 实现阿里云基础Client接口和自定义的添加客户端账号名称的接口
type AliyunClient interface {

	// 阿里云基本Client所拥有的方法
	DoAction(request requests.AcsRequest, response responses.AcsResponse) (err error)

	// 返回客户端账号名方法
	Name() string
}

// ResourceType 资源类型枚举体
type ResourceType int

const (

	// EcsType ecs资源类型
	EcsType ResourceType = iota

	// RdsType rds资源类型
	RdsType

	// DomainType 域名资源类型
	DomainType
)

// AcsResponseToEcsInfo 特例函数，针对ecs的信息查询，将response转为Info
func AcsResponseToEcsInfo(accountName string, response responses.AcsResponse) (result []Info, err error) {
	res, ok := response.(*ecs.DescribeInstancesResponse)
	if !ok {
		err = fmt.Errorf("response 类型不为 DescribeInstancesResponse")
		return
	}
	result = MydescribeInstancesResponse(*res).Info(accountName)
	return
}

// ResponseToResult 通用函数 将aliyun Response转为我们所需要Info
func ResponseToResult(accountName string, response responses.AcsResponse, resourceType ResourceType) (result []Info, err error) {

	switch resourceType {
	case EcsType:
		result, err = AcsResponseToEcsInfo(accountName, response)
		return

	case DomainType:

	default:
	}
	return nil, fmt.Errorf("资源类型传参错误")
}

// Describe 通用调用入口
func Describe(clients []AliyunClient, request requests.AcsRequest, response responses.AcsResponse, resourceType ResourceType) (result []Info, err error) {

	type res struct {
		infos []Info
		err   error
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(clients))
	ch := make(chan res, len(clients))

	for _, client := range clients {

		go func(
			wg *sync.WaitGroup,
			ch chan res,
			client AliyunClient,
			request requests.AcsRequest,
			response responses.AcsResponse,
		) {

			err = client.DoAction(request, response)
			if err != nil {
				return
			}
			i, e := ResponseToResult(client.Name(), response, resourceType)

			ch <- res{
				infos: i,
				err:   e,
			}
			wg.Done()
		}(wg, ch, client, request, response)
	}
	wg.Wait()
	close(ch)
	for info := range ch {
		if info.err != nil {
			err = info.err
		} else {
			result = append(result, info.infos...)
		}
	}

	return
}

type timeFormat string

const (
	ecsTimeFormat  timeFormat = "2006-01-02T15:04Z"
	certTimeFormat timeFormat = "2006-01-02"
)

func parseTime(timeString string, tFormat timeFormat) (s stat) {
	pTime, err := time.Parse(string(tFormat), timeString)
	if err != nil {
		log.Fatal(err)
	}

	s = green

	if time.Now().AddDate(0, 1, 0).After(pTime) {
		s = yellow
	}
	if time.Now().AddDate(0, 0, 7).After(pTime) {
		s = red
	}
	if time.Now().After(pTime) {
		s = nearDead
	}
	return s
}
