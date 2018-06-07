package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 审核客户
// https://cloud.tencent.com/document/api/563/16044

type AuditApplyClientRequest struct {
	// 审核结果，可能的取值：accept/reject
	AuditResult string `name:"AuditResult"`
	// 待审核客户账号ID
	ClientUin string `name:"ClientUin"`
	// 申请理由，B类客户审核通过时必须填写申请理由
	Note string `name:"Note"`
	// 区域
	Region string `name:"Region"`
}

func (req *AuditApplyClientRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AuditApplyClientResponse, error) {
	resp := &AuditApplyClientResponse{}
	err := client.Request("partners", "AuditApplyClient", "2018-03-21").Do(req, resp)
	return resp, err
}

type AuditApplyClientResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 关联时间对应的时间戳
	AgentTime int64 `json:"AgentTime"`
	// 审核结果，包括accept/reject/qcloudaudit（腾讯云审核）
	AuditResult string `json:"AuditResult"`
	// 客户账号ID
	ClientUin string `json:"ClientUin"`
	// 代理商账号ID
	Uin string `json:"Uin"`
}
