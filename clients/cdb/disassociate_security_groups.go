package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 安全组批量解绑云资源
// https://cloud.tencent.com/document/api/236/15851

type DisassociateSecurityGroupsRequest struct {
	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 安全组Id。
	SecurityGroupId string `name:"SecurityGroupId"`
}

func (req *DisassociateSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DisassociateSecurityGroupsResponse, error) {
	resp := &DisassociateSecurityGroupsResponse{}
	err := client.Request("cdb", "DisassociateSecurityGroups", "2017-03-20").Do(req, resp)
	return resp, err
}

type DisassociateSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
