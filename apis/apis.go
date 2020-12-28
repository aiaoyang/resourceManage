package apis

import (
	"encoding/json"
	"fmt"

	"github.com/aiaoyang/resourceManager/aliyun"
)

func GetResource() {

}

func GetECS() (string, error) {
	return errHandler(aliyun.GetECS())
}
func GetRDS() (string, error) {
	return errHandler(aliyun.GetRDS())
}

// func GetLifeTime() (string, error) {
// 	return aliyun.GetLifeTime()
// }
func GetAlarm() (string, error) {
	return errHandler(aliyun.GetAlarm())

}
func GetDomain() (string, error) {
	return errHandler(aliyun.GetDomain())
}

func GetRedis() {

}

func GetDelay() {

}

func GetCert() (string, error) {
	return errHandler(aliyun.GetCertList())
}

func errHandler(payload interface{}, err error) (res string, rerr error) {
	if err != nil {
		rerr = err
		return
	}
	tmpRes, rerr := json.Marshal(payload)
	if rerr != nil {
		return
	}
	res = fmt.Sprintf("%s", tmpRes)
	return
}
