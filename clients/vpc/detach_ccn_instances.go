package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 云联网解关联实例
// https://cloud.tencent.com/document/api/215/19198

type DetachCcnInstancesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 要解关联网络实例列表
	Instances []*CcnInstance `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *DetachCcnInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DetachCcnInstancesResponse, error) {
	resp := &DetachCcnInstancesResponse{}
	err := client.Request("vpc", "DetachCcnInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type DetachCcnInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
