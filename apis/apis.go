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

// func GetLifeTime() (string, error) {
// 	return aliyun.GetLifeTime()
// }
func GetAlarm() (string, error) {
	alarm, err := aliyun.GetAlarm()
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(alarm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", res), nil

}
func GetDomain() (string, error) {
	domain, err := aliyun.GetDomain()
	if err != nil {
		return "", err
	}

	res, err := json.Marshal(domain)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", res), nil
}

func GetRedis() {

}

func GetDelay() {

}
