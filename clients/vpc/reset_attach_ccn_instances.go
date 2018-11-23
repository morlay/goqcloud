package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重新申请关联实例
// https://cloud.tencent.com/document/api/215/20249

type ResetAttachCcnInstancesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// CCN所属UIN（根账号）。
	CcnUin string `name:"CcnUin"`
	// 重新申请关联网络实例列表。
	Instances []*CcnInstance `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResetAttachCcnInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetAttachCcnInstancesResponse, error) {
	resp := &ResetAttachCcnInstancesResponse{}
	err := client.Request("vpc", "ResetAttachCcnInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetAttachCcnInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
