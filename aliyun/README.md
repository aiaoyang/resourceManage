添加新资源需要添加的内容：
resource.go -> 
const (

	// EcsType ecs资源类型
	EcsType ResourceType = iota

	// RdsType rds资源类型
	RdsType

	// DomainType 域名资源类型
	DomainType

    // XXX XXX资源类型
    XXXType
)


request.go -> 

func NewXXXXRequest()XXXRequest

response.go -> 

type MyXXXResponse XXX.XXXResponse

func (x MyXXXResponse)Info(name string)(infos []Info){}

func AcsResponseToXXXInfo(accountName string, response responses.AcsResponse) (result []Info, err error)

资源名.go

// Client XXX请求客户端
type Client struct {
	*XXX.Client
	AccountName string
}

// Name 返回客户端的账号名
func (c Client) Name() string {
	return c.AccountName
}

// clients 客户端列表
var XXXClients []AliyunClient

func init() {

	for _, region := range config.GVC.Regions {
		for _, m := range config.GVC.Accounts {
			c, err := ecs.NewClientWithAccessKey(region, m.SecretID, m.SecretKEY)
			if err != nil {
				log.Fatal(err)
			}
			tmp := Client{c, m.Name}
			ecsClients = append(ecsClients, tmp)
		}
	}
}

// GetECS 查询ecs列表
func GetXXX() ([]Info, error) {
	var resp = XXX.DescribeXXXResponse{}
	return Describe(ecsClients, DescribeXXXRequest(), resp, XXXType)
}
