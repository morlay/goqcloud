package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 给客户转账
// https://cloud.tencent.com/document/api/563/19925

type AgentTransferMoneyRequest struct {
	// 转账金额，单位分
	Amount int64 `name:"Amount"`
	// 客户账号ID
	ClientUin string `name:"ClientUin"`
	// 区域
	Region string `name:"Region"`
}

func (req *AgentTransferMoneyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AgentTransferMoneyResponse, error) {
	resp := &AgentTransferMoneyResponse{}
	err := client.Request("partners", "AgentTransferMoney", "2018-03-21").Do(req, resp)
	return resp, err
}

type AgentTransferMoneyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
