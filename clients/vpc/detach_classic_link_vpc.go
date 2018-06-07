package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除基础网络互通
// https://cloud.tencent.com/document/api/215/15777

type DetachClassicLinkVpcRequest struct {
	// CVM实例ID查询。形如：ins-r8hr2upy。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
}

func (req *DetachClassicLinkVpcRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DetachClassicLinkVpcResponse, error) {
	resp := &DetachClassicLinkVpcResponse{}
	err := client.Request("vpc", "DetachClassicLinkVpc", "2017-03-12").Do(req, resp)
	return resp, err
}

type DetachClassicLinkVpcResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
