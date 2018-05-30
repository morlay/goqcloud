package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建基础网络互通
// https://cloud.tencent.com/document/api/215/15779
type AttachClassicLinkVpcRequest struct {
	// CVM实例ID
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// VPC实例ID
	VpcId string `name:"VpcId"`
}

func (req *AttachClassicLinkVpcRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AttachClassicLinkVpcResponse, error) {
	resp := &AttachClassicLinkVpcResponse{}
	err := client.Request("vpc", "AttachClassicLinkVpc", "2017-03-12").Do(req, resp)
	return resp, err
}

type AttachClassicLinkVpcResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
