package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除伸缩组
// https://cloud.tencent.com/document/api/377/20439

type DeleteAutoScalingGroupRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteAutoScalingGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAutoScalingGroupResponse, error) {
	resp := &DeleteAutoScalingGroupResponse{}
	err := client.Request("as", "DeleteAutoScalingGroup", "2018-04-19").Do(req, resp)
	return resp, err
}

type DeleteAutoScalingGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
