package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除VPC
// https://cloud.tencent.com/document/api/215/15775

type DeleteVpcRequest struct {
	// 区域
	Region string `name:"Region"`
	// VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
}

func (req *DeleteVpcRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteVpcResponse, error) {
	resp := &DeleteVpcResponse{}
	err := client.Request("vpc", "DeleteVpc", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteVpcResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
