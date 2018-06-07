package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 安全组批量绑定云资源
// https://cloud.tencent.com/document/api/236/15852

type AssociateSecurityGroupsRequest struct {
	// 实例ID列表，一个或者多个实例Id组成的数组。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 安全组Id。
	SecurityGroupId string `name:"SecurityGroupId"`
}

func (req *AssociateSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AssociateSecurityGroupsResponse, error) {
	resp := &AssociateSecurityGroupsResponse{}
	err := client.Request("cdb", "AssociateSecurityGroups", "2017-03-20").Do(req, resp)
	return resp, err
}

type AssociateSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
