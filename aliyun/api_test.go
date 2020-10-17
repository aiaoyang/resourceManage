package aliyun

import (
	"log"
	"testing"

	"github.com/aiaoyang/resourceManager/config"
)

func Test_ecs(t *testing.T) {
	testEcs()
}
func Test_alarm(t *testing.T) {

	account := Account{
		Name:      "yongshiwl",
		AccessID:  "LTAI4GDH3nRSaFMS2Twdkh8p",
		AccessKEY: "cUtiTejtC8qNfyjMcshctoiGsY7MDo",
	}
	res, err := getSingleAccountAlarm(account)
	if err != nil {
		log.Fatal(err)

	}
	log.Println(res)
}

func Test_lifeTime(t *testing.T) {
	getLifeTime()
	// log.Printf("123")
}

func Test_viper(t *testing.T) {
	config.LoadAliyunConfig()
	// v := viper.New()
	// v.AddConfigPath("./")
	// v.SetConfigType("yaml")
	// err := v.ReadInConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// keys := v.AllKeys()
	// log.Printf("all keys : %v\n", keys)
	// sliceMap := v.Get("accounts")
	// t1 := reflect.ValueOf(sliceMap)
	// switch t1.Kind() {
	// case reflect.Slice:
	// 	for i := 0; i < t1.Len(); i++ {
	// 		for k, v := range t1.Index(i).Interface().(map[interface{}]interface{}) {
	// 			fmt.Println(k.(string), v.(string))
	// 		}
	// 	}
	// }

}
