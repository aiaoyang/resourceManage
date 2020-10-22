package aliyun

// import (
// 	"sync"

// 	"github.com/aiaoyang/resourceManager/config"
// )

// type genClient func(region, account string) (client interface{})

// func NewClients(
// 	conf config.AliyunConfig,
// 	fn genClient,
// ) ([]interface{}, error) {

// 	num := len(conf.Regions) * len(conf.Accounts)

// 	clientsChan := make(chan interface{}, num)

// 	wg := &sync.WaitGroup{}

// 	wg.Add(num)

// 	for _, region := range conf.Regions {
// 		for _, account := range conf.Accounts {

// 			go func(
// 				wg *sync.WaitGroup,
// 				clientChan chan interface{},
// 				fn genClient,
// 				region, account string,
// 			) {

// 				clientChan <- fn(region, account)

// 				defer wg.Done()

// 			}(wg, clientsChan, fn, region, account.Name)

// 		}
// 	}

// 	wg.Wait()
// 	close(clientsChan)

// 	clients := make([]interface{}, 0)

// 	for client := range clientsChan {
// 		clients = append(clients, client)
// 	}

// 	return clients, nil
// }
