package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 停用伸缩组
// https://cloud.tencent.com/document/api/377/20435

type DisableAutoScalingGroupRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DisableAutoScalingGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DisableAutoScalingGroupResponse, error) {
	resp := &DisableAutoScalingGroupResponse{}
	err := client.Request("as", "DisableAutoScalingGroup", "2018-04-19").Do(req, resp)
	return resp, err
}

type DisableAutoScalingGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
