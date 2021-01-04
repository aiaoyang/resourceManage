package aliyun

import (
	ali "github.com/aiaoyang/resourceManager/aliyun/common"
)

func GetCert() (infos []ali.Info, err error) {
	return ali.GetCert()
}

func GetECS() (infos []ali.Info, err error) {
	return ali.GetECS()
}
func GetRDS() (infos []ali.Info, err error) {
	return ali.GetRDS()
}

func GetDomain() (infos []ali.Info, err error) {
	return ali.GetDomain()
}

func GetAlarm() (infos []ali.Info, err error) {
	return ali.GetAlarm()
}

// func GetSLB() (infos []common.Info, err error) {
// 	return common.GetSLB()
// }
