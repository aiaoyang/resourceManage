package aliyun

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aiaoyang/resourceManager/config"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

// NewSLBInfo 创建slb的信息
type NewSLBInfo struct {
	SLBName         string
	Region          string
	PayType         string
	FrontPort       string
	VServerGrupName string
	BackendPort     string
	Banwith         string
	Spec            string
	ECSInstanceIDs  []string
}

// BackendServer SLB后端服务器信息
type BackendServer struct {
	ServerID    string `json:"ServerId"`
	Type        string `json:"Type"`
	ServerIP    string `json:"ServerIp"`
	Port        string `json:"Port"`
	Description string `json:"Description"`
}

var client *slb.Client

// var err error

func init() {
	var err error
	// 临时测试用
	client, err = slb.NewClientWithAccessKey(config.GVC.Regions[0], config.GVC.Accounts[0].SecretID, config.GVC.Accounts[0].SecretKEY)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	errCreateSLB          = errors.New("create slb")
	errCreateTCPListener  = errors.New("create SLB TCPListener")
	errCreateVirtualGroup = errors.New("create VirtualGroup")
	errAddServerToGroup   = errors.New("add server to group")
)

// StartCreateSLB |
/*
	需要事务，如有失败，需要全部回退
	创建SLB顺序：
	1. 创建SLB
	2. 修改SLB监听端口与服务器端口
	3. 创建SLB虚拟服务器组
	4. 将服务器添加至虚拟服务器组
*/
func StartCreateSLB(slbInfo NewSLBInfo) (err error) {
	var slbInstanceID string

	createSLBResp, err := createSLB(slbInfo.SLBName, slbInfo.Region, slbInfo.PayType, slbInfo.Banwith, slbInfo.Spec)
	if err != nil {
		log.Println(err)
		return
	}
	if !createSLBResp.IsSuccess() {
		err = errCreateSLB
	}

	slbInstanceID = createSLBResp.LoadBalancerId
	defer func() {
		if err != nil {

			log.Printf("err detected, rollback!!!: err: %s\n", err)

			_, rollBackErr := rollBackDeleteSLB(slbInstanceID)

			if rollBackErr != nil {
				err = rollBackErr
			}
		}
	}()

	// 创建slb过程
	createSLBTCPListenerResp, err := CreateSLBTCPListener(slbInstanceID, slbInfo.FrontPort, slbInfo.BackendPort, slbInfo.Banwith)
	if err != nil {
		log.Println(err)
		return
	}
	if !createSLBTCPListenerResp.IsSuccess() {
		err = errCreateTCPListener
		return
	}

	// 创建slb后端虚拟服务器组
	createVSGroupResp, err := CreateVirtualGroupToSLB(slbInstanceID, slbInfo.VServerGrupName)
	if err != nil {
		log.Println(err)
		return
	}
	if !createVSGroupResp.IsSuccess() {
		err = errCreateVirtualGroup
		return
	}

	// 为slb后端虚拟服务器组添加ECSs
	addServerToVirtualGroupResp, err := AddServerToVirtualGroup(slbInfo.ECSInstanceIDs, slbInfo.BackendPort, createVSGroupResp.VServerGroupId)
	if err != nil {
		log.Println(err)
		return
	}
	if !addServerToVirtualGroupResp.IsSuccess() {
		err = errAddServerToGroup
		return
	}
	return nil

}

// rollBackDeleteSLB 创建slb过程出现错误需要回滚
func rollBackDeleteSLB(slbInstanceID string) (*slb.DeleteLoadBalancerResponse, error) {
	log.Println(slbInstanceID)
	request := slb.CreateDeleteLoadBalancerRequest()

	request.Scheme = "https"
	request.LoadBalancerId = slbInstanceID

	response, err := client.DeleteLoadBalancer(request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// log.Printf("response is %#v\n", response)
	return response, nil
}

// 传入参数：名称，地域，支付方式，实例带宽，实例规格
// 返回参数：实例ID，err
// createSLB 创建slb实例
func createSLB(name, region, payType, bandwith, spec string) (*slb.CreateLoadBalancerResponse, error) {

	request := slb.CreateCreateLoadBalancerRequest()
	request.Scheme = "https"
	request.RegionId = region
	request.InternetChargeType = "paybytraffic"
	request.PayType = payType
	request.Bandwidth = requests.Integer(bandwith)

	request.LoadBalancerName = name
	request.LoadBalancerSpec = spec

	response, err := client.CreateLoadBalancer(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func setSLBAttr(slbInstanceID string) {

}

// CreateSLBTCPListener |
// 传入参数：负载均衡实例ID，前端监听端口，后端监听端口，带宽
func CreateSLBTCPListener(loadBalancerID, frontPort, backendPort, bandwith string) (*slb.CreateLoadBalancerTCPListenerResponse, error) {

	request := slb.CreateCreateLoadBalancerTCPListenerRequest()
	request.Scheme = "https"
	request.LoadBalancerId = loadBalancerID

	request.BackendServerPort = requests.Integer(backendPort)
	request.ListenerPort = requests.Integer(frontPort)
	request.Bandwidth = requests.Integer(bandwith)

	response, err := client.CreateLoadBalancerTCPListener(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//CreateVirtualGroupToSLB |
// 传入参数：负载均衡实例ID，虚拟服务器组名称
// 返回参数：虚拟服务器组ID，err
func CreateVirtualGroupToSLB(loadBalancerID, vServerGroupName string) (*slb.CreateVServerGroupResponse, error) {

	request := slb.CreateCreateVServerGroupRequest()
	request.Scheme = "https"
	request.LoadBalancerId = loadBalancerID
	request.VServerGroupName = vServerGroupName

	response, err := client.CreateVServerGroup(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// AddServerToVirtualGroup |
// 传入参数：虚拟服务器组实例ID，后端服务器信息
// 返回参数：response，err
func AddServerToVirtualGroup(ecsInstanceIDs []string, ecsPort, vServerGroupID string) (*slb.AddVServerGroupBackendServersResponse, error) {

	request := slb.CreateAddVServerGroupBackendServersRequest()

	request.Scheme = "https"
	request.VServerGroupId = vServerGroupID

	servers := []BackendServer{}

	for _, ecsInstanceID := range ecsInstanceIDs {
		servers = append(servers, BackendServer{
			ServerID: ecsInstanceID,
			Type:     "ECS",
			Port:     ecsPort,
		})
	}
	serversString, err := json.Marshal(servers)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	request.BackendServers = fmt.Sprintf("%s", serversString)

	response, err := client.AddVServerGroupBackendServers(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
