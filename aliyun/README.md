添加新资源需要添加的内容：
request.go -> func NewXXXXRequest()XXXRequest

response.go -> 

type MyXXXResponse XXX.XXXResponse

func (x MyXXXResponse)Info(name string)(infos []Info){}

func AcsResponseToXXXInfo(accountName string, response responses.AcsResponse) (result []Info, err error)

资源名.go

