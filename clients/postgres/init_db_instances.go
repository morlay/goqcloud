package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 初始化实例
// https://cloud.tencent.com/document/api/409/16774

type InitDbInstancesRequest struct {
	// 实例根账号用户名。
	AdminName string `name:"AdminName"`
	// 实例根账号用户名对应的密码。
	AdminPassword string `name:"AdminPassword"`
	// 实例字符集，目前只支持：UTF8、LATIN1。
	Charset string `name:"Charset"`
	// 实例ID集合。
	DBInstanceIdSet []*string `name:"DBInstanceIdSet"`
	// 区域
	Region string `name:"Region"`
}

func (req *InitDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InitDbInstancesResponse, error) {
	resp := &InitDbInstancesResponse{}
	err := client.Request("postgres", "InitDBInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type InitDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例ID集合。
	DBInstanceIdSet []*string `json:"DBInstanceIdSet"`
}
