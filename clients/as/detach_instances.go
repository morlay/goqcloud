package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 从伸缩组移出 CVM 实例
// https://cloud.tencent.com/document/api/377/20436

type DetachInstancesRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// CVM实例ID列表
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DetachInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DetachInstancesResponse, error) {
	resp := &DetachInstancesResponse{}
	err := client.Request("as", "DetachInstances", "2018-04-19").Do(req, resp)
	return resp, err
}

type DetachInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
