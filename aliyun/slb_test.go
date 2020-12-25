package aliyun

import (
	"log"
	"testing"
)

func Test_GetSLB(t *testing.T) {
	log.SetFlags(log.Llongfile | log.Ldate)
	slbInfo := NewSLBInfo{
		SLBName:         "fortest",
		Region:          "cn-beijing",
		FrontPort:       "10000",
		VServerGrupName: "test_server",
		BackendPort:     "10000",
		Banwith:         "10",
		ECSInstanceIDs:  []string{"i-ecstest"},
	}
	StartCreateSLB(slbInfo)
}
