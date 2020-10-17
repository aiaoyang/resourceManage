package apis

import (
	"encoding/json"
	"fmt"

	"github.com/aiaoyang/resourceManager/aliyun"
)

func GetResource() {

}

func GetECS() (string, error) {
	ecsinfo, err := aliyun.GetECS()
	if err != nil {
		return "", err
	}
	res, err := json.Marshal(ecsinfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", res), nil
}
func GetRDS() (string, error) {
	rdsinfo, err := aliyun.GetRDS()
	if err != nil {
		return "", err
	}
	res, err := json.Marshal(rdsinfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", res), nil
}

func GetLifeTime() (string, error) {
	return aliyun.GetLifeTime()
}
func GetDomain() {}

func GetRedis() {

}

func GetAlarm() {}

func GetDelay() {

}
