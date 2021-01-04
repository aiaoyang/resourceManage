package common

import (
	"log"
	"sync"

	"github.com/aiaoyang/resourceManager/resource"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	resourceType resource.Type,

) (result []resource.Info, err error) {

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
	infos []resource.Info
	err   error
}

func doRequest(
	wg *sync.WaitGroup,
	ch chan res,
	client IClient,
	request requests.AcsRequest,
	response responses.AcsResponse,
	resourceType resource.Type,
) {
	defer wg.Done()
	err := client.DoAction(request, response)
	if err != nil {
		log.Println(err)
		ch <- res{err: err}
		return
	}

	i, e := ResponseToResult(client.Name(), response, resourceType)
	ch <- res{
		infos: i,
		err:   e,
	}

}
