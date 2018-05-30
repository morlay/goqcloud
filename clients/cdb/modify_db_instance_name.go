package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例名
// https://cloud.tencent.com/document/api/236/15877
type ModifyDbInstanceNameRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId string `name:"InstanceId"`
	// 实例名称。
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceNameRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceNameResponse, error) {
	resp := &ModifyDbInstanceNameResponse{}
	err := client.Request("cdb", "ModifyDBInstanceName", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceNameResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
