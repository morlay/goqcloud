package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 云联网接受关联实例
// https://cloud.tencent.com/document/api/215/20251

type AcceptAttachCcnInstancesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// 接受关联实例列表。
	Instances []*CcnInstance `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *AcceptAttachCcnInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AcceptAttachCcnInstancesResponse, error) {
	resp := &AcceptAttachCcnInstancesResponse{}
	err := client.Request("vpc", "AcceptAttachCcnInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type AcceptAttachCcnInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
