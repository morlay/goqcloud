package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 添加 CVM 实例到伸缩组
// https://cloud.tencent.com/document/api/377/20441

type AttachInstancesRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// CVM实例ID列表
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *AttachInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AttachInstancesResponse, error) {
	resp := &AttachInstancesResponse{}
	err := client.Request("as", "AttachInstances", "2018-04-19").Do(req, resp)
	return resp, err
}

type AttachInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
