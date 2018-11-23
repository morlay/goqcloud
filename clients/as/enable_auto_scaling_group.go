package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 启用伸缩组
// https://cloud.tencent.com/document/api/377/20434

type EnableAutoScalingGroupRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// 区域
	Region string `name:"Region"`
}

func (req *EnableAutoScalingGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*EnableAutoScalingGroupResponse, error) {
	resp := &EnableAutoScalingGroupResponse{}
	err := client.Request("as", "EnableAutoScalingGroup", "2018-04-19").Do(req, resp)
	return resp, err
}

type EnableAutoScalingGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
