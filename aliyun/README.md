添加新资源需要添加的内容：
resource.go -> 
```
const (

	// EcsType ecs资源类型
	EcsType ResourceType = iota

	// RdsType rds资源类型
	RdsType

	// DomainType 域名资源类型
	DomainType

    // XXX XXX资源类型
->  XXXType
)
```

timeFormator.go ->
```
const (
	ecsTimeFormat    timeFormat = "2006-01-02T15:04Z"
	certTimeFormat   timeFormat = "2006-01-02"
	domainTimeFormat timeFormat = "2006-01-02 15:04:05"
	rdsTimeFormat    timeFormat = "2006-01-02T15:04:05Z"

->	XXXTimeFormat    timeFormat = "XXXX"
)
```


response.go -> 
```
type MyXXXResponse XXX.XXXResponse

func (x MyXXXResponse)Info(name string)(infos []Info){}

func AcsResponseToXXXInfo(accountName string, response responses.AcsResponse) (result []Info, err error)
```
资源名.go
```
// GetXXX 查询XXX列表
func GetXXX() ([]Info, error) {
	var resp = XXX.CreateDescribeXXXResponse()
	var req = XXX.CreateDescribeXXXRequest()
	return Describe(ecsClients, DescribeXXXRequest(), resp, XXXType)
}
```
