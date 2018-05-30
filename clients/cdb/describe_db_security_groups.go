package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例安全组信息
// https://cloud.tencent.com/document/api/236/15854
type DescribeDbSecurityGroupsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbSecurityGroupsResponse, error) {
	resp := &DescribeDbSecurityGroupsResponse{}
	err := client.Request("cdb", "DescribeDBSecurityGroups", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组详情。
	Groups []*SecurityGroup `json:"Groups"`
}
