package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询项目安全组信息
// https://cloud.tencent.com/document/api/236/15850
type DescribeProjectSecurityGroupsRequest struct {
	// 项目ID。
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeProjectSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProjectSecurityGroupsResponse, error) {
	resp := &DescribeProjectSecurityGroupsResponse{}
	err := client.Request("cdb", "DescribeProjectSecurityGroups", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeProjectSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups"`
}
