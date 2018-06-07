package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库安全组
// https://cloud.tencent.com/document/api/236/15853

type ModifyDbInstanceSecurityGroupsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。
	SecurityGroupIds []*string `name:"SecurityGroupIds"`
}

func (req *ModifyDbInstanceSecurityGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceSecurityGroupsResponse, error) {
	resp := &ModifyDbInstanceSecurityGroupsResponse{}
	err := client.Request("cdb", "ModifyDBInstanceSecurityGroups", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceSecurityGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
