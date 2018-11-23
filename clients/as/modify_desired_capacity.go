package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改期望实例数
// https://cloud.tencent.com/document/api/377/20432

type ModifyDesiredCapacityRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// 期望实例数
	DesiredCapacity int64 `name:"DesiredCapacity"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDesiredCapacityRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDesiredCapacityResponse, error) {
	resp := &ModifyDesiredCapacityResponse{}
	err := client.Request("as", "ModifyDesiredCapacity", "2018-04-19").Do(req, resp)
	return resp, err
}

type ModifyDesiredCapacityResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
