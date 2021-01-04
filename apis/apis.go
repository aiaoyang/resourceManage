package apis

import (
	"encoding/json"
	"fmt"
	"strings"

	ali "github.com/aiaoyang/resourceManager/aliyun/common"
	"github.com/aiaoyang/resourceManager/resource"
)

// GetResourceFunc 获取资源信息的函数
type GetResourceFunc func() ([]resource.Info, error)

// GetECS 查询ecs
func GetECS() (string, error) {
	return errHandler(ali.GetECS)
}

// GetRDS 查询rds
func GetRDS() (string, error) {
	return errHandler(ali.GetRDS)
}

//GetAlarm 查询告警
func GetAlarm() (string, error) {
	return errHandler(ali.GetAlarm)

}

// GetDomain 查询域名
func GetDomain() (string, error) {
	return errHandler(ali.GetDomain)
}

// GetRedis 查询redis
func GetRedis() {

}

// GetCert 查询证书
func GetCert() (string, error) {
	return errHandler(ali.GetCert)
}

// func errHandlerOld(payload interface{}, err error) (res string, rerr error) {
// 	if err != nil {
// 		rerr = err
// 		return
// 	}
// 	tmpRes, rerr := json.Marshal(payload)
// 	if rerr != nil {
// 		return
// 	}
// 	res = fmt.Sprintf("%s", tmpRes)
// 	return
// }

func errHandler(funcs ...GetResourceFunc) (res string, errs error) {

	errSlice := []string{}
	resultSlice := []resource.Info{}

	for _, fn := range funcs {

		result, err := fn()

		if err != nil {

			errSlice = append(errSlice, err.Error())

			continue

		}

		resultSlice = append(resultSlice, result...)

	}
	if errSlice != nil && len(resultSlice) == 0 {

		errs = fmt.Errorf("%s", strings.Join(errSlice, "\n"))

		return
	}

	resultByte, err := json.Marshal(resultSlice)

	if err != nil {

		errs = err

		return

	}

	res = fmt.Sprintf("%s", resultByte)

	return
}
