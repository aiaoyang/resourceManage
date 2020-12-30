package aliyun

import (
	"log"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// Client 请求客户端结构体
type Client struct {
	iAliClient
	AccountName string
}

// Name 返回账号名
func (c Client) Name() string {
	return c.AccountName
}

// IClient 包内 客户端接口
type IClient interface {
	iAliClient
	Name() string
}

// iAliClient 实现阿里云基础Client接口和自定义的添加客户端账号名称的接口
type iAliClient interface {
	// 阿里云 Base Client所拥有的方法
	DoAction(request requests.AcsRequest, response responses.AcsResponse) (err error)
}

// GlobalClients 全局客户端
var GlobalClients []IClient

// NewClients 生成新的客户端列表
func NewClients() (clients []IClient) {
	clients = make([]IClient, 0)
	for _, region := range config.GVC.Regions {
		for _, m := range config.GVC.Accounts {
			iAliClient, err := sdk.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)
			if err != nil {
				log.Fatal(err)
			}
			tmp := Client{
				iAliClient,
				m.Name,
			}
			clients = append(clients, tmp)
		}
	}
	return
}

func init() {
	GlobalClients = NewClients()

}
