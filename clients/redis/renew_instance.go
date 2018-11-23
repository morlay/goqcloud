package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例
// https://cloud.tencent.com/document/api/239/20015

type RenewInstanceRequest struct {
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 购买时长，单位：月
	Period int64 `name:"Period"`
	// 区域
	Region string `name:"Region"`
}

func (req *RenewInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("redis", "RenewInstance", "2018-04-12").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 交易Id
	DealId string `json:"DealId"`
}
