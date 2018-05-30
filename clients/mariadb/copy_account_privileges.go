package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 复制账号权限
// https://cloud.tencent.com/document/api/237/16169
type CopyAccountPrivilegesRequest struct {
	// 目的用户允许的访问 host
	DstHost string `name:"DstHost"`
	// 目的账号的 ReadOnly 属性
	DstReadOnly *string `name:"DstReadOnly,omitempty"`
	// 目的用户名
	DstUserName string `name:"DstUserName"`
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 源用户允许的访问 host
	SrcHost string `name:"SrcHost"`
	// 源账号的 ReadOnly 属性
	SrcReadOnly *string `name:"SrcReadOnly,omitempty"`
	// 源用户名
	SrcUserName string `name:"SrcUserName"`
}

func (req *CopyAccountPrivilegesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CopyAccountPrivilegesResponse, error) {
	resp := &CopyAccountPrivilegesResponse{}
	err := client.Request("mariadb", "CopyAccountPrivileges", "2017-03-12").Do(req, resp)
	return resp, err
}

type CopyAccountPrivilegesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
