package aliyun

import (
	"github.com/aiaoyang/resourceManager/aliyun/common"
)

func GetCert() (infos []common.Info, err error) {
	return common.GetCert()
}

func GetECS() (infos []common.Info, err error) {
	return common.GetECS()
}
func GetRDS() (infos []common.Info, err error) {
	return common.GetRDS()
}

func GetDomain() (infos []common.Info, err error) {
	return common.GetDomain()
}

func GetAlarm() (infos []common.Info, err error) {
	return common.GetAlarm()
}

// func GetSLB() (infos []common.Info, err error) {
// 	return common.GetSLB()
// }
