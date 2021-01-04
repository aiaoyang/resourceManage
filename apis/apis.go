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

func errHandler(funcs ...GetResourceFunc) (res string, errs error) {

	errSlice := []string{}
	resultSlice := []resource.Info{}

	for _, fn := range funcs {

		result, err := fn()

		if err != nil && len(result) == 0 {

			errSlice = append(errSlice, err.Error())

			continue

		}

		resultSlice = append(resultSlice, result...)

	}

	// 如果都失败了，则将所有错误打包返回errs，令res为空
	if len(resultSlice) == 0 {
		errs = fmt.Errorf("%s", strings.Join(errSlice, "\n"))
		return
	}

	// 如果有成功也有失败，则错误和成功结果分别放入errs和res返回
	resultByte, err := json.Marshal(resultSlice)

	if err != nil {
		errSlice = append(errSlice, err.Error())
	}

	errs = fmt.Errorf("%s", strings.Join(errSlice, "\n"))

	res = fmt.Sprintf("%s", resultByte)

	return
}
