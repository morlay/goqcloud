package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级实例
// https://cloud.tencent.com/document/api/239/20013

type UpgradeInstanceRequest struct {
	// 升级的实例Id
	InstanceId string `name:"InstanceId"`
	// 规格 单位 MB
	MemSize int64 `name:"MemSize"`
	// 区域
	Region string `name:"Region"`
}

func (req *UpgradeInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeInstanceResponse, error) {
	resp := &UpgradeInstanceResponse{}
	err := client.Request("redis", "UpgradeInstance", "2018-04-12").Do(req, resp)
	return resp, err
}

type UpgradeInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单ID
	DealId string `json:"DealId"`
}
