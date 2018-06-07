package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库账户的权限信息
// https://cloud.tencent.com/document/api/236/17500

type DescribeAccountPrivilegesRequest struct {
	// 数据库的账号域名。
	Host string `name:"Host"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 数据库的账号名称。
	User string `name:"User"`
}

func (req *DescribeAccountPrivilegesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountPrivilegesResponse, error) {
	resp := &DescribeAccountPrivilegesResponse{}
	err := client.Request("cdb", "DescribeAccountPrivileges", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeAccountPrivilegesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据库表中的列权限数组。
	ColumnPrivileges []*ColumnPrivilege `json:"ColumnPrivileges"`
	// 数据库权限数组。
	DatabasePrivileges []*DatabasePrivilege `json:"DatabasePrivileges"`
	// 全局权限数组。
	GlobalPrivileges []*string `json:"GlobalPrivileges"`
	// 数据库中的表权限数组。
	TablePrivileges []*TablePrivilege `json:"TablePrivileges"`
}
