package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例所属项目
// https://cloud.tencent.com/document/api/557/19983

type ModifyDbInstancesProjectRequest struct {
	// 待修改的实例ID列表。实例 ID 形如：dcdbt-ow728lmc。
	InstanceIds []*string `name:"InstanceIds"`
	// 要分配的项目 ID，可以通过 DescribeProjects 查询项目列表接口获取。
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbInstancesProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstancesProjectResponse, error) {
	resp := &ModifyDbInstancesProjectResponse{}
	err := client.Request("dcdb", "ModifyDBInstancesProject", "2018-04-11").Do(req, resp)
	return resp, err
}

type ModifyDbInstancesProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
