package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例的所属项目
// https://cloud.tencent.com/document/api/236/15868

type ModifyDbInstanceProjectRequest struct {
	// 实例ID数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceIds []*string `name:"InstanceIds"`
	// 项目的ID。
	NewProjectId *int64 `name:"NewProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstanceProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceProjectResponse, error) {
	resp := &ModifyDbInstanceProjectResponse{}
	err := client.Request("cdb", "ModifyDBInstanceProject", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
