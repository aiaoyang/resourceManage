package common

import (
	"sync"

	"github.com/aiaoyang/resourceManager/resource"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type infoResult struct {
	infos []resource.Info
	err   error
}

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
	resultChan := make(chan infoResult, len(clients))

	for _, client := range clients {

		go doRequest(
			wg,
			resultChan,
			client,
			request,
			response,
			resourceType,
		)
	}

	wg.Wait()

	close(resultChan)

	for info := range resultChan {
		if info.err != nil {
			err = info.err
		} else {
			result = append(result, info.infos...)
		}
	}

	return
}

func doRequest(
	wg *sync.WaitGroup,
	ch chan infoResult,
	client IClient,
	request requests.AcsRequest,
	response responses.AcsResponse,
	resourceType resource.Type,
) {
	defer wg.Done()
	err := client.DoAction(request, response)
	if err != nil {
		ch <- infoResult{err: err}
		return
	}

	i, e := ResponseToResult(client.Name(), response, resourceType)
	ch <- infoResult{
		infos: i,
		err:   e,
	}

}
