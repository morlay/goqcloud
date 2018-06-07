package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建子网
// https://cloud.tencent.com/document/api/215/15782

type CreateSubnetRequest struct {
	// 子网网段，子网网段必须在VPC网段内，相同VPC内子网网段不能重叠。
	CidrBlock string `name:"CidrBlock"`
	// 区域
	Region string `name:"Region"`
	// 子网名称，最大长度不能超过60个字节。
	SubnetName string `name:"SubnetName"`
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
	// 子网所在的可用区ID，不同子网选择不同可用区可以做跨可用区灾备。
	Zone string `name:"Zone"`
}

func (req *CreateSubnetRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSubnetResponse, error) {
	resp := &CreateSubnetResponse{}
	err := client.Request("vpc", "CreateSubnet", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateSubnetResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 子网对象。
	Subnet Subnet `json:"Subnet"`
}
