package common

import (
	"fmt"
	"strings"
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
	requestFunc func() requests.AcsRequest,
	// 响应结构体
	responseFunc func() responses.AcsResponse,
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
			requestFunc,
			responseFunc,
			resourceType,
		)
	}

	wg.Wait()

	close(resultChan)
	var errSlice = []string{}
	for info := range resultChan {
		if info.err != nil {
			errSlice = append(errSlice, info.err.Error())
		} else {
			result = append(result, info.infos...)
		}
	}
	if len(errSlice) != 0 {
		err = fmt.Errorf("%s", strings.Join(errSlice, "\n"))
	}

	return
}

func doRequest(
	wg *sync.WaitGroup,
	resultChan chan infoResult,
	client IClient,
	requestFunc func() requests.AcsRequest,
	responseFunc func() responses.AcsResponse,
	resourceType resource.Type,
) {
	request := requestFunc()
	response := responseFunc()

	defer wg.Done()

	err := client.DoAction(request, response)

	if err != nil {

		resultChan <- infoResult{
			err: err,
		}

		return
	}

	infos, err := ResponseToResult(client.Name(), response, resourceType)

	resultChan <- infoResult{
		infos,
		err,
	}

}
