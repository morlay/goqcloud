package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 支付订单（合作伙伴使用）
// https://cloud.tencent.com/document/api/563/19186

type AgentPayDealsRequest struct {
	// 代付标志，1：代付；0：自付
	AgentPay int64 `name:"AgentPay"`
	// 订单号数组
	DealNames []*string `name:"DealNames"`
	// 订单所有者uin
	OwnerUin string `name:"OwnerUin"`
	// 区域
	Region string `name:"Region"`
}

func (req *AgentPayDealsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AgentPayDealsResponse, error) {
	resp := &AgentPayDealsResponse{}
	err := client.Request("partners", "AgentPayDeals", "2018-03-21").Do(req, resp)
	return resp, err
}

type AgentPayDealsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
