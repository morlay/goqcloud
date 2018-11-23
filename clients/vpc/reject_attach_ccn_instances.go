package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 云联网拒绝关联实例
// https://cloud.tencent.com/document/api/215/20250

type RejectAttachCcnInstancesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 拒绝关联实例列表。
	Instances []*CcnInstance `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *RejectAttachCcnInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RejectAttachCcnInstancesResponse, error) {
	resp := &RejectAttachCcnInstancesResponse{}
	err := client.Request("vpc", "RejectAttachCcnInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type RejectAttachCcnInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
