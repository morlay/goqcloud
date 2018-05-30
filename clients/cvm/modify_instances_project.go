package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例所属项目
// https://cloud.tencent.com/document/api/213/15746
type ModifyInstancesProjectRequest struct {
	// 一个或多个待操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `name:"InstanceIds"`
	// 项目ID。项目可以使用AddProject接口创建。后续使用DescribeInstances接口查询实例时，项目ID可用于过滤结果。
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyInstancesProjectRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstancesProjectResponse, error) {
	resp := &ModifyInstancesProjectResponse{}
	err := client.Request("cvm", "ModifyInstancesProject", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyInstancesProjectResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
