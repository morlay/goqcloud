package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 云联网关联实例
// https://cloud.tencent.com/document/api/215/19205

type AttachCcnInstancesRequest struct {
	// CCN实例ID。形如：ccn-f49l6u0z。
	CcnId string `name:"CcnId"`
	// CCN所属UIN（根账号），默认当前账号所属UIN
	CcnUin *string `name:"CcnUin,omitempty"`
	// 关联网络实例列表
	Instances []*CcnInstance `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *AttachCcnInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AttachCcnInstancesResponse, error) {
	resp := &AttachCcnInstancesResponse{}
	err := client.Request("vpc", "AttachCcnInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type AttachCcnInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
