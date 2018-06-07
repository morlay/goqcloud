package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例名称
// https://cloud.tencent.com/document/api/237/16190

type ModifyDbInstanceNameRequest struct {
	// 待修改的实例 ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 新的实例名称。允许的字符为字母、数字、下划线、连字符和中文。
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceNameRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceNameResponse, error) {
	resp := &ModifyDbInstanceNameResponse{}
	err := client.Request("mariadb", "ModifyDBInstanceName", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceNameResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
