package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看计算环境的创建信息
// https://cloud.tencent.com/document/api/599/15897
type DescribeComputeEnvCreateInfoRequest struct {
	// 计算环境ID
	EnvId string `name:"EnvId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComputeEnvCreateInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComputeEnvCreateInfoResponse, error) {
	resp := &DescribeComputeEnvCreateInfoResponse{}
	err := client.Request("batch", "DescribeComputeEnvCreateInfo", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeComputeEnvCreateInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 授权信息
	Authentications []*Authentication `json:"Authentications"`
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `json:"DesiredComputeNodeCount"`
	// 计算环境参数
	EnvData EnvData `json:"EnvData"`
	// 计算环境描述
	EnvDescription string `json:"EnvDescription"`
	// 计算环境 ID
	EnvId string `json:"EnvId"`
	// 计算环境名称
	EnvName string `json:"EnvName"`
	// 计算环境类型，仅支持“MANAGED”类型
	EnvType string `json:"EnvType"`
	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings"`
	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks"`
	// 通知信息
	Notifications []*Notification `json:"Notifications"`
}
