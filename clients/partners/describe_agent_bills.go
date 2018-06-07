package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询业务明细
// https://cloud.tencent.com/document/api/563/16042

type DescribeAgentBillsRequest struct {
	// 客户备注名称
	ClientRemark *string `name:"ClientRemark,omitempty"`
	// 客户账号ID
	ClientUin *string `name:"ClientUin,omitempty"`
	// 限制数目
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 预付费订单号
	OrderId *string `name:"OrderId,omitempty"`
	// 支付方式，prepay/postpay
	PayMode *string `name:"PayMode,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 支付月份，如2018-02
	SettleMonth *string `name:"SettleMonth,omitempty"`
}

func (req *DescribeAgentBillsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAgentBillsResponse, error) {
	resp := &DescribeAgentBillsResponse{}
	err := client.Request("partners", "DescribeAgentBills", "2018-03-21").Do(req, resp)
	return resp, err
}

type DescribeAgentBillsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 业务明细列表
	AgentBillSet []*AgentBillElem `json:"AgentBillSet"`
	// 符合查询条件列表总数量
	TotalCount int64 `json:"TotalCount"`
}
