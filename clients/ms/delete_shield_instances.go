package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量删除提交过的app信息
// https://cloud.tencent.com/document/api/283/17752

type DeleteShieldInstancesRequest struct {
	// 任务唯一标识ItemId的列表
	ItemIds []*string `name:"ItemIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteShieldInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteShieldInstancesResponse, error) {
	resp := &DeleteShieldInstancesResponse{}
	err := client.Request("ms", "DeleteShieldInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type DeleteShieldInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}
