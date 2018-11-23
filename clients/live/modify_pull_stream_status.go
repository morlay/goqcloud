package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改拉流配置状态
// https://cloud.tencent.com/document/api/267/30156

type ModifyPullStreamStatusRequest struct {
	// 配置id列表。
	ConfigIds []*string `name:"ConfigIds"`
	// 区域
	Region string `name:"Region"`
	// 目标状态。0无效，2正在运行，4暂停。
	Status string `name:"Status"`
}

func (req *ModifyPullStreamStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPullStreamStatusResponse, error) {
	resp := &ModifyPullStreamStatusResponse{}
	err := client.Request("live", "ModifyPullStreamStatus", "2018-08-01").Do(req, resp)
	return resp, err
}

type ModifyPullStreamStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
