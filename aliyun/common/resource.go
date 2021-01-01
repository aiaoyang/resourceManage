package common

import (
	"log"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ResourceType 资源类型枚举体
type ResourceType int

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

const (

	// EcsType ecs资源类型
	EcsType ResourceType = iota

	// RdsType rds资源类型
	RdsType

	// DomainType 域名资源类型
	DomainType

	// CertType 证书资源类型
	CertType

	// AlertType 告警资源类型
	AlertType
)

var (
	// ResourceMap 资源类型名称
	ResourceMap = map[int]string{
		0: "ECS",
		1: "RDS",
		2: "Domain",
		3: "Cert",
		4: "Alert",
	}
)

// Describe 通用调用入口
func Describe(
	// 客户端列表
	clients []IClient,
	// 请求结构体
	request requests.AcsRequest,
	// 响应结构体
	response responses.AcsResponse,
	// 资源类型
	resourceType ResourceType,

) (result []Info, err error) {

	wg := &sync.WaitGroup{}
	wg.Add(len(clients))
	ch := make(chan res, len(clients))

	for _, client := range clients {
		go doRequest(wg, ch, client, request, response, resourceType)
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

type res struct {
	infos []Info
	err   error
}

func doRequest(
	wg *sync.WaitGroup,
	ch chan res,
	client IClient,
	request requests.AcsRequest,
	response responses.AcsResponse,
	resourceType ResourceType,
) {
	defer wg.Done()
	err := client.DoAction(request, response)
	if err != nil {
		log.Println(err)
		ch <- res{err: err}
		return
	}

	i, e := ResponseToResult(client.Name(), response, resourceType)
	log.Println("123")
	ch <- res{
		infos: i,
		err:   e,
	}

}