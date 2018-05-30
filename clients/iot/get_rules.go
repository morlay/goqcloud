package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取转发规则列表
// https://cloud.tencent.com/document/api/568/16462
type GetRulesRequest struct {
	// 长度
	Length *int64 `name:"Length,omitempty"`
	// 偏移
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetRulesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetRulesResponse, error) {
	resp := &GetRulesResponse{}
	err := client.Request("iot", "GetRules", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetRulesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 规则列表
	Rules []*Rule `json:"Rules"`
}
