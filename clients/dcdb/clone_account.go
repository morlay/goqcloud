package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 克隆实例账户
// https://cloud.tencent.com/document/api/557/20255

type CloneAccountRequest struct {
	// 目的用户账户描述
	DstDesc *string `name:"DstDesc,omitempty"`
	// 目的用户HOST
	DstHost string `name:"DstHost"`
	// 目的用户账户名
	DstUser string `name:"DstUser"`
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 源用户HOST
	SrcHost string `name:"SrcHost"`
	// 源用户账户名
	SrcUser string `name:"SrcUser"`
}

func (req *CloneAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloneAccountResponse, error) {
	resp := &CloneAccountResponse{}
	err := client.Request("dcdb", "CloneAccount", "2018-04-11").Do(req, resp)
	return resp, err
}

type CloneAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务流程ID
	FlowId int64 `json:"FlowId"`
}
